package main

import (
	"log"
            "github.com/zabawaba99/firego"
	"fiscaluno-ws/models/rate/general"
            "fiscaluno-ws/models/institution"
	"fiscaluno-ws/models/student"
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

            institution := *institution.New( "f23fn9xsxx090nd3d", "Faculdade Impacta", 0.0)
	gRate := general.New("f134fqwefq4", "Muito Boaaa", "eu", 4.5, institution)
            student := *student.New("xxxxYYYYzzzz", "Subaru", institution, *gRate)

	//gRateJson, err := json.Marshal(gRate)

	pushedFirego, err := f.Push(gRate)
	if err != nil {
		log.Fatal(err)
	}

            pushedStudent, err := f.Push(student)
            if err != nil {
                log.Fatal(err)
            }

	var bar string
	if err := pushedFirego.Value(&bar); err != nil {
		log.Fatal(err)
	}

            var teste string
            if err := pushedStudent.Value(&teste); err != nil {
                log.Fatal(err)
            }

}
