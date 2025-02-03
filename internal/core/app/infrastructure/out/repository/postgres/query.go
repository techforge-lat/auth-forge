package postgres

import "github.com/techforge-lat/sqlcraft"

var table = "apps"

var sqlColumnByDomainField = map[string]string{
	"id":         "id",
	"name":       "name",
	"code":       "code",
	"created_at": "created_at",
	"updated_at": "updated_at",
}

var (
	insertQuery = sqlcraft.InsertInto(table).WithColumns("id", "name", "code", "created_at")
	updateQuery = sqlcraft.Update(table).WithColumns("name", "updated_at").SQLColumnByDomainField(sqlColumnByDomainField).WithPartialUpdate()
	deleteQuery = sqlcraft.DeleteFrom(table).SQLColumnByDomainField(sqlColumnByDomainField)
	selectQuery = sqlcraft.Select("id", "name", "code", "created_at", "updated_at").From(table).SQLColumnByDomainField(sqlColumnByDomainField)
)
