package institutiondao

import (
    "strconv"
    "fiscaluno-ws/database"
    "fiscaluno-ws/database/filter"
    "fiscaluno-ws/models/institution"
)

var db database.DBcontract = database.GetInstance()

func FindInstitutionById (id string) *institution.Institution {
    var filterList = [] filter.Filter{*filter.New("Id", id, "=")}
    institution_node := db.Find("Institution", filterList).(map[string] interface{})
    
    for key := range institution_node {
        institution_node = institution_node[key].(map[string] interface{})
    }

    institution_id := strconv.FormatFloat( institution_node["Id"].(float64), 'f', 1, 64 )
    return  institution.New(  institution_id, institution_node["Name"].(string), float32( institution_node["Rate"].(float64) ) )
}
