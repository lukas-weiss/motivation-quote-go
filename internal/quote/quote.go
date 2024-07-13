package quote

import (
	"math/rand"
)

type Quote struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func GetQuote() Quote {
	// TODO: calculate the range based on the number of quotes
	min := 1
	max := 3
	// generate a random number between min and max while min is the minimum value and not 0
	id := rand.Intn(max-min+1) + min
	return Quote{
		Id:     id,
		Author: "test",
		Quote:  "test",
	}
}
