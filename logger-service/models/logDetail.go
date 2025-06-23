package models

import (
	"context"
	"time"

	"logger-service/db"
)

type LogDetail struct {
	Id        string
	Name      string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (log *LogDetail) Insert() error {
	collection := db.Client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(context.TODO(), LogDetail{
		Name:      log.Name,
		Data:      log.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}
	return nil
}
