package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

func getOutput(resultType int, results []map[string]interface{}) string {
	var output = make(map[string]interface{}, 0)
	output["result"] = resultType
	output["message"] = results
	outputString := getJSONstring(output)
	return outputString
}

func showAllRanks(input []byte, db *sql.DB) string {
	var name, replayData, time string
	var score int
	resultSet, queryErr := db.Query("SELECT name, score, replay_data, time FROM view_ranking")
	if queryErr != nil {
		printError("query error", queryErr)
	}
	defer resultSet.Close()
	results := make([]map[string]interface{}, 0)
	resultType := -1
	for resultSet.Next() {
		scanErr := resultSet.Scan(&name, &score, &replayData, &time)
		if scanErr != nil {
			printError("scan error", scanErr)
		}
		var result = map[string]interface{} {
			"name": name,
			"score": score,
			"replay_data": replayData,
			"time": time,
		}
		results = append(results, result)
	}
	if len(results) > 0 {
		resultType = 2
	}
	return getOutput(resultType, results)
}

func addRank(input []byte, db *sql.DB) string {
	type Command struct {
		ID string
		Score int
		ReplayData string
	}
	command := Command{}
	if jsonErr := json.Unmarshal(input, &command); jsonErr != nil {
		printError("json error @AddRank ", jsonErr)
	}
	results := make([]map[string]interface{}, 0)
	results = append(results, map[string]interface{} { "insert": "fail"})
	resultType := -1
	result, queryErr := db.Exec("INSERT INTO ranks(score,replay_data,users_id,time) VALUES (?,?,?,now())", command.Score, command.ReplayData, command.ID)
	if queryErr != nil {
		printError("query error ", queryErr)
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		resultType = 1
		results[0]["insert"] = "success"
	}
	return getOutput(resultType, results)
}

