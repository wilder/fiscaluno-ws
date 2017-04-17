package database

type DBcontract interface {
	Save(object interface{})
	Find(node string, conditions[] interface{}) interface{}
	Update(newValue, conditions[] interface{})
	Delete(conditions[] interface{})
}