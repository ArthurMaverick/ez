package infra

import "github.com/ArthurMaverick/ez/src/domain/entity"

type DeleteIac interface {
	DeleteResource(entity.Template) (string, error)
	DeleteResources(entity.Template) (string, error)
}
