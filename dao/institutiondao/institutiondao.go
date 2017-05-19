package institutiondao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
)

var db = database.GetInstance()

func FindInstitutionById (id string) (interface {}, error) {
    var filterList = [] filter.Filter{*filter.New("Id", id, "=")}
    return db.Find("Institution", filterList)
}

//TODO: make return error
func CreateInstitution(institution interface{}) {
    db.Save(institution, false)
}

func UpdateInstitution(institution interface{}) error{
    return db.Save(institution, true)
}
