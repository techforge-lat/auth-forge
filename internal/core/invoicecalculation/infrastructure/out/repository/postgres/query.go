package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "invoicecalculations"

var sqlColumnByDomainField = map[string]string{
	"id":                    "id",
	"billing_customer_id":   "billing_customer_id",
	"consumer_customer_id":  "consumer_customer_id",
	"supplier_total_amount": "supplier_total_amount",
	"total_amount":          "total_amount",
	"period_begins_date":    "period_begins_date",
	"period_ends_date":      "period_ends_date",
	"hash":                  "hash",
	"notes":                 "notes",
	"file_url":              "file_url",
	"period_year":           "period_year",
	"period_month":          "period_month",
	"attachments":           "attachments",
	"created_at":            "created_at",
	"updated_at":            "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "billing_customer_id", "consumer_customer_id", "supplier_total_amount", "total_amount", "period_begins_date", "period_ends_date", "hash", "notes", "file_url", "period_year", "period_month", "attachments", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("billing_customer_id", "consumer_customer_id", "supplier_total_amount", "total_amount", "period_begins_date", "period_ends_date", "hash", "notes", "file_url", "period_year", "period_month", "attachments", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "billing_customer_id", "consumer_customer_id", "supplier_total_amount", "total_amount", "period_begins_date", "period_ends_date", "hash", "notes", "file_url", "period_year", "period_month", "attachments", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
