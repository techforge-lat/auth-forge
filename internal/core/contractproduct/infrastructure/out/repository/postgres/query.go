package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "contractproducts"

var sqlColumnByDomainField = map[string]string{
	"id":               "id",
	"contract_id":      "contract_id",
	"product_id":       "product_id",
	"product_price_id": "product_price_id",
	"supplier_price":   "supplier_price",
	"price":            "price",
	"price_type":       "price_type",
	"created_at":       "created_at",
	"updated_at":       "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "contract_id", "product_id", "product_price_id", "supplier_price", "price", "price_type", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("contract_id", "product_id", "product_price_id", "supplier_price", "price", "price_type", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "contract_id", "product_id", "product_price_id", "supplier_price", "price", "price_type", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
