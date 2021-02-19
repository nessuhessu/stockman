package main

import (
	"log"
	"html/template"
	"net/http"
	"os"
	"time"
)

var logInFile bool = true
var webPort = ":8080"

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	templateCache map[string]*template.Template
	webPort string
}

var app *application

func main() {
	createNewApp()
}

func createNewApp() *application {

	infoLog := &log.Logger{}
	errorLog := &log.Logger{}
	if logInFile == true {
		// Open log file
		f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		} else {
			infoLog = log.New(f, "INFO\t", log.Ldate|log.Ltime)
			errorLog = log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		}
	} else {
		infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	}

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app = &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		templateCache: templateCache,
		webPort:    webPort,
	}
	app.listen(app.webPort, app.webRoutes())
	return app
}

func (app *application) listen(addr string, handler http.Handler) {

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: app.errorLog,
		Handler:  handler,
		// Idle, Read and Write timeouts to the server.
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.infoLog.Printf("Starting server on: %v\n", addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
