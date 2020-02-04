package main
import (
	"database/sql"
)

var base_info map[string]string
var service_map map[string]func([]byte, *sql.DB)string

func initial() {
	base_info = GetPropertiesMap("properties")
	service_map = map[string]func([]byte, *sql.DB)string {
		"showAllRanks": ShowAllRanks,
		"addRank": AddRank,
	}
}

