package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

type Config struct {
	ModuleName   string  `json:"module_name"`
	DatabaseName string  `json:"database_name"`
	Fields       []Field `json:"fields"`
	PackagePath  string  `json:"package_path"`
}

type Field struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	IsNullable   bool   `json:"is_nullable"`
	NullType     string `json:"-"`
	AccessMethod string `json:"-"`
}

func main() {
	configFile := flag.String("config", "module-config.json", "Path to module configuration file")
	flag.Parse()

	config, err := loadConfig(*configFile)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	if err := generateModule(config); err != nil {
		fmt.Printf("Error generating module: %v\n", err)
		os.Exit(1)
	}

	if err := exec.Command("make", "fmt").Run(); err != nil {
		log.Fatal(err)
	}
}

func loadConfig(filepath string) (Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("error parsing config file: %w", err)
	}

	return config, nil
}

func generateModule(config Config) error {
	// Create directory structure
	dirs := []string{
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "application"),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "domain"),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "infrastructure", "in", "httprest"),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "infrastructure", "out", "repository", "postgres"),
		filepath.Join("pkg", "di"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	for i := range config.Fields {
		config.Fields[i].NullType = getNullType(config.Fields[i].Type, config.Fields[i].IsNullable, false)
		config.Fields[i].AccessMethod = getAccessMethod(config.Fields[i].Type)
	}

	// Generate files with templates
	files := map[string]string{
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "domain", "command.go"):                                        getDomainCommandTemplate(),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "domain", "query.go"):                                          getDomainQueryTemplate(),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "application", "usecase.go"):                                   getUseCaseTemplate(),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "infrastructure", "in", "httprest", "handler.go"):              getHandlerTemplate(),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "infrastructure", "out", "repository", "postgres", "psql.go"):  getRepositoryTemplate(),
		filepath.Join("internal", "core", strings.ToLower(config.ModuleName), "infrastructure", "out", "repository", "postgres", "query.go"): getSQLQueryTemplate(),
		filepath.Join("internal", "shared", "domain", "ports", "out", strings.ToLower(config.ModuleName)+".go"):                              getOutPortTemplate(),
		filepath.Join("internal", "shared", "domain", "ports", "in", strings.ToLower(config.ModuleName)+".go"):                               getInPortTemplate(),
		filepath.Join("pkg", "di", strings.ToLower(config.ModuleName)+".go"):                                                                 getDITemplate(),
	}

	for path, templateContent := range files {
		if err := generateFile(path, templateContent, config); err != nil {
			return fmt.Errorf("error generating file %s: %w", path, err)
		}
	}

	// Update dependency names
	if err := updateDependencyNames(config); err != nil {
		return fmt.Errorf("error updating dependency names: %w", err)
	}

	// Add routes
	if err := addRoutes(config); err != nil {
		return fmt.Errorf("error adding routes: %w", err)
	}

	return nil
}

func generateFile(path, templateContent string, config Config) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	funcs := template.FuncMap{
		"toLower":            strings.ToLower,
		"toUpper":            strings.ToUpper,
		"toSnakeCase":        toSnakeCase,
		"toCamelCase":        toCamelCase,
		"toUpperCase":        strings.ToUpper,
		"contains":           contains,
		"toKebabCase":        toKebabCase,
		"getValidatorRules":  getValidatorRules,
		"getValidatorMethod": getValidatorMethod,
		"getNullType": func(fieldType string, isNullable bool, forUpdate bool) string {
			return getNullType(fieldType, isNullable, forUpdate)
		},
	}

	tmpl, err := template.New(filepath.Base(path)).Funcs(funcs).Parse(templateContent)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, config)
}

