package institutiondao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
    //"fiscaluno-ws/models/institution"
)

var db database.DBcontract = database.GetInstance()

func FindInstitutionById (id string) interface{} {
    var filterList = []filter.Filter{*filter.New("Id", id, "=")}
    return db.Find("Institution", filterList)
}
