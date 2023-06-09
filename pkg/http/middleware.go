package http

import (
	"log"
	"net/http"
	"runtime"
)

func panicRecover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func(r *http.Request) {
			err, _ := recover().(error)
			if err != nil {

				stack := make([]byte, 4<<10) // 4kb
				length := runtime.Stack(stack, true)
				log.Println(string(stack[:length]))
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}(r)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}
