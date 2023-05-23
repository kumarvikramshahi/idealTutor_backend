package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Skills struct {
	MaxLevelClassTeach int32    `json:"max_level_class_teach,omitempty" bson:"max_level_class_teach,omitempty"`
	SubjectTeach       []string `json:"subject_teach,omitempty" bson:"subject_teach,omitempty"`
}

type Comments struct {
}

type TeacherProfile struct {
	ID                  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName           string             `json:"first_Name,omitempty" bson:"first_Name,omitempty"`
	LastName            string             `json:"last_Name,omitempty" bson:"last_Name,omitempty"`
	Email               string             `json:"email,omitempty" bson:"email,omitempty"`
	MobileNum           string             `json:"mobileNum,omitempty" bson:"mobileNum,omitempty"`
	Password            string             `json:"password,omitempty" bson:"password,omitempty"`
	ProfilePic          string             `json:"profile_Pic,omitempty" bson:"profile_Pic,omitempty"`
	Experience          int                `json:"experience,omitempty" bson:"experience,omitempty"`
	NumOfStudentTeach   int                `json:"num_of_student_teach,omitempty" bson:"num_of_student_teach,omitempty"`
	Skills              Skills             `json:"skills,omitempty" bson:"skills,omitempty"`
	Achievement         string             `json:"achievement,omitempty" bson:"achievement,omitempty"`
	Comments            []Comments         `json:"comments,omitempty" bson:"comments,omitempty"`
	RatingStar          int                `json:"rating_star,omitempty" bson:"rating_star,omitempty"`
	NumberOfRatingGiven int                `json:"num_of_rating_given,omitempty" bson:"num_of_rating_given,omitempty"`
	BookPrefrenceList   []string           `json:"book_prefrence_list,omitempty" bson:"book_prefrence_list,omitempty"`
	JoiningDate         primitive.DateTime `json:"joining_date,omitempty" bson:"joining_date,omitempty"`
}
