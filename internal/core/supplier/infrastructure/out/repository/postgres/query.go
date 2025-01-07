package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "suppliers"

var sqlColumnByDomainField = map[string]string{
	"id":             "id",
	"name":           "name",
	"description":    "description",
	"price":          "price",
	"stock_quantity": "stock_quantity",
	"category_i_d":   "category_i_d",
	"created_at":     "created_at",
	"updated_at":     "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "name", "description", "price", "stock_quantity", "category_i_d", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("name", "description", "price", "stock_quantity", "category_i_d", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "name", "description", "price", "stock_quantity", "category_i_d", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
