package passkit

// PassFieldsContent represents a single field in primaryFields, secondaryFields,
// auxiliaryFields, headerFields, backFields, or additionalInfoFields.

type PassFieldContent struct {

	// ---------------- REQUIRED ----------------
	Key   string      `json:"key"`
	Value interface{} `json:"value"`

	// ---------------- OPTIONAL ----------------
	Row               *int                 `json:"row,omitempty"`
	AttributedValue   *string              `json:"attributedValue,omitempty"`
	ChangeMessage     *string              `json:"changeMessage,omitempty"`
	CurrencyCode      *string              `json:"currencyCode,omitempty"`
	DataDetectorTypes *[]DataDetectorTypes `json:"dataDetectorTypes,omitempty"`
	DateStyle         *DateStyle           `json:"dateStyle,omitempty"`
	IgnoresTimeZone   *bool                `json:"ignoresTimeZone,omitempty"`
	IsRelative        *bool                `json:"isRelative,omitempty"`
	Label             *string              `json:"label,omitempty"`
	NumberStyle       *NumberStyle         `json:"numberStyle,omitempty"`
	TextAlignment     *TextAlignment       `json:"textAlignment,omitempty"`
	TimeStyle         *TimeStyle           `json:"timeStyle,omitempty"`
}
