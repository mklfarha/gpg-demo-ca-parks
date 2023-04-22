package mapper

import (
	"ca_parks/graph/model"
	"fmt"
	"github.com/gofrs/uuid"
	"time"

	"ca_parks/core/entity/park"
	parkmodule "ca_parks/core/module/park/types"

	"ca_parks/core/entity/recreation_area"
	recreation_areamodule "ca_parks/core/module/recreation_area/types"

	"ca_parks/core/entity/user"
	usermodule "ca_parks/core/module/user/types"

	"ca_parks/core/entity/feature"
	featuremodule "ca_parks/core/module/feature/types"

	"ca_parks/core/entity/park_has_feature"
	park_has_featuremodule "ca_parks/core/module/park_has_feature/types"

	"ca_parks/core/entity/event"
	eventmodule "ca_parks/core/module/event/types"
)

func MapPark(in []park.Park) []*model.Park {
	res := []*model.Park{}
	for _, item := range in {
		i := item
		res = append(res, &model.Park{
			ID:               i.ID.String(),
			Name:             i.Name,
			MainImage:        i.MainImage,
			Phone:            &i.Phone,
			Hours:            &i.Hours,
			AllowsDogs:       i.AllowsDogs,
			Links:            MapParkLinks(i.Links),
			Status:           i.Status.String(),
			CreatedAt:        FormatTimeToPointer(i.CreatedAt),
			UpdatedAt:        FormatTimeToPointer(i.UpdatedAt),
			RecreationAreaID: i.RecreationAreaID.String(),
			UserID:           i.UserID.String(),
		})
	}
	return res
}

func MapParkUpsert(i model.ParkInput) parkmodule.UpsertRequest {
	return parkmodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: i.Name,

		MainImage: i.MainImage,

		Phone: StringFromPointer(i.Phone),

		Hours: StringFromPointer(i.Hours),

		AllowsDogs: i.AllowsDogs,

		Links: MapParkLinksInput(i.Links),

		Status: park.StatusFromString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		RecreationAreaID: uuid.FromStringOrNil(i.RecreationAreaID),

		UserID: uuid.FromStringOrNil(i.UserID),
	}
}

func MapParkUpsertPartial(i model.ParkPartialInput) parkmodule.UpsertRequest {
	return parkmodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: StringFromPointer(i.Name),

		MainImage: StringFromPointer(i.MainImage),

		Phone: StringFromPointer(i.Phone),

		Hours: StringFromPointer(i.Hours),

		AllowsDogs: BoolFromPointer(i.AllowsDogs),

		Links: MapParkLinksInput(i.Links),

		Status: park.StatusFromPointerString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		RecreationAreaID: UuidFromPointerString(i.RecreationAreaID),

		UserID: UuidFromPointerString(i.UserID),
	}
}

func MapRecreationArea(in []recreation_area.RecreationArea) []*model.RecreationArea {
	res := []*model.RecreationArea{}
	for _, item := range in {
		i := item
		res = append(res, &model.RecreationArea{
			ID:        i.ID.String(),
			Name:      i.Name,
			Status:    i.Status.String(),
			CreatedAt: FormatTimeToPointer(i.CreatedAt),
			UpdatedAt: FormatTimeToPointer(i.UpdatedAt),
			UserID:    i.UserID.String(),
		})
	}
	return res
}

func MapRecreationAreaUpsert(i model.RecreationAreaInput) recreation_areamodule.UpsertRequest {
	return recreation_areamodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: i.Name,

		Status: recreation_area.StatusFromString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		UserID: uuid.FromStringOrNil(i.UserID),
	}
}

func MapRecreationAreaUpsertPartial(i model.RecreationAreaPartialInput) recreation_areamodule.UpsertRequest {
	return recreation_areamodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: StringFromPointer(i.Name),

		Status: recreation_area.StatusFromPointerString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		UserID: UuidFromPointerString(i.UserID),
	}
}

