package models

type Institution struct {
    Id string
    Name string
    Rate float32
}

func FindInstitution(id string) *Institution {
    institution := &Institution{Id: id, Name: "Faculdade Impacta", Rate: 5.2}
    return institution
}

func (institution *Institution) GetId() *string {
    return &institution.Id
}

func (institution *Institution) GetName() *string {
    return &institution.Name
}

func (institution *Institution) GetRate() *float32 {
    return &institution.Rate
}