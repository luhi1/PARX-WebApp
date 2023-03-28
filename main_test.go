package main

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestJoe(t *testing.T) {
	t.Logf(strings.TrimPrefix("http://localhost:8082/signup", "/"))
}
func TestHash(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var passwords []string
	for i := 0; i < 50; i++ {
		var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321!@#$%^&*()"
		str := make([]byte, rand.Intn(16)+1)
		for k := range str {
			str[k] = chars[rand.Intn(len(chars))]
		}
		passwords = append(passwords, string(str))
		result := hashPswd(passwords[i])
		expected := hash(passwords[i])
		if result != expected {
			t.Errorf("Incorrect hash for %s FAILED. Expected %s, got %s\n", passwords[i], expected, result)
			return
		}
	}
	t.Logf("All hashes as expected.")
}

func TestTplExec(t *testing.T) {
	templateNames := []string{"error.gohtml", "login.gohtml", "signup.gohtml", "home.gohtml"}
	for i := 0; i < len(templateNames); i++ {
		httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			err := tplExec(writer, templateNames[i], nil)
			if err != nil {
				t.Errorf("Unable to load template %s", templateNames[i])
				return
			}
		}))
	}
	t.Logf("All templates loaded.")
}

func TestDataValidation(t *testing.T) {
	rand.Seed(time.Now().Unix())
	var testData userData
	var expected bool
	requestMethod := "signup"

	for i := 0; i < 1000; i++ {
		if rand.Intn(2) != 0 {
			requestMethod = "login"
		}

		expected = false
		testData = userData{}

		var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321!@#$%^&*()"
		str := make([]byte, rand.Intn(16)+1)
		for k := range str {
			str[k] = chars[rand.Intn(len(chars))]
		}
		testData.passwordHash = hashPswd(string(str))
		testData.IdNumber = rand.Intn(9999999) + 1

		if requestMethod == "POST" {
			testData.Grade = rand.Intn(4) + 9
			testData.Name = string(str)
		}

		punishment := rand.Intn(11)

		if punishment%11 == 0 && requestMethod == "POST" {
			testData.Grade = 13
		} else if punishment%3 == 0 && requestMethod == "POST" {
			testData.Name = ""
		} else if punishment%5 == 0 {
			testData.passwordHash = hashPswd("")
		} else if punishment%7 == 0 {
			testData.IdNumber = 999999999999
		} else {
			expected = true
		}

		if checkData(requestMethod, &testData) != expected {
			t.Errorf("DID NOT STOP A BAD INPUT. Expected %t using method %s on iteration %d", expected, requestMethod, i)
			t.Logf(testData.passwordHash)
			t.Logf(testData.Name)
			t.Logf("%d", testData.Grade)
			t.Logf("%d", testData.IdNumber)
			return
		}
	}
	t.Logf("STOPPED BAD INPUTS")
}

func hash(pwd string) string {

	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}
