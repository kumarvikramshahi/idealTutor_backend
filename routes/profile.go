package routes

import (
	"github.com/gofiber/fiber/v2"
)

func StudentProfile(app *fiber.App) {

	// Middleware of type /profile/<method>/<profile name>
	profile := app.Group("/profile")

	// routes
	profile.Get("/fetch/student/", GetStudentProfile)
	// router.Patch("/profile/edit/student/", patchStudentProfile)

	// router.Get("/fetch/studentList/", getStudentList)

	// // Teaching Sessions
	// router.Post("/session/add/", postTeachSession)
	// router.Patch("/session/edit/", patchTeachSession)
	// // router.delete("/session/delete/", deleteTeachSession);
	// router.Get("/session/status/", fetchTeachSession)
}
