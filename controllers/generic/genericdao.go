package generic

import "fiscaluno-ws/database"

func getAllValues(node string) interface{} {
	var db = database.GetInstance()
	return db.Find(node, nil)
}
