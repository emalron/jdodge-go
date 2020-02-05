package service

import(
	"database/sql"
	// for mysql driver
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"net/http"
	"jdodge-go/util"
	"jdodge-go/model"
)

// ShowAllRanksService returns nothing
func ShowAllRanksService(w http.ResponseWriter, r *http.Request) {

	results := make([]map[string]interface{}, 0)
	resultType := -1

	results := model.ShowAllRanks()
	
	if len(results) > 0 {
		resultType = 2
	}
	output := util.GetOutputJSON(resultType, results)
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

