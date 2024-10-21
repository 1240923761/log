package hook

import "github.com/1240923761/log/entity"

type Hook interface {
	Process(e *entity.Entity) error
}
