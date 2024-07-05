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
	if(context.Query("email_id") != "") {
		obj.getOfferedCourseByProfessor(context)
		return
	} else if(context.Query("crn") != "" ) {
		obj.getOfferedCourseByCRN(context)
		return
	}

	result, err := obj.service.GetAllOfferedCourses()

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else if result != nil {
			context.JSON(http.StatusOK, result)
		}
}

func (obj *OfferedCourseController) getOfferedCourseByCRN(context *gin.Context) {
	crn, err := strconv.Atoi(context.Query("crn"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		result, err := obj.service.GetOfferedCourseByCRN(crn)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else if result != nil {
			context.JSON(http.StatusOK, result)
		}
	}
}

func (obj *OfferedCourseController) getOfferedCourseByProfessor(context *gin.Context) {
	
	result, err := obj.service.GetAllOfferedCourseByProfessor(context.Query("email_id"))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else if result != nil {
			context.JSON(http.StatusOK, result)
		}
}

func (obj *OfferedCourseController) AddOfferedCourse(context *gin.Context) {
	var offered_course models.OfferedCourse

	//Check if given JSON is valid
	if err := context.ShouldBindJSON(&offered_course); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		err := obj.service.CreateOfferedCourse(offered_course)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}

}

func (obj *OfferedCourseController) RemoveOfferedCourse(context *gin.Context) {
	crn, err := strconv.Atoi(context.Param("crn"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		err := obj.service.DeleteOfferedCourse(crn)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"response": err.Error()})
		} else {
			context.Status(http.StatusOK)
		}
	}
}

func (obj *OfferedCourseController) UpdateOffereddCourse(context *gin.Context) {
	crn, err := strconv.Atoi(context.Param("crn"))

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"response": err.Error()})
	} else {
		var offered_course models.OfferedCourse

		//Check if given JSON is valid
		if err := context.ShouldBindJSON(&offered_course); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
		} else {
			offered_course.CRN = crn

			err := obj.service.UpdateOfferedCourse(offered_course)

			if err != nil {
				context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
			}

			context.Status(http.StatusOK)
		}
	}
}

func (obj *OfferedCourseController) RegisterRoutes(rg *gin.RouterGroup) {
	offered_course_routes := rg.Group("offered_course")

	offered_course_routes.POST("", obj.AddOfferedCourse)
	offered_course_routes.GET("", obj.GetOfferedCourse)
	offered_course_routes.DELETE("/:crn", obj.RemoveOfferedCourse)
	offered_course_routes.PUT("/:crn", obj.UpdateOffereddCourse)
}
