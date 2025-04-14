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

	rows, err := db.Query("select EventID, EventName,Events.Points,EventDescription,EventDate,RoomNumber, Location, LocationDescription,Sports.SportName,Sports.SportDescription, Events.Advisors, Events.Active from Events left join Sports on Events.SportID = Sports.ID")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		*e = EventInfo{}
		a := StudentAttendance{}
		rows.Scan(&e.EventID, &e.EventName, &e.Points, &e.EventDescription, &e.EventDate, &e.RoomNumber, &e.Location, &e.LocationDescription, &e.Sport, &e.SportDescription, &e.AdvisorNames, &e.Active)
		attendanceQ, err := db.Query("select Users.UserID, Users.StudentName, UserEvents.Attended from UserEvents left join Users on UserEvents.UserID = Users.UserID where EventID = ?", e.EventID)
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
		check := db.QueryRow("select ID from Sports where SportName = ?", e.Sport)
		var sportID int
		err := check.Scan(&sportID)
		if err != nil {
			sportID = -1
		}
		if sportID == -1 {
			insert, _ := db.Exec("insert into Sports(SportName, SportDescription) values(?, ?);", e.Sport, e.SportDescription)
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

		result, err := db.Exec("update Events set Events.Points = ?, EventDescription = ?, EventDate = ?, RoomNumber = ?, Advisors = ?, Location = ?, LocationDescription = ?, SportID = ? where Events.EventName = ?",
			points, e.EventDescription, e.EventDate, roomNumber, e.AdvisorNames, e.Location, e.LocationDescription, sID, e.EventName,
		)
		if err != nil {
			fmt.Println(err)
		}

		insert := db.QueryRow("select EventID from Events where EventName = ?;", e.EventName)
		insert.Scan(&e.EventID)

		minion, _ := db.Exec("update UserEvents set Attended = 'false' where EventID = ?;", e.EventID)
		fmt.Println(minion.RowsAffected())
		for i := 0; i < len(e.Attendance); i++ {
			//Change it to add Points when the homies sign up for an event.
			vector, _ := db.Exec("update UserEvents set Attended = 'true' where EventID = ? and UserID = ?", e.EventID, e.Attendance[i].StudentNumber)
			fmt.Println(vector.RowsAffected())
		}
		fmt.Println(e.Attendance)
		fmt.Println(result.RowsAffected())
	}
	http.Redirect(writer, request, "../teacherEvents", 307)
}

func (e *EventInfo) dataVal(requestMethod string) bool {
	fmt.Println(e)
	if e.Points < 0 || e.EventDescription == "" || e.EventDate == "" || e.RoomNumber < 1 || e.AdvisorNames == "" || e.Location == "" {
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
	eventID := db.QueryRow("select EventID from Events where EventName = ?", e.EventName)
	eventID.Scan(&e.EventID)
	exec, err := db.Exec("update Events set Active = 0 where EventID = ?", e.EventID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())

	insert, _ := db.Query("select UserID from UserEvents where Attended = 'false' and EventID = ?", e.EventID)

	var subtracters []int
	for insert.Next() {
		var currentSubtracter int
		insert.Scan(&currentSubtracter)
		subtracters = append(subtracters, currentSubtracter)
	}
	for i := 0; i < len(subtracters); i++ {
		addition, _ := db.Exec("update Users set Points = Points-10 where UserID = ?", subtracters[i])
		fmt.Println(addition.RowsAffected())
	}
	http.Redirect(writer, request, "./teacherEvents", 307)
}

func (e *EventInfo) createEvent(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
		fmt.Println(request.Form)
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
	if e.dataVal("") {
		check := db.QueryRow("select ID from Sports where SportName = ?", e.Sport)
		var sportID int
		check.Scan(&sportID)
		if err != nil {
			sportID = -1
		}
		if sportID == -1 {
			insert, _ := db.Exec("insert into Sports(SportName, SportDescription) values(?, ?);", e.Sport, e.SportDescription)
			fmt.Println(insert.RowsAffected())
			getSportID, err := insert.LastInsertId()
			if err != nil {
				return
			}
			sportID = int(getSportID)
		}
		result, err := db.Exec("insert into Events(EventName, Points, EventDescription, EventDate, RoomNumber, Advisors, Location, LocationDescription, SportID, Active) VALUES (?,?,?,?,?,?,?,?,?,?)",
			e.EventName, e.Points, e.EventDescription, e.EventDate, e.RoomNumber, e.AdvisorNames, e.Location, e.LocationDescription, sportID, e.Active)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result.RowsAffected())

	}
	http.Redirect(writer, request, "../teacherEvents", 307)
}
