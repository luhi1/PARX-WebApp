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

//@TODO: Test the main.go function (tie it all together with a lot of mock requests)
func TestTplExec(t *testing.T) {
	testData := UserData{}
	templateNames := []string{"error.gohtml", "login.gohtml", "signup.gohtml"}
	for i := 0; i < len(templateNames); i++ {
		loginTemp := template.Must(template.ParseFiles(templateNames[i]))
		w := httptest.NewRecorder()
		err := loginTemp.Execute(w, testData.valid)

		if err != nil {
			t.Error(err)
		}
		expectedOutput := string(w.Body.Bytes())
		w = httptest.NewRecorder()
		err = tplExec(w, templateNames[i], nil)
		if err != nil {
			return
		}
		output := string(w.Body.Bytes())
		if expectedOutput != output {
			t.Logf("You done messed up your outputs. Bad tplEXEC!!!!")
			t.Logf("Expected %s: %s", templateNames[i], expectedOutput)
			t.Logf("Got: %s", output)
		}
	}
	t.Logf("Templates able to be correctly loaded.")
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

		hasher := sha256.New()
		hasher.Write([]byte(passwords[i]))
		expected := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		if result != expected {
			t.Errorf("Incorrect hash for %s FAILED. Expected %s, got %s\n", passwords[i], expected, result)
			return
		}
	}
	t.Logf("All hashes as expected.")
}
