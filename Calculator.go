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

type ResponseData struct {
	Operation string  `json:"operation"`
	Result    float64 `json:"result"`
}

func errorHandling(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func main() {

	http.HandleFunc("/calculate", handleCalculator) //  endpoint
	http.ListenAndServe("localhost:8080", nil)      // listen and serve
}

func handleCalculator(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data []Data
		var responseData []ResponseData

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil { // Decoding the request data
			errorHandling(w, err)
			http.Error(w, "Error decoding data", http.StatusBadRequest)
			return
		}

		for _, d := range data {
			var rd ResponseData
			rd.Operation = d.Operation
			switch d.Operation {
			case "sum":
				rd.Result = calculateAdd(d)
			case "subtract":
				rd.Result = calculateSubtract(d)
			case "multiplication":
				rd.Result = calculateMultiplication(d)
			case "division":
				rd.Result = calculateDivision(d)
			default:
				http.Error(w, "Invalid operation", http.StatusBadRequest)
				return
			}
			responseData = append(responseData, rd)
		}
		json.NewEncoder(w).Encode(responseData)
	}
}

func calculateAdd(data Data) float64 {
	return data.Num1 + data.Num2
}

func calculateSubtract(data Data) float64 {
	return data.Num1 - data.Num2
}

func calculateMultiplication(data Data) float64 {
	return data.Num1 * data.Num2
}

func calculateDivision(data Data) float64 {
	return data.Num1 / data.Num2
}
