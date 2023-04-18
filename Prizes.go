package main

import (
	"net/http"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type Prize struct {
	PrizeName    string
	PrizeWinners []StudentAttendance
}

func (p *Prize) GETHandler(writer http.ResponseWriter, request *http.Request) {
	var prizes []Prize
	prizes = append(prizes, Prize{
		PrizeName: "asdfsd",
		PrizeWinners: []StudentAttendance{
			{
				StudentName:   "asdf",
				StudentNumber: 12,
			},
		},
	})

	tplExec(writer, "prizes.gohtml", prizes)
}

func (p *Prize) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement.
}

func (p *Prize) valHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement Data Validation.
}

func (p *Prize) dataVal(requestMethod string) bool {
	//@todo: Implement Data Validation.
	return false
}
