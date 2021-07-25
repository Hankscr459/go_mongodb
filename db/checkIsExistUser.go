package db

import (
	"context"
	"time"
	"twitter/models"

	"go.mongodb.org/mongo-driver/bson"
)

func CheckIsExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("go_twitter")
	col := db.Collection("user")

	condition := bson.M{"email": email}
	var resultTodo models.User

	err := col.FindOne(ctx, condition).Decode(&resultTodo)
	ID := resultTodo.ID.Hex()
	if err != nil {
		return resultTodo, false, ID
	}
	return resultTodo, true, ID
}
