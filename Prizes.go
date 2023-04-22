package main

import (
	"fmt"
	"net/http"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type Prize struct {
	PrizeName    string
	PrizeWinners []StudentAttendance
}

func (p *Prize) GETHandler(writer http.ResponseWriter, request *http.Request) {
	var prizes []Prize
	//select PrizeName, StudentName from userprizes left join prizes on userprizes.PrizeID = prizes.ID left join users on userprizes.UserID = users.UserID
	insert, err := db.Query("select PrizeName from prizes")
	if err != nil {
		fmt.Println(err)
		return
	}
	for insert.Next() {
		*p = Prize{}
		insert.Scan(&p.PrizeName)
		prizes = append(prizes, *p)
	}

	rows, err := db.Query("select PrizeName, StudentName, users.UserID from userprizes left join prizes on userprizes.PrizeID = prizes.ID left join users on userprizes.UserID = users.UserID")
	if err != nil {
		fmt.Println(err)
		return
	}
	currentPrizeNumber := 0
	currentPrizeName := prizes[0].PrizeName
	for rows.Next() {
		*p = Prize{}
		a := StudentAttendance{}
		rows.Scan(&p.PrizeName, &a.StudentName, &a.StudentNumber)
		if p.PrizeName != currentPrizeName {
			currentPrizeNumber++
			currentPrizeName = prizes[currentPrizeNumber].PrizeName
		}
		prizes[currentPrizeNumber].PrizeWinners = append(prizes[currentPrizeNumber].PrizeWinners, a)
	}

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
