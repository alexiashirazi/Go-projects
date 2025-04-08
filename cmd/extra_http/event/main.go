package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Event struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	e.Date, err = time.Parse("02.01.2006", aux.Date)
	if err != nil {
		return fmt.Errorf("date parse error: %v", err)
	}

	return nil
}

func (e Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		Date string `json:"date"`
		*Alias
	}{
		Date:  e.Date.Format("02.01.2006"),
		Alias: (*Alias)(&e),
	})
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s \n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "12345" {
			http.Error(w, "Wrong API KEY", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is on\n"))
	})

	router.Route("/events", func(r chi.Router) {
		r.Use(logMiddleware)
		r.Use(apiKeyMiddleware)
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			w.Write([]byte(fmt.Sprintf("event with id : %s\n", id)))
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var event Event
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "failed to read body", http.StatusInternalServerError)
				return
			}

			err = json.Unmarshal(body, &event)
			if err != nil {
				http.Error(w, "invalid json", http.StatusBadRequest)
				return
			}
			marshaled, _ := json.Marshal(event)
			response := fmt.Sprintf("got body %s", marshaled)
			w.Write([]byte(response))

			w.Write([]byte("New event received\n"))
		})
	})

	fmt.Println("Server pornit pe http://localhost:3000")
	http.ListenAndServe(":3000", router)

	return
}
