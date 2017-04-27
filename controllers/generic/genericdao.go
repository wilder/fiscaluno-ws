package controllers

import "fiscaluno-ws/database"

func getAllValues(node string) interface{} {
	var db database.DBcontract = database.GetInstance()
	return db.Find(node, nil)
}
