package main
import (
	"database/sql"
)

var baseInfo map[string]string
var serviceMap map[string]func([]byte, *sql.DB)string

func initial() {
	baseInfo = getPropertiesMap("properties")
	serviceMap = map[string]func([]byte, *sql.DB)string {
		"showAllRanks": showAllRanks,
		"addRank": addRank,
	}
}

