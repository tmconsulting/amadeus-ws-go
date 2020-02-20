// generated by goxsd; DO NOT EDIT

package Fare_InformativePricingWithoutPNRReply_v19_1 // tipnrr191

// Fare_InformativePricingWithoutPNRReply is generated from an XSD element
type FareInformativePricingWithoutPNRReply struct {
	MessageDetails messageDetails `xml:"messageDetails"`
	ErrorGroup     *errorGroup    `xml:"errorGroup"`
	MainGroup      *mainGroup     `xml:"mainGroup"`
}

// messageDetails is generated from an XSD element
type messageDetails struct {
	MessageFunctionDetails messageFunctionDetails `xml:"messageFunctionDetails"`
	ResponseType           string                 `xml:"responseType"`
}

// messageFunctionDetails is generated from an XSD element
type messageFunctionDetails struct {
	BusinessFunction          string `xml:"businessFunction"`
	MessageFunction           string `xml:"messageFunction"`
	ResponsibleAgency         string `xml:"responsibleAgency"`
	AdditionalMessageFunction string `xml:"additionalMessageFunction"`
}

// errorGroup is generated from an XSD element
type errorGroup struct {
	ErrorOrWarningCodeDetails errorOrWarningCodeDetails `xml:"errorOrWarningCodeDetails"`
	ErrorWarningDescription   errorWarningDescription   `xml:"errorWarningDescription"`
}

// errorOrWarningCodeDetails is generated from an XSD element
type errorOrWarningCodeDetails struct {
	ErrorDetails errorDetails `xml:"errorDetails"`
}

// errorDetails is generated from an XSD element
type errorDetails struct {
	ErrorCode      string `xml:"errorCode"`
	ErrorCategory  string `xml:"errorCategory"`
	ErrorCodeOwner string `xml:"errorCodeOwner"`
}

// errorWarningDescription is generated from an XSD element
type errorWarningDescription struct {
	FreeTextDetails freeTextDetails `xml:"freeTextDetails"`
	FreeText        string          `xml:"freeText"`
}

// freeTextDetails is generated from an XSD element
type freeTextDetails struct {
	TextSubjectQualifier string `xml:"textSubjectQualifier"`
	InformationType      string `xml:"informationType"`
	Status               string `xml:"status"`
	CompanyID            string `xml:"companyId"`
	Language             string `xml:"language"`
	Source               string `xml:"source"`
	Encoding             string `xml:"encoding"`
}

// mainGroup is generated from an XSD element
type mainGroup struct {
	DummySegment           dummySegment             `xml:"dummySegment"`
	ConvertionRate         convertionRate           `xml:"convertionRate"`
	GeneralIndicatorsGroup generalIndicatorsGroup   `xml:"generalIndicatorsGroup"`
	PricingGroupLevelGroup []PricingGroupLevelGroup `xml:"PricingGroupLevelGroup"`
}

// dummySegment is generated from an XSD element
type dummySegment struct {
}

// convertionRate is generated from an XSD element
type convertionRate struct {
	ConversionRateDetails conversionRateDetails `xml:"conversionRateDetails"`
	OtherConvRateDetails  otherConvRateDetails  `xml:"otherConvRateDetails"`
}

// conversionRateDetails is generated from an XSD element
type conversionRateDetails struct {
	ConversionType          string  `xml:"conversionType"`
	Currency                string  `xml:"currency"`
	RateType                string  `xml:"rateType"`
	PricingAmount           float64 `xml:"pricingAmount"`
	ConvertedValueAmount    int     `xml:"convertedValueAmount"`
	DutyTaxFeeType          string  `xml:"dutyTaxFeeType"`
	MeasurementValue        int     `xml:"measurementValue"`
	MeasurementSignificance string  `xml:"measurementSignificance"`
}

// otherConvRateDetails is generated from an XSD element
type otherConvRateDetails struct {
	ConversionType          string  `xml:"conversionType"`
	Currency                string  `xml:"currency"`
	RateType                string  `xml:"rateType"`
	PricingAmount           float64 `xml:"pricingAmount"`
	ConvertedValueAmount    int     `xml:"convertedValueAmount"`
	DutyTaxFeeType          string  `xml:"dutyTaxFeeType"`
	MeasurementValue        int     `xml:"measurementValue"`
	MeasurementSignificance string  `xml:"measurementSignificance"`
}

// generalIndicatorsGroup is generated from an XSD element
type generalIndicatorsGroup struct {
	GeneralIndicators generalIndicators `xml:"generalIndicators"`
}

// generalIndicators is generated from an XSD element
type generalIndicators struct {
	PriceTicketDetails     priceTicketDetails     `xml:"priceTicketDetails"`
	PriceTariffType        string                 `xml:"priceTariffType"`
	ProductDateTimeDetails productDateTimeDetails `xml:"productDateTimeDetails"`
	CompanyDetails         companyDetails         `xml:"companyDetails"`
	CompanyNumberDetails   companyNumberDetails   `xml:"companyNumberDetails"`
	LocationDetails        locationDetails        `xml:"locationDetails"`
	OtherLocationDetails   otherLocationDetails   `xml:"otherLocationDetails"`
	IDNumber               string                 `xml:"idNumber"`
	MonetaryAmount         float64                `xml:"monetaryAmount"`
}

