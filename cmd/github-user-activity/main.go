package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Antonious-Stewart/github-user-activity/api"
	printer "github.com/Antonious-Stewart/github-user-activity/internal/print"
	"github.com/Antonious-Stewart/github-user-activity/internal/types"
)

func main() {
	if len(os.Args) < 2 {
		log.Panic("missing username flag")
		os.Exit(1)
	}
	username := os.Args[1]

	responseData := api.Response(username)
	var result []types.ApiResponse

	if err := json.Unmarshal(responseData, &result); err != nil {
		log.Fatal(err)
	}

	printer.PrintMessage(&result)
}
