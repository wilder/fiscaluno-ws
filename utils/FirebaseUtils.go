package utils

import (
	"github.com/mitchellh/mapstructure"
	"errors"
)

//removes firebase generated id
//params: firebaseData the data from firebase
//	  structTypeReference the reference of the struct type
func FirebaseToStruct(firebaseData interface{}, structTypeReference interface{}) error {

	var err error
	//removing the firebase node
	node :=  firebaseData.(map[string] interface{})

	//gets the value of the firebase data Ex: "-asdm24234mlka"
	keys := make([]interface{}, 1)
	i := 0

	//iterates over the values of the firebase response
	//and sets the response["-asdm24234mlka"] data to keys[0]
	for k := range node {
		keys[i] = node[k]
		i++
	}

	//Converts the firebase data to the struct
	firebaseVal := keys[0]

	if firebaseVal != nil {
		err = mapstructure.Decode(firebaseVal.(map[string] interface{}), structTypeReference)
	} else {
		//TODO: Maybe mobe to error generator class?
		err = errors.New("No Data Found")
	}

	return err

}
