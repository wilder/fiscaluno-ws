package student

import (
    "log"
    "fiscaluno-ws/models/student"
    "fiscaluno-ws/dao/studentdao"
)

func SaveNewStudent(student student.Student) (error) {
    err := studentdao.CreateNewStudent(student)
    log.Print("new student registered: ", student)

    return err
}

