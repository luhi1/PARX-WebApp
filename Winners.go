package main

import (
	"net/http"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type Winners struct {
	RandomNinthWinner    string
	RandomTenthWinner    string
	RandomEleventhWinner string
	RandomTwelvthWinner  string
	NinthWinners         []string
	TenthWinners         []string
	EleventhWinners      []string
	TwelvthWinners       []string
}

func (w *Winners) GETHandler(writer http.ResponseWriter, request *http.Request) {
	winners := Winners{
		RandomNinthWinner:    "a",
		RandomTenthWinner:    "b",
		RandomEleventhWinner: "c",
		RandomTwelvthWinner:  "d",
		NinthWinners:         []string{"asdf", "adafsd"},
		TenthWinners:         []string{"asdf"},
		EleventhWinners:      []string{"asdf"},
		TwelvthWinners:       []string{"asdf"},
	}
	tplExec(writer, "winners.gohtml", winners)
}

func (w *Winners) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement.
}

func (w *Winners) valHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement Data Validation.
}

func (w *Winners) dataVal(requestMethod string) bool {
	//@todo: Implement Data Validation.
	return false
}
