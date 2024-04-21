package main

import (
	"encoding/json"
	"fmt"

	"github.com/lyy0709/OpenAIAuth/auth"
)

func main() {
	auth := auth.NewAuthenticator("", "", "")
	err := auth.Begin()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		return
	}
	// if os.Getenv("PROXY") != "" {
	_, err = auth.GetPUID()
	if err != nil {
		println("Error: " + err.Details)
		println("Location: " + err.Location)
		println("Status code: " + fmt.Sprint(err.StatusCode))
		return
	}
	// }
	// JSON encode auth.GetAuthResult()
	result := auth.GetAuthResult()
	result_json, _ := json.Marshal(result)
	println(string(result_json))
}