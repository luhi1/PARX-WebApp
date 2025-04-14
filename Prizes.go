package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Prize struct {
	PrizeName    string
	PrizeWinners []StudentAttendance
	Points       int
}

func (p *Prize) GETHandler(writer http.ResponseWriter, request *http.Request) {
	var prizes []Prize
	//select PrizeName, StudentName from userprizes left join prizes on userprizes.PrizeID = prizes.ID left join users on userprizes.UserID = users.UserID
	insert, err := db.Query("select PrizeName from Prizes")
	if err != nil {
		fmt.Println(err)
		return
	}
	for insert.Next() {
		*p = Prize{}
		insert.Scan(&p.PrizeName)
		prizes = append(prizes, *p)
	}

	rows, err := db.Query("select PrizeName, StudentName, Attended, Users.UserID from UserPrizes left join Prizes on UserPrizes.PrizeID = Prizes.ID left join Users on UserPrizes.UserID = Users.UserID")
	if err != nil {
		fmt.Println(err)
		return
	}
	currentPrizeNumber := 0
	currentPrizeName := prizes[0].PrizeName
	for rows.Next() {
		*p = Prize{}
		a := StudentAttendance{}
		rows.Scan(&p.PrizeName, &a.StudentName, &a.Attended, &a.StudentNumber)
		if a.Attended == "true" {
			a.Attended = "checked"
		}
		if p.PrizeName != currentPrizeName {
			currentPrizeNumber++
			currentPrizeName = prizes[currentPrizeNumber].PrizeName
		}
		prizes[currentPrizeNumber].PrizeWinners = append(prizes[currentPrizeNumber].PrizeWinners, a)
	}

	tplExec(writer, "prizes.gohtml", prizes)
}

func (p *Prize) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "createPrize.gohtml", p)
	if err != nil {
		return
	}
}

func (p *Prize) valHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	p.PrizeName = request.FormValue("PrizeName")
	var prizeID int
	insert := db.QueryRow("select ID from Prizes where PrizeName = ?", p.PrizeName)
	insert.Scan(&prizeID)
	for i := 0; i < len(request.Form["prizeWinner"]); i++ {
		currentHomie, _ := strconv.Atoi(request.Form["prizeWinner"][i])
		p.PrizeWinners = append(p.PrizeWinners, StudentAttendance{StudentNumber: currentHomie, Attended: "true"})
	}
	update, _ := db.Exec("update UserPrizes set Attended = 'false' where PrizeID = ?;", prizeID)
	fmt.Println(update.RowsAffected())
	for i := 0; i < len(p.PrizeWinners); i++ {
		vector, _ := db.Exec("update UserPrizes set Attended = 'true' where PrizeID = ? and UserID = ?", prizeID, p.PrizeWinners[i].StudentNumber)
		fmt.Println(vector.RowsAffected())
	}
	http.Redirect(writer, request, "./prizes", 307)
}

func (p *Prize) dataVal(requestMethod string) bool {
	if p.PrizeName == "" || p.Points < 1 {
		return false
	}
	return true
}

func (p *Prize) createPrize(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}
	p.PrizeName = request.FormValue("PrizeName")
	p.Points, _ = strconv.Atoi(request.FormValue("Points"))
	fmt.Println(p.dataVal(""))
	if p.dataVal("") {
		insert, _ := db.Exec("insert into Prizes(PrizeName, PointThreshold) values(?,?)", p.PrizeName, p.Points)
		fmt.Println(insert.RowsAffected())
	}
	http.Redirect(writer, request, "./prizes", 307)
}
