package user

import (
	entity "ca_parks/core/entity/user"
	ca_parksdb "ca_parks/core/repository/gen"
	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []ca_parksdb.User) []entity.User {
	result := []entity.User{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model ca_parksdb.User) entity.User {
	return entity.User{
		ID:        uuid.FromStringOrNil(model.ID),
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		Status:    entity.Status(model.Status),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
