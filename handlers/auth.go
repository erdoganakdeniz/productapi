package handlers

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"productapi/models"
	"productapi/utils"
)

type AuthHandlerInterface interface {
	Login(ctx *fiber.Ctx) interface{}
	Signup(ctx *fiber.Ctx) interface{}
}
type AuthHandler struct {
	UsersColl *mongo.Collection
}

func (a AuthHandler) Login(c *fiber.Ctx)  {
	u:=new(models.LoginInputs)
	if err:=c.BodyParser(u);err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}
	user:=new(models.User)
	filter:=bson.M{"email":u.Email}
	err:=a.UsersColl.FindOne(c.Fasthttp,filter).Decode(user)
	if err != nil {
		c.Status(fiber.StatusUnauthorized).Send(fiber.Map{"message":"Invalid Credentials"})
		return
	}
	isMatch:=utils.Password{Password: u.Password}.Compare(user.Password)
	if !isMatch{
		c.Status(fiber.StatusUnauthorized).Send(fiber.Map{"message":"Invalid Credentials"})
		return
	}
	accessToken,err:=utils.CreateJWTToken(map[string]interface{}{"username":user.UserName,"email":user.Email,"id":user.ID,})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successfully",
		"data":    fiber.Map{"token": accessToken},
	})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}

}
