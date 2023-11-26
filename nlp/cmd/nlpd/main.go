package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chrisbradleydev/practical-go-foundations/nlp"
)

func main() {
	// routing
	// /healthz is an exact match
	// /healthz/ is a prefix match
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)

	// run server
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Disallow unsupported methods
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	// Step 2: Get, convert, and validate the data
	defer r.Body.Close()
	lr := io.LimitReader(r.Body, 1_000_000) // limit request body size
	data, err := io.ReadAll(lr)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		http.Error(w, "no data", http.StatusBadRequest)
		return
	}
	text := string(data)

	// Step 3: Do work
	tokens := nlp.Tokenize(text)

	// Step 4: Encode and emit output
	resp := map[string]any{
		"tokens": tokens,
	}
	// You can also do:
	// err = json.NewEncoder(w).Encode(resp)
	data, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, "can't encode", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// run health check
	fmt.Fprintln(w, "ok")
}