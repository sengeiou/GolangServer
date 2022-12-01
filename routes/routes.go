package routes

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse
import (
	"main/controllers"
	"main/middleware"
	"main/validation"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// auth Start
	// manual signup
	app.Post("api/user/manual-signup",
		validation.ValidateUser,
		controllers.ManualRegister)
	// manual signin
	app.Post("api/user/manual-signin",
		validation.ValidateUserManualLogin,
		controllers.ManualLogin)
	// ------
	// Social login & Register google
	app.Post("api/user/social-Google-Apple-Facebook",
		validation.ValidateUserSocialAuth,
		controllers.SocialAuth)
	// Auth login & Register With Phone Number
	app.Post("api/user/Auth-Phone-Number",
		validation.ValidateUserPhoneAuth,
		controllers.AuthWithPhoneNumber)

	// Update User Data Any Time or After Siginup
	// we should have the token bearar and send it as Authorization header & user ID
	app.Patch("api/user/UpdateInfo/:id",
		middleware.AuthMiddleware,
		controllers.UpdateUserInfo)
	// update user sug like or hate
	app.Patch("api/user/UpdateSug/:mid/LikeOrHate/:nid",
		controllers.UpUserSug)
	// suggested partner
	app.Get("api/user/Suggested/:id",
		controllers.SuggestedPartner)

	// Message
	// Get User Message Room
	app.Get("api/message/:FuserID/:SuserID", controllers.GetPrivateRoomeID)

	// SendMessageToPrivateRoom
	app.Post("api/message", controllers.SendMessageToPrivateRoom)

	// GetMessageByNumbers // get the copy of the user
	app.Post("api/message/GetMessages/:roomId/:userId", controllers.GetMessageByNumbers)

	// testmessage
	app.Get("api/message/TestMessage", controllers.TestMessage)

	// PrivateFunGetRoomData
	app.Get("Private/:RoomID", controllers.PrivateFunGetRoomData)

	// ---------------------------- Group ----------------------- //

	// GetGroupRoomID
	app.Post("api/GroupMessage/:GroupName/:CreatorID", controllers.GetGroupRoomID)

	// SendMessageToGroupChat
	app.Post("api/SendMessageToGroup/:GroupID/:SenerId", controllers.SendMessageToGroupChat)

	// GetGroupMessageByNumbers
	app.Post("api/GetGroupMessageByNumbers/:roomId/:userId", controllers.GetGroupMessageByNumbers)

	// AddNewUserToGroup
	app.Post("api/AddNewUserToGroup/:roomId/:JoinUserId", middleware.AuthMiddleware, controllers.AddNewUserToGroup)

	// RemoveMemberFromchatGroup
	app.Delete("api/RemoveMemberFromchatGroup/:roomId/:JoinUserId", middleware.AuthMiddleware, controllers.RemoveMemberFromchatGroup)

	// ------ Story ------- //

	app.Post("api/AddStory/:userId", controllers.AddStory) // Add New Story

	app.Get("api/GetStoryes/:userId", controllers.GetStory) // return all User Storys

	app.Delete("api/RemoveStory/:userId/:StoryNumber", controllers.RemoveStory) // return Remove One Story

}
