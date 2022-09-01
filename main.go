package main

import (
	"encoding/json"
	"fmt"
	"github.com/StillFantastic/bullshit/generator"
	"github.com/rs/cors"
	"net/http"
	"os"
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
	if len(os.Args) < 2 {
		addr = "0.0.0.0:10000"
	} else {
		addr = "0.0.0.0:" + os.Args[1]
	}
	err := http.ListenAndServe(addr, limit(handler))
	if err != nil {
		fmt.Println(err)
	}
}
