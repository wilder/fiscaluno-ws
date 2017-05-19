package student

import (
    "fiscaluno-ws/models/institution"
)

type Student struct {
            Id string
            Name string
            Institution string
            Course string
}


// Student Constructor
func New(Id string, Name string, Institution institution.Institution, course string) *Student {
            return &Student{
                Id: Id,
                Name: Name,
                Institution: Institution,
                Course: course,
            }
}