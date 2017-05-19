package institution

import (
    "fiscaluno-ws/models/institution"
    "fiscaluno-ws/dao/institutiondao"
    "fiscaluno-ws/utils"
)

func findInstitution(id string) (interface{}){
    val, err := GetInstitutionById(id)
    if err != nil {
	    //TODO create error struct
	    return err.Error()
    }
    return val
}

//TODO: Return map instead of single value
func getRate(id string) (float32) {
    inst, _ := GetInstitutionById(id)
    return inst.Rate
}

func GetInstitutionById(id string) (*institution.Institution, error){
    institution_interface, err := institutiondao.FindInstitutionById(id)

    inst := &institution.Institution{}
    if err == nil {
	err = utils.FirebaseToStruct(institution_interface, inst)

    }
    return inst, err
}

func CreateNewInstitution(id string, name string, rate float32) {
    institutiondao.CreateInstitution(institution.New(id, name, rate) )
}
