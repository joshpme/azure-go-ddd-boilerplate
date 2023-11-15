package instruments

type IInstrumentHandler interface {
	GetInstruments() ([]InstrumentAggregate, error)
}

func GetInstrumentList(handler IInstrumentHandler) ([]InstrumentAggregate, error) {
	instruments, err := handler.GetInstruments()
	if err != nil {
		return make([]InstrumentAggregate, 0), err
	}
	return instruments, nil
}
