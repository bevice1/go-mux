package main

import (
        "net/http"
        "net/http/httptest"
        "os"
        "testing"
        "time"
        "strconv"
        "crypto/sha256"
        "math/rand"
        "fmt"
       )

var app App

func TestMain(m *testing.M) {
    app.Initialize()

        code := m.Run()
        os.Exit(code)
}


func TestGreeting(m *testing.T) {



    req, _ := http.NewRequest("GET", "/greeting", nil)
        req.Header.Set("HashREST", generateHashValue("http://localhost:8010/greeting", 1))
        response := executeRequest(req)

        checkResponseCode(m, http.StatusOK, response.Code)

        if body := response.Body.String(); body != `Hello, World!` {
            m.Errorf("Expected 'Hello World'. Got '%s'", body)
        }
}

func TestList(m *testing.T) {
    req, _ := http.NewRequest("GET", "/list", nil)
        req.Header.Set("HashREST", generateHashValue("http://localhost:8010/list", 2))
        response := executeRequest(req)

        checkResponseCode(m, http.StatusOK, response.Code)
        var string = `["a","b","c"]`

        if body := response.Body.String(); body != string {
            m.Errorf(string + "got ", body)
        }

    req.Header.Set("HashREST", "01234567890")
        response = executeRequest(req)
        checkResponseCode(m, http.StatusBadRequest, response.Code)
}

func TestUploadImage(m *testing.T) {
    req, _ := http.NewRequest("POST", "/upload", nil)
        req.Header.Set("HashREST", generateHashValue("http://localhost:8010/upload", 3))
        response := executeRequest(req)

        checkResponseCode(m, http.StatusOK, response.Code)

        if body := response.Body.String(); body != `Image uploaded` {
            m.Errorf("Expected 'Image uploaded'. Got '%s'", body)
        }
}

func generateHashValue(url string, difficulty int) string {
timestamp := time.Now().Unix()
               randString := randString(6)
               counter := 0
               hashString := ""
               for{
data := strconv.FormatInt(timestamp, 10) + ";" + url + ";" + randString + ";" + strconv.Itoa(counter)
          hashBytes := sha256.Sum256([]byte(data))
          hashString = fmt.Sprintf("%x", hashBytes)
          if startsWithZeroes(hashString, difficulty) {
              break
          }
      counter++


               }
           fmt.Println(hashString)
               return hashString
}
func randString(n int) string {
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

        b := make([]rune, n)
        for i := range b {
            b[i] = letter[rand.Intn(len(letter))]
        }
    return string(b)
}


func executeRequest(req *http.Request) *httptest.ResponseRecorder {
rr := httptest.NewRecorder()
        app.Router.ServeHTTP(rr, req)
        return rr
}

func checkResponseCode(t *testing.T, expected int, actual int) {
    if expected != actual {
        t.Errorf("Expected response Code %d. Got %d\n", expected, actual)
    }
}
