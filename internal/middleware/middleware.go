package middleware

import (
	"context"
	"log"
	"net/http"
)

type PugContext struct {
	Proxy     string
	Useragent string
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a response writer:
		ctx := context.WithValue(r.Context(), "pugCtx", PugContext{
			Proxy:     "lol",
			Useragent: "kek",
		})
		// Here we are passing our custom response writer to the next http handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (p *PugContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func Context(handler func(ctx PugContext, w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a response writer:
		v := r.Context().Value("pugCtx")
		log.Println(v)
		// Here we are passing our custom response writer to the next http handler.
		handler(v.(PugContext), w, r)
		return
	})
}
