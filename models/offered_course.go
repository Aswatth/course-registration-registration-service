package models

type OfferedCourse struct {
	Course_id int      `json:"course_id" bson:"course_id"`
	CRN       int      `json:"crn" bson:"crn"`
	OfferedBy string   `json:"offered_by" bson:"offered_by"`
	Days      []string `json:"days" bson:"days"`
	Timings   []string `json:"timings" bson:"timings"`
}
