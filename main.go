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
	db, err := sql.Open("mysql", base_info["url"])
	if err != nil {
		Print_error("db error: ", err)
	}
	defer db.Close()
	output := service_map[api](command, db)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, output)
}

func main() {
	initial()
	router := mux.NewRouter()
	router.HandleFunc("/v1/{first}", controller)
	http.Handle("/", router)
	https_err := http.ListenAndServeTLS(":8443", base_info["cert"], base_info["key"], nil)
	if https_err != nil {
		Print_error("https error: ", https_err)
	}
	http_err := http.ListenAndServe(":8001", nil)
	if http_err != nil {
		Print_error("http error: ", http_err)
	}
}
