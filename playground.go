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


func main() {

	f := firego.New("https://***REMOVED***", nil)
	v := "bar"
	pushedFirego, err := f.Push(v)
	if err != nil {
		log.Fatal(err)
	}

	var bar string
	if err := pushedFirego.Value(&bar); err != nil {
		log.Fatal(err)
	}

}