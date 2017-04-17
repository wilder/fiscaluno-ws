package filter

//struct used as params in 'where' clauses
type filter struct {
	name string
	value interface{}
}

func New(name string, value interface{}) *filter {
	return &filter{name:name, value:value}
}
