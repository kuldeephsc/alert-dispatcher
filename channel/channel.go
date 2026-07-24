package channel

import (
	"github.com/kuldeephsc/model"
)

type Channel interface {
	Send(alert model.Alert) error
	Name() string
}
