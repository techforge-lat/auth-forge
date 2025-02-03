package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "tenants"

var sqlColumnByDomainField = map[string]string{
	"id":         "id",
	"name":       "name",
	"code":       "code",
	"app_id":     "app_id",
	"created_at": "created_at",
	"updated_at": "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "name", "code", "app_id", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("name", "code", "app_id", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "name", "code", "app_id", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
