package core

type WorldTimeError struct {
	IsWorldTimeError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewWorldTimeError(code string, msg string, ctx *Context) *WorldTimeError {
	return &WorldTimeError{
		IsWorldTimeError: true,
		Sdk:              "WorldTime",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *WorldTimeError) Error() string {
	return e.Msg
}
