package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"main/database"
	"main/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SuggestedPartner
func SuggestedPartner(c *fiber.Ctx) error {
	var UserSchema = database.DB.Collection("users")
	var SuggestedSchema = database.DB.Collection("suggested")

	var ctx, _ = context.WithTimeout(context.Background(), 300*time.Second)

	var MainUser models.UserModel
	var users []models.UserModel
	var Foundeduser models.UserModel

	// var Suggestedes []models.SuggestedModel
	var Suggested models.SuggestedModel

	userid, _ := primitive.ObjectIDFromHex(c.Params("id"))

	err := UserSchema.FindOne(ctx, bson.M{"_id": userid}).Decode(&MainUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
	}

	// Searching
	// 1 Start By Searching of Parteners With in Same city and Location  Fo The User
	City := MainUser.UserLocation
	SCity := strings.Join(City, " ")

	filterCityLocation := bson.M{}

	findOptionsCity := options.Find()

	filterCityLocation = bson.M{
		"$or": []bson.M{
			{
				"userLocation": bson.M{
					"$regex": primitive.Regex{
						Pattern: SCity,
						Options: "i",
					},
				},
			},
		},
	}

	cursorUsers, _ := UserSchema.Find(ctx, filterCityLocation, findOptionsCity)

	defer cursorUsers.Close(ctx)

	if cursorUsers.RemainingBatchLength() <= 1 {
		cursorUsers, _ = UserSchema.Find(ctx, bson.M{}, options.Find())
		fmt.Println("Cursor", cursorUsers.RemainingBatchLength())
	}

	userIdStr := c.Params("id")

	for cursorUsers.Next(ctx) {
		// var sug models.SugListModel
		var user models.UserModel
		cursorUsers.Decode(&user)

		if userIdStr != user.ID.Hex() {
			users = append(users, user)

		}

		// sug.SugUserID = user.ID.Hex()
		// sug.LoveOrHate = true
		// sug.Score = 0

		// Suggested.MainUid = userIdStr
		// Suggested.SuggestedList = append(Suggested.SuggestedList, sug)
	}

	// Add many to Suggested Collection

	err = SuggestedSchema.FindOne(ctx, bson.M{"mainUid": userIdStr}).Decode(&Suggested)

	compeatedTime := time.Now().Sub(Suggested.CreatedAt).Hours()

	MaxNumber := 0
	var FinalSelectedUser string

	// -----------
	MSCORE := 0
	// var SelectedOne string
	for i := range Suggested.SuggestedList {
		sc := Suggested.SuggestedList[i].Score
		if sc > MSCORE && Suggested.SuggestedList[i].LoveOrHate {
			MSCORE = sc
			FinalSelectedUser = Suggested.SuggestedList[i].SugUserID
		}
	}

	if (err != nil && compeatedTime > 24) && MSCORE == 0 {
		for _, user := range users {
			var sug models.SugListModel
			sug.SugUserID = user.ID.Hex()
			sug.LoveOrHate = true
			sug.Score = 0
			//------------------------------
			for _, item := range user.PhyslcalAttraction {
				for _, i := range MainUser.UserBody {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserDrink {
				for _, i := range MainUser.UserDrink {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserHobbyes {
				for _, i := range MainUser.UserHobbyes {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserLanguages {
				for _, i := range MainUser.UserLanguages {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserPets {
				for _, i := range MainUser.UserPets {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserZodlac {
				for _, i := range MainUser.UserZodlac {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			if user.PlaceOfJob == MainUser.PlaceOfJob {
				sug.Score = sug.Score + 1
			}

			if user.School == MainUser.School {
				sug.Score = sug.Score + 1
			}

			if user.UserHeight == MainUser.UserHeight {
				sug.Score = sug.Score + 1
			}

			if user.UserLookFor == MainUser.UserLookFor {
				sug.Score = sug.Score + 1
			}

			if user.UserRelationship == MainUser.UserRelationship {
				sug.Score = sug.Score + 1
			}

			if user.UserRole != MainUser.UserRole {
				sug.Score = sug.Score + 1
			}

			if user.IsUserSmoking == MainUser.IsUserSmoking {
				sug.Score = sug.Score + 1
			}

			//------------------------------

			if sug.Score > MaxNumber {
				FinalSelectedUser = user.ID.Hex()
			}

			MaxNumber = sug.Score
			//-----------------------------
			Suggested.MainUid = userIdStr
			Suggested.SuggestedList = append(Suggested.SuggestedList, sug)
		}

		GetTimeNow := time.Now()
		Suggested.CreatedAt = GetTimeNow
		_, err = SuggestedSchema.InsertOne(ctx, &Suggested)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

	} else if MSCORE == 0 {

		SuggestedSchema.DeleteOne(ctx, bson.M{"mainUid": userIdStr})
		Suggested.SuggestedList = nil
		for _, user := range users {
			var sug models.SugListModel
			sug.SugUserID = user.ID.Hex()
			sug.LoveOrHate = true
			sug.Score = 0
			//------------------------------
			for _, item := range user.PhyslcalAttraction {
				for _, i := range MainUser.UserBody {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserDrink {
				for _, i := range MainUser.UserDrink {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserHobbyes {
				for _, i := range MainUser.UserHobbyes {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserLanguages {
				for _, i := range MainUser.UserLanguages {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserPets {
				for _, i := range MainUser.UserPets {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			for _, item := range user.UserZodlac {
				for _, i := range MainUser.UserZodlac {
					if item == i {
						sug.Score = sug.Score + 1
					}
				}
			}

			if user.PlaceOfJob == MainUser.PlaceOfJob {
				sug.Score = sug.Score + 1
			}

			if user.School == MainUser.School {
				sug.Score = sug.Score + 1
			}

			if user.UserHeight == MainUser.UserHeight {
				sug.Score = sug.Score + 1
			}

			if user.UserLookFor == MainUser.UserLookFor {
				sug.Score = sug.Score + 1
			}

			if user.UserRelationship == MainUser.UserRelationship {
				sug.Score = sug.Score + 1
			}

			if user.UserRole != MainUser.UserRole {
				sug.Score = sug.Score + 1
			}

			if user.IsUserSmoking == MainUser.IsUserSmoking {
				sug.Score = sug.Score + 1
			}

			//------------------------------

			if sug.Score > MaxNumber {
				FinalSelectedUser = user.ID.Hex()
			}

			MaxNumber = sug.Score
			//-----------------------------
			Suggested.MainUid = userIdStr
			Suggested.SuggestedList = append(Suggested.SuggestedList, sug)
		}

		GetTimeNow := time.Now()
		Suggested.CreatedAt = GetTimeNow
		SuggestedSchema.DeleteOne(ctx, bson.M{"mainUid": userIdStr})
		_, err = SuggestedSchema.InsertOne(ctx, Suggested)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		// return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Erro": "No Available Suggested For Now Try Again"})

	}
	fmt.Println("F A S", FinalSelectedUser)
	if FinalSelectedUser != "" {
		fid, _ := primitive.ObjectIDFromHex(FinalSelectedUser)

		userResult := UserSchema.FindOne(ctx, bson.M{"_id": fid})

		if userResult.Err() != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Or user Posts Not found",
			})
		}

		userResult.Decode(&Foundeduser)
		// fmt.Println("MSSSS", Foundeduser)

	}

	// ------

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"User": Foundeduser})

}

// Update user Sug LikeSugOrNot
func UpUserSug(c *fiber.Ctx) error {
	var SuggestedSchema = database.DB.Collection("suggested")
	var UserSchema = database.DB.Collection("users")

	var ctx, _ = context.WithTimeout(context.Background(), 300*time.Second)

	var sug models.SuggestedModel

	mainUsserid := c.Params("mid")
	SecondUser := c.Params("nid")

	err := SuggestedSchema.FindOne(ctx, bson.M{"mainUid": mainUsserid}).Decode(&sug)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
	}

	for i := range sug.SuggestedList {
		// fmt.Println(sug.SuggestedList[i])
		if sug.SuggestedList[i].SugUserID == SecondUser {
			sug.SuggestedList[i].LoveOrHate = !sug.SuggestedList[i].LoveOrHate
		}
	}

	result, err := SuggestedSchema.UpdateOne(ctx, bson.M{"mainUid": mainUsserid}, bson.M{"$set": sug})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
	}
	// err = SuggestedSchema.UpdateOne(ctx, bson.M{"mainUid": mainUsserid},{sug}).Decode(&sug)
	var updatedSug models.SuggestedModel
	if result.MatchedCount == 1 {
		err := SuggestedSchema.FindOne(ctx, bson.M{"mainUid": mainUsserid}).Decode(&updatedSug)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
		}
	}

	// Get the Suggested
	// MaxNumber := 0
	var FinalSelectedUser string
	var Foundeduser models.UserModel

	// -----------
	MSCORE := 0
	// var SelectedOne string
	for i := range updatedSug.SuggestedList {
		sc := updatedSug.SuggestedList[i].Score
		if sc > MSCORE && updatedSug.SuggestedList[i].LoveOrHate {
			MSCORE = sc
			FinalSelectedUser = updatedSug.SuggestedList[i].SugUserID
		}
	}

	if FinalSelectedUser != "" {
		fid, _ := primitive.ObjectIDFromHex(FinalSelectedUser)

		userResult := UserSchema.FindOne(ctx, bson.M{"_id": fid})

		if userResult.Err() != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Or user Posts Not found",
			})
		}

		userResult.Decode(&Foundeduser)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": Foundeduser})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"data": "No Matched Users Founded!"})
	}

}

// Update User Informatin

func UpdateUserInfo(c *fiber.Ctx) error {
	var UserSchema = database.DB.Collection("users")
	var ctx, _ = context.WithTimeout(context.Background(), 120*time.Second)

	var user models.UserModel
	c.BodyParser(&user)

	userid, _ := primitive.ObjectIDFromHex(c.Params("id"))

	var getingUser models.UserModel
	err := UserSchema.FindOne(ctx, bson.M{"_id": userid}).Decode(&getingUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
	}

	b, err := json.Marshal(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
	}

	var jsonMap map[string]string
	json.Unmarshal([]byte(string(b)), &jsonMap)

	var Next map[string][]string
	json.Unmarshal([]byte(string(b)), &Next)

	update := bson.M{}

	for key, val := range jsonMap { // strings
		if val != "" && key != "_id" {
			update[key] = val
			// fmt.Println("k", key, "val", val)
		}
	}

	for key, val := range Next { // arrays
		if val != nil && key != "_id" {
			if key == "userMediaPhoto" {
				update[key] = append(getingUser.UserMediaPhoto, string(val[0]))
			}

			if key == "userMediaVideo" {
				update[key] = append(getingUser.UserMediaVideo, string(val[0]))
			}

			if key == "userLanguages" {
				update[key] = append(getingUser.UserLanguages, string(val[0]))
			}

			if key == "userLocation" {
				update[key] = append(getingUser.UserLocation, string(val[0]))
			}

			if key == "locationDetails" {
				update[key] = append(getingUser.LocationDetails, string(val[0]))
			}

			if key == "userHobbyes" {
				update[key] = append(getingUser.UserHobbyes, string(val[0]))
			}

			if key == "userZodlac" {
				update[key] = append(getingUser.UserZodlac, string(val[0]))
			}

			if key == "userPets" {
				update[key] = append(getingUser.UserPets, string(val[0]))
			}

			if key == "physlcalAttraction" {
				update[key] = append(getingUser.PhyslcalAttraction, string(val[0]))
			}

			if key == "userTurnON" {
				update[key] = append(getingUser.UserTurnON, string(val[0]))
			}

			if key == "userStyle" {
				update[key] = append(getingUser.UserStyle, string(val[0]))
			}

			if key == "userBody" {
				update[key] = append(getingUser.UserBody, string(val[0]))
			}

			if key == "userSesson" {
				update[key] = append(getingUser.UserSesson, string(val[0]))
			}

			if key == "userMovies" {
				update[key] = append(getingUser.UserMovies, string(val[0]))
			}

			if key == "userDrink" {
				update[key] = append(getingUser.UserDrink, string(val[0]))
			}
		}
	}

	// fmt.Println(update)

	result, err := UserSchema.UpdateOne(ctx, bson.M{"_id": userid}, bson.M{"$set": update})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
	}
	//get updated post details
	var updatedUser models.UserModel
	if result.MatchedCount == 1 {
		err := UserSchema.FindOne(ctx, bson.M{"_id": userid}).Decode(&updatedUser)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"data": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": updatedUser})

}
