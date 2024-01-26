package domain

import incomingports "myapp/internal/ports/incoming_ports"

type ProcessService struct {
	processUseCase incomingports.ProcessUseCase
}

func NewProcessService(processUseCase incomingports.ProcessUseCase) *ProcessService {
	return &ProcessService{
		processUseCase: processUseCase,
	}
}

func (s *ProcessService) Find(data string) bool {
	return s.processUseCase.FindData(data)
}

func (s *ProcessService) Create(data string) bool {
	return s.processUseCase.CreateData(data)
}
