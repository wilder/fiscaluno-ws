package studentdao

import (
    "fiscaluno-ws/database"
    _"fiscaluno-ws/database/filter"
    "fiscaluno-ws/models/student"
)

var db = database.GetInstance()

func CreateNewStudent(student student.Student) (error){
    return db.Save(student, false, student.Institution)
}