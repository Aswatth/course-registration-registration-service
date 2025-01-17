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

func (obj *RegisteredCourseService) GetRegisteredCourseByCRN(crn int) ([]models.RegisteredCourse, error) {
	filter := bson.D{bson.E{ Key: "registered_course_crns", Value: crn}}

	var registered_course_list []models.RegisteredCourse
	result, err := obj.collection.Find(obj.context, filter)

	if (err != nil) {
		return nil, err
	} else if (result.Err() != nil){
		return nil, result.Err()
	} else {
		result.All(obj.context, &registered_course_list)

		return registered_course_list, nil
	}
}

func (obj *RegisteredCourseService) GetRegisteredCourseByEmailId(student_email_id string) (*models.RegisteredCourse, error) {
	filter := bson.D{bson.E{Key: "student_email_id", Value: student_email_id}}

	var registered_course *models.RegisteredCourse
	err := obj.collection.FindOne(obj.context, filter).Decode(&registered_course)

	return registered_course, err
}

func (obj *RegisteredCourseService) RegisterCourses(register_course models.RegisteredCourse) error {
	if register_course.Student_Email_id == "" || register_course.Registered_course_crns == nil {
		return errors.New("invalid data")
	}

	record, _ := obj.GetRegisteredCourseByEmailId(register_course.Student_Email_id)

	if record != nil {
		obj.UpdateRegisteredCourses(register_course)
		// return errors.New(register_course.Student_Email_id + " already exists")
	} else {
		_, err := obj.collection.InsertOne(obj.context, register_course)
	
		return err
	}
	
	return nil
}

func (obj *RegisteredCourseService) UpdateRegisteredCourses(register_course models.RegisteredCourse) error {

	if register_course.Registered_course_crns == nil {
		return errors.New("invalid data")
	}

	filter := bson.D{bson.E{Key: "student_email_id", Value: register_course.Student_Email_id}}

	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "registered_course_crns", Value: register_course.Registered_course_crns}}}}

	result, err := obj.collection.UpdateOne(obj.context, filter, update)

	if result.MatchedCount == 0 {
		return errors.New("record not found")
	}

	return err
}

func (obj *RegisteredCourseService) DeleteRegisteredCourses(student_email_id string) error {

	filter := bson.D{bson.E{Key: "student_email_id", Value: student_email_id}}

	result, err := obj.collection.DeleteOne(obj.context, filter)

	if result.DeletedCount == 0 {
		return errors.New("record not found")
	}

	return err
}
