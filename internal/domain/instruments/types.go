package instruments

import "regexp"

type InstrumentId string

func (id InstrumentId) IsValid() bool {
	regexpPattern := regexp.MustCompile("^[A-Z0-9-]+$")
	return string(id) != "" && regexpPattern.MatchString(string(id))
}

type InstrumentType string

const (
	MetStation InstrumentType = "Met Station"
)

var InstrumentTypes = []InstrumentType{MetStation}

func (s InstrumentType) IsValid() bool {
	for _, instrumentType := range InstrumentTypes {
		if instrumentType == s {
			return true
		}
	}
	return false
}

type CommissionStatus string

const (
	Production CommissionStatus = "Production"
	Testing    CommissionStatus = "Testing"
	Historical CommissionStatus = "Historical"
)

var CommissionStatuses = []CommissionStatus{Production, Testing, Historical}

func (s CommissionStatus) IsValid() bool {
	for _, status := range CommissionStatuses {
		if status == s {
			return true
		}
	}
	return false
}

type Container string

type Location struct {
	Name string
	Lat  float64
	Long float64
}
