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

type location struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type date struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

// put structs into variables for unmarshelling.
// Unmarshalling takes data from the struct and puts it into the variable
var artists artist
var loc location
var datesvar date

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

	//2. location url get, put json into variable and unmarshall

	locationResp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		print("Can't get Artist URL!", err)
	}
	defer locationResp.Body.Close()
	locbody, err := ioutil.ReadAll(locationResp.Body)
	if err != nil {
		print("Error copying location json into body", err)
	}

	err3 := json.Unmarshal(locbody, &loc)
	if err3 != nil {
		log.Fatalln("error unmarhsalling location", err3)
	}

	//3. dates url get, put json into variable and unmarshall

	datesResp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		print("Can't get Artist URL!", err)
	}
	defer locationResp.Body.Close()
	datesbody, err := ioutil.ReadAll(datesResp.Body)
	if err != nil {
		print("Error copying dates json into body", err)
	}

	err4 := json.Unmarshal(datesbody, &datesvar)
	if err4 != nil {
		log.Fatalln("error unmarshalling dates", err4)
	}

	http.Handle("/Users/ashamohamed/Documents/GitHub/ghgroupie-tracker/assets/", http.StripPrefix("/Users/ashamohamed/Documents/GitHub/ghgroupie-tracker/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/loc", locHandler)
	http.HandleFunc("/dates", datesHandler)
	http.HandleFunc("/locart", locArtHandler)
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

func locHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Location handler running")
	err := tpl.ExecuteTemplate(writer, "loc.html", loc)
	if err != nil {
		log.Fatalln("Location Handler Error!", err)
	}
}

func datesHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Dates handler running")
	err := tpl.ExecuteTemplate(writer, "dates.html", datesvar)
	if err != nil {
		log.Fatalln("Dates Handler Error!", err)
	}

}

func locArtHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Locart handler running")
	err := tpl.ExecuteTemplate(writer, "locart.html", artists)
	if err != nil {
		log.Fatalln("Dates Handler Error!", err)
	}

}

// func GetDates(url string) []string {
// 	var dates artist
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	err = json.NewDecoder(resp.Body).Decode(&dates)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return
// }
