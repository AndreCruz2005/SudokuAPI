package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func factorial(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func factorialRoute(write http.ResponseWriter, request *http.Request) {
	number := request.URL.Query().Get("number")

	// Convert the number to an integer
	n, err := strconv.Atoi(number)
	if err != nil {
		// Handle the error if conversion fails
		http.Error(write, "Invalid number", http.StatusBadRequest)
		return
	}

	// Calculate the factorial
	result := factorial(n)

	// Create a struct to hold the response
	response := map[string]interface{}{
		"number":    n,
		"factorial": result,
	}

	// Set the header to indicate a JSON response
	write.Header().Set("Content-Type", "application/json")

	// Encode and send the response as JSON
	if err := json.NewEncoder(write).Encode(response); err != nil {
		http.Error(write, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/factorial", factorialRoute)
	fmt.Print(factorial(5))
	log.Fatal(http.ListenAndServe(":6000", nil))
}
