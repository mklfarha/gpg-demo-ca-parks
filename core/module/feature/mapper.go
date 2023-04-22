package feature

import (
	entity "ca_parks/core/entity/feature"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []ca_parksdb.Feature) []entity.Feature {
	result := []entity.Feature{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model ca_parksdb.Feature) entity.Feature {
	return entity.Feature{
		ID:        uuid.FromStringOrNil(model.ID),
		Name:      model.Name,
		Status:    entity.Status(model.Status),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		UserID:    uuid.FromStringOrNil(model.UserID),
	}
}
