// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: updates.sql

package ca_parksdb

import (
	"context"
	"encoding/json"
	"time"
)

const updateEvent = `-- name: UpdateEvent :exec
UPDATE event
SET
name = ?, description = ?, main_image = ?, start_date = ?, end_date = ?, status = ?, created_at = ?, updated_at = ?, park_id = ?, user_id = ?
WHERE id = ?
`

type UpdateEventParams struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MainImage   string    `json:"main_image"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      int32     `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ParkID      string    `json:"park_id"`
	UserID      string    `json:"user_id"`
	ID          string    `json:"id"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) error {
	_, err := q.db.ExecContext(ctx, updateEvent,
		arg.Name,
		arg.Description,
		arg.MainImage,
		arg.StartDate,
		arg.EndDate,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ParkID,
		arg.UserID,
		arg.ID,
	)
	return err
}

const updateFeature = `-- name: UpdateFeature :exec
UPDATE feature
SET
name = ?, status = ?, created_at = ?, updated_at = ?, user_id = ?
WHERE id = ?
`

type UpdateFeatureParams struct {
	Name      string    `json:"name"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	ID        string    `json:"id"`
}

func (q *Queries) UpdateFeature(ctx context.Context, arg UpdateFeatureParams) error {
	_, err := q.db.ExecContext(ctx, updateFeature,
		arg.Name,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.ID,
	)
	return err
}

const updatePark = `-- name: UpdatePark :exec
UPDATE park
SET
name = ?, main_image = ?, phone = ?, hours = ?, allows_dogs = ?, links = ?, status = ?, created_at = ?, updated_at = ?, recreation_area_id = ?, user_id = ?
WHERE id = ?
`

type UpdateParkParams struct {
	Name             string          `json:"name"`
	MainImage        string          `json:"main_image"`
	Phone            string          `json:"phone"`
	Hours            string          `json:"hours"`
	AllowsDogs       bool            `json:"allows_dogs"`
	Links            json.RawMessage `json:"links"`
	Status           int32           `json:"status"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	RecreationAreaID string          `json:"recreation_area_id"`
	UserID           string          `json:"user_id"`
	ID               string          `json:"id"`
}

func (q *Queries) UpdatePark(ctx context.Context, arg UpdateParkParams) error {
	_, err := q.db.ExecContext(ctx, updatePark,
		arg.Name,
		arg.MainImage,
		arg.Phone,
		arg.Hours,
		arg.AllowsDogs,
		arg.Links,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.RecreationAreaID,
		arg.UserID,
		arg.ID,
	)
	return err
}

const updateParkHasFeature = `-- name: UpdateParkHasFeature :exec
UPDATE park_has_feature
SET
details = ?, status = ?, created_at = ?, updated_at = ?, park_id = ?, feature_id = ?
WHERE id = ?
`

type UpdateParkHasFeatureParams struct {
	Details   string    `json:"details"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ParkID    string    `json:"park_id"`
	FeatureID string    `json:"feature_id"`
	ID        string    `json:"id"`
}

func (q *Queries) UpdateParkHasFeature(ctx context.Context, arg UpdateParkHasFeatureParams) error {
	_, err := q.db.ExecContext(ctx, updateParkHasFeature,
		arg.Details,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ParkID,
		arg.FeatureID,
		arg.ID,
	)
	return err
}

const updateRecreationArea = `-- name: UpdateRecreationArea :exec
UPDATE recreation_area
SET
name = ?, status = ?, created_at = ?, updated_at = ?, user_id = ?
WHERE id = ?
`

type UpdateRecreationAreaParams struct {
	Name      string    `json:"name"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	ID        string    `json:"id"`
}

func (q *Queries) UpdateRecreationArea(ctx context.Context, arg UpdateRecreationAreaParams) error {
	_, err := q.db.ExecContext(ctx, updateRecreationArea,
		arg.Name,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE user
SET
name = ?, email = ?, password = ?, status = ?, created_at = ?, updated_at = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