func MapUser(in []user.User) []*model.User {
	res := []*model.User{}
	for _, item := range in {
		i := item
		res = append(res, &model.User{
			ID:        i.ID.String(),
			Name:      i.Name,
			Email:     i.Email,
			Password:  i.Password,
			Status:    i.Status.String(),
			CreatedAt: FormatTimeToPointer(i.CreatedAt),
			UpdatedAt: FormatTimeToPointer(i.UpdatedAt),
		})
	}
	return res
}

func MapUserUpsert(i model.UserInput) usermodule.UpsertRequest {
	return usermodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: i.Name,

		Email: i.Email,

		Password: i.Password,

		Status: user.StatusFromString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),
	}
}

func MapUserUpsertPartial(i model.UserPartialInput) usermodule.UpsertRequest {
	return usermodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: StringFromPointer(i.Name),

		Email: StringFromPointer(i.Email),

		Password: StringFromPointer(i.Password),

		Status: user.StatusFromPointerString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),
	}
}

func MapFeature(in []feature.Feature) []*model.Feature {
	res := []*model.Feature{}
	for _, item := range in {
		i := item
		res = append(res, &model.Feature{
			ID:        i.ID.String(),
			Name:      i.Name,
			Status:    i.Status.String(),
			CreatedAt: FormatTimeToPointer(i.CreatedAt),
			UpdatedAt: FormatTimeToPointer(i.UpdatedAt),
			UserID:    i.UserID.String(),
		})
	}
	return res
}

func MapFeatureUpsert(i model.FeatureInput) featuremodule.UpsertRequest {
	return featuremodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: i.Name,

		Status: feature.StatusFromString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		UserID: uuid.FromStringOrNil(i.UserID),
	}
}

func MapFeatureUpsertPartial(i model.FeaturePartialInput) featuremodule.UpsertRequest {
	return featuremodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: StringFromPointer(i.Name),

		Status: feature.StatusFromPointerString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		UserID: UuidFromPointerString(i.UserID),
	}
}

func MapParkHasFeature(in []park_has_feature.ParkHasFeature) []*model.ParkHasFeature {
	res := []*model.ParkHasFeature{}
	for _, item := range in {
		i := item
		res = append(res, &model.ParkHasFeature{
			ID:        i.ID.String(),
			Details:   &i.Details,
			Status:    i.Status.String(),
			CreatedAt: FormatTimeToPointer(i.CreatedAt),
			UpdatedAt: FormatTimeToPointer(i.UpdatedAt),
			ParkID:    i.ParkID.String(),
			FeatureID: i.FeatureID.String(),
		})
	}
	return res
}

func MapParkHasFeatureUpsert(i model.ParkHasFeatureInput) park_has_featuremodule.UpsertRequest {
	return park_has_featuremodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Details: StringFromPointer(i.Details),

		Status: park_has_feature.StatusFromString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		ParkID: uuid.FromStringOrNil(i.ParkID),

		FeatureID: uuid.FromStringOrNil(i.FeatureID),
	}
}

func MapParkHasFeatureUpsertPartial(i model.ParkHasFeaturePartialInput) park_has_featuremodule.UpsertRequest {
	return park_has_featuremodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Details: StringFromPointer(i.Details),

		Status: park_has_feature.StatusFromPointerString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		ParkID: UuidFromPointerString(i.ParkID),

		FeatureID: UuidFromPointerString(i.FeatureID),
	}
}

func MapEvent(in []event.Event) []*model.Event {
	res := []*model.Event{}
	for _, item := range in {
		i := item
		res = append(res, &model.Event{
			ID:          i.ID.String(),
			Name:        i.Name,
			Description: &i.Description,
			MainImage:   i.MainImage,
			StartDate:   i.StartDate.Format("2006-01-02"),
			EndDate:     i.EndDate.Format("2006-01-02"),
			Status:      i.Status.String(),
			CreatedAt:   FormatTimeToPointer(i.CreatedAt),
			UpdatedAt:   FormatTimeToPointer(i.UpdatedAt),
			ParkID:      i.ParkID.String(),
			UserID:      i.UserID.String(),
		})
	}
	return res
}

