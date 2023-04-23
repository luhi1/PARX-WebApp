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
		attendanceQ, err := db.Query("select users.UserID, users.StudentName, userevents.Attended from userevents left join users on userevents.UserID = users.UserID where EventID = ?", eventID)
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
		}
		if e.Active {
			e.Attendance = append(e.Attendance, a)
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
	e.Points, _ = strconv.Atoi(request.FormValue("Points"))
	e.EventDescription = request.FormValue("EventDescription")
	e.EventDate = request.FormValue("EventDate")
	e.RoomNumber, _ = strconv.Atoi(request.FormValue("RoomNumber"))
	e.AdvisorNames = request.FormValue("AdvisorNames")
	e.Location = request.FormValue("Location")
	e.LocationDescription = request.FormValue("LocationDescription")
	e.Sport = request.FormValue("Sport")
	e.SportDescription = request.FormValue("SportDescription")
	attendanceNum, _ := strconv.Atoi(request.FormValue("Atendee")[7:])
	for attendanceNum >= 0 {
		attendanceBruh := strconv.Itoa(attendanceNum)
		e.Attendance[attendanceNum].StudentNumber, _ = strconv.Atoi(request.FormValue("Attendee" + attendanceBruh))
		e.Attendance[attendanceNum].Attended = "true"
		attendanceNum--
	}

	if e.dataVal(strings.TrimPrefix(request.URL.Path, "/eventValidation/")) {
		insert, _ := db.Exec("insert into sports(sportname, sportdescription) values(?, ?);", e.Sport, e.SportDescription)
		fmt.Println(insert.RowsAffected())
		sportID, err := insert.LastInsertId()
		if err != nil {
			return
		}
		result, err := db.Exec(
			"update events set Points = ?,EventDescription = ?, EventDate = ?, RoomNumber = ?, Advisors = ?, Location = ?, LocationDescription = ?, SportID = ?",
			e.Points,
			e.EventDescription,
			e.EventDate,
			e.RoomNumber,
			e.AdvisorNames,
			e.Location,
			e.Location,
			sportID,
		)
		if err != nil {
			return
		}
		fmt.Println(result.RowsAffected())

		//Write the code to make attendance struct -> database;
	}
}

func (e *EventInfo) dataVal(requestMethod string) bool {
	if e.Points < 1 || e.EventDescription == "" || e.EventDate == "" || e.RoomNumber >= 1 || e.AdvisorNames == "" || e.Location == "" {
		return false
	}
	for i := 0; i < len(e.Attendance); i++ {
		if e.Attendance[i].StudentNumber < 1 || e.Attendance[i].StudentName == "" {
			return false
		}
	}
	return true
}
