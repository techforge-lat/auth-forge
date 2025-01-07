package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "invoicepayments"

var sqlColumnByDomainField = map[string]string{
	"id":                    "id",
	"invoice_id":            "invoice_id",
	"payment_account_id":    "payment_account_id",
	"payment_method_id":     "payment_method_id",
	"amount":                "amount",
	"notes":                 "notes",
	"exchange_rate":         "exchange_rate",
	"payment_date":          "payment_date",
	"hash":                  "hash",
	"reference_code":        "reference_code",
	"is_detraction":         "is_detraction",
	"pen_usd_exchange_rate": "pen_usd_exchange_rate",
	"created_at":            "created_at",
	"updated_at":            "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "invoice_id", "payment_account_id", "payment_method_id", "amount", "notes", "exchange_rate", "payment_date", "hash", "reference_code", "is_detraction", "pen_usd_exchange_rate", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("invoice_id", "payment_account_id", "payment_method_id", "amount", "notes", "exchange_rate", "payment_date", "hash", "reference_code", "is_detraction", "pen_usd_exchange_rate", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "invoice_id", "payment_account_id", "payment_method_id", "amount", "notes", "exchange_rate", "payment_date", "hash", "reference_code", "is_detraction", "pen_usd_exchange_rate", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
