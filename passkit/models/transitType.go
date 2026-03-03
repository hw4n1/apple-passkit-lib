package models

type TransitType string

const (
	TransitTypeAir     TransitType = "PKTransitTypeAir"
	TransitTypeBus     TransitType = "PKTransitTypeBus"
	TransitTypeTrain   TransitType = "PKTransitTypeTrain"
	TransitTypeBoat    TransitType = "PKTransitTypeBoat"
	TransitTypeGeneric TransitType = "PKTransitTypeGeneric"
)
