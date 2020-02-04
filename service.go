package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

func GetOutput(result_type int, results []map[string]interface{}) string {
	var output = make(map[string]interface{}, 0)
	output["result"] = result_type
	output["message"] = results
	output_string := GetJSONstring(output)
	return output_string
}

func ShowAllRanks(input []byte, db *sql.DB) string {
	var name, replay_data, time string
	var score int
	result_set, query_err := db.Query("SELECT name, score, replay_data, time FROM view_ranking")
	if query_err != nil {
		Print_error("query error", query_err)
	}
	defer result_set.Close()
	results := make([]map[string]interface{}, 0)
	result_type := -1
	for result_set.Next() {
		scan_err := result_set.Scan(&name, &score, &replay_data, &time)
		if scan_err != nil {
			Print_error("scan error", scan_err)
		}
		var result = map[string]interface{} {
			"name": name,
			"score": score,
			"replay_data": replay_data,
			"time": time,
		}
		results = append(results, result)
	}
	if len(results) > 0 {
		result_type = 2
	}
	return GetOutput(result_type, results)
}

func AddRank(input []byte, db *sql.DB) string {
	type Command struct {
		Id string
		Score int
		Replay_data string
	}
	command := Command{}
	if json_err := json.Unmarshal(input, &command); json_err != nil {
		Print_error("json error @AddRank ", json_err)
	}
	results := make([]map[string]interface{}, 0)
	results = append(results, map[string]interface{} { "insert": "fail"})
	result_type := -1
	result, query_err := db.Exec("INSERT INTO ranks(score,replay_data,users_id,time) VALUES (?,?,?,now())", command.Score, command.Replay_data, command.Id)
	if query_err != nil {
		Print_error("query error ", query_err)
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		result_type = 1
		results[0]["insert"] = "success"
	}
	return GetOutput(result_type, results)
}

