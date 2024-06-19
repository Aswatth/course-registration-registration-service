package services

import (
	"context"
	"course-registration-system/registration-service/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegisteredCourseService struct {
	database   MongoDatabase
	context    context.Context
	collection mongo.Collection
}

func (obj *RegisteredCourseService) Init(database MongoDatabase) {
	obj.database = database
	obj.database.CreateCollection("registered_courses")
	obj.context, obj.collection = obj.database.GetCollection("registered_courses")
}

func (obj *RegisteredCourseService) GetRegisteredCourse(student_email_id string) (*models.RegisteredCourse, error) {
	filter := bson.D{bson.E{Key: "student_email_id", Value: student_email_id}}

	var registered_course *models.RegisteredCourse
	err := obj.collection.FindOne(obj.context, filter).Decode(&registered_course)

	return registered_course, err
}

func (obj *RegisteredCourseService) RegisterCourses(register_course models.RegisteredCourse) error {
	record, _ := obj.GetRegisteredCourse(register_course.Student_Email_id)

	if record != nil {
		return errors.New(register_course.Student_Email_id + " already exists")
	}

	_, err := obj.collection.InsertOne(obj.context, register_course)

	return err
}

func (obj *RegisteredCourseService) UpdateRegisteredCourses(register_course models.RegisteredCourse) error {
	filter := bson.D{bson.E{Key: "student_email_id", Value: register_course.Student_Email_id}}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "registered_courses_crns", Value: register_course.Registered_courses_crns}}}}

	_, err := obj.collection.UpdateOne(obj.context, filter, update)

	return err
}

func (obj *RegisteredCourseService) DeleteRegisteredCourses(student_email_id string) error {

	filter := bson.D{bson.E{Key: "student_email_id", Value: student_email_id}}

	_, err := obj.collection.DeleteOne(obj.context, filter)

	return err
}