package main

import (
	"net/http"
)

type Credentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
} 

func (app *application) ping(w http.ResponseWriter, r *http.Request) {

	app.infoLog.Println("Ping received")
	w.Write([]byte("OK"))
}
