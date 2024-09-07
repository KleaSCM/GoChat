package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile handles file uploads
func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to upload file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the destination file in the "uploads" directory
	dst, err := os.Create(filepath.Join("uploads", header.Filename))
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// Respond with the file URL
	fileURL := fmt.Sprintf("/uploads/%s", header.Filename)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "File uploaded successfully: %s", fileURL)
}
