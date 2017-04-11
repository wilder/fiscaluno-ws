package student

import (
    "fiscaluno-ws/models/institution"
    "fiscaluno-ws/models/rate/general"
)

type Student struct {
            Id string
            Name string
            Institution institution.Institution
            GeneralRate general.GeneralRate
}

// Student Constructor
func New(Id string, Name string, Institution institution.Institution, GeneralRate general.GeneralRate) *Student {
            return &Student{
                Id: Id,
                Name: Name,
                Institution: Institution,
                GeneralRate: GeneralRate,
            }
}