func updateDependencyNames(config Config) error {
	const dependencyFile = "pkg/dependency/name.go"
	content := fmt.Sprintf(`
const (
	%sHandler    = "%s.handler"
	%sUseCase    = "%s.usecase"
	%sRepository = "%s.repository"
)`,
		config.ModuleName, strings.ToLower(config.ModuleName),
		config.ModuleName, strings.ToLower(config.ModuleName),
		config.ModuleName, strings.ToLower(config.ModuleName))

	// Append to existing file
	f, err := os.OpenFile(dependencyFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func addRoutes(config Config) error {
	routerPath := filepath.Join("cmd", "api", "router", fmt.Sprintf("%s.go", strings.ToLower(config.ModuleName)))

	routerContent := fmt.Sprintf(`package router

import (
	"%s/pkg/server"
	"%s/pkg/dependency"

	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/linkit"

	"cloud-crm-backend/internal/core/%s/infrastructure/in/httprest"
)

func %sRoutes(server *server.Server) error {
	handler, err := linkit.Resolve[httprest.Handler](server.Container, dependency.%sHandler)
	if err != nil {
		return errortrace.OnError(err)
	}

	group := server.Echo.Group("v1/%ss")

	group.POST("", handler.Create)
	group.PUT("", handler.Update)
	group.PUT("/:id", handler.UpdateByID)
	group.DELETE("/:id", handler.DeleteByID)
	group.DELETE("", handler.Delete)
	group.GET("/:id", handler.FindOneByID)
	group.GET("", handler.FindOne)
	group.GET("/all", handler.FindAll)

	return nil
}`, config.PackagePath, config.PackagePath, strings.ToLower(config.ModuleName), config.ModuleName, config.ModuleName, strings.ToLower(config.ModuleName))

	return os.WriteFile(routerPath, []byte(routerContent), 0644)
}

// Helper functions
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}

func getDomainCommandTemplate() string {
	return `package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/valid"
	"gopkg.in/guregu/null.v4"
)

var (
    {{range .Fields -}}
	{{if not (or (eq .Type "bool") (or (eq .Type "null.Bool") (eq .Type "json.RawMessage")))}}{{$.ModuleName | toCamelCase}}{{.Name}}Rules = {{getValidatorRules .Type .IsNullable}}{{end}}
    {{end -}}
)

// {{.ModuleName}}CreateRequest represents the request to create a {{.ModuleName}}
type {{.ModuleName}}CreateRequest struct {
    ID uint ` + "`json:\"id\"`" + `
{{range .Fields -}}
    {{.Name}} {{if .IsNullable}}{{getNullType .Type .IsNullable true}}{{else}}{{.Type}}{{end}} ` + "`json:\"{{.Name | toSnakeCase}}\"`" + `
{{end -}}
    CreatedAt time.Time ` + "`json:\"created_at\"`" + `
}

// Validate validates the fields of {{.ModuleName}}CreateRequest
func (c {{.ModuleName}}CreateRequest) Validate() error {
    v := valid.New()

    {{range .Fields -}}
		{{if not (or (eq .Type "bool") (or (eq .Type "null.Bool") (or (eq .Type "uuid.UUID") (eq .Type "json.RawMessage"))))}}
		{{if .IsNullable}}
				if c.{{.Name}}.Valid {
					v.{{getValidatorMethod .Type .IsNullable}}("{{.Name | toSnakeCase}}", {{if or (contains .Type "int") (contains .Type "float")}}{{.Type}}({{end}}c.{{.Name}}{{if .AccessMethod}}.{{.AccessMethod}}{{end}}{{if or (contains .Type "int") (contains .Type "float")}}){{end}}, {{$.ModuleName | toCamelCase}}{{.Name}}Rules...)
				}
			{{else}}
				v.{{getValidatorMethod .Type .IsNullable}}("{{.Name | toSnakeCase}}", c.{{.Name}}, {{$.ModuleName | toCamelCase}}{{.Name}}Rules...)
			{{end}}
		{{end}}
    {{end}}

    if v.HasErrors() {
        return errortrace.OnError(v.Errors())
    }

    return nil
}

// {{.ModuleName}}UpdateRequest represents the request to update a {{.ModuleName}}
type {{.ModuleName}}UpdateRequest struct {
{{- range .Fields}}
    {{.Name}} {{getNullType .Type .IsNullable true}} ` + "`json:\"{{.Name | toSnakeCase}}\"`" + `
{{- end}}
    UpdatedAt null.Time ` + "`json:\"updated_at\"`" + `
}

// Validate validates the fields of {{.ModuleName}}UpdateRequest
func (u {{.ModuleName}}UpdateRequest) Validate() error {
    v := valid.New()

    {{range .Fields -}}
	{{if not (or (eq .Type "bool") (or (eq .Type "null.Bool") (or (eq .Type "uuid.UUID") (eq .Type "json.RawMessage"))))}}
    if u.{{.Name}}.Valid {
        v.{{getValidatorMethod .Type true}}("{{.Name | toSnakeCase}}", {{if or (contains .Type "int") (contains .Type "float")}}{{.Type}}({{end}}u.{{.Name}}{{if .AccessMethod}}.{{.AccessMethod}}{{end}}{{if or (contains .Type "int") (contains .Type "float")}}){{end}}, {{$.ModuleName | toCamelCase}}{{.Name}}Rules...)
    }
	{{end}}
    {{end}}

    if v.HasErrors() {
        return errortrace.OnError(v.Errors())
    }

    return nil
}`
}

func getDomainQueryTemplate() string {
	return `package domain

import (
    "time"
    "github.com/google/uuid"
    "gopkg.in/guregu/null.v4"
)

type {{.ModuleName}} struct {
    ID        uuid.UUID ` + "`json:\"id\"`" + `
    {{range .Fields}}{{.Name}} {{.Type}} ` + "`json:\"{{.Name | toSnakeCase}}\"`" + `
    {{end -}}
    CreatedAt time.Time  ` + "`json:\"created_at\"`" + `
    UpdatedAt null.Time ` + "`json:\"updated_at\"`" + `
}`
}

func getUseCaseTemplate() string {
	return `package application

import (
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/domain"
	"{{.PackagePath}}/internal/shared/domain/ports/out"
	"context"

	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type UseCase struct {
	repo out.{{.ModuleName}}Repository
}

func NewUseCase(repo out.{{.ModuleName}}Repository) UseCase {
	return UseCase{repo: repo}
}

func (uc UseCase) Create(ctx context.Context, entity domain.{{.ModuleName}}CreateRequest) error {
	if err := entity.Validate(); err != nil {
		return errortrace.OnError(err)
	}

	err := uc.repo.Create(ctx, entity)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) Update(ctx context.Context, entity domain.{{.ModuleName}}UpdateRequest, filters ...dafi.Filter) error {
	if err := entity.Validate(); err != nil {
		return errortrace.OnError(err)
	}

	err := uc.repo.Update(ctx, entity)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) Delete(ctx context.Context, filters ...dafi.Filter) error {
	err := uc.repo.Delete(ctx, filters...)
	if err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (uc UseCase) FindOne(ctx context.Context, criteria dafi.Criteria) (domain.{{.ModuleName}}, error) {
	result, err := uc.repo.FindOne(ctx, criteria)
	if err != nil {
		return domain.{{.ModuleName}}{}, errortrace.OnError(err)
	}

	return result, nil
}

func (uc UseCase) FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.{{.ModuleName}}, error) {
	result, err := uc.repo.FindAll(ctx, criteria)
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	return result, nil
}`
}

func getHandlerTemplate() string {
	return `package httprest

import (
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/domain"
	"{{.PackagePath}}/internal/shared/domain/ports/in"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
	"github.com/techforge-lat/errortrace/v2/errtype"
	"github.com/techforge-lat/rapi"
)

type Handler struct {
	useCase in.{{.ModuleName}}UseCase
}

func New(useCase in.{{.ModuleName}}UseCase) Handler {
	return Handler{useCase: useCase}
}

func (h Handler) Create(c echo.Context) error {
	entity := domain.{{.ModuleName}}CreateRequest{}

	if err := c.Bind(&entity); err != nil {
		return errortrace.OnError(err).WithCode(errtype.UnprocessableEntity)
	}

	if err := h.useCase.Create(c.Request().Context(), entity); err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusCreated, rapi.Created(entity))
}

func (h Handler) UpdateByID(c echo.Context) error {
	entity := domain.{{.ModuleName}}UpdateRequest{}

	if err := c.Bind(&entity); err != nil {
		return errortrace.OnError(err).WithCode(errtype.UnprocessableEntity)
	}

	if err := h.useCase.Update(c.Request().Context(), entity, dafi.FilterBy("id", dafi.Equal, c.Param("id"))...); err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Updated())
}

func (h Handler) Update(c echo.Context) error {
	entity := domain.{{.ModuleName}}UpdateRequest{}

	if err := c.Bind(&entity); err != nil {
		return errortrace.OnError(err).WithCode(errtype.UnprocessableEntity)
	}

	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	if err := h.useCase.Update(c.Request().Context(), entity, criteria.Filters...); err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Updated())
}

func (h Handler) DeleteByID(c echo.Context) error {
	if err := h.useCase.Delete(c.Request().Context(), dafi.FilterBy("id", dafi.Equal, c.Param("id"))...); err != nil {
		return errortrace.OnError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h Handler) Delete(c echo.Context) error {
	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	if err := h.useCase.Delete(c.Request().Context(), criteria.Filters...); err != nil {
		return errortrace.OnError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h Handler) FindOneByID(c echo.Context) error {
	result, err := h.useCase.FindOne(c.Request().Context(), dafi.Where("id", dafi.Equal, c.Param("id")))
	if err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Ok(result))
}

func (h Handler) FindOne(c echo.Context) error {
	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	result, err := h.useCase.FindOne(c.Request().Context(), criteria)
	if err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Ok(result))
}

func (h Handler) FindAll(c echo.Context) error {
	criteria, err := dafi.NewQueryParser().Parse(c.QueryParams())
	if err != nil {
		return errortrace.OnError(err)
	}

	result, err := h.useCase.FindAll(c.Request().Context(), criteria)
	if err != nil {
		return errortrace.OnError(err)
	}

	return c.JSON(http.StatusOK, rapi.Ok(result))
}`
}

func getRepositoryTemplate() string {
	return `package postgres

import (
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/domain"
	"{{.PackagePath}}/internal/shared/domain/ports/out"
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/techforge-lat/dafi/v2"
	"github.com/techforge-lat/errortrace/v2"
)

type Repository struct {
	db out.Database
	tx out.Tx
}

func NewRepository(db out.Database) Repository {
	return Repository{db: db}
}

// WithTx returns a new instance of the repository with the transaction set
func (r Repository) WithTx(tx out.Transaction) out.{{.ModuleName}}Repository {
	return Repository{
		db: r.db,
		tx: tx.GetTx(),
	}
}

func (r Repository) Create(ctx context.Context, entity domain.{{.ModuleName}}CreateRequest) error {
	result, err := insertQuery.WithValues({{range $i, $f := .Fields}}{{if $i}}, {{end}}entity.{{$f.Name}}{{end}}, entity.CreatedAt).ToSQL()
	if err != nil {
		return errortrace.OnError(err)
	}

	if _, err := r.conn().Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (r Repository) Update(ctx context.Context, entity domain.{{.ModuleName}}UpdateRequest, filters ...dafi.Filter) error {
	if !entity.UpdatedAt.Valid {
		entity.UpdatedAt.SetValid(time.Now())
	}

	result, err := updateQuery.WithValues({{range $i, $f := .Fields}}{{if $i}}, {{end}}entity.{{$f.Name}}{{end}}, entity.UpdatedAt).Where(filters...).ToSQL()
	if err != nil {
		return errortrace.OnError(err)
	}

	if _, err := r.conn().Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, filters ...dafi.Filter) error {
	result, err := deleteQuery.Where(filters...).ToSQL()
	if err != nil {
		return errortrace.OnError(err)
	}

	if _, err := r.conn().Exec(ctx, result.Sql, result.Args...); err != nil {
		return errortrace.OnError(err)
	}

	return nil
}

func (r Repository) FindOne(ctx context.Context, criteria dafi.Criteria) (domain.{{.ModuleName}}, error) {
	result, err := selectQuery.Where(criteria.Filters...).OrderBy(criteria.Sorts...).Limit(1).RequiredColumns(criteria.SelectColumns...).ToSQL()
	if err != nil {
		return domain.{{.ModuleName}}{}, errortrace.OnError(err)
	}

	var m domain.{{.ModuleName}}
	if err := pgxscan.Get(ctx, r.conn(), &m, result.Sql, result.Args...); err != nil {
		return domain.{{.ModuleName}}{}, errortrace.OnError(err)
	}

	return m, nil
}

func (r Repository) FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.{{.ModuleName}}, error) {
	result, err := selectQuery.Where(criteria.Filters...).OrderBy(criteria.Sorts...).Limit(criteria.Pagination.PageSize).Page(criteria.Pagination.PageNumber).RequiredColumns(criteria.SelectColumns...).ToSQL()
	if err != nil {
		return nil, errortrace.OnError(err)
	}

	var ms []domain.{{.ModuleName}}
	if err := pgxscan.Select(ctx, r.conn(), &ms, result.Sql, result.Args...); err != nil {
		return nil, errortrace.OnError(err)
	}

	return ms, nil
}

// conn returns the database connection to use
// if there is a transaction, it returns the transaction connection
func (r Repository) conn() out.Database {
	if r.tx != nil {
		return r.tx
	}

	return r.db
}`
}

func getSQLQueryTemplate() string {
	return `package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "{{.ModuleName | toLower}}s"

var sqlColumnByDomainField = map[string]string{
	"id": "id",
	{{range .Fields}}"{{.Name | toSnakeCase}}": "{{.Name | toSnakeCase}}",
	{{end}}"created_at": "created_at",
	"updated_at": "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", {{range .Fields}}"{{.Name | toSnakeCase}}", {{end}}"created_at")
	updateQuery = sqlcraft.Update(table).WithColumns({{range .Fields}}"{{.Name | toSnakeCase}}", {{end}}"updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", {{range .Fields}}"{{.Name | toSnakeCase}}", {{end}}"created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)`
}

func getDITemplate() string {
	return `package di

import (
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/application"
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/infrastructure/in/httprest"
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/infrastructure/out/repository/postgres"
	"{{.PackagePath}}/pkg/database"
	"{{.PackagePath}}/pkg/dependency"
)

func provide{{.ModuleName}}Dependencies(container *linkit.DependencyContainer, db *database.Adapter) {
	repo := postgres.NewRepository(db)
	container.Provide(dependency.{{.ModuleName}}Repository, repo)

	useCase := application.NewUseCase(repo)
	container.Provide(dependency.{{.ModuleName}}UseCase, useCase)

	handler := httprest.New(useCase)
	container.Provide(dependency.{{.ModuleName}}Handler, handler)
}
`
}

// Helper function to generate interface files
func getPortsTemplate() string {
	return `package ports

import (
	"{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/domain"
	"context"
	"github.com/techforge-lat/dafi/v2"
)

type {{.ModuleName}}Repository interface {
	Create(ctx context.Context, entity domain.{{.ModuleName}}CreateRequest) error
	Update(ctx context.Context, entity domain.{{.ModuleName}}UpdateRequest, filters ...dafi.Filter) error
	Delete(ctx context.Context, filters ...dafi.Filter) error
	FindOne(ctx context.Context, criteria dafi.Criteria) (domain.{{.ModuleName}}, error)
	FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.{{.ModuleName}}, error)
	WithTx(tx Transaction) {{.ModuleName}}Repository
}

type {{.ModuleName}}UseCase interface {
	Create(ctx context.Context, entity domain.{{.ModuleName}}CreateRequest) error
	Update(ctx context.Context, entity domain.{{.ModuleName}}UpdateRequest, filters ...dafi.Filter) error
	Delete(ctx context.Context, filters ...dafi.Filter) error
	FindOne(ctx context.Context, criteria dafi.Criteria) (domain.{{.ModuleName}}, error)
	FindAll(ctx context.Context, criteria dafi.Criteria) ([]domain.{{.ModuleName}}, error)
}`
}

func getOutPortTemplate() string {
	return `package out

import "{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/domain"

type {{.ModuleName}}Repository interface {
    RepositoryTx[{{.ModuleName}}Repository]
    RepositoryCommand[domain.{{.ModuleName}}CreateRequest, domain.{{.ModuleName}}UpdateRequest]
    RepositoryQuery[domain.{{.ModuleName}}]
}`
}

func getInPortTemplate() string {
	return `package in

import "{{.PackagePath}}/internal/core/{{.ModuleName | toLower}}/domain"

type {{.ModuleName}}UseCase interface {
    UseCaseCommand[domain.{{.ModuleName}}CreateRequest, domain.{{.ModuleName}}UpdateRequest]
    UseCaseQuery[domain.{{.ModuleName}}]
}`
}

func getValidatorRules(fieldType string, isNullable bool) string {
	switch fieldType {
	case "string", "null.String":
		if isNullable {
			return "valid.StringRules().Build()"
		}
		return "valid.StringRules().Required().Build()"
	case "null.Int":
		return fmt.Sprintf("valid.NumberRules[%s]().Required().Build()", getPrimitiveByNull(fieldType))
	case "null.Float":
		return fmt.Sprintf("valid.FloatRules[%s]().Required().Build()", getPrimitiveByNull(fieldType))
	case "uint", "uint64", "int", "int64":
		if isNullable {
			return fmt.Sprintf("valid.NumberRules[%s]().Build()", fieldType)
		}
		return fmt.Sprintf("valid.NumberRules[%s]().Required().Build()", fieldType)
	case "float32", "float64":
		if isNullable {
			return fmt.Sprintf("valid.FloatRules[%s]().Build()", fieldType)
		}
		return fmt.Sprintf("valid.FloatRules[%s]().Required().Build()", fieldType)
	case "time.Time", "null.Time":
		if isNullable {
			return "valid.TimeRules().Build()"
		}
		return "valid.TimeRules().Required().Build()"
	case "bool":
		return ""
	default:
		return "valid.StringRules().Build()"
	}
}

func getValidatorMethod(fieldType string, isNullable bool) string {
	switch fieldType {
	case "string", "null.String":
		return "String"
	case "uint", "uint64":
		return "Uint"
	case "int", "null.Int":
		return "Int"
	case "float64", "float32", "null.Float":
		return "Float64"
	case "time.Time", "null.Time":
		return "Time"
	case "uuid.UUID":
		return "NullUUID"
	case "bool":
		return ""
	default:
		return "String"
	}
}

func getPrimitiveByNull(fieldType string) string {
	switch fieldType {
	case "null.String":
		return "string"
	case "null.Int":
		return "int64"
	case "null.Float":
		return "float64"
	case "null.Time":
		return "time.Time"
	case "null.Bool":
		return "bool"
	default:
		return "string"
	}
}

func getNullType(fieldType string, isNullable bool, forUpdate bool) string {
	// Si es para la estructura Update, siempre retornamos el tipo nulo
	if forUpdate {
		switch fieldType {
		case "string":
			return "null.String"
		case "int", "uint", "int64":
			return "null.Int"
		case "float64":
			return "null.Float"
		case "bool":
			return "null.Bool"
		case "time.Time":
			return "null.Time"
		case "json.RawMessage":
			return "json.RawMessage"
		default:
			return fieldType
		}
	}

	// Para otras estructuras, respetamos el isNullable
	if !isNullable {
		return fieldType
	}

	switch fieldType {
	case "string":
		return "null.String"
	case "int", "uint":
		return "null.Int"
	case "float64":
		return "null.Float"
	case "bool":
		return "null.Bool"
	case "time.Time":
		return "null.Time"
	case "json.RawMessage":
		return "json.RawMessage"
	default:
		return fieldType
	}
}

// contains es una función auxiliar que verifica si una cadena contiene otra
func contains(s, substr string) bool {
	if strings.Contains(strings.ToLower(s), "null.") {
		return false
	}

	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func toKebabCase(s string) string {
	// First convert snake_case to space separated
	s = strings.ReplaceAll(s, "_", " ")

	// Convert camelCase or PascalCase to space separated
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			if !unicode.IsUpper(rune(s[i-1])) {
				result.WriteRune(' ')
			}
		}
		result.WriteRune(r)
	}
	s = result.String()

	// Convert to lowercase
	s = strings.ToLower(s)

	// Replace spaces with hyphens
	s = strings.ReplaceAll(s, " ", "-")

	// Remove any non-alphanumeric characters (except hyphens)
	reg := regexp.MustCompile(`[^a-z0-9\-]`)
	s = reg.ReplaceAllString(s, "")

	// Replace multiple hyphens with single hyphen
	reg = regexp.MustCompile(`-+`)
	s = reg.ReplaceAllString(s, "-")

	// Trim hyphens from start and end
	return strings.Trim(s, "-")
}

func getAccessMethod(fieldType string) string {
	switch fieldType {
	case "string", "null.String":
		return "String"
	case "int", "uint", "int64", "null.Int":
		return "Int64"
	case "float64", "null.Float":
		return "Float64"
	case "bool", "null.Bool":
		return "Bool"
	case "time.Time", "null.Time":
		return "Time"
	case "json.RawMessage":
		return ""
	default:
		return ""
	}
}
