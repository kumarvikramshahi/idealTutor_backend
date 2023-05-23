// first_Name: {
// 	type: String,
// 	required: true,
// 	maxlength: 15,
// 	minlength: 2,
// 	trim: true
// },
// last_Name: { type: String, required: true, maxlength: 15, minlength: 2, trim: true },
// email: { type: String, required: true, unique: true, trim: true },
// mobileNo: { type: Number, required: true, unique: true, },
// password: { type: String, required: true },
// profile_Pic: { type: String },
// educational_details: {
// 	education_level: { type: String },
// 	degree: { type: String },
// },
// address: {
// 	cur_addr: { type: String, maxlength: 100, },
// 	pincode: { type: Number },
// 	city: { type: Number, },
// 	state: {
// 		type: [Number, "state field must have state code "],
// 	},
// },
// teach_sessions: [{
// 	type: String,
// 	trim: true,
// 	minlength: [24, "Invalid user ID"],
// 	maxlength: [24, "Invalid user ID"]
// }],
// joining_date: { type: Date, required: true, default: Date.now },
// });

package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type address struct {
	CurAddr string `json:"cur_addr,omitempty" bson:"cur_addr,omitempty"`
	Pincode int    `json:"pincode,omitempty" bson:"pincode,omitempty"`
	City    int16  `json:"city,omitempty" bson:"city,omitempty"`
	State   int16  `json:"state,omitempty" bson:"state,omitempty"`
}

type educationalDetails struct {
	EducationLevel string `json:"education_level,omitempty" bson:"education_level,omitempty"`
	Degree         string `json:"degree,omitempty" bson:"degree,omitempty"`
}

type StudentrProfile struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName          string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName           string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email              string             `json:"email,omitempty" bson:"email,omitempty"`
	MobileNum          string             `json:"mobile_num,omitempty" bson:"mobile_num,omitempty"`
	Password           string             `json:"password,omitempty" bson:"password,omitempty"`
	ProfilePic         string             `json:"profile_pic,omitempty" bson:"profile_pic,omitempty"`
	EducationalDetails educationalDetails `json:"educational_details,omitempty" bson:"educational_details,omitempty"`
	Address            address            `json:"address,omitempty" bson:"address,omitempty"`
	TeachSessions      []string           `json:"teach_sessions,omitempty" bson:"teach_sessions,omitempty"`
	JoiningDate        string             `json:"joining_date,omitempty" bson:"joining_date,omitempty"`
}
