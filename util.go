package main

import(
	"bufio"
	"os"
	"fmt"
	"strings"
	"encoding/json"
)

func GetPropertiesMap(filename string) map[string]string {
	file1, err1 := os.Open(filename)
	if err1 != nil {
		Print_error("file open error: ", err1)
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

func Print_error(message string, err error) {
	fmt.Fprintf(os.Stderr, message, err, "\n")
}

func GetJSONstring(mapObject map[string]interface{}) string {
	jsonBytes, err := json.Marshal(mapObject)
	if err != nil {
		Print_error("json error: ", err)
	}
	retString := string(jsonBytes)
	return retString
}
