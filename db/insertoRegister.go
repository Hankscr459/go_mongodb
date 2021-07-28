package db

import (
	"context"
	"time"
	"twitter/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("go_twitter")
	col := db.Collection("users")
	u.Password, _ = EncriptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
