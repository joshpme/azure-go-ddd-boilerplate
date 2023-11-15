package location

type ILocationHandler interface {
	GetLocationNames() (LocationNames, error)
}

func GetLocationNames(handler ILocationHandler) (LocationResponse, error) {
	locationNames, err := handler.GetLocationNames()
	if err != nil {
		return LocationResponse{}, err
	}
	locationResponse := LocationResponse{LocationNames: locationNames}
	return locationResponse, nil
}
