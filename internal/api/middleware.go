package api

import (
	"fmt"
	"net"
	"net/http"
)

type middlewareFunc func(http.Handler) http.Handler

// go-swagger responders panic on error while writing response to client,
// this shouldn't result in crash - unlike a real, reasonable panic.
//
// Usually it should be second middleware (after logger).
func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			const code = http.StatusInternalServerError
			var stack interface{}
			stack = recover()
			switch err := stack.(type) {
			default:
				fmt.Println(err, "got error")
				w.WriteHeader(code)
			case nil:
			case net.Error:
				fmt.Println(err, "got error")
				w.WriteHeader(code)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
