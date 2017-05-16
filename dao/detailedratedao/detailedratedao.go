package detailedratedao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
    "fiscaluno-ws/models/rate/general"
)

var db = database.GetInstance()

func CreateDetailedRate(detailedRate interface{}) {
    db.Save(detailedRate, false)
}

func GetRatingsByUserId(conditions[] filter.Filter) (interface{}, error) {
    return db.Find("DetailedRate", conditions)
}

func NewGeneralRate(rate general.GeneralRate) (error){
    return db.Save(rate, false, rate.Institution)
}