package controllers

import (
	"course-registration-system/registration-service/models"
	"course-registration-system/registration-service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfferedCourseController struct {
	service services.OfferedCourseService
}

func (obj *OfferedCourseController) Init(service services.OfferedCourseService) {
	obj.service = service
}

func (obj *OfferedCourseController) GetOfferedCourse(context *gin.Context) {
	crn, _ := strconv.Atoi(context.Param("crn"))

	result, err := obj.service.GetOfferedCourse(crn)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, result)
}

func (obj *OfferedCourseController) AddOfferedCourse(context *gin.Context) {
	var offered_course models.OfferedCourse

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&offered_course); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := obj.service.CreateOfferedCourse(offered_course)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
}

func (obj *OfferedCourseController) RemoveOfferedCourse(context *gin.Context) {
	crn, _ := strconv.Atoi(context.Param("crn"))

	err := obj.service.DeleteOfferedCourse(crn)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, nil)
}

func (obj *OfferedCourseController) UpdateOffereddCourse(context *gin.Context) {
	crn, _ := strconv.Atoi(context.Param("crn"))

	var offered_course models.OfferedCourse

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&offered_course); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	offered_course.CRN = crn

	err := obj.service.UpdateOfferedCourse(offered_course)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, nil)
}

func (obj *OfferedCourseController) RegisterRoutes(rg *gin.RouterGroup) {
	offered_course_routes := rg.Group("offered_course")

	offered_course_routes.POST("", obj.AddOfferedCourse)
	offered_course_routes.GET("/:crn", obj.GetOfferedCourse)
	offered_course_routes.DELETE("/:crn", obj.RemoveOfferedCourse)
	offered_course_routes.PUT("/:crn", obj.UpdateOffereddCourse)
}
