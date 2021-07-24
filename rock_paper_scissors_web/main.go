package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsWeb/rps"
)

func homePage(writer http.ResponseWriter, request *http.Request) {
	log.Println(request)
	renderTemplate(writer, "index.html")
}

func playRound(writer http.ResponseWriter, request *http.Request) {

	log.Println(request)

	choiceString := request.URL.Query().Get("c")
	round := rps.PlayRound(choiceString)
	out, err := json.MarshalIndent(round, "", "    ")
	check(err)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}

func main() {

	log.Println("Starting server")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)
	http.ListenAndServe(":8080", nil)

}

func renderTemplate(writer http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	check(err)
	err = t.Execute(writer, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
