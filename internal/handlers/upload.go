package handlers

import (
	"net/http"
	"log"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
	log.Println("upload handler")
}
