package recreation_area

import (
	entity "ca_parks/core/entity/recreation_area"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []ca_parksdb.RecreationArea) []entity.RecreationArea {
	result := []entity.RecreationArea{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model ca_parksdb.RecreationArea) entity.RecreationArea {
	return entity.RecreationArea{
		ID:        uuid.FromStringOrNil(model.ID),
		Name:      model.Name,
		Status:    entity.Status(model.Status),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		UserID:    uuid.FromStringOrNil(model.UserID),
	}
}
