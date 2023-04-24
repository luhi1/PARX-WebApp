package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type studentEventInfo struct {
	EventID             int
	EventName           string
	Points              int
	EventDescription    string
	EventDate           string
	RoomNumber          int
	AdvisorNames        string
	Location            string
	LocationDescription string
	Sport               string
	SportDescription    string
	Active              bool
	U                   *UserData
	Attended            bool
}

func (se *studentEventInfo) GETStudentHandler(writer http.ResponseWriter, request *http.Request) {
	userInfo := se.U
	events := []studentEventInfo{}
	rows, err := db.Query("select EventID, EventName,events.Points,EventDescription,EventDate,RoomNumber, Location, LocationDescription,sports.SportName,sports.SportDescription, events.Advisors, events.Active from events left join sports on events.SportID = sports.ID")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		*se = studentEventInfo{U: userInfo}
		rows.Scan(&se.EventID, &se.EventName, &se.Points, &se.EventDescription, &se.EventDate, &se.RoomNumber, &se.Location, &se.LocationDescription, &se.Sport, &se.SportDescription, &se.AdvisorNames, &se.Active)
		se.Attended = true
		insert := db.QueryRow("select * from userevents where EventID = ? and UserID = ?", se.EventID, se.U.IdNumber)
		err := insert.Scan()
		if err.Error() == "sql: no rows in result set" {
			se.Attended = false
			fmt.Println(err)
		}
		if se.Active {
			events = append(events, *se)
		}
	}
	err = multiTplExec(writer, "studentEvents.gohtml", events, "home.gohtml")
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func (se *studentEventInfo) dropOutHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}
	se.EventID, err = strconv.Atoi(request.FormValue("EventID"))
	if err != nil {
		return
	}
	exec, err := db.Exec("delete from userevents where userevents.userID = ? and userevents.eventID = ?", se.U.IdNumber, se.EventID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())

	addition, _ := db.Exec("update users set Points = Points-10 where userID = ?", se.U.IdNumber)
	fmt.Println(addition.RowsAffected())
	insert := db.QueryRow("select Points from users where userID = ?", se.U.IdNumber)
	insert.Scan(&se.U.Points)
	http.Redirect(writer, request, "./studentEvents", 307)
}
func (se *studentEventInfo) studentSignupEventHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}
	se.EventID, err = strconv.Atoi(request.FormValue("EventID"))
	if err != nil {
		return
	}
	exec, err := db.Exec("insert into userevents(userid, eventid, attended) values(?,?,'false')", se.U.IdNumber, se.EventID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())

	addition, _ := db.Exec("update users set Points = Points+10 where userID = ?", se.U.IdNumber)
	insert := db.QueryRow("select Points from users where userID = ?", se.U.IdNumber)
	insert.Scan(&se.U.Points)
	fmt.Println(addition.RowsAffected())
	http.Redirect(writer, request, "./studentEvents", 307)
}
