package models

type OfferedCourse struct {
	Course_id int       `json:"course_id" bson:"course_id"`
	CRN       int       `json:"crn" bson:"crn"`
	OfferedBy string    `json:"offered_by" bson:"offered_by"`
	DayTime   []DayTime `json:"day_time" bson:"day_time"`
}
