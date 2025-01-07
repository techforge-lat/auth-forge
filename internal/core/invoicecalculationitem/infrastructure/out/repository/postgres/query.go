package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "invoicecalculationitems"

var sqlColumnByDomainField = map[string]string{
	"id":                     "id",
	"invoice_calculation_id": "invoice_calculation_id",
	"contract_product_id":    "contract_product_id",
	"period_begins_date":     "period_begins_date",
	"period_ends_date":       "period_ends_date",
	"quantity":               "quantity",
	"supplier_unit_amount":   "supplier_unit_amount",
	"supplier_total_amount":  "supplier_total_amount",
	"unit_amount":            "unit_amount",
	"total_amount":           "total_amount",
	"created_at":             "created_at",
	"updated_at":             "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "invoice_calculation_id", "contract_product_id", "period_begins_date", "period_ends_date", "quantity", "supplier_unit_amount", "supplier_total_amount", "unit_amount", "total_amount", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("invoice_calculation_id", "contract_product_id", "period_begins_date", "period_ends_date", "quantity", "supplier_unit_amount", "supplier_total_amount", "unit_amount", "total_amount", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "invoice_calculation_id", "contract_product_id", "period_begins_date", "period_ends_date", "quantity", "supplier_unit_amount", "supplier_total_amount", "unit_amount", "total_amount", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
