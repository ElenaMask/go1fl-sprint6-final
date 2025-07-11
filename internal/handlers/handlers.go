package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ElenaMask/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File index.html not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error reading index.html", http.StatusInternalServerError)
		return
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
		return
	}
	input := strings.TrimSpace(string(data))
	result, err := service.ConvertMorseOrText(input)
	if err != nil {
		http.Error(w, "Error converting data: "+err.Error(), http.StatusInternalServerError)
		return
	}
	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("output_%s%s", time.Now().UTC().Format("20060102T150405Z"), ext)
	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Error creating output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Error writing to output file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, result)
}
