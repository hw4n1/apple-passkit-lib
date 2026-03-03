package models

type BoardingPass struct {
	TransitType TransitType `json:"transitType"`
	PassFields  *PassFields `json:"passFields,omitempty"`
}
