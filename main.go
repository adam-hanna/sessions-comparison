package main

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"

	redisStore "gopkg.in/boj/redistore.v1"
)

var (
	sesh    *redisStore.RediStore
	seshErr error
)

var issueSession = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Get a session.
	session, err := sesh.Get(r, "session-key")
	if err != nil {
		log.Printf("Err getting user session: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if session == nil {
		session = &sessions.Session{}
	}

	// Set some session values.
	session.Values["UserID"] = "fakeUserID"
	session.Values["JSON"] = ""

	// Save.
	if err = sessions.Save(r, w); err != nil {
		log.Printf("Err saving user session: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(w.Header())

	w.WriteHeader(http.StatusOK)
})

var requiresSession = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Get a session.
	session, err := sesh.Get(r, "session-key")
	if err != nil {
		log.Printf("Err getting user session: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session == nil || session.Values["UserID"] == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
})

func main() {
	// Fetch new store.
	sesh, seshErr = redisStore.NewRediStore(0, "tcp", ":6379", "", []byte("DOZDgBdMhGLImnk0BGYgOUI+h1n7U+OdxcZPctMbeFCsuAom2aFU4JPV4Qj11hbcb5yaM4WDuNP/3B7b+BnFhw=="))
	if seshErr != nil {
		panic(seshErr)
	}
	defer sesh.Close()

	http.HandleFunc("/issue", issueSession)
	http.HandleFunc("/require", requiresSession)

	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", context.ClearHandler(http.DefaultServeMux)))
}
