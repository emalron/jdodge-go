package main
import(
    "jdodge-go/util"
    "jdodge-go/usecases"
    "jdodge-go/interfaces"
    "jdodge-go/infrastructure"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

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

var baseInfo map[string]string
func Ignite() {
	baseInfo = util.GetPropertiesMap("properties")
}

func main() {
    Ignite()
    dbHandler := infrastructure.NewMysqlHandler(baseInfo["url"])
    handlers := make(map[string]interfaces.DBHandler, 0)
    handlers["DBRankRepo"] = dbHandler
    handlers["DBUserRepo"] = dbHandler
    rankInteractor := new(usecases.RankInteractor)
    rankInteractor.RankRepository = interfaces.NewDBRankRepo(handlers)
    rankInteractor.UserRepository = interfaces.NewDBUserRepo(handlers)
    webService := interfaces.WebserviceHandler{}
    webService.RankInteractor = rankInteractor

    router := mux.NewRouter()
	router.HandleFunc("/v1/showAllRanks", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("you hit the router")
        webService.ShowAll(w, r)
    })
	http.Handle("/", headerMiddleWare(router))
	httpsErr := http.ListenAndServeTLS(":8443", baseInfo["cert"], baseInfo["key"], nil)
    if httpsErr != nil {
        util.PrintError("https error: ", httpsErr)
    }
	httpErr := http.ListenAndServe(":8001", nil)
	if httpErr != nil {
		util.PrintError("http error: ", httpErr)
	}
}
