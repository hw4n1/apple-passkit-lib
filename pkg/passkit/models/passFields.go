package passkit

type PassFields struct {
	AdditionalInfoFields *[]PassFieldContent `json:"additionalInfoFields,omitempty"`
	AuxiliaryFields      *[]PassFieldContent `json:"auxiliaryFields,omitempty"`
	BackFields           *[]PassFieldContent `json:"backFields,omitempty"`
	HeaderFields         *[]PassFieldContent `json:"headerFields,omitempty"`
	PrimaryFields        *[]PassFieldContent `json:"primaryFields,omitempty"`
	SecondaryFields      *[]PassFieldContent `json:"secondaryFields,omitempty"`
}
