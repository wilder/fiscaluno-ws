package main

import (
	"log"
	"github.com/zabawaba99/firego"
	//"fiscaluno-ws/models/rate/general"
	"fiscaluno-ws/models/institution"
	//"fiscaluno-ws/models/student"
	"fmt"
	"fiscaluno-ws/database"
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

	/*institution := *institution.New("f23fn9xsxx090nd3d", "Faculdade Impacta", 0.0)
	gRate := general.New("f134fqwefq4", "Muito Boaaa", "eu", 4.5, institution)
	student := *student.New("xxxxYYYYzzzz", "Subaru", institution, *gRate)*/

	/*_, err := f.Push(gRate)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Push(student)
	if err != nil {
		log.Fatal(err)
	}*/

	//Testing if Firebase implements DBcontract
	var db database.DBcontract = database.Firebase{}
	db.Save(*institution.New("a","a",4.4))

	var v map[string]interface{}
	if err := f.Value(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", v)

}
