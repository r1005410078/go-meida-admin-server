package shared

type IEventBus interface {
	Dispatch(event any) error
	Register(handler any)
}
