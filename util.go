package main

import(
	"bufio"
	"os"
	"fmt"
	"strings"
	"encoding/json"
)

func getPropertiesMap(filename string) map[string]string {
	file1, err1 := os.Open(filename)
	if err1 != nil {
		printError("file open error: ", err1)
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

func printError(message string, err error) {
	fmt.Fprintf(os.Stderr, message, err, "\n")
}

func getJSONstring(mapObject map[string]interface{}) string {
	jsonBytes, err := json.Marshal(mapObject)
	if err != nil {
		printError("json error: ", err)
	}
	retString := string(jsonBytes)
	return retString
}
