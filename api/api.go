package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Response(username string) []byte {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v/events", username))

	if err != nil {
		log.Fatal(err)
	}

	switch statusCode := resp.StatusCode; statusCode {
	case http.StatusBadRequest:
		log.Fatal("provided username is not valid")
	case http.StatusNotFound:
		log.Fatal("username was not found")
	case http.StatusBadGateway:
		log.Fatal("bad gateway")
	case http.StatusGatewayTimeout:
		log.Fatal("gateway timeout")
	case http.StatusInternalServerError:
		log.Fatal("Internal server error")
	}

	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
