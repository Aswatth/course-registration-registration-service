package controllers

import (
	"course-registration-system/registration-service/models"
	"course-registration-system/registration-service/services"
	"encoding/json"
	"io"
	"net/http"

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
		context.AbortWithError(http.StatusBadRequest, err)
	}

	err := obj.service.RegisterCourses(register_course)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.Status(http.StatusOK)
}
func (obj *RegisteredCourseController) GetRegisteredCourses(context *gin.Context) {
	student_email_id := context.Param("student_email_id")
	result, err := obj.service.GetRegisteredCourse(student_email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, result)
}
func (obj *RegisteredCourseController) DeleteRegisteredCourses(context *gin.Context) {
	student_email_id := context.Param("student_email_id")

	err := obj.service.DeleteRegisteredCourses(student_email_id)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	context.Status(http.StatusOK)
}

func (obj *RegisteredCourseController) UpdateRegisteredCourses(context *gin.Context) {
	student_email_id := context.Param("student_email_id")

	req_body, _ := io.ReadAll(context.Request.Body)

	updated_courses := make(map[string][]int)

	json.Unmarshal(req_body, &updated_courses)

	err := obj.service.UpdateRegisteredCourses(models.RegisteredCourse{Student_Email_id: student_email_id, Registered_courses_crns: updated_courses["registered_course_crns"]})

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	context.Status(http.StatusOK)
}

func (obj *RegisteredCourseController) RegisterRoutes(rg *gin.RouterGroup) {
	offered_course_routes := rg.Group("register_course")

	offered_course_routes.POST("", obj.RegisterCourses)
	offered_course_routes.GET("/:student_email_id", obj.GetRegisteredCourses)
	offered_course_routes.DELETE("/:student_email_id", obj.DeleteRegisteredCourses)
	offered_course_routes.PUT("/:student_email_id", obj.UpdateRegisteredCourses)
}
