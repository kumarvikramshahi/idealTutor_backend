package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kumarvikramshahi/idealTutor_backend/controllers/profiles"
)

func Profiles(app *fiber.App) {
	// Middleware of type /profile/<method>/<profile name>
	profilesRoute := app.Group("/profile")

	// routes
	profilesRoute.Get("/fetch/student/", profiles.GetStudentProfile)
	// router.Patch("/profile/edit/student/", patchStudentProfile)

	// router.Get("/fetch/studentList/", getStudentList)

	// // Teaching Sessions
	// router.Post("/session/add/", postTeachSession)
	// router.Patch("/session/edit/", patchTeachSession)
	// // router.delete("/session/delete/", deleteTeachSession);
	// router.Get("/session/status/", fetchTeachSession)
}
