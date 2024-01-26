package incomingports

type ProcessUseCase interface {
	FindData(data string) bool
	CreateData(data string) bool
}
