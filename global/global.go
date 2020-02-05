package global
import (
	"net/http"
	"jdodge-go/service"
	"jdodge-go/util"
)
// BaseInfo has security information
var BaseInfo map[string]string
// ServiceMap has 
var ServiceMap map[string]func(http.ResponseWriter, *http.Request)string

// Ignite sets basic information
func Ignite() {
	BaseInfo = util.GetPropertiesMap("properties")
	ServiceMap = map[string]func(http.ResponseWriter, *http.Request)string {
		"showAllRanks": service.ShowAllRanks,
	}
}

