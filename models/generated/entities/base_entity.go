package entities

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID          uuid.UUID `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
	UpdatedAt  	time.Time `json:"updatedAt"`
	CreatedAt 	time.Time `json:"createdAt"`
}

func (b BaseEntity) GetID() uuid.UUID {
	return b.ID
}

func (b BaseEntity) GetName() string {
	return b.Name
}

func (b BaseEntity) GetDisplayName() string {
	return b.DisplayName
}

func (b BaseEntity) GetDescription() string {
	return b.Description
}

func (b BaseEntity) GetUpdatedAt() time.Time {
	return b.UpdatedAt
}

func (b BaseEntity) GetCreatedAt() time.Time {
	return b.CreatedAt
}

func (b BaseEntity) ToJson() ([]byte, error) {
	return json.Marshal(b)
}

func (b BaseEntity) JsonToString() (string, error) {
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}



