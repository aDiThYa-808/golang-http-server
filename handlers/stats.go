package handlers 

import (
	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
	"net/http"
	"encoding/json"
)

type Stats struct{
	Total int64 `json:"total"`
	Success int64 `json:"success"`
	Client_Error int64 `json:"client-error"`
	Server_Error int64 `json:"server-error"`
	Unauthorized int64 `json:"unauthorized"`
}

func StatsHandler(w http.ResponseWriter, r *http.Request){
	stats := Stats{
		Total : middlewares.GetStats("total"),
		Success : middlewares.GetStats("success"),
		Client_Error : middlewares.GetStats("client-error"),
		Server_Error : middlewares.GetStats("server-error"),
		Unauthorized : middlewares.GetStats("unauthorized"),
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(stats)
}
