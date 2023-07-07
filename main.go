package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abelemlih/manta-ray/messenger"
	"github.com/google/uuid"
)

func main() {

	message := messenger.Message{Id: uuid.New(), Content: "Manta Ray is listening on port 3000", SentAt: time.Now(), DeliveredAt: time.Now()}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message.Content)

	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
