package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	port := "4080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
	log.Printf("starting up on http://localhost:4080/ui/search", port)

	r := chi.NewRouter()

	type Email struct {
		Name         string
		SentMail     string
		AllDocuments string
		Contacts     string
		DeletedItems string
		Discussion   string
		Inbox        string
		NotesInbox   string
		Sents        string
		SensItems    string
		Straw        string
	}

	type CreateEmailCMD struct {
		Name         string `json:"name"`
		SentMail     string `json: "sent_mail"`
		AllDocuments string `json: "alldocuments"`
		Contacts     string `json: "contacts"`
		DeletedItems string `json: "deleted_items"`
		Discussion   string `json: "discussion"`
		Inbox        string `json: "inbox"`
		NotesInbox   string `json: "notes_inbox"`
		Sents        string `json: "sents"`
		SensItems    string `json: "sens_items"`
		Straw        string `json: "straw"`
	}

	//r.Get(http://localhost:4080/api/_bulk -i -u admin:Complexpass#123 --data-binary "@minidata.ndjson")
	//r.Mount(minidata.ndjson)

}