// priceTicketDetails is generated from an XSD element
type priceTicketDetails struct {
	Indicators string `xml:"indicators"`
}

// productDateTimeDetails is generated from an XSD element
type productDateTimeDetails struct {
	DepartureDate string `xml:"departureDate"`
	DepartureTime *int   `xml:"departureTime"`
	ArrivalDate   string `xml:"arrivalDate"`
	ArrivalTime   int    `xml:"arrivalTime"`
	DateVariation int    `xml:"dateVariation"`
}

// companyDetails is generated from an XSD element
type companyDetails struct {
	MarketingCompany string `xml:"marketingCompany"`
	OperatingCompany string `xml:"operatingCompany"`
	OtherCompany     string `xml:"otherCompany"`
}

// companyNumberDetails is generated from an XSD element
type companyNumberDetails struct {
	IDentifier      string `xml:"identifier"`
	OtherIDentifier string `xml:"otherIdentifier"`
}

// locationDetails is generated from an XSD element
type locationDetails struct {
	City    string `xml:"city"`
	Country string `xml:"country"`
}

// otherLocationDetails is generated from an XSD element
type otherLocationDetails struct {
	City    string `xml:"city"`
	Country string `xml:"country"`
}

// PricingGroupLevelGroup is generated from an XSD element
type PricingGroupLevelGroup struct {
	NumberOfPax   numberOfPax   `xml:"numberOfPax"`
	PassengersID  passengersID  `xml:"passengersID"`
	FareInfoGroup fareInfoGroup `xml:"fareInfoGroup"`
}

// numberOfPax is generated from an XSD element
type numberOfPax struct {
	SegmentControlDetails segmentControlDetails `xml:"segmentControlDetails"`
}

// segmentControlDetails is generated from an XSD element
type segmentControlDetails struct {
	Quantity           int `xml:"quantity"`
	NumberOfUnits      int `xml:"numberOfUnits"`
	TotalNumberOfItems int `xml:"totalNumberOfItems"`
}

// passengersID is generated from an XSD element
type passengersID struct {
	TravellerDetails []travellerDetails `xml:"travellerDetails"`
	Dummy_NET        Dummy_NET          `xml:"Dummy_NET"`
}

// travellerDetails is generated from an XSD element
type travellerDetails struct {
	ReferenceNumber  string `xml:"referenceNumber"`
	MeasurementValue int    `xml:"measurementValue"`
	FirstDate        string `xml:"firstDate"`
	Surname          string `xml:"surname"`
	FirstName        string `xml:"firstName"`
}

// Dummy_NET is generated from an XSD element
type Dummy_NET struct {
}

// fareInfoGroup is generated from an XSD element
type fareInfoGroup struct {
	EmptySegment              emptySegment                `xml:"emptySegment"`
	PricingIndicators         *pricingIndicators          `xml:"pricingIndicators"`
	FareAmount                fareAmount                  `xml:"fareAmount"`
	TextData                  []textData                  `xml:"textData"`
	OfferReferences           offerReferences             `xml:"offerReferences"`
	SurchargesGroup           surchargesGroup             `xml:"surchargesGroup"`
	CorporateGroup            corporateGroup              `xml:"corporateGroup"`
	NegoFareGroup             negoFareGroup               `xml:"negoFareGroup"`
	SegmentLevelGroup         []segmentLevelGroup         `xml:"segmentLevelGroup"`
	StructuredFareCalcGroup   structuredFareCalcGroup     `xml:"structuredFareCalcGroup"`
	CarrierFeeGroup           carrierFeeGroup             `xml:"carrierFeeGroup"`
	FareComponentDetailsGroup []fareComponentDetailsGroup `xml:"fareComponentDetailsGroup"`
}

// emptySegment is generated from an XSD element
type emptySegment struct {
	ValueQualifier   string           `xml:"valueQualifier"`
	Value            int              `xml:"value"`
	FareDetails      fareDetails      `xml:"fareDetails"`
	IDentityNumber   string           `xml:"identityNumber"`
	FareTypeGrouping fareTypeGrouping `xml:"fareTypeGrouping"`
	RateCategory     string           `xml:"rateCategory"`
}

// fareDetails is generated from an XSD element
type fareDetails struct {
	Qualifier    string  `xml:"qualifier"`
	Rate         float64 `xml:"rate"`
	Country      string  `xml:"country"`
	FareCategory string  `xml:"fareCategory"`
}

// fareTypeGrouping is generated from an XSD element
type fareTypeGrouping struct {
	PricingGroup string `xml:"pricingGroup"`
}

