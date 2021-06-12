package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de ususario en la BD en mongo*/
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre, omitempty" json:"nombre"`
	Apellidos       string             `bson:"apellidos" json:"apellidos"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password, omitempty" json:"password"`
	Avatar          string             `bson:"avatar, omitempty" json:"avatar"`
	Banner          string             `bson:"banner" json:"banner"`
	Biografia       string             `bson:"biografia" json:"biografia"`
	Ubicacion       string             `bson:"ubicacion, omitempty" json:"ubicacion"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioWeb"`
}
