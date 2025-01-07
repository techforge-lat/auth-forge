package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "products"

var sqlColumnByDomainField = map[string]string{
	"id":          "id",
	"name":        "name",
	"description": "description",
	"supplier":    "supplier",
	"metadata":    "metadata",
	"sku":         "sku",
	"slug":        "slug",
	"created_at":  "created_at",
	"updated_at":  "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "name", "description", "supplier", "metadata", "sku", "slug", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("name", "description", "supplier", "metadata", "sku", "slug", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "name", "description", "supplier", "metadata", "sku", "slug", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
