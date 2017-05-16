package detailedratedao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
    "fiscaluno-ws/models/rate/general"
    "fiscaluno-ws/models/rate/detailed"
)

var db = database.GetInstance()

func CreateDetailedRate(detailedRate specific.DetailedRate) (error){
    return db.Save(detailedRate, false, detailedRate.Institution)
}

func GetRatingsByUserId(conditions[] filter.Filter) (interface{}, error) {
    return db.Find("DetailedRate", conditions)
}

func NewGeneralRate(rate general.GeneralRate) (error){
    return db.Save(rate, false, rate.Institution)
}