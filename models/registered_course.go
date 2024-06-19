package models

type RegisteredCourse struct {
	Student_Email_id       string `json:"student_email_id" bson:"student_email_id"`
	Registered_course_crns []int  `json:"registered_course_crns" bson:"registered_course_crns"`
}
