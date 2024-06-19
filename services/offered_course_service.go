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

func (obj *OfferedCourseService) GetOfferedCourse(crn int) (*models.OfferedCourse, error) {
	var offered_course *models.OfferedCourse

	query := bson.D{bson.E{Key: "crn", Value: crn}}

	err := obj.collection.FindOne(obj.context, query).Decode(&offered_course)

	return offered_course, err
}

func (obj *OfferedCourseService) CreateOfferedCourse(offered_course models.OfferedCourse) error {
	record, _ := obj.GetOfferedCourse(offered_course.CRN)

	fmt.Println(record)

	if record != nil {
		return errors.New(string(rune(offered_course.CRN)) + " already exists")
	}

	_, err := obj.collection.InsertOne(obj.context, offered_course)

	return err
}

func (obj *OfferedCourseService) UpdateOfferedCourse(offered_course models.OfferedCourse) error {
	filter := bson.D{bson.E{Key: "crn", Value: offered_course.CRN}}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "days", Value: offered_course.Days}, bson.E{Key: "timings", Value: offered_course.Timings}}}}

	_, err := obj.collection.UpdateOne(obj.context, filter, update)

	return err
}

func (obj *OfferedCourseService) DeleteOfferedCourse(crn int) error {
	filter := bson.D{bson.E{Key: "crn", Value: crn}}

	_, err := obj.collection.DeleteOne(obj.context, filter)

	return err
}