// pricingIndicators is generated from an XSD element
type pricingIndicators struct {
	PriceTicketDetails     priceTicketDetails      `xml:"priceTicketDetails"`
	PriceTariffType        string                  `xml:"priceTariffType"`
	ProductDateTimeDetails *productDateTimeDetails `xml:"productDateTimeDetails"`
	CompanyDetails         companyDetails          `xml:"companyDetails"`
	CompanyNumberDetails   companyNumberDetails    `xml:"companyNumberDetails"`
	LocationDetails        locationDetails         `xml:"locationDetails"`
	OtherLocationDetails   otherLocationDetails    `xml:"otherLocationDetails"`
	IDNumber               string                  `xml:"idNumber"`
	MonetaryAmount         float64                 `xml:"monetaryAmount"`
}

// fareAmount is generated from an XSD element
type fareAmount struct {
	MonetaryDetails      monetaryDetails        `xml:"monetaryDetails"`
	OtherMonetaryDetails []otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// monetaryDetails is generated from an XSD element
type monetaryDetails struct {
	TypeQualifier string `xml:"typeQualifier"`
	Amount        string `xml:"amount"`
	Currency      string `xml:"currency"`
	Location      string `xml:"location"`
}

// otherMonetaryDetails is generated from an XSD element
type otherMonetaryDetails struct {
	TypeQualifier string `xml:"typeQualifier"`
	Amount        string `xml:"amount"`
	Currency      string `xml:"currency"`
	Location      string `xml:"location"`
}

// textData is generated from an XSD element
type textData struct {
	FreeTextQualification *freeTextQualification `xml:"freeTextQualification"`
	FreeText              []string               `xml:"freeText"`
}

// freeTextQualification is generated from an XSD element
type freeTextQualification struct {
	TextSubjectQualifier string `xml:"textSubjectQualifier"`
	InformationType      string `xml:"informationType"`
	Status               string `xml:"status"`
	CompanyID            string `xml:"companyId"`
	Language             string `xml:"language"`
}

// offerReferences is generated from an XSD element
type offerReferences struct {
	OfferIDentifier offerIDentifier `xml:"offerIdentifier"`
	References      references      `xml:"references"`
}

// offerIDentifier is generated from an XSD element
type offerIDentifier struct {
	Reference            string `xml:"reference"`
	OfferID              string `xml:"offerId"`
	UniqueOfferReference string `xml:"uniqueOfferReference"`
}

// references is generated from an XSD element
type references struct {
	ReferenceDetails referenceDetails `xml:"referenceDetails"`
	Dummy_NET        Dummy_NET        `xml:"Dummy_NET"`
}

// referenceDetails is generated from an XSD element
type referenceDetails struct {
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

// surchargesGroup is generated from an XSD element
type surchargesGroup struct {
	TaxesAmount     taxesAmount      `xml:"taxesAmount"`
	PenaltiesAmount *penaltiesAmount `xml:"penaltiesAmount"`
	PfcAmount       pfcAmount        `xml:"pfcAmount"`
}

// taxesAmount is generated from an XSD element
type taxesAmount struct {
	TaxCategory string       `xml:"taxCategory"`
	TaxDetails  []taxDetails `xml:"taxDetails"`
}

// taxDetails is generated from an XSD element
type taxDetails struct {
	Rate         string `xml:"rate"`
	CountryCode  string `xml:"countryCode"`
	CurrencyCode string `xml:"currencyCode"`
	Type         string `xml:"type"`
}

// penaltiesAmount is generated from an XSD element
type penaltiesAmount struct {
	DiscountPenaltyQualifier string                   `xml:"discountPenaltyQualifier"`
	DiscountPenaltyDetails   []discountPenaltyDetails `xml:"discountPenaltyDetails"`
}

// discountPenaltyDetails is generated from an XSD element
type discountPenaltyDetails struct {
	Function   string   `xml:"function"`
	AmountType string   `xml:"amountType"`
	Amount     *float64 `xml:"amount"`
	Rate       string   `xml:"rate"`
	Currency   string   `xml:"currency"`
}

// pfcAmount is generated from an XSD element
type pfcAmount struct {
	MonetaryDetails      monetaryDetails      `xml:"monetaryDetails"`
	OtherMonetaryDetails otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// corporateGroup is generated from an XSD element
type corporateGroup struct {
	CorporateData corporateData `xml:"corporateData"`
}

// corporateData is generated from an XSD element
type corporateData struct {
	ChargeCategory    string  `xml:"chargeCategory"`
	Amount            float64 `xml:"amount"`
	LocationCode      string  `xml:"locationCode"`
	OtherLocationCode string  `xml:"otherLocationCode"`
	Rate              float64 `xml:"rate"`
}

// negoFareGroup is generated from an XSD element
type negoFareGroup struct {
	NegoFareIndicators    negoFareIndicators    `xml:"negoFareIndicators"`
	ExtNegoFareIndicators extNegoFareIndicators `xml:"extNegoFareIndicators"`
	NegoFareAmount        negoFareAmount        `xml:"negoFareAmount"`
	NegoFareText          negoFareText          `xml:"negoFareText"`
}

// negoFareIndicators is generated from an XSD element
type negoFareIndicators struct {
	ItemNumber                 string           `xml:"itemNumber"`
	FareBasisDetails           fareBasisDetails `xml:"fareBasisDetails"`
	FareValue                  int              `xml:"fareValue"`
	PriceType                  string           `xml:"priceType"`
	SpecialCondition           string           `xml:"specialCondition"`
	OtherSpecialCondition      string           `xml:"otherSpecialCondition"`
	AdditionalSpecialCondition string           `xml:"additionalSpecialCondition"`
	TaxCategory                string           `xml:"taxCategory"`
}

// fareBasisDetails is generated from an XSD element
type fareBasisDetails struct {
	RateTariffClass          string `xml:"rateTariffClass"`
	RateTariffIndicator      string `xml:"rateTariffIndicator"`
	OtherRateTariffClass     string `xml:"otherRateTariffClass"`
	OtherRateTariffIndicator string `xml:"otherRateTariffIndicator"`
}

// extNegoFareIndicators is generated from an XSD element
type extNegoFareIndicators struct {
	MovementType          string                `xml:"movementType"`
	FareCategories        fareCategories        `xml:"fareCategories"`
	FareDetails           fareDetails           `xml:"fareDetails"`
	AdditionalFareDetails additionalFareDetails `xml:"additionalFareDetails"`
	DiscountDetails       discountDetails       `xml:"discountDetails"`
}

// fareCategories is generated from an XSD element
type fareCategories struct {
	FareType      string `xml:"fareType"`
	OtherFareType string `xml:"otherFareType"`
}

// additionalFareDetails is generated from an XSD element
type additionalFareDetails struct {
	RateClass         string `xml:"rateClass"`
	CommodityCategory string `xml:"commodityCategory"`
	PricingGroup      string `xml:"pricingGroup"`
	SecondRateClass   string `xml:"secondRateClass"`
}

// discountDetails is generated from an XSD element
type discountDetails struct {
	FareQualifier string  `xml:"fareQualifier"`
	RateCategory  string  `xml:"rateCategory"`
	Amount        float64 `xml:"amount"`
	Percentage    int     `xml:"percentage"`
}

// negoFareAmount is generated from an XSD element
type negoFareAmount struct {
	DiscountPenaltyQualifier string                 `xml:"discountPenaltyQualifier"`
	DiscountPenaltyDetails   discountPenaltyDetails `xml:"discountPenaltyDetails"`
}

// negoFareText is generated from an XSD element
type negoFareText struct {
	FreeTextQualification freeTextQualification `xml:"freeTextQualification"`
	FreeText              string                `xml:"freeText"`
}

// segmentLevelGroup is generated from an XSD element
type segmentLevelGroup struct {
	SegmentInformation    segmentInformation    `xml:"segmentInformation"`
	AdditionalInformation additionalInformation `xml:"additionalInformation"`
	FareBasis             *fareBasis            `xml:"fareBasis"`
	CabinGroup            []cabinGroup          `xml:"cabinGroup"`
	BaggageAllowance      baggageAllowance      `xml:"baggageAllowance"`
	PtcSegment            ptcSegment            `xml:"ptcSegment"`
	CouponInformation     couponInformation     `xml:"couponInformation"`
}

// segmentInformation is generated from an XSD element
type segmentInformation struct {
	FlightDate           flightDate           `xml:"flightDate"`
	BoardPointDetails    boardPointDetails    `xml:"boardPointDetails"`
	OffpointDetails      offpointDetails      `xml:"offpointDetails"`
	CompanyDetails       companyDetails       `xml:"companyDetails"`
	FlightIDentification flightIDentification `xml:"flightIdentification"`
	FlightTypeDetails    flightTypeDetails    `xml:"flightTypeDetails"`
	ItemNumber           *int                 `xml:"itemNumber"`
	SpecialSegment       string               `xml:"specialSegment"`
	MarriageDetails      marriageDetails      `xml:"marriageDetails"`
}

// flightDate is generated from an XSD element
type flightDate struct {
	DepartureDate string `xml:"departureDate"`
	DepartureTime int    `xml:"departureTime"`
	ArrivalDate   string `xml:"arrivalDate"`
	ArrivalTime   int    `xml:"arrivalTime"`
	DateVariation int    `xml:"dateVariation"`
}

// boardPointDetails is generated from an XSD element
type boardPointDetails struct {
	TrueLocationID string `xml:"trueLocationId"`
	TrueLocation   string `xml:"trueLocation"`
}

// offpointDetails is generated from an XSD element
type offpointDetails struct {
	TrueLocationID string `xml:"trueLocationId"`
	TrueLocation   string `xml:"trueLocation"`
}

// flightIDentification is generated from an XSD element
type flightIDentification struct {
	FlightNumber      string `xml:"flightNumber"`
	BookingClass      string `xml:"bookingClass"`
	OperationalSuffix string `xml:"operationalSuffix"`
	Modifier          string `xml:"modifier"`
}

// flightTypeDetails is generated from an XSD element
type flightTypeDetails struct {
	FlightIndicator string `xml:"flightIndicator"`
}

// marriageDetails is generated from an XSD element
type marriageDetails struct {
	Relation           string `xml:"relation"`
	MarriageIDentifier int    `xml:"marriageIdentifier"`
	LineNumber         int    `xml:"lineNumber"`
	OtherRelation      string `xml:"otherRelation"`
	CarrierCode        string `xml:"carrierCode"`
}

// additionalInformation is generated from an XSD element
type additionalInformation struct {
	PriceTicketDetails     priceTicketDetails     `xml:"priceTicketDetails"`
	PriceTariffType        string                 `xml:"priceTariffType"`
	ProductDateTimeDetails productDateTimeDetails `xml:"productDateTimeDetails"`
	CompanyDetails         companyDetails         `xml:"companyDetails"`
	CompanyNumberDetails   companyNumberDetails   `xml:"companyNumberDetails"`
	LocationDetails        locationDetails        `xml:"locationDetails"`
	OtherLocationDetails   otherLocationDetails   `xml:"otherLocationDetails"`
	IDNumber               string                 `xml:"idNumber"`
	MonetaryAmount         float64                `xml:"monetaryAmount"`
}

// fareBasis is generated from an XSD element
type fareBasis struct {
	MovementType          string                 `xml:"movementType"`
	FareCategories        fareCategories         `xml:"fareCategories"`
	FareDetails           fareDetails            `xml:"fareDetails"`
	AdditionalFareDetails *additionalFareDetails `xml:"additionalFareDetails"`
	DiscountDetails       discountDetails        `xml:"discountDetails"`
}

// cabinGroup is generated from an XSD element
type cabinGroup struct {
	CabinSegment cabinSegment `xml:"cabinSegment"`
}

// cabinSegment is generated from an XSD element
type cabinSegment struct {
	ProductDetailsQualifier string                `xml:"productDetailsQualifier"`
	BookingClassDetails     []bookingClassDetails `xml:"bookingClassDetails"`
}

// bookingClassDetails is generated from an XSD element
type bookingClassDetails struct {
	Designator         string   `xml:"designator"`
	AvailabilityStatus string   `xml:"availabilityStatus"`
	SpecialService     string   `xml:"specialService"`
	Option             []string `xml:"option"`
}

// baggageAllowance is generated from an XSD element
type baggageAllowance struct {
	ExcessBaggageDetails excessBaggageDetails `xml:"excessBaggageDetails"`
	BaggageDetails       baggageDetails       `xml:"baggageDetails"`
	OtherBaggageDetails  otherBaggageDetails  `xml:"otherBaggageDetails"`
	ExtraBaggageDetails  extraBaggageDetails  `xml:"extraBaggageDetails"`
	BagTagDetails        bagTagDetails        `xml:"bagTagDetails"`
}

// excessBaggageDetails is generated from an XSD element
type excessBaggageDetails struct {
	Currency         string  `xml:"currency"`
	Amount           float64 `xml:"amount"`
	ProcessIndicator string  `xml:"processIndicator"`
}

// baggageDetails is generated from an XSD element
type baggageDetails struct {
	FreeAllowance    int    `xml:"freeAllowance"`
	Measurement      int    `xml:"measurement"`
	QuantityCode     string `xml:"quantityCode"`
	UnitQualifier    string `xml:"unitQualifier"`
	ProcessIndicator string `xml:"processIndicator"`
}

// otherBaggageDetails is generated from an XSD element
type otherBaggageDetails struct {
	FreeAllowance    int    `xml:"freeAllowance"`
	Measurement      int    `xml:"measurement"`
	QuantityCode     string `xml:"quantityCode"`
	UnitQualifier    string `xml:"unitQualifier"`
	ProcessIndicator string `xml:"processIndicator"`
}

// extraBaggageDetails is generated from an XSD element
type extraBaggageDetails struct {
	FreeAllowance    int    `xml:"freeAllowance"`
	Measurement      int    `xml:"measurement"`
	QuantityCode     string `xml:"quantityCode"`
	UnitQualifier    string `xml:"unitQualifier"`
	ProcessIndicator string `xml:"processIndicator"`
}

// bagTagDetails is generated from an XSD element
type bagTagDetails struct {
	Company            string `xml:"company"`
	IDentifier         string `xml:"identifier"`
	Number             int    `xml:"number"`
	Location           string `xml:"location"`
	CompanyNumber      string `xml:"companyNumber"`
	Indicator          string `xml:"indicator"`
	Characteristic     string `xml:"characteristic"`
	SpecialRequirement string `xml:"specialRequirement"`
	Measurement        int    `xml:"measurement"`
	UnitQualifier      string `xml:"unitQualifier"`
	Description        string `xml:"description"`
}

// ptcSegment is generated from an XSD element
type ptcSegment struct {
	QuantityDetails      quantityDetails      `xml:"quantityDetails"`
	OtherQuantityDetails otherQuantityDetails `xml:"otherQuantityDetails"`
}

// quantityDetails is generated from an XSD element
type quantityDetails struct {
	NumberOfUnit  *int   `xml:"numberOfUnit"`
	UnitQualifier string `xml:"unitQualifier"`
}

// otherQuantityDetails is generated from an XSD element
type otherQuantityDetails struct {
	NumberOfUnit  int    `xml:"numberOfUnit"`
	UnitQualifier string `xml:"unitQualifier"`
}

// couponInformation is generated from an XSD element
type couponInformation struct {
	QuantityDetails      quantityDetails      `xml:"quantityDetails"`
	OtherquantityDetails otherquantityDetails `xml:"otherquantityDetails"`
}

// otherquantityDetails is generated from an XSD element
type otherquantityDetails struct {
	Qualifier string `xml:"qualifier"`
	Value     int    `xml:"value"`
	Unit      string `xml:"unit"`
}

// structuredFareCalcGroup is generated from an XSD element
type structuredFareCalcGroup struct {
	StructureFareCalcRoot structureFareCalcRoot `xml:"structureFareCalcRoot"`
	Group27               group27               `xml:"group27"`
}

// structureFareCalcRoot is generated from an XSD element
type structureFareCalcRoot struct {
	FareComponentDetails fareComponentDetails `xml:"fareComponentDetails"`
	TicketNumber         string               `xml:"ticketNumber"`
}

// fareComponentDetails is generated from an XSD element
type fareComponentDetails struct {
	DataType        string `xml:"dataType"`
	Count           int    `xml:"count"`
	PricingDate     string `xml:"pricingDate"`
	AccountCode     string `xml:"accountCode"`
	InputDesignator string `xml:"inputDesignator"`
}

// group27 is generated from an XSD element
type group27 struct {
	StructuredFareCalcG27EQN structuredFareCalcG27EQN `xml:"structuredFareCalcG27EQN"`
	Group28                  group28                  `xml:"group28"`
	DummySegmentGroup27      dummySegmentGroup27      `xml:"dummySegmentGroup27"`
	StructuredFareCalcG27MON structuredFareCalcG27MON `xml:"structuredFareCalcG27MON"`
	StructuredFareCalcG27TXD structuredFareCalcG27TXD `xml:"structuredFareCalcG27TXD"`
	StructuredFareCalcG27CVR structuredFareCalcG27CVR `xml:"structuredFareCalcG27CVR"`
}

// structuredFareCalcG27EQN is generated from an XSD element
type structuredFareCalcG27EQN struct {
	QuantityDetails      quantityDetails      `xml:"quantityDetails"`
	OtherQuantityDetails otherQuantityDetails `xml:"otherQuantityDetails"`
}

// group28 is generated from an XSD element
type group28 struct {
	StructuredFareCalcG28ITM structuredFareCalcG28ITM `xml:"structuredFareCalcG28ITM"`
	Group29                  group29                  `xml:"group29"`
	StructuredFareCalcG28MON structuredFareCalcG28MON `xml:"structuredFareCalcG28MON"`
	StructuredFareCalcG28PTS structuredFareCalcG28PTS `xml:"structuredFareCalcG28PTS"`
	StructuredFareCalcG28FCC structuredFareCalcG28FCC `xml:"structuredFareCalcG28FCC"`
	StructuredFareCalcG28PTK structuredFareCalcG28PTK `xml:"structuredFareCalcG28PTK"`
	StructuredFareCalcG28FRU structuredFareCalcG28FRU `xml:"structuredFareCalcG28FRU"`
}

// structuredFareCalcG28ITM is generated from an XSD element
type structuredFareCalcG28ITM struct {
	ItemNumberDetails itemNumberDetails `xml:"itemNumberDetails"`
}

// itemNumberDetails is generated from an XSD element
type itemNumberDetails struct {
	Number            string `xml:"number"`
	Type              string `xml:"type"`
	Qualifier         string `xml:"qualifier"`
	ResponsibleAgency string `xml:"responsibleAgency"`
}

// group29 is generated from an XSD element
type group29 struct {
	StructuredFareCalcG28ADT structuredFareCalcG28ADT `xml:"structuredFareCalcG28ADT"`
	StructuredFareCalcG28TVL structuredFareCalcG28TVL `xml:"structuredFareCalcG28TVL"`
}

// structuredFareCalcG28ADT is generated from an XSD element
type structuredFareCalcG28ADT struct {
	NumberOfItemsDetails numberOfItemsDetails `xml:"numberOfItemsDetails"`
	LastItemsDetails     lastItemsDetails     `xml:"lastItemsDetails"`
}

// numberOfItemsDetails is generated from an XSD element
type numberOfItemsDetails struct {
	ActionQualifier    string `xml:"actionQualifier"`
	ReferenceQualifier string `xml:"referenceQualifier"`
	NumberOfItems      string `xml:"numberOfItems"`
}

// lastItemsDetails is generated from an XSD element
type lastItemsDetails struct {
	NumberOfItems      string `xml:"numberOfItems"`
	LastItemIDentifier string `xml:"lastItemIdentifier"`
}

// structuredFareCalcG28TVL is generated from an XSD element
type structuredFareCalcG28TVL struct {
	FlightDate           flightDate           `xml:"flightDate"`
	BoardPointDetails    boardPointDetails    `xml:"boardPointDetails"`
	OffpointDetails      offpointDetails      `xml:"offpointDetails"`
	CompanyDetails       companyDetails       `xml:"companyDetails"`
	FlightIDentification flightIDentification `xml:"flightIdentification"`
	FlightTypeDetails    flightTypeDetails    `xml:"flightTypeDetails"`
	ItemNumber           int                  `xml:"itemNumber"`
	SpecialSegment       string               `xml:"specialSegment"`
	MarriageDetails      marriageDetails      `xml:"marriageDetails"`
}

// structuredFareCalcG28MON is generated from an XSD element
type structuredFareCalcG28MON struct {
	MonetaryDetails      monetaryDetails      `xml:"monetaryDetails"`
	OtherMonetaryDetails otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// structuredFareCalcG28PTS is generated from an XSD element
type structuredFareCalcG28PTS struct {
	ItemNumber                 string           `xml:"itemNumber"`
	FareBasisDetails           fareBasisDetails `xml:"fareBasisDetails"`
	FareValue                  int              `xml:"fareValue"`
	PriceType                  string           `xml:"priceType"`
	SpecialCondition           string           `xml:"specialCondition"`
	OtherSpecialCondition      string           `xml:"otherSpecialCondition"`
	AdditionalSpecialCondition string           `xml:"additionalSpecialCondition"`
	TaxCategory                string           `xml:"taxCategory"`
}

// structuredFareCalcG28FCC is generated from an XSD element
type structuredFareCalcG28FCC struct {
	ChargeCategory    string  `xml:"chargeCategory"`
	Amount            float64 `xml:"amount"`
	LocationCode      string  `xml:"locationCode"`
	OtherLocationCode string  `xml:"otherLocationCode"`
	Rate              float64 `xml:"rate"`
}

// structuredFareCalcG28PTK is generated from an XSD element
type structuredFareCalcG28PTK struct {
	PriceTicketDetails     priceTicketDetails     `xml:"priceTicketDetails"`
	PriceTariffType        string                 `xml:"priceTariffType"`
	ProductDateTimeDetails productDateTimeDetails `xml:"productDateTimeDetails"`
	CompanyDetails         companyDetails         `xml:"companyDetails"`
	CompanyNumberDetails   companyNumberDetails   `xml:"companyNumberDetails"`
	LocationDetails        locationDetails        `xml:"locationDetails"`
	OtherLocationDetails   otherLocationDetails   `xml:"otherLocationDetails"`
	IDNumber               string                 `xml:"idNumber"`
	MonetaryAmount         float64                `xml:"monetaryAmount"`
}

// structuredFareCalcG28FRU is generated from an XSD element
type structuredFareCalcG28FRU struct {
	TariffClassID  string         `xml:"tariffClassId"`
	CompanyDetails companyDetails `xml:"companyDetails"`
	RuleSectionID  string         `xml:"ruleSectionId"`
}

// dummySegmentGroup27 is generated from an XSD element
type dummySegmentGroup27 struct {
}

// structuredFareCalcG27MON is generated from an XSD element
type structuredFareCalcG27MON struct {
	MonetaryDetails      monetaryDetails      `xml:"monetaryDetails"`
	OtherMonetaryDetails otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// structuredFareCalcG27TXD is generated from an XSD element
type structuredFareCalcG27TXD struct {
	TaxCategory string     `xml:"taxCategory"`
	TaxDetails  taxDetails `xml:"taxDetails"`
}

// structuredFareCalcG27CVR is generated from an XSD element
type structuredFareCalcG27CVR struct {
	ConversionRateDetails conversionRateDetails `xml:"conversionRateDetails"`
	OtherConvRateDetails  otherConvRateDetails  `xml:"otherConvRateDetails"`
}

// carrierFeeGroup is generated from an XSD element
type carrierFeeGroup struct {
	FeeType    feeType    `xml:"feeType"`
	FeeDetails feeDetails `xml:"feeDetails"`
}

// feeType is generated from an XSD element
type feeType struct {
	SelectionDetails      selectionDetails      `xml:"selectionDetails"`
	OtherSelectionDetails otherSelectionDetails `xml:"otherSelectionDetails"`
}

// selectionDetails is generated from an XSD element
type selectionDetails struct {
	Option            string `xml:"option"`
	OptionInformation string `xml:"optionInformation"`
}

// otherSelectionDetails is generated from an XSD element
type otherSelectionDetails struct {
	Option            string `xml:"option"`
	OptionInformation string `xml:"optionInformation"`
}

// feeDetails is generated from an XSD element
type feeDetails struct {
	FeeInfo        feeInfo        `xml:"feeInfo"`
	FeeAmounts     feeAmounts     `xml:"feeAmounts"`
	FeeTaxes       feeTaxes       `xml:"feeTaxes"`
	FeeDescription feeDescription `xml:"feeDescription"`
}

// feeInfo is generated from an XSD element
type feeInfo struct {
	DataTypeInformation dataTypeInformation `xml:"dataTypeInformation"`
	DataInformation     dataInformation     `xml:"dataInformation"`
}

// dataTypeInformation is generated from an XSD element
type dataTypeInformation struct {
	Type        string `xml:"type"`
	StatusEvent string `xml:"statusEvent"`
}

// dataInformation is generated from an XSD element
type dataInformation struct {
	Indicator string `xml:"indicator"`
	Value     int    `xml:"value"`
	Unit      string `xml:"unit"`
}

// feeAmounts is generated from an XSD element
type feeAmounts struct {
	MonetaryDetails      monetaryDetails      `xml:"monetaryDetails"`
	OtherMonetaryDetails otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// feeTaxes is generated from an XSD element
type feeTaxes struct {
	TaxCategory string     `xml:"taxCategory"`
	TaxDetails  taxDetails `xml:"taxDetails"`
}

// feeDescription is generated from an XSD element
type feeDescription struct {
	FreeTextQualification freeTextQualification `xml:"freeTextQualification"`
	FreeText              string                `xml:"freeText"`
}

// fareComponentDetailsGroup is generated from an XSD element
type fareComponentDetailsGroup struct {
	FareComponentID      fareComponentID      `xml:"fareComponentID"`
	MarketFareComponent  marketFareComponent  `xml:"marketFareComponent"`
	MonetaryInformation  monetaryInformation  `xml:"monetaryInformation"`
	ComponentClassInfo   componentClassInfo   `xml:"componentClassInfo"`
	FareQualifiersDetail fareQualifiersDetail `xml:"fareQualifiersDetail"`
	FareFamilyDetails    fareFamilyDetails    `xml:"fareFamilyDetails"`
	FareFamilyOwner      fareFamilyOwner      `xml:"fareFamilyOwner"`
	CouponDetailsGroup   []couponDetailsGroup `xml:"couponDetailsGroup"`
}

// fareComponentID is generated from an XSD element
type fareComponentID struct {
	ItemNumberDetails itemNumberDetails `xml:"itemNumberDetails"`
}

// marketFareComponent is generated from an XSD element
type marketFareComponent struct {
	BoardPointDetails boardPointDetails `xml:"boardPointDetails"`
	OffpointDetails   offpointDetails   `xml:"offpointDetails"`
}

// monetaryInformation is generated from an XSD element
type monetaryInformation struct {
	MonetaryDetails      monetaryDetails      `xml:"monetaryDetails"`
	OtherMonetaryDetails otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// componentClassInfo is generated from an XSD element
type componentClassInfo struct {
	FareBasisDetails fareBasisDetails `xml:"fareBasisDetails"`
}

// fareQualifiersDetail is generated from an XSD element
type fareQualifiersDetail struct {
	DiscountDetails discountDetails `xml:"discountDetails"`
}

// fareFamilyDetails is generated from an XSD element
type fareFamilyDetails struct {
	FareFamilyname          string                  `xml:"fareFamilyname"`
	Hierarchy               int                     `xml:"hierarchy"`
	CommercialFamilyDetails commercialFamilyDetails `xml:"commercialFamilyDetails"`
}

// commercialFamilyDetails is generated from an XSD element
type commercialFamilyDetails struct {
	CommercialFamily string `xml:"commercialFamily"`
}

// fareFamilyOwner is generated from an XSD element
type fareFamilyOwner struct {
	CompanyIDentification companyIDentification `xml:"companyIdentification"`
}

// companyIDentification is generated from an XSD element
type companyIDentification struct {
	OtherCompany string `xml:"otherCompany"`
}

// couponDetailsGroup is generated from an XSD element
type couponDetailsGroup struct {
	ProductID             productID             `xml:"productId"`
	FlightConnectionType  flightConnectionType  `xml:"flightConnectionType"`
	CouponTaxDetailsGroup couponTaxDetailsGroup `xml:"couponTaxDetailsGroup"`
}

// productID is generated from an XSD element
type productID struct {
	ReferenceDetails referenceDetails `xml:"referenceDetails"`
}

// flightConnectionType is generated from an XSD element
type flightConnectionType struct {
	BoardPointDetails boardPointDetails `xml:"boardPointDetails"`
	OffpointDetails   offpointDetails   `xml:"offpointDetails"`
	FlightTypeDetails flightTypeDetails `xml:"flightTypeDetails"`
}

// couponTaxDetailsGroup is generated from an XSD element
type couponTaxDetailsGroup struct {
	TaxTriggerInfo taxTriggerInfo `xml:"taxTriggerInfo"`
	TaxDetails     taxDetails     `xml:"taxDetails"`
	MonetaryInfo   monetaryInfo   `xml:"monetaryInfo"`
	LocationInfo   locationInfo   `xml:"locationInfo"`
}

// taxTriggerInfo is generated from an XSD element
type taxTriggerInfo struct {
	TaxQualifier string `xml:"taxQualifier"`
}

// monetaryInfo is generated from an XSD element
type monetaryInfo struct {
	MonetaryDetails      monetaryDetails      `xml:"monetaryDetails"`
	OtherMonetaryDetails otherMonetaryDetails `xml:"otherMonetaryDetails"`
}

// locationInfo is generated from an XSD element
type locationInfo struct {
	LocationType        string              `xml:"locationType"`
	LocationDescription locationDescription `xml:"locationDescription"`
}

// locationDescription is generated from an XSD element
type locationDescription struct {
	Code string `xml:"code"`
}
