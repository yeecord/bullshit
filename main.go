package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/StillFantastic/bullshit/generator"
	"github.com/rs/cors"
)

type Data struct {
	Topic  string
	MinLen int
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ip, _, err := net.SplitHostPort(r.RemoteAddr)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		// 	return
		// }

		// limiter := getVisitor(ip)
		// if limiter.Allow() == false {
		// 	http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
func bullshitHandler(w http.ResponseWriter, r *http.Request) {
	var d Data
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// logRequest(d.Topic, d.MinLen)
	ret := generator.Generate(d.Topic, d.MinLen)
	w.Write([]byte(ret))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/bullshit", bullshitHandler)
	handler := cors.Default().Handler(mux)
	var addr string

	port := os.Getenv("PORT")

	if len(port) == 0 {
		addr = "0.0.0.0:80"
	} else {
		addr = "0.0.0.0:" + port
	}
	err := http.ListenAndServe(addr, limit(handler))
	if err != nil {
		fmt.Println(err)
	}
}
