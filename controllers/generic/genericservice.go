package generic

import "log"

//Returns all node values. Equivalent to Select * from xtablex;
func getAll(node string) (interface{}) {
	values, err := getAllValues(node)
	if err != nil {
		log.Fatal("GenericService getAll() ", err)
		//TODO: Create error response struct
		return err
	}
	return values
}