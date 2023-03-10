package main

import (
	"net/http"
)

type Data struct {
	Operation string  `json:"Operation"`
	Num1      float64 `json:"Num1"`
	Num2      float64 `json:"Num2"`
}

var data Data

func main() {

	http.HandleFunc("/calculate", handleCalculator) //  endpoint
	http.ListenAndServe("localhost:8080", nil)      // listen and serve
}

func handleCalculator(w http.ResponseWriter, r *http.Request) {

}
