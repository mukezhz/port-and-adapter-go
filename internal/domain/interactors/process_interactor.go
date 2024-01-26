package interactors

import outgoingports "myapp/internal/ports/outgoing_ports"

type ProcessInteractor struct {
	processFinder  outgoingports.ProcessFinder
	processCreator outgoingports.ProcessCreator
}

func NewProcessInteractor(processFinder outgoingports.ProcessFinder, processCreator outgoingports.ProcessCreator) *ProcessInteractor {
	return &ProcessInteractor{
		processFinder:  processFinder,
		processCreator: processCreator,
	}
}

func (i *ProcessInteractor) FindData(data string) bool {
	return i.processFinder.Find(data)
}

func (i *ProcessInteractor) CreateData(data string) bool {
	return i.processCreator.Create(data)
}
