package student

import (
    "log"
    "github.com/emicklei/go-restful"
    "fiscaluno-ws/models/student"
)

func CreateNewStudent(request *restful.Request, response *restful.Response) {
    student := new( student.Student )
    err := request.ReadEntity(&student)

    if err == nil {
        err = SaveNewStudent(*student)
        if err != nil {
            response.WriteEntity(err)
        } else {
            response.WriteEntity(response.StatusCode())
        }
    } else{
        log.Panic(err)
    }
}