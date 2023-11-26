package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/chrisbradleydev/practical-go-foundations/nlp"
	"github.com/chrisbradleydev/practical-go-foundations/nlp/stemmer"
	"github.com/gorilla/mux"
)

var (
	numTok = expvar.NewInt("tokenize.call")
)

func main() {
	logger := log.New(log.Writer(), "[nlpd] ", log.LstdFlags|log.Lshortfile)
	// Create server
	s := Server{
		logger: logger, // dependency injection
	}
	/* Before gorilla/mux
	routing
	/healthz is an exact match
	/healthz/ is a prefix match
	*/
	r := mux.NewRouter()
	r.HandleFunc("/healthz", s.healthzHandler).Methods(http.MethodGet)
	r.HandleFunc("/stem/{word}", s.stemHandler).Methods(http.MethodGet)
	r.HandleFunc("/tokenize", s.tokenizeHandler).Methods(http.MethodPost)
	http.Handle("/", r)

	// run server
	addr := os.Getenv("NLPD_ADDR")
	if addr == "" {
		addr = ":8000"
	}
	s.logger.Printf("server running on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error: %s", err)
	}
}

type Server struct {
	logger *log.Logger
}

func (s *Server) stemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]
	stem := stemmer.Stem(word)
	fmt.Fprintln(w, stem)
}

func (s *Server) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	/* Before gorilla/mux
	// Disallow unsupported methods
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
	*/
	numTok.Add(1)

	// Step 1: Get, convert, and validate the data
	lr := io.LimitReader(r.Body, 1_000_000) // limit request body size
	data, err := io.ReadAll(lr)
	if err != nil {
		s.logger.Printf("error: can't read - %s", err)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}

	if len(data) == 0 {
		s.logger.Printf("error: no data - %s", err)
		http.Error(w, "no data", http.StatusBadRequest)
		return
	}
	text := string(data)

	// Step 2: Do work
	tokens := nlp.Tokenize(text)

	// Step 3: Encode and emit output
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

func (s *Server) healthzHandler(w http.ResponseWriter, r *http.Request) {
	// run health check
	fmt.Fprintln(w, "ok")
}