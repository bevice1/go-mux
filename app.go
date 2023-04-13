package main

import (
        "fmt"
        "log"
        "net/http"
        "encoding/json"

        "github.com/gorilla/mux"
        _ "github.com/lib/pq"
       )

type App struct {
    Router *mux.Router
}

func (a *App) Initialize() {
    a.Router = mux.NewRouter()
        a.initializeRoutes()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(":8010", a.Router))
}


func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(code)
        w.Write(response)
}
func checkDifficulty(hashHeader string, neccesarryDifficulty int, w http.ResponseWriter) bool {

    if hashHeader == "" {
        fmt.Println("HashREST: ", hashHeader)
            http.Error(w, "Missing Authorization header", http.StatusBadRequest)
            return false
    } else if (startsWithZeroes(hashHeader, neccesarryDifficulty)) {
        return true
    } else {

        fmt.Println("HashREST: ", hashHeader)
            http.Error(w, "Difficulty not correct", http.StatusBadRequest)
            return false
    }

}

func startsWithZeroes(str string, n int) bool {
    if len(str) < n {
        return false 
    }
    for i := 0; i < n; i++ {
        if str[i] != '0' {
            return false 
        }
    }
    return true 
}

func (a *App) makeGreeting(w http.ResponseWriter, r *http.Request) {
hashHeader := r.Header.Get("HashREST")
                if(checkDifficulty(hashHeader, 1, w)) {
                    w.Write([]byte("Hello, World!"))
                }
}

func (a *App) getList(w http.ResponseWriter, r *http.Request) {
hashHeader := r.Header.Get("HashREST")
                if(checkDifficulty(hashHeader, 2, w)) {
                    w.Write([]byte(`["a","b","c"]`))
                }
}

func (a *App) uploadImage(w http.ResponseWriter, r *http.Request) {
hashHeader := r.Header.Get("HashREST")
                if(checkDifficulty(hashHeader, 3, w)) {
                    w.Write([]byte("Image uploaded"))
                }
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/greeting", a.makeGreeting).Methods("GET")
        a.Router.HandleFunc("/list", a.getList).Methods("GET")
        a.Router.HandleFunc("/upload", a.uploadImage).Methods("POST")
}
