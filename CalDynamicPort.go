package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Host struct {
	Network struct {
		Port int `json:"port"`
	} `json:"network"`
}
type InputData struct {
	Operation string    `json:"operation"`
	Num       []float64 `json:"num"`
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

	portno := config()
	fmt.Println("port number", portno)
	http.HandleFunc("/calculate/dynamicArray", handleCalculate) //  endpoint
	fmt.Println("Starting Server...")
	//http.ListenAndServe("localhost:8080", nil) // listen and serve
	http.ListenAndServe(portno, nil) // listen and serve
}

func config() string {

	// Open  jsonFile
	jsonFile, err := os.Open("setting.json")
	//returns an error so handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of jsonFile so that parse it later
	defer jsonFile.Close()

	// read jsonFile as a byte array
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(byteValue))

	portt := Host{}

	//unmarshal byteArray which contains jsonFile's content into 'portt'
	err = json.Unmarshal(byteValue, &portt)
	if err != nil {
		panic(err)
	}
	//fmt.Println(portt)
	//fmt.Println(portt.Network.port_number)
	port := fmt.Sprintf(":%v", portt.Network.Port)
	return port
}

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Function")
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
	}
	return total
}

func calculateDivide(data InputData) float64 {
	total := data.Num[0]
	for i := 1; i < len(data.Num); i++ {
		total = total / data.Num[i]
	}
	return total
}
