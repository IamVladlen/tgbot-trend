package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type countryRepo struct {
	mg *mongo.Collection
}

// ChangeCountry updates the country value in the chat document
// or creates new document if there's no records.
func (db *countryRepo) ChangeCountry(id int, country string) error {
	chat := entity.Chat{
		ChatId:    id,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		Country:   country,
	}

	ctx, cancel := context.WithTimeout(context.Background(), _mgdbRequestTimeout*time.Second)
	defer cancel()

	filter := bson.D{{Key: "chat_id", Value: chat.ChatId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "country", Value: chat.Country}, {Key: "updated_at", Value: chat.UpdatedAt},
		}},
	}
	opts := options.Update().SetUpsert(true)

	_, err := db.mg.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("repository - ChangeCountry: %w", err)
	}

	return nil
}

// GetCountry gets ISO code of the country set in chat.
func (db *countryRepo) GetCountry(id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _mgdbRequestTimeout*time.Second)
	defer cancel()

	var chat entity.Chat

	filter := bson.D{{Key: "chat_id", Value: id}}
	if err := db.mg.FindOne(ctx, filter).Decode(&chat); err != nil {
		return "", err
	}

	return chat.Country, nil
}

func newCountryRepo(mg *mongodb.DB) *countryRepo {
	return &countryRepo{
		mg: mg.Collection(_chatMgdbCollection),
	}
}
