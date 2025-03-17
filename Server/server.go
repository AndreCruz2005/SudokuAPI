package Server

import (
	"SudokuAPI/Generator"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func getSudoku(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	hints := request.URL.Query().Get("hints")

	h, err := strconv.Atoi(hints)
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}

	sudoku := Generator.GenerateSudoku(h)
	response, err := json.Marshal(sudoku)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
	}
	write.Write(response)
}

func Serve() {
	http.HandleFunc("/sudoku", getSudoku)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
