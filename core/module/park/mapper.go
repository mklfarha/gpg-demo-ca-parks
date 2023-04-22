package park

import (
	entity "ca_parks/core/entity/park"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []ca_parksdb.Park) []entity.Park {
	result := []entity.Park{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model ca_parksdb.Park) entity.Park {
	return entity.Park{
		ID:               uuid.FromStringOrNil(model.ID),
		Name:             model.Name,
		MainImage:        model.MainImage,
		Phone:            model.Phone,
		Hours:            model.Hours,
		AllowsDogs:       model.AllowsDogs,
		Links:            entity.LinksFromJSON(model.Links),
		Status:           entity.Status(model.Status),
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		RecreationAreaID: uuid.FromStringOrNil(model.RecreationAreaID),
		UserID:           uuid.FromStringOrNil(model.UserID),
	}
}
