package detailedratedao

import (
    "fiscaluno-ws/database"
)

var db = database.GetInstance()

func CreateDetailedRate(detailedRate interface{}) {
    db.Save(detailedRate, false)
}
