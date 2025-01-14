package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	chatcontroller "github.com/vishalpandhare01/holypeer_backend/internal/controller/chat_controller"
	usercontroller "github.com/vishalpandhare01/holypeer_backend/internal/controller/user_controller"
	websocketcontroller "github.com/vishalpandhare01/holypeer_backend/internal/controller/web_socket_controller"
	"github.com/vishalpandhare01/holypeer_backend/internal/middleware"
)

func SetUpRouts(app *fiber.App) {
	app.Get("/", usercontroller.Server)
	app.Post("/register", usercontroller.RegisterUser)
	app.Post("/send_otp", usercontroller.SendOtp)
	app.Patch("/verify_otp", usercontroller.VeryfyOtp)

	//middleware.Authentication, middleware.SecureRoom,websocket.New(websocketcontroller.SocketConnection)
	// app.Get("/ws/*",)
	app.Get("/ws/:id", middleware.Authentication, middleware.SecureRoom, websocket.New(websocketcontroller.SocketConnection))

	//member
	memebr := app.Group("/member")
	memebr.Post("/add", middleware.Authentication, middleware.IsMember, usercontroller.AddMemberProfile)
	memebr.Get("/get", middleware.Authentication, middleware.IsMember, usercontroller.GetMemberById)
	memebr.Get("/request", middleware.Authentication, middleware.IsMember, chatcontroller.RequestForchat)
	memebr.Delete("/close", middleware.Authentication, middleware.IsMember, chatcontroller.CloseChat)

	//listener
	listner := app.Group("/listner")
	listner.Post("/add", middleware.Authentication, middleware.IsListener, usercontroller.AddListnerProfile)
	listner.Get("/get", middleware.Authentication, middleware.IsListener, usercontroller.GetListnerById)
	listner.Get("/requests", middleware.Authentication, middleware.IsListener, chatcontroller.GetCahtRequests)
	listner.Put("/accept/:chatId", middleware.Authentication, middleware.IsListener, chatcontroller.AcceptRequest)

}
