package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "invoices"

var sqlColumnByDomainField = map[string]string{
	"id":                        "id",
	"billing_customer_id":       "billing_customer_id",
	"supplier_total_price":      "supplier_total_price",
	"total_price":               "total_price",
	"period_month":              "period_month",
	"period_year":               "period_year",
	"owner_customer_id":         "owner_customer_id",
	"hash":                      "hash",
	"notes":                     "notes",
	"state":                     "state",
	"invoice_calculation_id":    "invoice_calculation_id",
	"has_been_sent_to_customer": "has_been_sent_to_customer",
	"due_date":                  "due_date",
	"tax_amount":                "tax_amount",
	"detraction_amount":         "detraction_amount",
	"sub_total":                 "sub_total",
	"is_detraction_paid":        "is_detraction_paid",
	"pen_usd_exchange_rate":     "pen_usd_exchange_rate",
	"attachments":               "attachments",
	"created_at":                "created_at",
	"updated_at":                "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "billing_customer_id", "supplier_total_price", "total_price", "period_month", "period_year", "owner_customer_id", "hash", "notes", "state", "invoice_calculation_id", "has_been_sent_to_customer", "due_date", "tax_amount", "detraction_amount", "sub_total", "is_detraction_paid", "pen_usd_exchange_rate", "attachments", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("billing_customer_id", "supplier_total_price", "total_price", "period_month", "period_year", "owner_customer_id", "hash", "notes", "state", "invoice_calculation_id", "has_been_sent_to_customer", "due_date", "tax_amount", "detraction_amount", "sub_total", "is_detraction_paid", "pen_usd_exchange_rate", "attachments", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "billing_customer_id", "supplier_total_price", "total_price", "period_month", "period_year", "owner_customer_id", "hash", "notes", "state", "invoice_calculation_id", "has_been_sent_to_customer", "due_date", "tax_amount", "detraction_amount", "sub_total", "is_detraction_paid", "pen_usd_exchange_rate", "attachments", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
