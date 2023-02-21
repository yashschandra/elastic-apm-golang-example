package main

import (
	"fmt"
	"net/http"
)

const (
	N = 100
)

func main() {
	for i:=0; i<N; i++ {
		sendRequest()
	}
	fmt.Println("client done")
}

func sendRequest() {
	_, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("received error ", err.Error())
	}
}
