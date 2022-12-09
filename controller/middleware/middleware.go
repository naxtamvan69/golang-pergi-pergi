package middleware

import (
	"encoding/json"
	"net/http"
	"pergipergi/model"
	"time"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.FailedResponse{
				TimeStamp: time.Now(),
				Status:    http.StatusMethodNotAllowed,
				Error:     http.StatusText(http.StatusMethodNotAllowed),
				Message:   "method is not allowed!",
				Path:      r.RequestURI,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.FailedResponse{
				TimeStamp: time.Now(),
				Status:    http.StatusMethodNotAllowed,
				Error:     http.StatusText(http.StatusMethodNotAllowed),
				Message:   "method is not allowed!",
				Path:      r.RequestURI,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Patch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.FailedResponse{
				TimeStamp: time.Now(),
				Status:    http.StatusMethodNotAllowed,
				Error:     http.StatusText(http.StatusMethodNotAllowed),
				Message:   "method is not allowed!",
				Path:      r.RequestURI,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Put(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.FailedResponse{
				TimeStamp: time.Now(),
				Status:    http.StatusMethodNotAllowed,
				Error:     http.StatusText(http.StatusMethodNotAllowed),
				Message:   "method is not allowed!",
				Path:      r.RequestURI,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(model.FailedResponse{
				TimeStamp: time.Now(),
				Status:    http.StatusMethodNotAllowed,
				Error:     http.StatusText(http.StatusMethodNotAllowed),
				Message:   "method is not allowed!",
				Path:      r.RequestURI,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
