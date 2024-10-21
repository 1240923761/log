package formatter

import "github.com/1240923761/log/entity"

type Formatter interface {
	Format(entity *entity.Entity) string
}
