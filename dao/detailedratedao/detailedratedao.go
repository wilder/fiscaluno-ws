package detailedratedao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
    "fiscaluno-ws-bkp/models/rate/general"
)

var db = database.GetInstance()

func CreateDetailedRate(detailedRate interface{}) {
    db.Save(detailedRate, false)
}

func GetRatingsByUserId(conditions[] filter.Filter) (interface{}, error) {
    return db.Find("DetailedRate", conditions)
}

func NewGeneralRate(rate general.GeneralRate) {
    db.Save(rate, false)


    db.Save()
}