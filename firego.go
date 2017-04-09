package main

import (
	"github.com/zabawaba99/firego"
	"log"
)


/*
    apiKey: "AIzaSyBrjeCg34msto2W3mfokcoORxThGZCq_04",
    authDomain: "fiscaluno-4ec50.firebaseapp.com",
    databaseURL: "https://***REMOVED***",
    projectId: "fiscaluno-4ec50",
    storageBucket: "fiscaluno-4ec50.appspot.com",
    messagingSenderId: "284590836348"
 */

type PersonName struct {
	First string
	Last string
}

type Person struct {
	Name PersonName
}

func main() {

	f := firego.New("https://***REMOVED***", nil)
	v := map[string]string{"foo":"bar"}
	if err := f.Set(v); err != nil {
		log.Fatal(err)
	}

}