package filter

//struct used as params in 'where' clauses
type Filter struct {
	Name string //name of the field
	Value string //value of the field
	Operation string // =, <=, >=
}

func New(name string, value string, operation string) *Filter {
	return &Filter{Name:name, Value:value, Operation:operation}
}
