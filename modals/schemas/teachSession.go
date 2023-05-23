// tid: {
// 	type: String,
// 	required: true,
// 	trim: true,
// 	minlength: [24, "Invalid user ID"],
// 	maxlength: [24, "Invalid user ID"]
// },
// booked_by_id: {
// 	type: String,
// 	required: true,
// 	trim: true,
// 	minlength: [24, "Invalid user ID"],
// 	maxlength: [24, "Invalid user ID"]
// },
// is_grp_study: { type: Boolean, required: true },
// s_ids: [{ type: String, trim: true, minlength: 24, maxlength: 24, required: false }],
// sess_type: {
// 	type: String,
// 	required: true,
// 	lowercase: true,
// 	trim: true,
// 	enum: ["test", "teach", "discuss", "doubt"]
// },
// start_date: { type: Date, required: true, },
// end_data: { type: Date, required: true, },
// status: {
// 	type: String,
// 	required: true,
// 	lowercase: true,
// 	trim: true,
// 	enum: ["scheduled", "pending", "completed","cancelled"]
// },
// sub_teach: [{ type: String, required: true }],
// is_teacher_visit: { type: Boolean, required: true, },
// total_charge: { type: Number, required: true },
// pay_modes: [{
// 	type: String,
// 	required: true,
// 	lowercase: true,
// 	trim: true,
// 	enum: ["cod", "net banking", "upi"]
// }], // it is array coz, in partial payment will be multiple pay id
// pay_ids: [{
// 	type: String,
// 	trim: true,
// 	minlength: [24, "Invalid payment ID"],
// 	maxlength: [24, "Invalid payment ID"]
// }]
// });

package schemas

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeachSessions struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BookedById  string             `json:"booked_by_id,omitempty" bson:"booked_by_id,omitempty"`
	IsGrpStudy  bool               `json:"is_grp_study,omitempty" bson:"is_grp_study,omitempty"`
	StudentsIds []string           `json:"students_ids,omitempty" bson:"students_ids,omitempty"`
	SessionType string             `json:"session_type,omitempty" bson:"session_type,omitempty"`
	// 	enum: ["test", "teach", "discuss", "doubt"]
	StartDate string `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Status    string `bson:"omitempty"`
	// 	enum: ["scheduled", "pending", "completed","cancelled"]
	SubTeach       []string `json:"sub_teach,omitempty" bson:"sub_teach,omitempty"`
	IsTeacherVisit bool     `json:"is_teacher_visit,omitempty" bson:"is_teacher_visit,omitempty"`
	TotalCharge    int      `json:"total_charge,omitempty" bson:"total_charge,omitempty"`
	PayModes       []string `json:"pay_modes,omitempty" bson:"pay_modes,omitempty"`
	// 	enum: ["cod", "net banking", "upi"]
	PayIds []string `json:"pay_ids,omitempty" bson:"pay_ids,omitempty"`
}
