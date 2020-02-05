package main
import(
	G "jdodge-go/global"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
)

func controller(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	api := param["first"]
	var body map[string]interface{}
	_ = json.NewDecoder(r.Body).Decode(&body)
	command, _ := json.Marshal(body)


	output := serviceMap[api](command, db)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, output)
}

func headerMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w,r)
	})
}

func init() {
	G.Ignite()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/{first}", controller)
	http.Handle("/", headerMiddleWare(router))
	httpsErr := http.ListenAndServeTLS(":8443", baseInfo["cert"], baseInfo["key"], nil)
	if httpsErr != nil {
		printError("https error: ", httpsErr)
	}
	httpErr := http.ListenAndServe(":8001", nil)
	if httpErr != nil {
		printError("http error: ", httpErr)
	}
}
