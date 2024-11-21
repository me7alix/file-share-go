package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

const (
	uploadPath = "./uploads"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
        return
    }

    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    filename := header.Filename
    filepath := filepath.Join(uploadPath, filename)

    out, err := os.Create(filepath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer out.Close()

    _, err = io.Copy(out, file)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "File %s uploaded successfully", filename)
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)
    fs := http.FileServer(http.Dir("./"))
    http.Handle("/", fs)
    http.ListenAndServe(":5005", nil)
}

func main() {
    fmt.Println("The server is running on port 5005...")
    setupRoutes()
}
