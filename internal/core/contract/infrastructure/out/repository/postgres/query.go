package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "contracts"

var sqlColumnByDomainField = map[string]string{
	"id":                  "id",
	"owner_customer_id":   "owner_customer_id",
	"billing_customer_id": "billing_customer_id",
	"begins_at":           "begins_at",
	"ends_at":             "ends_at",
	"commitment":          "commitment",
	"payment_frequency":   "payment_frequency",
	"renewal_date":        "renewal_date",
	"state":               "state",
	"attachments":         "attachments",
	"created_at":          "created_at",
	"updated_at":          "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "owner_customer_id", "billing_customer_id", "begins_at", "ends_at", "commitment", "payment_frequency", "renewal_date", "state", "attachments", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("owner_customer_id", "billing_customer_id", "begins_at", "ends_at", "commitment", "payment_frequency", "renewal_date", "state", "attachments", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "owner_customer_id", "billing_customer_id", "begins_at", "ends_at", "commitment", "payment_frequency", "renewal_date", "state", "attachments", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
