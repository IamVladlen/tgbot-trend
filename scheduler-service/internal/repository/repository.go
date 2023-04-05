package repository

import (
	"context"
	"fmt"

	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/postgres"
)

type Repository struct {
	*postgres.DB
}

func New(pg *postgres.DB) *Repository {
	return &Repository{
		pg,
	}
}

func (db *Repository) GetScheduledMessages(ctx context.Context, interval string) ([]int64, error) {
	chatIds := make([]int64, 0)

	q := `SELECT id
			FROM chats
			WHERE newsletter_interval = $1;`
	rows, err := db.Query(ctx, q, interval)
	if err != nil {
		return []int64{}, fmt.Errorf("repository - GetScheduledMessages: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var chatId int64
		if err := rows.Scan(&chatId); err != nil {
			return chatIds, fmt.Errorf("repository - GetScheduledMessages: %w", err)
		}

		chatIds = append(chatIds, chatId)
	}

	return chatIds, nil
}

func (db *Repository) SetChatSchedule(ctx context.Context, chatId int64, interval string) error {
	q := `INSERT INTO chats (id, newsletter_interval)
			VALUES ($1, $2)
			ON CONFLICT (id)
			DO UPDATE
			SET newsletter_interval = EXCLUDED.newsletter_interval;`
	_, err := db.Exec(ctx, q, chatId, interval)
	if err != nil {
		return fmt.Errorf("repository - SetChatSchedule: %w", err)
	}

	return nil
}
