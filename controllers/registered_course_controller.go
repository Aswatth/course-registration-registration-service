package controllers

import (
	"course-registration-system/registration-service/models"
	"course-registration-system/registration-service/services"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RegisteredCourseController struct {
	service services.RegisteredCourseService
}

func (obj *RegisteredCourseController) Init(service services.RegisteredCourseService) {
	obj.service = service
}

func (obj *RegisteredCourseController) RegisterCourses(context *gin.Context) {
	var register_course models.RegisteredCourse

	if err := context.ShouldBindJSON(&register_course); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		err := obj.service.RegisterCourses(register_course)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}
}
func (obj *RegisteredCourseController) getRegisteredCoursesByCRN(context *gin.Context) {
	
	crn, err := strconv.Atoi(context.Query("crn"))

	if (err != nil) {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		result, err := obj.service.GetRegisteredCourseByCRN(crn)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			context.JSON(http.StatusOK, result)
		}
	}	
}

func (obj *RegisteredCourseController) getRegisteredCoursesByEmailId(context *gin.Context) {
	student_email_id := context.Query("email_id")
	result, err := obj.service.GetRegisteredCourseByEmailId(student_email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.JSON(http.StatusOK, result)
	}
}

func (obj *RegisteredCourseController) GetRegisteredCourses(context *gin.Context) {
	if(context.Query("email_id") != "" ){
		obj.getRegisteredCoursesByEmailId(context)
		return
	} else if(context.Query("crn") != "") {
		obj.getRegisteredCoursesByCRN(context)
		return
	}
}

func (obj *RegisteredCourseController) DeleteRegisteredCourses(context *gin.Context) {
	student_email_id := context.Query("email_id")

	err := obj.service.DeleteRegisteredCourses(student_email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
	} else {
		context.Status(http.StatusOK)
	}
}

func (obj *RegisteredCourseController) UpdateRegisteredCourses(context *gin.Context) {
	student_email_id := context.Query("email_id")

	req_body, err := io.ReadAll(context.Request.Body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		updated_courses := make(map[string][]int)

		err := json.Unmarshal(req_body, &updated_courses)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			err := obj.service.UpdateRegisteredCourses(models.RegisteredCourse{Student_Email_id: student_email_id, Registered_course_crns: updated_courses["registered_course_crns"]})

			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
			} else {
				context.Status(http.StatusOK)
			}
		}
	}
}

func (obj *RegisteredCourseController) RegisterRoutes(rg *gin.RouterGroup) {
	offered_course_routes := rg.Group("register_course")

	offered_course_routes.POST("", obj.RegisterCourses)
	offered_course_routes.GET("", obj.GetRegisteredCourses)
	offered_course_routes.DELETE("", obj.DeleteRegisteredCourses)
	offered_course_routes.PUT("", obj.UpdateRegisteredCourses)
}
