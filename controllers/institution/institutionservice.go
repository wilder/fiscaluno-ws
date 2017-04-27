package institution

import (
    "fiscaluno-ws/models/institution"
    "fiscaluno-ws/dao/institutiondao"
    "fiscaluno-ws/utils"
)

func findInstitution(id string) *institution.Institution{
    return getInstitutionById(id)
}

func getRate(id string) float32 {
    inst := getInstitutionById(id)
    return inst.Rate
}

func getInstitutionById(id string) *institution.Institution{
    institution_interface := institutiondao.FindInstitutionById(id)
    inst := &institution.Institution{}
    utils.FirebaseToStruct(institution_interface, inst)
    return inst
}

func CreateNewInstitution(id string, name string, rate float32) {
    institutiondao.CreateInstitution( institution.New(id, name, rate) )
}