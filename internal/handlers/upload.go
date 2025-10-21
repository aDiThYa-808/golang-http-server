package handlers

import (
	"net/http"
	"log"
	"os"
	"io"
)

const uploadDir = "./uploads"

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


	tempFile, tempErr := os.CreateTemp(uploadDir,"upload-*.tmp")

	if tempErr != nil {
		http.Error(w,"failed to create temp file",500)
		return
	}

	defer tempFile.Close()

	_,copyErr:= io.Copy(tempFile,file)

	if copyErr != nil{
		http.Error(w,"failed to copy the file into a temp file",500)
		return
	}

	log.Printf("%v\n",tempFile.Name())
}
