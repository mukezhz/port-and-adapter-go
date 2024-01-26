package outgoingports

type ProcessFinder interface {
	Find(data string) bool
}

type ProcessCreator interface {
	Create(data string) bool
}
