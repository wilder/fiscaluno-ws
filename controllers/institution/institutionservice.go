package institution

import (
    "strconv"
    "fiscaluno-ws/models/institution"
    "fiscaluno-ws/dao/institutiondao"
)

func findInstitution(id string) *institution.Institution{
    institution_interface := institutiondao.FindInstitutionById(id)

    //removing the firebase node
    institution_node :=  institution_interface.(map[string] interface{})
    for key := range institution_node {
        institution_node = institution_node[key].(map[string] interface{})
    }

    institution_id := strconv.FormatFloat( institution_node["Id"].(float64), 'f', 1, 64 )
    return institution.New( institution_id, institution_node["Name"].(string), float32( institution_node["Rate"].(float64) ) )
}