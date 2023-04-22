// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package ca_parksdb

import (
	"context"
	"database/sql"
)

type Querier interface {
	FetchEventByID(ctx context.Context, id string) ([]Event, error)
	FetchEventByParkID(ctx context.Context, arg FetchEventByParkIDParams) ([]Event, error)
	FetchEventByParkIDAndStatus(ctx context.Context, arg FetchEventByParkIDAndStatusParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByCreatedAtASCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByCreatedAtDESCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByEndDateASC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByEndDateASCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByEndDateDESC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByEndDateDESCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByStartDateASC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByStartDateASCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByStartDateDESC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByStartDateDESCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByUpdatedAtASCParams) ([]Event, error)
	FetchEventByParkIDAndStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchEventByParkIDAndStatusOrderedByUpdatedAtDESCParams) ([]Event, error)
	FetchEventByParkIDOrderedByCreatedAtASC(ctx context.Context, arg FetchEventByParkIDOrderedByCreatedAtASCParams) ([]Event, error)
	FetchEventByParkIDOrderedByCreatedAtDESC(ctx context.Context, arg FetchEventByParkIDOrderedByCreatedAtDESCParams) ([]Event, error)
	FetchEventByParkIDOrderedByEndDateASC(ctx context.Context, arg FetchEventByParkIDOrderedByEndDateASCParams) ([]Event, error)
	FetchEventByParkIDOrderedByEndDateDESC(ctx context.Context, arg FetchEventByParkIDOrderedByEndDateDESCParams) ([]Event, error)
	FetchEventByParkIDOrderedByStartDateASC(ctx context.Context, arg FetchEventByParkIDOrderedByStartDateASCParams) ([]Event, error)
	FetchEventByParkIDOrderedByStartDateDESC(ctx context.Context, arg FetchEventByParkIDOrderedByStartDateDESCParams) ([]Event, error)
	FetchEventByParkIDOrderedByUpdatedAtASC(ctx context.Context, arg FetchEventByParkIDOrderedByUpdatedAtASCParams) ([]Event, error)
	FetchEventByParkIDOrderedByUpdatedAtDESC(ctx context.Context, arg FetchEventByParkIDOrderedByUpdatedAtDESCParams) ([]Event, error)
	FetchEventByStatus(ctx context.Context, arg FetchEventByStatusParams) ([]Event, error)
	FetchEventByStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchEventByStatusOrderedByCreatedAtASCParams) ([]Event, error)
	FetchEventByStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchEventByStatusOrderedByCreatedAtDESCParams) ([]Event, error)
	FetchEventByStatusOrderedByEndDateASC(ctx context.Context, arg FetchEventByStatusOrderedByEndDateASCParams) ([]Event, error)
	FetchEventByStatusOrderedByEndDateDESC(ctx context.Context, arg FetchEventByStatusOrderedByEndDateDESCParams) ([]Event, error)
	FetchEventByStatusOrderedByStartDateASC(ctx context.Context, arg FetchEventByStatusOrderedByStartDateASCParams) ([]Event, error)
	FetchEventByStatusOrderedByStartDateDESC(ctx context.Context, arg FetchEventByStatusOrderedByStartDateDESCParams) ([]Event, error)
	FetchEventByStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchEventByStatusOrderedByUpdatedAtASCParams) ([]Event, error)
	FetchEventByStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchEventByStatusOrderedByUpdatedAtDESCParams) ([]Event, error)
	FetchFeatureByID(ctx context.Context, id string) ([]Feature, error)
	FetchFeatureByStatus(ctx context.Context, arg FetchFeatureByStatusParams) ([]Feature, error)
	FetchFeatureByStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchFeatureByStatusOrderedByCreatedAtASCParams) ([]Feature, error)
	FetchFeatureByStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchFeatureByStatusOrderedByCreatedAtDESCParams) ([]Feature, error)
	FetchFeatureByStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchFeatureByStatusOrderedByUpdatedAtASCParams) ([]Feature, error)
	FetchFeatureByStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchFeatureByStatusOrderedByUpdatedAtDESCParams) ([]Feature, error)
	FetchParkByID(ctx context.Context, id string) ([]Park, error)
	FetchParkByRecreationAreaID(ctx context.Context, arg FetchParkByRecreationAreaIDParams) ([]Park, error)
	FetchParkByRecreationAreaIDAndStatus(ctx context.Context, arg FetchParkByRecreationAreaIDAndStatusParams) ([]Park, error)
	FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtASCParams) ([]Park, error)
	FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtDESCParams) ([]Park, error)
	FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtASCParams) ([]Park, error)
	FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtDESCParams) ([]Park, error)
	FetchParkByRecreationAreaIDOrderedByCreatedAtASC(ctx context.Context, arg FetchParkByRecreationAreaIDOrderedByCreatedAtASCParams) ([]Park, error)
	FetchParkByRecreationAreaIDOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkByRecreationAreaIDOrderedByCreatedAtDESCParams) ([]Park, error)
	FetchParkByRecreationAreaIDOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkByRecreationAreaIDOrderedByUpdatedAtASCParams) ([]Park, error)
	FetchParkByRecreationAreaIDOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkByRecreationAreaIDOrderedByUpdatedAtDESCParams) ([]Park, error)
	FetchParkByStatus(ctx context.Context, arg FetchParkByStatusParams) ([]Park, error)
	FetchParkByStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchParkByStatusOrderedByCreatedAtASCParams) ([]Park, error)
	FetchParkByStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkByStatusOrderedByCreatedAtDESCParams) ([]Park, error)
	FetchParkByStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkByStatusOrderedByUpdatedAtASCParams) ([]Park, error)
	FetchParkByStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkByStatusOrderedByUpdatedAtDESCParams) ([]Park, error)
	FetchParkHasFeatureByFeatureID(ctx context.Context, arg FetchParkHasFeatureByFeatureIDParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDAndStatus(ctx context.Context, arg FetchParkHasFeatureByFeatureIDAndStatusParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDOrderedByCreatedAtASC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDOrderedByCreatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDOrderedByCreatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByID(ctx context.Context, id string) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkID(ctx context.Context, arg FetchParkHasFeatureByParkIDParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDAndStatus(ctx context.Context, arg FetchParkHasFeatureByParkIDAndStatusParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDOrderedByCreatedAtASC(ctx context.Context, arg FetchParkHasFeatureByParkIDOrderedByCreatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByParkIDOrderedByCreatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkHasFeatureByParkIDOrderedByUpdatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByParkIDOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByParkIDOrderedByUpdatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByStatus(ctx context.Context, arg FetchParkHasFeatureByStatusParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchParkHasFeatureByStatusOrderedByCreatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByStatusOrderedByCreatedAtDESCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchParkHasFeatureByStatusOrderedByUpdatedAtASCParams) ([]ParkHasFeature, error)
	FetchParkHasFeatureByStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchParkHasFeatureByStatusOrderedByUpdatedAtDESCParams) ([]ParkHasFeature, error)
	FetchRecreationAreaByID(ctx context.Context, id string) ([]RecreationArea, error)
	FetchRecreationAreaByStatus(ctx context.Context, arg FetchRecreationAreaByStatusParams) ([]RecreationArea, error)
	FetchRecreationAreaByStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchRecreationAreaByStatusOrderedByCreatedAtASCParams) ([]RecreationArea, error)
	FetchRecreationAreaByStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchRecreationAreaByStatusOrderedByCreatedAtDESCParams) ([]RecreationArea, error)
	FetchRecreationAreaByStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchRecreationAreaByStatusOrderedByUpdatedAtASCParams) ([]RecreationArea, error)
	FetchRecreationAreaByStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchRecreationAreaByStatusOrderedByUpdatedAtDESCParams) ([]RecreationArea, error)
	FetchUserByEmail(ctx context.Context, arg FetchUserByEmailParams) ([]User, error)
	FetchUserByEmailAndStatus(ctx context.Context, arg FetchUserByEmailAndStatusParams) ([]User, error)
	FetchUserByEmailAndStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchUserByEmailAndStatusOrderedByCreatedAtASCParams) ([]User, error)
	FetchUserByEmailAndStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchUserByEmailAndStatusOrderedByCreatedAtDESCParams) ([]User, error)
	FetchUserByEmailAndStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchUserByEmailAndStatusOrderedByUpdatedAtASCParams) ([]User, error)
	FetchUserByEmailAndStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchUserByEmailAndStatusOrderedByUpdatedAtDESCParams) ([]User, error)
	FetchUserByEmailOrderedByCreatedAtASC(ctx context.Context, arg FetchUserByEmailOrderedByCreatedAtASCParams) ([]User, error)
	FetchUserByEmailOrderedByCreatedAtDESC(ctx context.Context, arg FetchUserByEmailOrderedByCreatedAtDESCParams) ([]User, error)
	FetchUserByEmailOrderedByUpdatedAtASC(ctx context.Context, arg FetchUserByEmailOrderedByUpdatedAtASCParams) ([]User, error)
	FetchUserByEmailOrderedByUpdatedAtDESC(ctx context.Context, arg FetchUserByEmailOrderedByUpdatedAtDESCParams) ([]User, error)
	FetchUserByID(ctx context.Context, id string) ([]User, error)
	FetchUserByStatus(ctx context.Context, arg FetchUserByStatusParams) ([]User, error)
	FetchUserByStatusOrderedByCreatedAtASC(ctx context.Context, arg FetchUserByStatusOrderedByCreatedAtASCParams) ([]User, error)
	FetchUserByStatusOrderedByCreatedAtDESC(ctx context.Context, arg FetchUserByStatusOrderedByCreatedAtDESCParams) ([]User, error)
	FetchUserByStatusOrderedByUpdatedAtASC(ctx context.Context, arg FetchUserByStatusOrderedByUpdatedAtASCParams) ([]User, error)
	FetchUserByStatusOrderedByUpdatedAtDESC(ctx context.Context, arg FetchUserByStatusOrderedByUpdatedAtDESCParams) ([]User, error)
	InsertEvent(ctx context.Context, arg InsertEventParams) (sql.Result, error)
	InsertFeature(ctx context.Context, arg InsertFeatureParams) (sql.Result, error)
	InsertPark(ctx context.Context, arg InsertParkParams) (sql.Result, error)
	InsertParkHasFeature(ctx context.Context, arg InsertParkHasFeatureParams) (sql.Result, error)
	InsertRecreationArea(ctx context.Context, arg InsertRecreationAreaParams) (sql.Result, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (sql.Result, error)
	SearchFeature(ctx context.Context, arg SearchFeatureParams) ([]Feature, error)
	SearchPark(ctx context.Context, arg SearchParkParams) ([]Park, error)
	SearchRecreationArea(ctx context.Context, arg SearchRecreationAreaParams) ([]RecreationArea, error)
	UpdateEvent(ctx context.Context, arg UpdateEventParams) error
	UpdateFeature(ctx context.Context, arg UpdateFeatureParams) error
	UpdatePark(ctx context.Context, arg UpdateParkParams) error
	UpdateParkHasFeature(ctx context.Context, arg UpdateParkHasFeatureParams) error
	UpdateRecreationArea(ctx context.Context, arg UpdateRecreationAreaParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)