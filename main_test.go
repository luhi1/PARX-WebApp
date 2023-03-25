package main

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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
	t.Logf("PASSED. All hashes as expected.")
}

func TestTplExec(t *testing.T) {
	templateNames := []string{"error.gohtml", "login.gohtml", "signup.gohtml"}
	for i := 0; i < len(templateNames); i++ {
		httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			err := tplExec(writer, templateNames[i], nil)
			if err != nil {
				t.Errorf("FAILED. Unable to load template %s", templateNames[i])
				return
			}
		}))
	}
	t.Logf("PASSED. All templates loaded.")
}

func hash(pwd string) string {

	hasher := sha256.New()
	hasher.Write([]byte(pwd))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return string(sha)
}
