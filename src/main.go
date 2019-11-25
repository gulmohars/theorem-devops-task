
package main

import (
	"encoding/json"
	"net/http"
)

type Devops struct {
  Name    string
  Hobbies []string
}

func getDevops(w http.ResponseWriter, r *http.Request) {
	answer := Devops{"Alberto Eduardo", []string{"AWS", "Kubernetes"}}

	js, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", getDevops)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":80", nil)
}
