package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "paymentaccounts"

var sqlColumnByDomainField = map[string]string{
	"id":                  "id",
	"currency_id":         "currency_id",
	"name":                "name",
	"account_holder_name": "account_holder_name",
	"account_number":      "account_number",
	"created_at":          "created_at",
	"updated_at":          "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "currency_id", "name", "account_holder_name", "account_number", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("currency_id", "name", "account_holder_name", "account_number", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "currency_id", "name", "account_holder_name", "account_number", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
