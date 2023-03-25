package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
)

type userData struct {
}

// Start server run, files, and other shit.
func main() {

	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		tplExec(writer, "login.gohtml", nil)
	})
	http.HandleFunc("/signup", func(writer http.ResponseWriter, request *http.Request) {
		tplExec(writer, "signup.gohtml", nil)
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			tplExec(writer, "error.gohtml", nil)
		} else {
			http.Redirect(writer, request, "./login", 301)
		}
	})

	/*@todo: Add this to the setup wizard eventually */
	fmt.Println("Server is running on port 8081")

	err := http.ListenAndServe(":8081", nil)
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

func checkData(r *http.Request) userData {
	// Use this function to check if login/signup data is working
	return userData{}
}
func hashPswd(pwd string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return string(sha)
}
