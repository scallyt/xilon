// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Chat struct {
	ID        pgtype.UUID `json:"id"`
	ProjectID pgtype.UUID `json:"project_id"`
}

type Message struct {
	ID        pgtype.UUID        `json:"id"`
	ChatID    pgtype.UUID        `json:"chat_id"`
	Sender    string             `json:"sender"`
	Content   string             `json:"content"`
	Timestamp pgtype.Timestamptz `json:"timestamp"`
}

type Project struct {
	ID           pgtype.UUID    `json:"id"`
	CustomerName string         `json:"customer_name"`
	CustomerID   pgtype.UUID    `json:"customer_id"`
	ProjectName  string         `json:"project_name"`
	ChatID       pgtype.UUID    `json:"chat_id"`
	Description  pgtype.Text    `json:"description"`
	DeveloperID  pgtype.UUID    `json:"developer_id"`
	Status       string         `json:"status"`
	StartDate    pgtype.Date    `json:"start_date"`
	EndDate      pgtype.Date    `json:"end_date"`
	Budget       pgtype.Numeric `json:"budget"`
}

type Task struct {
	ID        pgtype.UUID `json:"id"`
	ProjectID pgtype.UUID `json:"project_id"`
	TaskName  string      `json:"task_name"`
	Todo      pgtype.Text `json:"todo"`
	Status    string      `json:"status"`
}

type User struct {
	ID              string           `json:"id"`
	FirstName       string           `json:"first_name"`
	LastName        string           `json:"last_name"`
	Email           string           `json:"email"`
	PhoneNumber     pgtype.Text      `json:"phone_number"`
	ProfileImageUrl pgtype.Text      `json:"profile_image_url"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
	UpdatedAt       pgtype.Timestamp `json:"updated_at"`
	ExternalID      pgtype.Text      `json:"external_id"`
}
