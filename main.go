package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Response is a struct that holds the response data.
type Response struct {
    Message string `json:"message"`
}

// helloHandler responds to HTTP GET requests with a JSON response.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    response := Response{Message: "Hello, World!"}

    // Set the Content-Type header to application/json
    w.Header().Set("Content-Type", "application/json")

    // Marshal the response struct to JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Write the JSON response
    w.Write(jsonResponse)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    fmt.Println("Server is running on port 8000...")
    if err := http.ListenAndServe(":8000", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
