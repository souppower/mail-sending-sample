package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/mail", mailHandler)

	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func mailHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	msg := &mail.Message{
		Sender:  "soup <r.kubota92@gmail.com>",
		To:      []string{"r.kubota92+1@gmail.com"},
		Subject: "Test email",
		Body:    "hey",
	}

	if err := mail.Send(ctx, msg); err != nil {
		log.Fatalf("Couldn't send email: %v", err)
	}
}
