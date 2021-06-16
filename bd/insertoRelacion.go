package bd

import (
	"context"
	"time"

	"github.com/DiegoEscCab/twider/models"
)

/* InsertoRelacion graba la relacion en la Base de Datos*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twider")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
