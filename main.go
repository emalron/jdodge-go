package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
)

func controller(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	param := mux.Vars(r)
	api := param["first"]
	var body map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&body)
	command, _ := json.Marshal(body)
	fmt.Println(string(command))
	db, err := sql.Open("mysql", baseInfo["url"])
	if err != nil {
		printError("db error: ", err)
	}
	defer db.Close()
	output := serviceMap[api](command, db)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, output)
}

func main() {
	initial()
	router := mux.NewRouter()
	router.HandleFunc("/v1/{first}", controller)
	http.Handle("/", router)
	httpsErr := http.ListenAndServeTLS(":8443", baseInfo["cert"], baseInfo["key"], nil)
	if httpsErr != nil {
		printError("https error: ", httpsErr)
	}
	httpErr := http.ListenAndServe(":8001", nil)
	if httpErr != nil {
		printError("http error: ", httpErr)
	}
}
