package event

import (
	entity "ca_parks/core/entity/event"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []ca_parksdb.Event) []entity.Event {
	result := []entity.Event{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model ca_parksdb.Event) entity.Event {
	return entity.Event{
		ID:          uuid.FromStringOrNil(model.ID),
		Name:        model.Name,
		Description: model.Description,
		MainImage:   model.MainImage,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		Status:      entity.Status(model.Status),
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		ParkID:      uuid.FromStringOrNil(model.ParkID),
		UserID:      uuid.FromStringOrNil(model.UserID),
	}
}
