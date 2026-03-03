package passkit

type TransitType string

const (
	TransitTypeAir     TransitType = "PKTransitTypeAir"
	TransitTypeBus     TransitType = "PKTransitTypeBus"
	TransitTypeTrain   TransitType = "PKTransitTypeTrain"
	TransitTypeBoat    TransitType = "PKTransitTypeBoat"
	TransitTypeGeneric TransitType = "PKTransitTypeGeneric"
)

type BoardingPass struct {
	TransitType TransitType `json:"transitType"`
	PassFields  *PassFields `json:"passFields,omitempty"`
}
