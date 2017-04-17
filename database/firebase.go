package database

import (
	"fmt"
	"github.com/zabawaba99/firego"
	"reflect"
	"log"
)

type Firebase struct {
}

//creates a new node on firebase of the object passed as argument
//parameter: Any struct
func (fb Firebase) Save(object interface{}) {
	nodeName := getType(object)
	ref := firego.New("https://***REMOVED***/"+nodeName, nil)

	_, err := ref.Push(object)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("saving...")
	log.Print("saved new "+nodeName)
}

func (fb Firebase) Update(newValue, conditions[] interface{}) {
	fmt.Print("updating...")
}

func (fb Firebase) Find(node string, conditions[] interface{}) interface{} {
	fmt.Print("finding...")
	return map[string]string{"foo": "1", "bar": "2"}
}

func (fb Firebase) Delete(conditions[] interface{}) {
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