package entities

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id"`
	FirstName       string             `json:"firstname" bson:"firstname" binding:"required"`
	LastName        string             `json:"lastname" bson:"lastname" binding:"required"`
	Age             int                `json:"age" bson:"age" binding:"required"`
	Email           string             `json:"email" bson:"email" binding:"required"`
	Password        string             `json:"password" bson:"password" binding:"required,min=8"`
	ConfirmPassword string             `json:"confirmPassword" bson:"confirmPassword,omitempty" binding:"required"`
	Role            string             `json:"role" bson:"role"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
type RegisterResponse struct {
	FirstName string    `json:"firstname" bson:"firstname"`
	LastName  string    `json:"lastname" bson:"lastname"`
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
type Login struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required,min=8"`
}
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh"`
}

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}
