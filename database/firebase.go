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