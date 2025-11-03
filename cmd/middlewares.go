package main

import (
	"net/http"
)

func authenticate(next http.Handler) http.Handler {
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	sessionid, err := cookies.Read(r, "sessionid")
	// 	if err != nil {
	// 	}
	//
	// })
	return nil
}
