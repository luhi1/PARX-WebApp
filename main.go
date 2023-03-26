package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type userData struct {
	Name         string
	Grade        int
	IdNumber     int
	passwordHash string
}

type DisplayError struct {
	ErrorDescription string
}

var userInfo = userData{
	"",
	-1,
	-1,
	"",
}

// Start server run, files, and other shit.
func main() {
	credentialCheck := ""

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

	http.HandleFunc("/validation", func(writer http.ResponseWriter, request *http.Request) {
		var err error
		err = request.ParseForm()
		if err != nil {
			http.Redirect(writer, request, "./error", 303)
			return
		}
		userInfo.Name = request.FormValue("name")
		userInfo.Grade, err = strconv.Atoi(request.FormValue("grade"))
		userInfo.IdNumber, err = strconv.Atoi(request.FormValue("studentNumber"))
		userInfo.passwordHash = hashPswd(request.FormValue("password"))

		//Temporary error handling, fix one day
		if err != nil {
			userInfo = userData{
				"",
				-1,
				-1,
				"",
			}
			credentialCheck = "Invalid Credentials"
			if request.Method == "POST" {
				http.Redirect(writer, request, "./signup", 303)
			} else {
				http.Redirect(writer, request, "./login", 303)
			}
			return
		}

		if checkData(request.Method) {
			http.Redirect(writer, request, "./home", 307)
		} else {
			credentialCheck = "Invalid Credentials"
			if request.Method == "POST" {
				http.Redirect(writer, request, "./signup", 303)
			} else {
				http.Redirect(writer, request, "./login", 303)
			}
			userInfo = userData{
				"",
				-1,
				-1,
				"",
			}
		}
	})

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		if userInfo.IdNumber != -1 && userInfo.passwordHash != "" {
			//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

			//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
			//Also, we need to find out how to get signup to upload to db and login to get
			//We can probably just do different interactions for get/post requests to the home, same way we did
			err := tplExec(writer, "home.gohtml", userInfo)
			if err != nil {
				return
			}
		} else {
			http.Redirect(writer, request, "./login", 303)
		}
	})

	http.HandleFunc("/logout", func(writer http.ResponseWriter, request *http.Request) {
		userInfo = userData{
			"",
			-1,
			-1,
			"",
		}
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
	fmt.Println("Server is running on port 8081")

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

func checkData(requestMethod string) bool {

	//Check if ID Number is blank or out of bounds
	//Check if password is blank
	if userInfo.IdNumber <= 0 ||
		userInfo.IdNumber >= 9999999 ||
		userInfo.passwordHash == hashPswd("") {
		return false
	}

	//If on signup screen, make sure all info is filled out
	//Check if grade is 9-12
	if requestMethod == "POST" {
		if userInfo.Name == "" || userInfo.Grade == -1 || userInfo.IdNumber == -1 || userInfo.passwordHash == "" {
			return false
		}
		if userInfo.Grade != 9 && userInfo.Grade != 10 &&
			userInfo.Grade != 11 && userInfo.Grade != 12 {
			return false
		}
	}
	//here we can check if the password matches the one in our database
	return true
}
func hashPswd(pwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
