package generic

//Returns all node values. Equivalent to Select * from xtablex;
func getAll(node string) interface{} {
	values := getAllValues(node)
	return values
}