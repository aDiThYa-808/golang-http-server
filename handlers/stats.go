package handlers 

import (
	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
	"net/http"
	"encoding/json"
)

type Stats struct{
	Total int `json:"total"`
	Success int `json:"success"`
	ClientErr int `json:"client-error"`
	ServerErr int `json:"server-error"`
	Unauthorized int `json:"unauthorized"`
}

func StatsHandler(w http.ResponseWriter, r *http.Request){
	stats := map[string]int64{
		"total":middlewares.GetStats("total"),
		"success":middlewares.GetStats("success"),
		"client_error":middlewares.GetStats("client-error"),
		"server-error":middlewares.GetStats("server-error"),
		"unauthorized":middlewares.GetStats("unauthorized"),
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(stats)
}
