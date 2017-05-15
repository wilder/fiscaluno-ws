package detailedratedao

import (
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
)

var db = database.GetInstance()

func CreateDetailedRate(detailedRate interface{}) {
    db.Save(detailedRate, false)
}

func GetRatingsByUserId(conditions[] filter.Filter) (interface{}, error) {
    return db.Find("DetailedRate", conditions)
}