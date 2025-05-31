package main

import (
	"encoding/json" // Package for encoding and decoding JSON
	"fmt"           // Package for formatted I/O
	"log"           // Package for logging messages
	"math/rand"     // Package for generating pseudo-random numbers
	"net/http"      // Package for building HTTP servers and clients
	"time"          // Package for time-related functions
)

// Global variable to store the secret number for the current game session.
// This is simplified for a single-user demo. For a multi-user application,
// proper session management (e.g., using cookies and a map for secrets) would be required.
var secretNumber int

// Global random number generator instance.
var r *rand.Rand

func init() {
	// Initialize the random number generator once when the package is loaded.
	// Use rand.NewSource and rand.New for modern Go random number generation.
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
	// Generate the initial secret number.
	generateNewSecretNumber()
}

// generateNewSecretNumber generates a new random number between 0 and 5 (inclusive)
// and assigns it to the global secretNumber variable.
func generateNewSecretNumber() {
	secretNumber = r.Intn(6)                                    // Generates a number from 0 to 5
	log.Printf("New secret number generated: %d", secretNumber) // Log the new secret for debugging
}

// GuessRequest struct defines the expected structure of the JSON request body
// when a user submits a guess.
type GuessRequest struct {
	Guess int `json:"guess"`
}

// GuessResponse struct defines the structure of the JSON response sent back to the client.
type GuessResponse struct {
	Status       string `json:"status"`                 // "correct", "too_low", "too_high", "error"
	Message      string `json:"message"`                // User-friendly message
	SecretNumber *int   `json:"secretNumber,omitempty"` // The secret number, only sent on correct guess
}

// guessHandler handles the POST requests to the /guess endpoint.
// It processes the user's guess, compares it to the secret number, and sends back a JSON response.
func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json for all responses.
	w.Header().Set("Content-Type", "application/json")

	// Ensure the request method is POST.
	if r.Method != http.MethodPost {
		http.Error(w, `{"message": "Only POST method is allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req GuessRequest
	// Decode the JSON request body into the GuessRequest struct.
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// If there's an error decoding the JSON, send an error response.
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, `{"status": "error", "message": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Compare the user's guess with the global secret number.
	if req.Guess == secretNumber {
		// If the guess is correct, generate a new secret number for the next round.
		// This makes the game playable multiple times without restarting the server.
		correctNumber := secretNumber // Store the current correct number before generating a new one
		generateNewSecretNumber()     // Generate a new secret for the next game
		err := json.NewEncoder(w).Encode(GuessResponse{
			Status:       "correct",
			Message:      fmt.Sprintf("Congratulations! You guessed the correct number: %d", correctNumber),
			SecretNumber: &correctNumber, // Send the correct number back to the client
		})
		if err != nil {
			return
		}
	} else if req.Guess < secretNumber {
		err := json.NewEncoder(w).Encode(GuessResponse{
			Status:  "too_low",
			Message: "Your guess is too low! Try again.",
		})
		if err != nil {
			return
		}
	} else { // req.Guess > secretNumber
		err2 := json.NewEncoder(w).Encode(GuessResponse{
			Status:  "too_high",
			Message: "Your guess is too high! Try again.",
		})
		if err2 != nil {
			return
		}
	}
}

func main() {
	// Create a file server to serve static files from the "http" directory.
	// This will serve index.html when the root path is requested.
	fs := http.FileServer(http.Dir("./http"))
	http.Handle("/", fs) // Handle all requests by serving files from the "http" directory

	// Handle the /guess API endpoint for game logic.
	http.HandleFunc("/guess", guessHandler)

	// Start the HTTP server on port 8080.
	log.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
