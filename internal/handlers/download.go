package handlers

import (
	"net/http"
	"path/filepath"
	"os"
)

const uploadDirectory = "./uploads"

func DownloadHandler(w http.ResponseWriter , r *http.Request){
	fileName := r.URL.Query().Get("file")
	if fileName == ""{
		http.Error(w,"Missing file parameter",http.StatusBadRequest)
	}

	filePath := filepath.Join(uploadDirectory,filepath.Base(fileName))

	if _,err := os.Stat(filePath); err != nil{
		if os.IsNotExist(err){
			http.Error(w,"File doesn't exist",http.StatusNotFound)
			return
		}
		http.Error(w,"Internal server error",http.StatusInternalServerError)
		return
	}

	http.ServeFile(w,r,filePath)
}