func MapEventUpsert(i model.EventInput) eventmodule.UpsertRequest {
	return eventmodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: i.Name,

		Description: StringFromPointer(i.Description),

		MainImage: i.MainImage,

		StartDate: ParseDate(i.StartDate),

		EndDate: ParseDate(i.EndDate),

		Status: event.StatusFromString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		ParkID: uuid.FromStringOrNil(i.ParkID),

		UserID: uuid.FromStringOrNil(i.UserID),
	}
}

func MapEventUpsertPartial(i model.EventPartialInput) eventmodule.UpsertRequest {
	return eventmodule.UpsertRequest{

		ID: UuidFromPointerString(i.ID),

		Name: StringFromPointer(i.Name),

		Description: StringFromPointer(i.Description),

		MainImage: StringFromPointer(i.MainImage),

		StartDate: ParseDateFromPointer(i.StartDate),

		EndDate: ParseDateFromPointer(i.EndDate),

		Status: event.StatusFromPointerString(i.Status),

		CreatedAt: ParseTimeFromPointer(i.CreatedAt),

		UpdatedAt: ParseTimeFromPointer(i.UpdatedAt),

		ParkID: UuidFromPointerString(i.ParkID),

		UserID: UuidFromPointerString(i.UserID),
	}
}

func MapParkLinks(items []park.Links) []*model.ParkLinks {
	res := []*model.ParkLinks{}
	for _, item := range items {
		i := item
		res = append(res, &model.ParkLinks{
			Type: i.Type.String(),
			Link: i.Link,
		})
	}
	return res
}
func MapParkLinksInput(items []*model.ParkLinksInput) []park.Links {

	if items == nil {
		return []park.Links{}
	}

	res := []park.Links{}
	for _, item := range items {
		i := item
		res = append(res, park.Links{
			Type: park.LinksTypeFromString(i.Type),
			Link: i.Link,
		})
	}
	return res
}

func ParseTime(in string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", in)
	if err != nil {
		fmt.Printf("error parsing date %v", err)
	}
	return t
}

func ParseTimeFromPointer(in *string) time.Time {
	if in == nil {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02 15:04:05", *in)
	if err != nil {
		fmt.Printf("error parsing date %v", err)
	}
	return t
}

func ParseDate(in string) time.Time {
	t, err := time.Parse("2006-01-02", in)
	if err != nil {
		fmt.Printf("error parsing date %v", err)
	}
	return t
}

func ParseDateFromPointer(in *string) time.Time {
	if in == nil {
		return time.Time{}
	}
	t, err := time.Parse("2006-01-02", *in)
	if err != nil {
		fmt.Printf("error parsing date %v", err)
	}
	return t
}

func FormatDateToPointer(date time.Time) *string {
	str := date.Format("2006-01-02")
	return &str
}

func FormatTimeToPointer(date time.Time) *string {
	str := date.Format("2006-01-02 15:04:05")
	return &str
}

func StringFromPointer(in *string) string {
	if in == nil {
		return ""
	}
	return *in
}

func IntFromPointer(in *int) int32 {
	if in == nil {
		return int32(0)
	}
	return int32(*in)
}

func IntPointer(in int32) *int {
	res := int(in)
	return &res
}

func FloatFromPointer(in *float64) float64 {
	if in == nil {
		return 0.0
	}
	return *in
}

func UuidToPointerString(u uuid.UUID) *string {
	if u == uuid.Nil {
		return nil
	}
	str := u.String()
	return &str
}

func UuidFromPointerString(i *string) uuid.UUID {
	if i == nil {
		return uuid.Nil
	}
	return uuid.FromStringOrNil(*i)
}

func BoolFromPointer(i *bool) bool {
	if i == nil {
		return false
	}
	return *i
}
