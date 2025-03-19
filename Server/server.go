package Server

import (
	"SudokuAPI/Generator"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}

func getSudoku(write http.ResponseWriter, request *http.Request) {
	enableCors(&write)
	hints := request.URL.Query().Get("hints")
	boxes := request.URL.Query().Get("boxes")
	if boxes == "" {
		boxes = "false"
	}

	h, err := strconv.Atoi(hints)
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}

	sudoku := Generator.GenerateSudoku(h)
	if boxes == "true" {
		sudoku = Generator.OrganizeIntoBoxes(&sudoku)
	}
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
