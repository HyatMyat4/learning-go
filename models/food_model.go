package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Create_at  time.Time          `json:"create_at"`
	Update_at  time.Time          `json:"update_at"`
	Food_id    string             `json:"food_id" validate:"required"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}
