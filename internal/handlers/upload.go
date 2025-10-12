package handlers

import (
	"net/http"
	"log"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
	file,header,err := r.FormFile("file")

	if err != nil{
		log.Printf("Error: %v\n",err)
		http.Error(w,"Failed to get file",400)
		return
	}

	defer file.Close()

	log.Println("Recieved file")
	log.Printf("File name: %s\n",header.Filename)
	log.Printf("File Size: %d bytes\n",header.Size)

}
