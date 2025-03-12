package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/scallyt/neko/internal/db"
	"github.com/scallyt/neko/internal/repository"
	svix "github.com/svix/svix-webhooks/go"
)

func ClerkWebhookHandler(c *gin.Context) {
	secret := "whsec_Sv5lyrtrihEa/T2/QJrvQRV00RjkyBUx"
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Webhook secret not set"})
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	svixHeaders := c.Request.Header
	wh, err := svix.NewWebhook(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create webhook verifier"})
		return
	}

	err = wh.Verify(body, svixHeaders)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid webhook signature"})
		return
	}

	err = ParseClerkWebhook(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to process webhook: %v", err)})
		return
	}

	c.Status(http.StatusOK)
	c.Writer.WriteString("Webhook processed successfully")
}

func ParseClerkWebhook(webhookData []byte) error {
	repo := repository.New(db.Conn)

	var webhook struct {
		Data struct {
			ID             string `json:"id"`
			FirstName      string `json:"first_name"`
			LastName       string `json:"last_name"`
			EmailAddresses []struct {
				EmailAddress string `json:"email_address"`
				ID           string `json:"id"`
			} `json:"email_addresses"`
			PhoneNumbers []struct {
				PhoneNumber string `json:"phone_number"`
				ID          string `json:"id"`
			} `json:"phone_numbers"`
			PrimaryEmailAddressID string `json:"primary_email_address_id"`
			PrimaryPhoneNumberID  string `json:"primary_phone_number_id"`
			ProfileImageURL       string `json:"profile_image_url"`
			CreatedAt             int64  `json:"created_at"`
			UpdatedAt             int64  `json:"updated_at"`
			ExternalID            string `json:"external_id"`
		} `json:"data"`
		Type string `json:"type"`
	}

	if err := json.Unmarshal(webhookData, &webhook); err != nil {
		return fmt.Errorf("error parsing webhook data: %v", err)
	}

	switch webhook.Type {
	case "user.created", "user.updated":
		return handleUserCreatedOrUpdated(repo, webhook.Data)
	case "user.deleted":
		return handleUserDeleted(repo, webhook.Data.ID)
	default:
		return nil
	}
}

func handleUserCreatedOrUpdated(repo *repository.Queries, data struct {
	ID             string `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	EmailAddresses []struct {
		EmailAddress string `json:"email_address"`
		ID           string `json:"id"`
	} `json:"email_addresses"`
	PhoneNumbers []struct {
		PhoneNumber string `json:"phone_number"`
		ID          string `json:"id"`
	} `json:"phone_numbers"`
	PrimaryEmailAddressID string `json:"primary_email_address_id"`
	PrimaryPhoneNumberID  string `json:"primary_phone_number_id"`
	ProfileImageURL       string `json:"profile_image_url"`
	CreatedAt             int64  `json:"created_at"`
	UpdatedAt             int64  `json:"updated_at"`
	ExternalID            string `json:"external_id"`
}) error {
	var email, phone string
	for _, e := range data.EmailAddresses {
		if e.ID == data.PrimaryEmailAddressID {
			email = e.EmailAddress
			break
		}
	}

	for _, p := range data.PhoneNumbers {
		if p.ID == data.PrimaryPhoneNumberID {
			phone = p.PhoneNumber
			break
		}
	}

	return repo.CreateWebhookUser(db.Ctx, repository.CreateWebhookUserParams{
		ID:              data.ID,
		FirstName:       data.FirstName,
		LastName:        data.LastName,
		Email:           email,
		PhoneNumber:     pgtype.Text{String: phone, Valid: phone != ""},
		ProfileImageUrl: pgtype.Text{String: data.ProfileImageURL, Valid: data.ProfileImageURL != ""},
		CreatedAt:       pgtype.Timestamp{Time: time.Unix(0, data.CreatedAt*1000000), Valid: true},
		UpdatedAt:       pgtype.Timestamp{Time: time.Unix(0, data.UpdatedAt*1000000), Valid: true},
		ExternalID:      pgtype.Text{String: data.ExternalID, Valid: data.ExternalID != ""},
	})
}

func handleUserDeleted(repo *repository.Queries, userID string) error {
	return repo.DeleteWebhookUser(db.Ctx, userID)
}
