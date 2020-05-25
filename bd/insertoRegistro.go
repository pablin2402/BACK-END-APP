package bd

import (
	"context"
	"time"

	"github.com/snetwork/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro inserta los datos en la BD*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}
	ObjID := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
