package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// URL struct for storing in-memory database
type URL struct {
	ID           string    `json:"id"`
	LongURL      string    `json:"long_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

// generateShortURL generates a short URL based on MD5 hash of the long URL
func generateShortURL(LongURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(LongURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:6] // Use the first 6 characters of the hash as short URL
}

// createURL generates a short URL, stores it in urlDB, and returns it
func createURL(longURL string) string {
	shortURL := generateShortURL(longURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		LongURL:      longURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	fmt.Println("urlDB after storing:", urlDB)
	return shortURL
}

// getURL retrieves a URL object from urlDB based on the short code
func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

// handler is a simple handler function
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Service is up and running!")
}

// ShortURLHandler handles the /shorturl endpoint to shorten a long URL
func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	// Decode JSON request body into data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Create short URL based on the provided long URL
	shortURL := createURL(data.URL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL}

	// Encode response as JSON and return it
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// redirectURLHandler redirects users to the original long URL based on the short code
func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Extract short code from URL path
	shortCode := r.URL.Path[len("/redirect/"):]
	fmt.Println("Redirecting with shortCode:", shortCode)

	// Retrieve URL from database
	url, err := getURL(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		fmt.Println("Error:", err)
		return
	}

	// Log the redirection
	fmt.Printf("Redirecting from short URL '%s' to long URL '%s'\n", shortCode, url.LongURL)

	// Perform the redirection to the original long URL
	http.Redirect(w, r, url.LongURL, http.StatusFound)
}

func main() {
	// Register handlers for different endpoints
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorturl", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler) // Ensure it ends with a slash for correct path handling

	// Start the HTTP server
	fmt.Println("Server running on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error while starting server:", err)
	}
}
