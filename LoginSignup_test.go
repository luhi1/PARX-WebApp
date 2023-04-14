package main

import (
	"crypto/sha256"
	"encoding/base64"
	"html/template"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//@TODO: TO IMPROVE COVERAGE, MAKE IT SO THAT THERE IS NO
/*
	if err != nil {
		return
	}
*/
//@TODO: IN THE MAIN CODE
func TestUserData_GETHandler(t *testing.T) {
	loginTemp := template.Must(template.ParseFiles("login.gohtml"))
	expectedValidity := ""
	testData := UserData{}
	testData.valid = DisplayError{"Not Empty"}
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/login", nil)
	if err != nil {
		t.Error(err)
	}
	err = loginTemp.Execute(w, testData.valid)

	if err != nil {
		t.Error(err)
	}
	expectedHTML := string(w.Body.Bytes())

	w = httptest.NewRecorder()

	testData.GETHandler(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected Status 200, got %d", w.Code)
	}

	if testData.valid.ErrorDescription != expectedValidity {
		t.Errorf("Expected Cleared Error Description, got %s", testData.valid.ErrorDescription)
	}
	bodyStr := string(w.Body.Bytes())
	if len(bodyStr) <= 0 {
		t.Errorf("Expected a more interesting response body. Error: 0 size body")
	}
	if bodyStr != expectedHTML {
		t.Errorf("Wow you didn't boot up the right thing. Error: Template not parsed or executed correctly.")
		t.Errorf("Expected %s", expectedHTML)
		t.Errorf("Got %s", bodyStr)
	}
}

func TestUserData_POSTHandler(t *testing.T) {
	signupTemp := template.Must(template.ParseFiles("signup.gohtml"))
	expectedValidity := ""
	testData := UserData{}
	testData.valid = DisplayError{"Not Empty"}
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/signup", nil)
	if err != nil {
		t.Error(err)
	}
	err = signupTemp.Execute(w, testData.valid)

	if err != nil {
		t.Error(err)
	}
	expectedHTML := string(w.Body.Bytes())

	w = httptest.NewRecorder()

	testData.POSTHandler(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected Status 200, got %d", w.Code)
	}

	if testData.valid.ErrorDescription != expectedValidity {
		t.Errorf("Expected Cleared Error Description, got %s", testData.valid.ErrorDescription)
	}
	bodyStr := string(w.Body.Bytes())
	if len(bodyStr) <= 0 {
		t.Errorf("Expected a more interesting response body. Error: 0 size body")
	}
	if bodyStr != expectedHTML {
		t.Errorf("Wow you didn't boot up the right thing. Error: Template not parsed or executed correctly.")
		t.Errorf("Expected %s", expectedHTML)
		t.Errorf("Got %s", bodyStr)
	}
}

func TestUserData_ValHandler(t *testing.T) {
	//@TODO: Figure out how to emulate postforms.
	//@TODO: Implement.
}

func TestUserData_dataVal(t *testing.T) {
	rand.Seed(time.Now().Unix())
	testData := UserData{}
	var expected bool
	requestMethod := "signup"

	for i := 0; i < 1000000; i++ {
		if rand.Intn(2) != 0 {
			requestMethod = "login"
		}

		expected = false

		var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321!@#$%^&*()"
		str := make([]byte, rand.Intn(16)+1)
		for k := range str {
			str[k] = chars[rand.Intn(len(chars))]
		}
		testData.passwordHash = hashPswd(string(str))
		testData.IdNumber = rand.Intn(9999998) + 1

		if requestMethod == "signup" {
			testData.Grade = rand.Intn(4) + 9
			testData.Name = string(str)
		}

		punishment := rand.Intn(11)

		if punishment == 1 && requestMethod == "signup" {
			testData.Grade = 13
		} else if punishment == 3 && requestMethod == "signup" {
			testData.Name = ""
		} else if punishment == 5 {
			testData.passwordHash = hashPswd("")
		} else if punishment == 7 {
			testData.IdNumber = 999999999999
		} else {
			expected = true
		}

		if testData.dataVal(requestMethod) != expected {
			t.Errorf("DID NOT STOP A BAD INPUT. Expected %t using method %s on iteration %d using random number %d", expected, requestMethod, i, punishment)
			return
		}
	}
	t.Logf("STOPPED BAD INPUTS")
}