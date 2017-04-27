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
	fmt.Println("https://"+config.FirebaseUrl()+"/"+nodeName)
	var err error
	var v map[string]interface{}
	if conditions != nil{
		var mainCondition = conditions[0]
		fmt.Printf("maincondition: ", mainCondition)
		if(conditions[0].Operation == "="){
			//When the operation is =, the endAt and StartAt are equal
			//Also, only one element will be retrieved, hence LimitToFirst(1)
			err = ref.StartAt(mainCondition.Value).EndAt(mainCondition.Value).LimitToFirst(1).OrderBy(mainCondition.Name).Value(&v);
		}
	}else{
		err = ref.Value(&v)
		fmt.Println("v ", v)
	}

	if err != nil {
		log.Panic(err)
		return err
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