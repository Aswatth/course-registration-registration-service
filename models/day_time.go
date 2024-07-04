package models

type DayTime struct {
	Day       string `json:"day" bson:"day"`
	StartTime string `json:"start_time" bson:"start_time"`
	EndTime   string `json:"end_time" bson:"end_time"`
}
