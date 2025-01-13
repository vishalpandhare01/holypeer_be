package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	usercontroller "github.com/vishalpandhare01/holypeer_backend/internal/controller/user_controller"
	websocketcontroller "github.com/vishalpandhare01/holypeer_backend/internal/controller/web_socket_controller"
	"github.com/vishalpandhare01/holypeer_backend/internal/middleware"
)

func SetUpRouts(app *fiber.App) {
	app.Get("/", usercontroller.Server)
	app.Get("/ws/:id", websocket.New(websocketcontroller.SocketConnection))
	app.Post("/register", usercontroller.RegisterUser)
	app.Post("/send_otp", usercontroller.SendOtp)
	app.Patch("/verify_otp", usercontroller.VeryfyOtp)

	memebr := app.Group("/member")
	memebr.Post("/add", middleware.Authentication, middleware.IsMember, usercontroller.AddMemberProfile)
	memebr.Get("/get", middleware.Authentication, middleware.IsMember, usercontroller.GetMemberById)

	listner := app.Group("/listner")
	listner.Post("/add", middleware.Authentication, middleware.IsListener, usercontroller.AddListnerProfile)
	listner.Get("/get", middleware.Authentication, middleware.IsListener, usercontroller.GetListnerById)

}
