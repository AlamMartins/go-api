package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name                 string             `bson:"name" json:"name"`
	Email                string             `bson:"email" json:"email"`
	Password             string             `bson:"password" json:"password"`
	ProfileImage         *string            `bson:"profileImage,omitempty" json:"profileImage,omitempty"`
	Bio                  *string            `bson:"bio,omitempty" json:"bio,omitempty"`
	ResetPasswordToken   *string            `bson:"resetPasswordToken,omitempty" json:"resetPasswordToken,omitempty"`
	ResetPasswordExpires *time.Time         `bson:"resetPasswordExpires,omitempty" json:"resetPasswordExpires,omitempty"`
	Role                 string             `bson:"role" json:"role"`
	CreatedAt            time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt            time.Time          `bson:"updatedAt" json:"updatedAt"`
}
