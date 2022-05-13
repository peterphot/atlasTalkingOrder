package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type teamOrder struct {
	DateGenerated string   `json:"date_order_generated"`
	TalkingOrder  []string `json:"talking_order"`
}

type httpError struct {
	Name   string
	Errors []error
}

func GenerateTalkingOrder(w http.ResponseWriter, r *http.Request) {
	team := []string{"Michael Wilczynska", "Peter Photinos", "Shraddha Gupta", "Aaron Tiang", "Kiril Boyadzhiev", "Nikhil Agrawal", "Michael Ebner", "Liam Pearson", "Edoardo Tonelli", "Mayank Singh", "Praveen Rajan"}
	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	rand.Seed(rounded.Unix())

	rand.Shuffle(len(team), func(i, j int) {
		team[i], team[j] = team[j], team[i]
	})

	order := teamOrder{time.Now().String(), team}
	resp, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		marshalledError, _ := json.Marshal(httpError{"can not randomise", []error{err}})
		_, _ = w.Write(marshalledError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}
