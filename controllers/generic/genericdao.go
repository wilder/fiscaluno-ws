package generic

import "fiscaluno-ws/database"

func getAllValues(node string) (interface{}, error) {
	var db = database.GetInstance()
	return db.Find(node, nil)
}
