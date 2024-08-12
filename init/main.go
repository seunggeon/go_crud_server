package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("에러가 떴네용!")
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}
