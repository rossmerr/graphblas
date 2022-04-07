package constraints

type None interface {
	comparable
	Ordered | bool
}
