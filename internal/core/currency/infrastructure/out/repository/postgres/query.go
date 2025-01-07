package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "currencys"

var sqlColumnByDomainField = map[string]string{
	"id":            "id",
	"name":          "name",
	"symbol":        "symbol",
	"exchange_rate": "exchange_rate",
	"created_at":    "created_at",
	"updated_at":    "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "name", "symbol", "exchange_rate", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("name", "symbol", "exchange_rate", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "name", "symbol", "exchange_rate", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
