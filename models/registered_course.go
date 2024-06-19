package models

type RegisteredCourse struct {
	Student_Email_id        string `json:"student_email_id" bson:"student_email_id"`
	Registered_courses_crns []int  `json:"registered_courses_crns" bson:"registered_courses_crns"`
}
