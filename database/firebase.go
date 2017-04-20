package database

import (
	"sync"
	"fmt"
	"github.com/zabawaba99/firego"
	"reflect"
	"log"
	"fiscaluno-ws/database/filter"
	"fiscaluno-ws/config"
)

type firebase struct {
}

var instance firebase
var once sync.Once

func GetInstance() firebase {
	once.Do(func() {
		instance = firebase{}
	})
	return instance
}

//creates a new node on firebase of the object passed as argument
//parameter: Any struct
func (fb firebase) Save(object interface{}) {
	nodeName := getType(object)
	//TODO: improve
	ref := firego.New("https://"+config.FirebaseUrl()+"/"+nodeName, nil)

	_, err := ref.Push(object)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("saving...")
	log.Print("saved new "+nodeName)
}

func (fb firebase) Update(newValue, conditions[] filter.Filter) {
	log.Print("updating...")
}

//Creates a reference with the passed node name
//iterates over the filter list and creates the
//proper node reference
func (fb firebase) Find(nodeName string, conditions[] filter.Filter) interface{} {
	log.Print("finding..."+nodeName+"\n")
	//TODO: improve
	ref := firego.New("https://"+config.FirebaseUrl()+"/"+nodeName, nil)

	var v map[string]interface{}
	for _, element := range conditions {
		fmt.Printf("operation: "+element.Operation)
		if(element.Operation == "="){
			//When the operation is =, the endAt and StartAt are equal
			//Also, only one element will be retrieved, hence LimitToFirst(1)
			//todo: find a way to add the Value() out of the loop
			if err := ref.StartAt(element.Value).EndAt(element.Value).LimitToFirst(1).OrderBy(element.Name).Value(&v); err != nil {
				log.Print("error")
				log.Fatal(err)
			}
		}
		//todo implement other filters
	}

	return v
}

func (fb firebase) Delete(conditions[] interface{}) {
	fmt.Print("deleting...")
}

/**
   gets the node name from a type passed as argument
   used to get the correct firebase refenrence
 */
func getType(myvar interface{}) string {
	valueOf := reflect.ValueOf(myvar)
	if valueOf.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(valueOf).Type().Name()
	} else {
		return valueOf.Type().Name()
	}
}