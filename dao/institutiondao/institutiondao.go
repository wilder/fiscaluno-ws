package institutiondao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
)

var db database.DBcontract = database.GetInstance()

func FindInstitutionById (id string) (interface {}, error) {
    var filterList = [] filter.Filter{*filter.New("Id", id, "=")}
    return db.Find("Institution", filterList)
}
