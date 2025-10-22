package main

import   (
	"log"
	"net/http"
	"os"

	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
	handlers "github.com/aDiThYa-808/golang-http-server/internal/handlers"
)

const (
	uploadDir = "./uploads"
	uploadsMaxBytes = 50 << 20 // this mean 50 * 1,048,576 = 52,428,800 which is 50mb
)

func init(){
	if err := os.MkdirAll(uploadDir,0755); err != nil{
		log.Fatal("Could'nt create uploads directory: \n %v",err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.HomeHandler)))
	mux.Handle("/stats",middlewares.AuthMiddleware(http.HandlerFunc(handlers.StatsHandler)))
	mux.Handle("/work",http.HandlerFunc(handlers.WorkHandler))
	mux.Handle("/upload", middlewares.MaxBodySize(uploadsMaxBytes)(http.HandlerFunc(handlers.UploadHandler)))
	mux.Handle("/download",http.HandlerFunc(handlers.DownloadHandler))

	handler := middlewares.StatsRecorderMiddleware(mux)
	handler = middlewares.LoggingMiddleware(handler)
	handler = middlewares.RequestIdMiddleware(handler)

	log.Print("Server running on port 4000.")
	log.Fatal(http.ListenAndServe(":4000", handler))
}


