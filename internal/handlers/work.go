package handlers

import (
	"fmt"
	"net/http"
	"log"
	"strconv"
)

func WorkHandler(w http.ResponseWriter, r *http.Request){
	limitStr := r.URL.Query().Get("limit")

	limit,err := strconv.Atoi(limitStr)

	if err != nil || limit < 0{
		limit = 10000 //default limit will be 1000 if limit query was not mentioned.
	}


	w.Write([]byte("Work Simulation has begun."))
	w.Write([]byte("Printing Prime numbers starting from 2\n"))

	num := 0
	count:= 0
	for count < limit{
		select {
		case <- r.Context().Done():
			log.Println("Client disconnected. %d prime numbers found",count)
			return
		
		default:
			if isPrime(num){
				w.Write([]byte(fmt.Sprintf("%d\n",num)))
				count++
				if f,ok := w.(http.Flusher); ok{
					f.Flush()
				}
			}
			num++
		}
	}

	log.Println("Work simulated successfully.")
	log.Println("%d prime numbers requested",limit)
	log.Println("%d prime numbers computed",count)
	w.Write([]byte("Work Completed successfully.\n"))
	w.Write([]byte(fmt.Sprintf("Prime Numbers computed: %d\n",count)))
}

func isPrime(num int) bool {
	if num == 2{
		return true
	}

	if num < 2{
		return false
	}

	if num % 2 == 0{
		return false
	}

	for i:= 3; i*i <= num; i+=2{
		if num % i == 0{
			return false
		}
	}
	return true
}
