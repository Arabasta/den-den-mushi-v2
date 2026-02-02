package handler

import (
	"den-den-mushi-v2/internal/protocol"
)

var (
	inputHandler  Handler = &Input{}
	resizeHandler Handler = &Resize{}
)

var Get = map[protocol.Header]Handler{
	protocol.Input:  inputHandler,
	protocol.Resize: resizeHandler,
}
