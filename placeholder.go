package main

import (
	"fmt"
	"log"
	"net/http"
)

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

var loc location
var datesvar date

//////////////GET DATE FUNCTION///////////////

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

/////////////////////////////////////////////////////////////////////////////////////

//////// DATE AND LOCATIONS UNMARSHALLING AND HTTP.GET//////////

//2. location url get, put json into variable and unmarshall

// locationResp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
// if err != nil {
// 	print("Can't get Artist URL!", err)
// }
// defer locationResp.Body.Close()
// locbody, err := ioutil.ReadAll(locationResp.Body)
// if err != nil {
// 	print("Error copying location json into body", err)
// }

// err3 := json.Unmarshal(locbody, &loc)
// if err3 != nil {
// 	log.Fatalln("error unmarhsalling location", err3)
// }

// //3. dates url get, put json into variable and unmarshall

// datesResp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
// if err != nil {
// 	print("Can't get Artist URL!", err)
// }
// defer locationResp.Body.Close()
// datesbody, err := ioutil.ReadAll(datesResp.Body)
// if err != nil {
// 	print("Error copying dates json into body", err)
// }

// err4 := json.Unmarshal(datesbody, &datesvar)
// if err4 != nil {
// 	log.Fatalln("error unmarshalling dates", err4)
// }

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
