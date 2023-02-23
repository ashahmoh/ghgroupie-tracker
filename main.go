package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/404", Err404)
	http.HandleFunc("/400", Err400)
	http.HandleFunc("/500", Err500)
	fmt.Println("listening...")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("index handler running")
	if request.URL.Path != "/" {
		tpl.ExecuteTemplate(writer, "404.html", nil)
		return
	}
	tpl.ExecuteTemplate(writer, "index.html", artists)
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	artistName := request.FormValue("artist")
	lower := strings.ToLower(artistName)
	for i, artist := range artists {
		artistlower := strings.ToLower(artist.Name)
		for _, v := range artist.Members {
			var mem []string
			mem = append(mem, v)
			strMem := strings.Join(mem, ", ")
			strMemLow := strings.ToLower(strMem)
			creaDate := artist.CreationDate
			conDate := artist.FirstAlbum
			if artistlower == lower || strMemLow == lower || strconv.Itoa(creaDate) == lower || conDate == lower {
				tpl.ExecuteTemplate(writer, "artist.html", artists[i])
				return

			}
		}
	}

	http.Error(writer, "500 Internal Server Error", 500)
}

func Err404(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "404.html", nil)
}

func Err500(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "500.html", nil)
}

func Err400(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "400.html", nil)
}
