package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "productprices"

var sqlColumnByDomainField = map[string]string{
	"id":                  "id",
	"product_id":          "product_id",
	"begins_at":           "begins_at",
	"ends_at":             "ends_at",
	"is_active":           "is_active",
	"supplier_price":      "supplier_price",
	"price":               "price",
	"contract_commitment": "contract_commitment",
	"payment_frequency":   "payment_frequency",
	"created_at":          "created_at",
	"updated_at":          "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "product_id", "begins_at", "ends_at", "is_active", "supplier_price", "price", "contract_commitment", "payment_frequency", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("product_id", "begins_at", "ends_at", "is_active", "supplier_price", "price", "contract_commitment", "payment_frequency", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "product_id", "begins_at", "ends_at", "is_active", "supplier_price", "price", "contract_commitment", "payment_frequency", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
