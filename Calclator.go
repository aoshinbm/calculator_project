package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Operation string  `json:"Operation"`
	Num1      float64 `json:"Num1"`
	Num2      float64 `json:"Num2"`
}

var data Data

func errorHandling(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func main() {

	http.HandleFunc("/calculate", handleCalculator) //  endpoint
	http.ListenAndServe("localhost:8080", nil)      // listen and serve
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil { // Decoding the request data
			errorHandling(w, err)
			http.Error(w, "Error decoding data", http.StatusBadRequest)
			return
		}

		switch data.Operation {
		case "sum":
			result := calculateSum(data)
		case "subtract":
			result := calculateSubtract(data)
		case "multiplication":
			result := calculateMultiplication(data)
		case "division":
			result := calculateDivision(data)
		default:
			http.Error(w, "Invalid operation", http.StatusBadRequest)
			return
		}
	}
}
