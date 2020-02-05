package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// GetPropertiesMap returns map[string]string from a binary file
func GetPropertiesMap(filename string) map[string]string {
	file1, err1 := os.Open(filename)
	if err1 != nil {
		PrintError("file open error: ", err1)
	}
	defer file1.Close()
	scanner := bufio.NewScanner(file1)
	output := make(map[string]string)
	for scanner.Scan() {
		tmp1 := scanner.Text()
		tmp2 := strings.Split(tmp1, " ")
		output[tmp2[0]] = tmp2[1]
	}
	return output
}

// PrintError prints message and error to os.Stderr
func PrintError(message string, err error) {
	fmt.Fprintf(os.Stderr, message, err, "\n")
}

// GetJSONstring returns stringified JSON from a map[string]interface{}
func GetJSONstring(mapObject map[string]interface{}) string {
	jsonBytes, err := json.Marshal(mapObject)
	if err != nil {
		PrintError("json error: ", err)
	}
	retString := string(jsonBytes)
	return retString
}

// GetOutputJSON returns stringified JSON as the final response of the request
func GetOutputJSON(resultType int, results []map[string]interface{}) string {
	var output = make(map[string]interface{}, 0)
	output["result"] = resultType
	output["message"] = results
	outputString := GetJSONstring(output)
	return outputString
}