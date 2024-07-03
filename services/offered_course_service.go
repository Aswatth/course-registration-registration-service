package services

import (
	"context"
	"course-registration-system/registration-service/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OfferedCourseService struct {
	database   MongoDatabase
	context    context.Context
	collection mongo.Collection
}

func (obj *OfferedCourseService) Init(database MongoDatabase) {
	obj.database = database
	obj.context, obj.collection = obj.database.GetCollection("offered_courses")
}

func (obj *OfferedCourseService) GetOfferedCourseByCRN(crn int) (*models.OfferedCourse, error) {
	var offered_course *models.OfferedCourse
	
	query := bson.D{bson.E{Key: "crn", Value: crn}}

	result := obj.collection.FindOne(obj.context, query)
	
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(&offered_course)

	return offered_course, err
}

func (obj *OfferedCourseService) GetAllOfferedCourseByProfessor(email_id string) ([]models.OfferedCourse, error) {
	var offered_course []models.OfferedCourse

	query := bson.D{bson.E{Key: "offered_by", Value: email_id}}

	result, _ := obj.collection.Find(obj.context, query)
	
	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.All(obj.context, &offered_course)

	return offered_course, err
}


func (obj *OfferedCourseService) CreateOfferedCourse(offered_course models.OfferedCourse) error {
	if offered_course.CRN == 0 || offered_course.Course_id == 0 || offered_course.OfferedBy == "" || offered_course.Days == nil || offered_course.Timings == nil {
		return errors.New("invalid data")
	}

	record, _ := obj.GetOfferedCourseByCRN(offered_course.CRN)

	if record != nil {
		return errors.New(fmt.Sprint(offered_course.CRN) + " already exists")
	}

	_, err := obj.collection.InsertOne(obj.context, offered_course)

	return err
}

func (obj *OfferedCourseService) UpdateOfferedCourse(offered_course models.OfferedCourse) error {
	if offered_course.Days == nil || offered_course.Timings == nil {
		return errors.New("invalid data")
	}

	filter := bson.D{bson.E{Key: "crn", Value: offered_course.CRN}}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "days", Value: offered_course.Days}, bson.E{Key: "timings", Value: offered_course.Timings}}}}

	result, err := obj.collection.UpdateOne(obj.context, filter, update)

	if result.MatchedCount == 0 {
		return errors.New("record not found")
	}

	return err
}

func (obj *OfferedCourseService) DeleteOfferedCourse(crn int) error {
	filter := bson.D{bson.E{Key: "crn", Value: crn}}

	result, err := obj.collection.DeleteOne(obj.context, filter)

	if result.DeletedCount == 0 {
		return errors.New("record not found")
	}

	return err
}
