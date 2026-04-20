package converter

import (
	"github.com/halyph/go-service-blueprint/pkg/model"
	"github.com/halyph/go-service-blueprint/pkg/repository/entity"
)

// UserEntityConverter converts between UserEntity (DB) and User (domain model)
// goverter:converter
type UserEntityConverter interface {
	// EntityToModel converts database entity to domain model
	EntityToModel(source entity.UserEntity) model.User

	// EntityListToModelList converts list of entities to domain models
	EntityListToModelList(source []entity.UserEntity) []model.User

	// ModelToEntity converts domain model to database entity
	// goverter:ignore CreatedAt BaseModel
	ModelToEntity(source model.User) entity.UserEntity

	// ModelListToEntityList converts list of models to database entities
	ModelListToEntityList(source []model.User) []entity.UserEntity
}
