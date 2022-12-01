package infra

import "github.com/ArthurMaverick/ez/src/domain/entity"

type GetIac interface {
	GetResource(entity.Template) (string, error)
	GetResources(entity.Template) (string, error)
}
