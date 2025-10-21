package handlers

import (
	"net/http"
	"path/filepath"
	"regexp"
	"log"
	"os"
	"io"
	"encoding/json"
	"fmt"
	"strings"
	"github.com/google/uuid"
)

type response struct{
	Success bool `json:"success"`
	FileName string `json:"file_name"`
	FileSize int64 `json:"file_size"`
}

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

	log.Println("Temp file created in uploads directory")
	log.Printf("Name of the temp file: %v\n",tempFile.Name())

	finalName := filepath.Join(uploadDir,generateSafeFileName(header.Filename))

	renameErr := os.Rename(tempFile.Name(),finalName)

	if renameErr != nil{
		http.Error(w,"Failed to rename temp file",500)
		return
	}
  
	res := response{
		Success : true,
		FileName : finalName,
		FileSize : header.Size, 
	}

	fmt.Println("Upload complete.")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(res)
}

func generateSafeFileName(name string) string{

	baseName := filepath.Base(name)
	extension := filepath.Ext(baseName)

	baseName = strings.ToLower(baseName)

	baseName = strings.TrimSuffix(baseName,extension)

	baseName = regexp.MustCompile(`[^a-z0-9._-]+`).ReplaceAllString(baseName, "_")

	if len(baseName) > 50{
		baseName = baseName[:50]
	}

	allowedExtensions := map[string]bool{
		".jpg": true,
		".jpeg": true,
		".png": true,
		".webp": true,
		".pdf": true,
		".mp3": true,
		".mp4":true,
		".mov":true,
	}

	if _,found := allowedExtensions[extension]; !found{
		extension = ".bin"
	}

	safeName := fmt.Sprintf("%s_%s%s",uuid.New().String(),baseName,extension)

	return safeName
}
