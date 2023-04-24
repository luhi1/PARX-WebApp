package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type StudentAttendance struct {
	StudentName   string
	StudentNumber int
	Attended      string
}

type EventInfo struct {
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
	Attendance          []StudentAttendance
	Active              bool
}

func (e *EventInfo) GETHandler(writer http.ResponseWriter, request *http.Request) {
	events := []EventInfo{}

	rows, err := db.Query("select EventID, EventName,events.Points,EventDescription,EventDate,RoomNumber, Location, LocationDescription,sports.SportName,sports.SportDescription, events.Advisors, events.Active from events left join sports on events.SportID = sports.ID")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		*e = EventInfo{}
		a := StudentAttendance{}
		rows.Scan(&e.EventID, &e.EventName, &e.Points, &e.EventDescription, &e.EventDate, &e.RoomNumber, &e.Location, &e.LocationDescription, &e.Sport, &e.SportDescription, &e.AdvisorNames, &e.Active)
		attendanceQ, err := db.Query("select users.UserID, users.StudentName, userevents.Attended from userevents left join users on userevents.UserID = users.UserID where EventID = ?", e.EventID)
		if err != nil {
			fmt.Println(err)
			return
		}
		for attendanceQ.Next() {
			attendanceQ.Scan(&a.StudentNumber, &a.StudentName, &a.Attended)
			if e.Active {
				if a.Attended == "true" {
					a.Attended = "checked"
				}
			}
			e.Attendance = append(e.Attendance, a)
		}
		if e.Active {
			events = append(events, *e)
		}
	}

	//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

	//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
	//Also, we need to find out how to get signup to upload to db and login to get
	//We can probably just do different interactions for get/post requests to the home, same way we did
	err = tplExec(writer, "teacher_events.gohtml", events)
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func (e *EventInfo) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "teacher_create_event.gohtml", nil)
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func (e *EventInfo) valHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	err = request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "./error", 303)
		return
	}
	e.EventName = request.FormValue("EventName")
	e.Points, _ = strconv.Atoi(request.FormValue("Points"))
	e.EventDescription = request.FormValue("EventDescription")
	e.EventDate = request.FormValue("EventDate")
	e.RoomNumber, _ = strconv.Atoi(request.FormValue("RoomNumber"))
	e.AdvisorNames = request.FormValue("AdvisorNames")
	e.Location = request.FormValue("Location")
	e.LocationDescription = request.FormValue("LocationDescription")
	e.Sport = request.FormValue("Sport")
	e.SportDescription = request.FormValue("SportDescription")
	e.Active = true
	for i := 0; i < len(request.Form["Attendee"]); i++ {
		currentHomie, _ := strconv.Atoi(request.Form["Attendee"][i])
		e.Attendance = append(e.Attendance, StudentAttendance{StudentNumber: currentHomie, Attended: "true"})
	}
	if e.dataVal(strings.TrimPrefix(request.URL.Path, "/eventValidation/")) {
		check := db.QueryRow("select ID from sports where SportName = ?", e.Sport)
		var sportID int
		err := check.Scan(&sportID)
		if err != nil {
			sportID = -1
		}
		if sportID == -1 {
			insert, _ := db.Exec("insert into sports(sportname, sportdescription) values(?, ?);", e.Sport, e.SportDescription)
			fmt.Println(insert.RowsAffected())
			getSportID, err := insert.LastInsertId()
			if err != nil {
				return
			}
			sportID = int(getSportID)
		}
		sID := strconv.Itoa(sportID)
		points := strconv.Itoa(e.Points)
		roomNumber := strconv.Itoa(e.RoomNumber)

		result, err := db.Exec("update events set events.Points = ?, EventDescription = ?, EventDate = ?, RoomNumber = ?, Advisors = ?, Location = ?, LocationDescription = ?, SportID = ? where events.EventName = ?",
			points, e.EventDescription, e.EventDate, roomNumber, e.AdvisorNames, e.Location, e.LocationDescription, sID, e.EventName,
		)
		if err != nil {
			fmt.Println(err)
		}

		insert := db.QueryRow("select EventID from events where EventName = ?;", e.EventName)
		insert.Scan(&e.EventID)
		minion, _ := db.Exec("update userevents set Attended = 'false' where EventID = ?;", e.EventID)
		fmt.Println(minion.RowsAffected())
		for i := 0; i < len(e.Attendance); i++ {
			fmt.Println(e.Attendance[i].StudentNumber)
			vector, _ := db.Exec("update userevents set Attended = 'true' where EventID = ? and UserID = ?", e.EventID, e.Attendance[i].StudentNumber)
			fmt.Println(vector.RowsAffected())
		}
		fmt.Println(e.Attendance)
		fmt.Println(result.RowsAffected())
	}
	http.Redirect(writer, request, "../teacherEvents", 307)
}

func (e *EventInfo) dataVal(requestMethod string) bool {
	if e.Points < 0 || e.EventDescription == "" || e.EventDate == "" || e.RoomNumber > 1 || e.AdvisorNames == "" || e.Location == "" {
		return false
	}
	for i := 0; i < len(e.Attendance); i++ {
		if e.Attendance[i].StudentNumber < 1 {
			return false
		}
	}
	return true
}

func (e *EventInfo) removeHandler(writer http.ResponseWriter, request *http.Request) {
	e.EventName = request.FormValue("EventName")
	eventID := db.QueryRow("select EventID from events where EventName = ?", e.EventName)
	eventID.Scan(&e.EventID)
	exec, err := db.Exec("update events set Active = 0 where EventID = ?", e.EventID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())
	http.Redirect(writer, request, "./teacherEvents", 307)
}
