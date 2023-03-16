package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InputData struct {
	Operation string `json:"operation"`
	Num       []float64
}

type OutputData struct {
	Operation string    `json:"operation"`
	Values    []float64 `json:"values"`
	Result    float64   `json:"result"`
}

func errorHandling(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func main() {

	http.HandleFunc("/calculate/dynamicArray", handleCalculate) //  endpoint
	fmt.Println("Starting Server...")
	http.ListenAndServe("localhost:8080", nil) // listen and serve
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var input []InputData
		var outData []OutputData

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil { // Decoding the request data
			errorHandling(w, err)
			http.Error(w, "Error decoding data", http.StatusBadRequest)
			return
		}

		for _, dt := range input {
			var op OutputData
			//var rd ResponseData
			op.Operation = dt.Operation
			switch dt.Operation {
			case "sum":
				//w.WriteHeader()
				op.Result = calculateSumm(dt)
			case "subtract":
				op.Result = calculateSubtractt(dt)
			case "multiplication":
				op.Result = calculateMultiply(dt)
			case "division":
				op.Result = calculateDivide(dt)
			default:
				http.Error(w, "Invalid operation", http.StatusBadRequest)
				return
			}
			op.Values = dt.Num
			outData = append(outData, op)
		}
		json.NewEncoder(w).Encode(outData)
	}
}

func calculateSumm(data InputData) float64 {
	total := 0.0
	for i := 0; i < len(data.Num); i++ {
		total = total + data.Num[i]
	}
	return total
}

func calculateSubtractt(data InputData) float64 {
	total := data.Num[0]
	for i := 1; i < len(data.Num); i++ {
		total = total - data.Num[i]
	}
	return total
}

func calculateMultiply(data InputData) float64 {
	total := data.Num[0]
	for i := 1; i < len(data.Num); i++ {
		total = total * data.Num[i]
		if total*data.Num[i] == 0 {
			fmt.Println("Encountered a ZERO")
		}
	}
	return total
}

func calculateDivide(data InputData) float64 {
	total := data.Num[0]
	for i := 1; i < len(data.Num); i++ {
		total = total / data.Num[i]
		if total/data.Num[i] == 0 {
			fmt.Println("Encountered a ZERO")
		}
	}
	return total
}
