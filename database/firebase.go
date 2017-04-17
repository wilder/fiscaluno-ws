package database

import (
	"fmt"
)

type Firebase struct {
}

func (fb Firebase) Save(object interface{}) {
	fmt.Print("saving...")
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