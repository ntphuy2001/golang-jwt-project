package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"First_name" validate:"required,min=2,max=100"`
	LastName     *string            `json:"Last_name" validate:"required,min=2,max=100"`
	Password     *string            `json:"Password" validate:"required,min=2,max=100"`
	Email        *string            `json:"Email" validate:"email,required"`
	Phone        *string            `json:"Phone" validate:"required"`
	Token        *string            `json:"Token"`
	UserType     *string            `json:"User_type" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"Refresh_token"`
	CreatedAt    time.Time          `json:"Created_at"`
	UpdatedAt    time.Time          `json:"Updated_at"`
	UserId       string             `json:"User_id"`
}
