package institution

import (
    "fiscaluno-ws/models/institution"
    "fiscaluno-ws/dao/institutiondao"
    "fiscaluno-ws/utils"
)

func findInstitution(id string) *institution.Institution{
    institution_interface := institutiondao.FindInstitutionById(id)
    inst := &institution.Institution{}
    utils.FirebaseToStruct(institution_interface, inst)
    return inst
}