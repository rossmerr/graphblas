package constraints

type Logical interface {
	comparable
}

type Number interface {
	Logical
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

type String interface {
	Logical
}

type Ordered interface {
	Number | string
}

type Type interface {
	Number | bool | string
	String
	Logical
}
