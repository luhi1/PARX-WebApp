package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
)

type UserData struct {
	Name         string
	Grade        int
	IdNumber     int
	passwordHash string
	valid        DisplayError
}

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type EventInfo struct {
	Points              int
	EventDescription    string
	EventDate           string
	RoomNumber          int
	AdvisorNames        string
	Location            string
	LocationDescription string
	Sport               string
	SportDescription    string
	EventImage          string
	StudentName         string
	StudentNumber       int
	StudentAttended     bool
	inputImage          fs.File
}

type DisplayError struct {
	ErrorDescription string
}

// TeacherPageHandlers Consider creating a generic interface for both teacher and student to implement.
type TeacherPageHandlers interface {
	GETHandler(writer http.ResponseWriter, request *http.Request)
	POSTHandler(writer http.ResponseWriter, request *http.Request)
	valHandler(writer http.ResponseWriter, request *http.Request)
	dataVal(requestMethod string) bool
}

// USE POINTERS INSTEAD OF PACKAGE LEVEL STATE
var userInfo = UserData{}
var eventInfo = EventInfo{}

func (u *UserData) GETHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "login.gohtml", u.valid)
	//@TODO: REMOVE
	if err != nil {
		return
	}
	u.valid = DisplayError{""}
}

func (u *UserData) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "signup.gohtml", u.valid)
	//@TODO: REMOVE
	if err != nil {
		return
	}
	u.valid = DisplayError{""}
}

func (u *UserData) valHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	err = request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "./error", 303)
		return
	}
	u.Name = request.FormValue("name")
	u.Grade, err = strconv.Atoi(request.FormValue("grade"))
	u.IdNumber, err = strconv.Atoi(request.FormValue("IdNumber"))
	u.passwordHash = hashPswd(request.FormValue("password"))

	if err != nil || u.dataVal(strings.TrimPrefix(request.URL.Path, "/userValidation/")) {
		http.Redirect(writer, request, "../teacher_events", 307)
	} else {
		u.valid = DisplayError{"Invalid Credentials"}
		if strings.TrimPrefix(request.URL.Path, "/userValidation/") == "signup" {
			http.Redirect(writer, request, "../signup", 303)
		} else {
			http.Redirect(writer, request, "../login", 303)
		}
	}
}

func (u *UserData) dataVal(requestMethod string) bool {
	valid := false
	if (*u != UserData{}) &&
		(u.IdNumber > 0 &&
			u.IdNumber < 9999999 &&
			u.passwordHash != hashPswd("")) {

		valid = true
	}

	if requestMethod == "signup" && ((u.Grade != 9 && u.Grade != 10 &&
		u.Grade != 11 && u.Grade != 12) || u.Name == "") {
		valid = false
	}
	//here we can check if the password matches the one in our database
	if !valid {
		*u = UserData{}
	}
	return valid
}

func (e *EventInfo) GETHandler(writer http.ResponseWriter, request *http.Request) {
	if (userInfo != UserData{}) {
		//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

		//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
		//Also, we need to find out how to get signup to upload to db and login to get
		//We can probably just do different interactions for get/post requests to the home, same way we did
		var Events []EventInfo
		for i := 0; i < 3; i++ {
			Events = append(Events, EventInfo{
				Points:              0,
				EventDescription:    "asdf",
				EventDate:           "2017-06-01",
				RoomNumber:          0,
				AdvisorNames:        "asdf",
				Location:            "asdf",
				LocationDescription: "asdf",
				Sport:               "asdf",
				SportDescription:    "asdf",
				EventImage:          "https://imgs.search.brave.com/ToRVheIVFOHdWRebW6v6BriMZf_slwrqoAXvU-I62CY/rs:fit:1200:1200:1/g:ce/aHR0cHM6Ly90aGV3/b3dzdHlsZS5jb20v/d3AtY29udGVudC91/cGxvYWRzLzIwMTUv/MDEvbmF0dXJlLWlt/YWdlcy4uanBn",
				StudentName:         "asdf",
				StudentNumber:       0,
				StudentAttended:     true,
			})
		}
		err := tplExec(writer, "teacher_events.gohtml", Events)
		//@TODO: REMOVE
		if err != nil {
			return
		}
	} else {
		http.Redirect(writer, request, "./login", 303)
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
	//@todo: Implement Data Validation.
}

func (e *EventInfo) dataVal(requestMethod string) bool {
	//@todo: Implement Data Validation.
	return false
}

// Start server run, files, and other shit.
func main() {
	http.HandleFunc("/login", userInfo.GETHandler)

	http.HandleFunc("/signup", userInfo.POSTHandler)

	http.HandleFunc("/userValidation/", userInfo.valHandler)

	http.HandleFunc("/teacher_events", eventInfo.GETHandler)

	http.HandleFunc("/teacher_create_event", eventInfo.POSTHandler)

	http.HandleFunc("/eventValidation/", eventInfo.POSTHandler)

	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		userInfo = UserData{}
		http.Redirect(writer, request, "./login", 307)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			err := tplExec(writer, "error.gohtml", nil)
			//@TODO: REMOVE
			if err != nil {
				return
			}
		} else {
			http.Redirect(writer, request, "./login", 301)
		}
	})

	/*@todo: Add this to the setup wizard eventually */
	fmt.Println("Server is running on port 8082")

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		fmt.Println("Error starting server, aborting tasks")
		panic(err)
	}
}

func tplExec(w http.ResponseWriter, filename string, information any) error {
	temp := template.Must(template.ParseFiles(filename))

	err := temp.Execute(w, information)
	//@TODO: REMOVE
	if err != nil {
		return err
	}
	return nil
}

func hashPswd(pwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
