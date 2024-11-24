package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	listenAddress string // adresa i port na kojima ce server slusati
}

func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

// ova funkcija definise alijas za funkcije koje uptavljaju http requestovima, tehnicki omogucava bolji obrazac za rad sa api korisnicima
type apiFunction func(http.ResponseWriter, *http.Request) error

// Omotava apiFunction u HTTP handler kompatibilan sa Golangom
func makeHTTPHandleFunc(f apiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			panic(err)
		}
	}
}

// Kreira novi API server sa definisanom adresom slušanja
func newAPIServer(listenAddress string) *APIserver {
	return &APIserver{
		listenAddress: listenAddress,
	}
}

// funkcija za pokretanje servera
func (s *APIserver) Run() {

	// rutiranje
	r := mux.NewRouter()

	r.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	fmt.Println("JSON API server running on port: ", s.listenAddress)
	// log.Println("JSON API server running on port: ", s.listenAddress)	// ili ovaj zapis

	http.ListenAndServe(s.listenAddress, r)
}

// funckija za upravljanje akauntima
func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("methods are not allowed %s !", r.Method)
}

// funkcija za ucitavanje postojecih akaunta
func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// funckcija koja pravi novi akaunt
func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// funckija koja brise akaunt
func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// funkcija koja prebacuje akaunt
func (s *APIserver) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
