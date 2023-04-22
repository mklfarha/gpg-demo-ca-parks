package park_has_feature

import (
	entity "ca_parks/core/entity/park_has_feature"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []ca_parksdb.ParkHasFeature) []entity.ParkHasFeature {
	result := []entity.ParkHasFeature{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model ca_parksdb.ParkHasFeature) entity.ParkHasFeature {
	return entity.ParkHasFeature{
		ID:        uuid.FromStringOrNil(model.ID),
		Details:   model.Details,
		Status:    entity.Status(model.Status),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		ParkID:    uuid.FromStringOrNil(model.ParkID),
		FeatureID: uuid.FromStringOrNil(model.FeatureID),
	}
}
