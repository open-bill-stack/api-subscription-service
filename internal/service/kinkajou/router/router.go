package router

import (
	"github.com/open-bill-stack/kinkajou"
)

type Router interface {
	Register(*kinkajou.App)
}
