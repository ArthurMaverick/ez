package infra

import "github.com/ArthurMaverick/ez/package/domain/entity"

type GetIac interface {
	GetResource(entity.Template) (string, error)
	GetResources(entity.Template) (string, error)
}
