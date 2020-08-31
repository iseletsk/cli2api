package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func StartHTTP(config Config) {
	r := mux.NewRouter()
	for _, cli := range config.CLI {
		r.HandleFunc("/cli/"+cli.ID, getCLIHandler(cli))
	}
	srv := &http.Server{
		Handler: r,
		Addr:    config.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServeTLS(config.Server.CertFile, config.Server.KeyFile))
}

type CLIParams struct {
	Args []string `json:"args"`
}

type cliHandler func(w http.ResponseWriter, r *http.Request)

func getCLIHandler(c CLIConfig) cliHandler {
	cli := c
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != cli.APIKey {
			http.Error(w, "X-API-Key header doesn't match API Key for the command", http.StatusForbidden)
			return
		}
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Invalid request content type (expecting application/json): "+contentType,
				http.StatusInternalServerError)
			return
		}

		var params CLIParams
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cmd := exec.Command(cli.Command, params.Args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", cli.OutputContentType)
		_, err = w.Write(out)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
