package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "invoiceitems"

var sqlColumnByDomainField = map[string]string{
	"id":                   "id",
	"invoice_id":           "invoice_id",
	"quantity":             "quantity",
	"supplier_unit_price":  "supplier_unit_price",
	"supplier_total_price": "supplier_total_price",
	"unit_price":           "unit_price",
	"total_price":          "total_price",
	"contract_product_id":  "contract_product_id",
	"product_id":           "product_id",
	"description":          "description",
	"created_at":           "created_at",
	"updated_at":           "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "invoice_id", "quantity", "supplier_unit_price", "supplier_total_price", "unit_price", "total_price", "contract_product_id", "product_id", "description", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("invoice_id", "quantity", "supplier_unit_price", "supplier_total_price", "unit_price", "total_price", "contract_product_id", "product_id", "description", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "invoice_id", "quantity", "supplier_unit_price", "supplier_total_price", "unit_price", "total_price", "contract_product_id", "product_id", "description", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
