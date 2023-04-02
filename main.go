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

type userData struct {
	Name         string
	Grade        int
	IdNumber     int
	passwordHash string
}

type eventInfo struct {
	Points           int
	EventDescription string
	EventDate        string
	RoomNumber       int
	//Might Need to be for array somewhere along the way
	AdvisorNames        string
	Location            string
	LocationDescription string
	Sport               string
	SportDescription    string
	//Consider Changing the Types of these last 2 over here!
	EventImage      string
	StudentName     string
	StudentNumber   int
	StudentAttended bool
	inputImage      fs.File
}

type DisplayError struct {
	ErrorDescription string
}

//USE POINTERS INSTEAD OF PACKAGE LEVEL STATE
//ANY IF ERR != NIL RETURN STATEMENTS SHOULD BE FIXED LATER!

// Start server run, files, and other shit.
func main() {
	credentialCheck := ""
	userInfo := userData{}

	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		err := tplExec(writer, "login.gohtml", DisplayError{credentialCheck})
		if err != nil {
			return
		}
		credentialCheck = ""
	})

	http.HandleFunc("/signup", func(writer http.ResponseWriter, request *http.Request) {
		err := tplExec(writer, "signup.gohtml", DisplayError{credentialCheck})
		if err != nil {
			return
		}
		credentialCheck = ""
	})

	http.HandleFunc("/validation/", func(writer http.ResponseWriter, request *http.Request) {
		var err error
		err = request.ParseForm()
		if err != nil {
			http.Redirect(writer, request, "./error", 303)
			return
		}
		userInfo.Name = request.FormValue("name")
		userInfo.Grade, err = strconv.Atoi(request.FormValue("grade"))
		userInfo.IdNumber, err = strconv.Atoi(request.FormValue("IdNumber"))
		userInfo.passwordHash = hashPswd(request.FormValue("password"))

		if err != nil || checkData(strings.TrimPrefix(request.URL.Path, "/validation/"), &userInfo) {
			http.Redirect(writer, request, "../teacher_events", 307)
		} else {
			credentialCheck = "Invalid Credentials"
			if strings.TrimPrefix(request.URL.Path, "/validation/") == "signup" {
				http.Redirect(writer, request, "../signup", 303)
			} else {
				http.Redirect(writer, request, "../login", 303)
			}
		}
	})

	http.HandleFunc("/teacher_events", func(writer http.ResponseWriter, request *http.Request) {
		if (userInfo != userData{}) {
			//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

			//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
			//Also, we need to find out how to get signup to upload to db and login to get
			//We can probably just do different interactions for get/post requests to the home, same way we did
			var Events []eventInfo
			for i := 0; i < 3; i++ {
				Events = append(Events, eventInfo{
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
			if err != nil {
				return
			}
		} else {
			http.Redirect(writer, request, "./login", 303)
		}
	})

	http.HandleFunc("/teacher_create_event", func(writer http.ResponseWriter, request *http.Request) {
		err := tplExec(writer, "teacher_create_event.gohtml", nil)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		userInfo = userData{}
		http.Redirect(writer, request, "./login", 307)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			err := tplExec(writer, "error.gohtml", nil)
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
	if err != nil {
		return err
	}
	return nil
}

func checkData(requestMethod string, userInfo *userData) bool {

	//Check if ID Number is blank or out of bounds
	//Check if password is blank
	valid := false
	if (*userInfo != userData{}) &&
		(userInfo.IdNumber > 0 &&
			userInfo.IdNumber < 9999999 &&
			userInfo.passwordHash != hashPswd("")) {

		valid = true
	}

	if requestMethod == "signup" && ((userInfo.Grade != 9 && userInfo.Grade != 10 &&
		userInfo.Grade != 11 && userInfo.Grade != 12) || userInfo.Name == "") {
		valid = false
	}
	//here we can check if the password matches the one in our database
	if !valid {
		*userInfo = userData{}
	}
	return valid
}

func hashPswd(pwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
