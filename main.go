package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type HomePageData struct {
	ProjectName string
	DateTime string
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()

	dateTime := fmt.Sprintf("%d-%d-%d %d:%d:%d\n",
	currentTime.Year(),
	currentTime.Month(),
	currentTime.Day(),
	currentTime.Hour(),
	currentTime.Hour(),
	currentTime.Second())

	projectName := os.Getenv("PROJECT_NAME")
	if projectName == "" {
		projectName = "Example"
	}


	data := HomePageData{ProjectName: projectName, DateTime: dateTime}
	tpl.Execute(w, data)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)

	fmt.Println("Running...")
	http.ListenAndServe(":"+port, mux)
}