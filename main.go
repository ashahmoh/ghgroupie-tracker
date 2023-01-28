package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type artist []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// put structs into variables for unmarshelling.
// Unmarshalling takes data from the struct and puts it into the variable
var artists artist
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	//1 . artists url get, put json into a variable and umarshall
	artistResp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		print("Can't get Artist URL!", err)
	}
	defer artistResp.Body.Close()
	artistsbody, err := ioutil.ReadAll(artistResp.Body)
	if err != nil {
		print("Error copying artist json into body", err)
	}

	err2 := json.Unmarshal(artistsbody, &artists) //interface MUST pass unmarshal a pointer!
	if err2 != nil {
		log.Fatalln("error unmarshalling artist", err2)
	}

	http.Handle("/Users/ashamohamed/Documents/GitHub/ghgroupie-tracker/assets/", http.StripPrefix("/Users/ashamohamed/Documents/GitHub/ghgroupie-tracker/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	fmt.Println("listening...")
	http.ListenAndServe(":8100", nil)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("index handler running")
	err := tpl.ExecuteTemplate(writer, "index.html", artists)
	if err != nil {
		log.Fatalln("Index Handler Error!", err)
	}
}
