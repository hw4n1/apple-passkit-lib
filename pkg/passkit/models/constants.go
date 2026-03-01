package passkit

type DataDetectorType string
type DateStyle string
type NumberStyle string
type TextAlignment string
type TimeStyle string
type BarcodeFormat string
type PassengerCapability string
type EventType string
type TransitSecurityProgram string

type DataDetectorTypes []DataDetectorType

const (
	FormatQR      BarcodeFormat = "PKBarcodeFormatQR"
	FormatPDF417  BarcodeFormat = "PKBarcodeFormatPDF417"
	FormatAztec   BarcodeFormat = "PKBarcodeFormatAztec"
	FormatCode128 BarcodeFormat = "PKBarcodeFormatCode128"
)

const (
	DataDetectorTypePhoneNumber   DataDetectorType = "PKDataDetectorTypePhoneNumber"
	DataDetectorTypeLink          DataDetectorType = "PKDataDetectorTypeLink"
	DataDetectorTypeAddress       DataDetectorType = "PKDataDetectorTypeAddress"
	DataDetectorTypeCalendarEvent DataDetectorType = "PKDataDetectorTypeCalendarEvent"
)

const (
	DateStyleNone   DateStyle = "PKDateStyleNone"
	DateStyleShort  DateStyle = "PKDateStyleShort"
	DateStyleMedium DateStyle = "PKDateStyleMedium"
	DateStyleLong   DateStyle = "PKDateStyleLong"
	DateStyleFull   DateStyle = "PKDateStyleFull"
)

const (
	NumberStyleDecimal    NumberStyle = "PKNumberStyleDecimal"
	NumberStylePercent    NumberStyle = "PKNumberStylePercent"
	NumberStyleScientific NumberStyle = "PKNumberStyleScientific"
	NumberStyleSpellOut   NumberStyle = "PKNumberStyleSpellOut"
)

const (
	TextAlignmentLeft    TextAlignment = "PKTextAlignmentLeft"
	TextAlignmentCenter  TextAlignment = "PKTextAlignmentCenter"
	TextAlignmentRight   TextAlignment = "PKTextAlignmentRight"
	TextAlignmentNatural TextAlignment = "PKTextAlignmentNatural"
)

const (
	TimeStyleNone   TimeStyle = "PKDateStyleNone"
	TimeStyleShort  TimeStyle = "PKDateStyleShort"
	TimeStyleMedium TimeStyle = "PKDateStyleMedium"
	TimeStyleLong   TimeStyle = "PKDateStyleLong"
	TimeStyleFull   TimeStyle = "PKDateStyleFull"
)

const (
	TransitSecurityProgramTSAPreCheck          TransitSecurityProgram = "PKTransitSecurityProgramTSAPreCheck"
	TransitSecurityProgramTSAPreCheckTouchless TransitSecurityProgram = "PKTransitSecurityProgramTSAPreCheckTouchlessID"
	TransitSecurityProgramOSS                  TransitSecurityProgram = "PKTransitSecurityProgramOSS"
	TransitSecurityProgramITI                  TransitSecurityProgram = "PKTransitSecurityProgramITI"
	TransitSecurityProgramITD                  TransitSecurityProgram = "PKTransitSecurityProgramITD"
	TransitSecurityProgramGlobalEntry          TransitSecurityProgram = "PKTransitSecurityProgramGlobalEntry"
	TransitSecurityProgramCLEAR                TransitSecurityProgram = "PKTransitSecurityProgramCLEAR"
)

const (
	EventTypeGeneric         EventType = "PKEventTypeGeneric"
	EventTypeLivePerformance EventType = "PKEventTypeLivePerformance"
	EventTypeMovie           EventType = "PKEventTypeMovie"
	EventTypeSports          EventType = "PKEventTypeSports"
	EventTypeConference      EventType = "PKEventTypeConference"
	EventTypeConvention      EventType = "PKEventTypeConvention"
	EventTypeWorkshop        EventType = "PKEventTypeWorkshop"
	EventTypeSocialGathering EventType = "PKEventTypeSocialGathering"
)

const (
	PassengerCapabilityPreboarding      PassengerCapability = "PKPassengerCapabilityPreboarding"
	PassengerCapabilityPriorityBoarding PassengerCapability = "PKPassengerCapabilityPriorityBoarding"
	PassengerCapabilityCarryon          PassengerCapability = "PKPassengerCapabilityCarryon"
	PassengerCapabilityPersonalItem     PassengerCapability = "PKPassengerCapabilityPersonalItem"
)
