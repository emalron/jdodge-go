package model

import(
	"database/sql"
	// for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"jdodge-go/util"
	G "jdodge-go/global"
)

// ShowAllRanks returns the []map[string]interface{} having records of view_ranking
func ShowAllRanks() []map[string]interface{} {
	type Rank struct {
		Name string
		Score int
		ReplayData string
		Time string
	}
	rank := Rank{}
	db, err := sql.Open("mysql", baseInfo["url"])
	if err != nil {
		util.PrintError("db error: ", err)
	}
	defer db.Close()
	resultSet, queryErr := db.Query("SELECT name, score, replay_data, time FROM view_ranking")
	if queryErr != nil {
		util.PrintError("query error", queryErr)
	}
	defer resultSet.Close()
	results := make([]map[string]interface{}, 0)
	for resultSet.Next() {
		scanErr := resultSet.Scan(rank)
		if scanErr != nil {
			util.PrintError("scan error", scanErr)
		}
		var result = map[string]interface{} {
			"name": rank.Name,
			"score": rank.Score,
			"replay_data": rank.ReplayData,
			"time": rank.Time,
		}
		results = append(results, result)
	}
	
	return results
}

func addRank(input []byte, db *sql.DB) string {
	type Command struct {
		ID string
		Score int
		ReplayData string
	}
	command := Command{}
	if jsonErr := json.Unmarshal(input, &command); jsonErr != nil {
		util.PrintError("json error @AddRank ", jsonErr)
	}
	results := make([]map[string]interface{}, 0)
	results = append(results, map[string]interface{} { "insert": "fail"})
	resultType := -1
	result, queryErr := db.Exec("INSERT INTO ranks(score,replay_data,users_id,time) VALUES (?,?,?,now())", command.Score, command.ReplayData, command.ID)
	if queryErr != nil {
		util.PrintError("query error ", queryErr)
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		resultType = 1
		results[0]["insert"] = "success"
	}
	return util.GetOutputJSON(resultType, results)
}

