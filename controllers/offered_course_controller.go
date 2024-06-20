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
	crn, err := strconv.Atoi(context.Param("crn"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		result, err := obj.service.GetOfferedCourse(crn)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else if result != nil {
			context.JSON(http.StatusOK, result)
		}
	}
}

func (obj *OfferedCourseController) AddOfferedCourse(context *gin.Context) {
	var offered_course models.OfferedCourse

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&offered_course); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		err := obj.service.CreateOfferedCourse(offered_course)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}

}

func (obj *OfferedCourseController) RemoveOfferedCourse(context *gin.Context) {
	crn, err := strconv.Atoi(context.Param("crn"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		err := obj.service.DeleteOfferedCourse(crn)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}
}

func (obj *OfferedCourseController) UpdateOffereddCourse(context *gin.Context) {
	crn, err := strconv.Atoi(context.Param("crn"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var offered_course models.OfferedCourse

		//Check if given JSON is valid
		if err := context.ShouldBindJSON(&offered_course); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		} else {
			offered_course.CRN = crn

			err := obj.service.UpdateOfferedCourse(offered_course)

			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			}

			context.Status(http.StatusOK)
		}
	}
}

func (obj *OfferedCourseController) RegisterRoutes(rg *gin.RouterGroup) {
	offered_course_routes := rg.Group("offered_course")

	offered_course_routes.POST("", obj.AddOfferedCourse)
	offered_course_routes.GET("/:crn", obj.GetOfferedCourse)
	offered_course_routes.DELETE("/:crn", obj.RemoveOfferedCourse)
	offered_course_routes.PUT("/:crn", obj.UpdateOffereddCourse)
}
