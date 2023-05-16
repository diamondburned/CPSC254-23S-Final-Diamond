package main

import (
	"log"
	"net/http"
	"q11_answer/backend"
	"q11_answer/bank"
)

func main() {
	mockBank := &bankService{
		users: map[string]bank.User{
			"alice": {
				Username: "alice",
				Password: "password",
				Activity: []bank.UserActivity{
					{Type: "deposit", Amount: 100},
					{Type: "withdraw", Amount: 50},
				},
				Balance: 50,
			},
		},
	}

	handler := backend.New(mockBank)

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
