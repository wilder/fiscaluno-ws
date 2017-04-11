package main

import (
	"github.com/zabawaba99/firego"
	"log"
	"fiscaluno-ws/models/rate/general"
	"fiscaluno-ws/models/institution"
	//"encoding/json"
)


/*
    apiKey: "AIzaSyBrjeCg34msto2W3mfokcoORxThGZCq_04",
    authDomain: "fiscaluno-4ec50.firebaseapp.com",
    databaseURL: "https://***REMOVED***",
    projectId: "fiscaluno-4ec50",
    storageBucket: "fiscaluno-4ec50.appspot.com",
    messagingSenderId: "284590836348"
 */


func main() {

	f := firego.New("https://***REMOVED***", nil)

	gRate := general.New("f134fqwefq4", "Muito Boaaa", 4.5, "", *institution.New( "f23fn9xsxx090nd3d,", "Faculdade Impacta", 0.0))
	//gRateJson, err := json.Marshal(gRate)

	pushedFirego, err := f.Push(gRate)
	if err != nil {
		log.Fatal(err)
	}

	var bar string
	if err := pushedFirego.Value(&bar); err != nil {
		log.Fatal(err)
	}

}
