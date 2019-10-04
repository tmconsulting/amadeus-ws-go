package PNR_Retrieve_v11_3_request

import (
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs/formats"
	"github.com/tmconsulting/amadeus-golang-sdk/v2/structs/pnr/retrieve"
)

func MakeRequest(request *PNR_Information.Request) *Request {
	return &Request{
		RetrievalFacts: &RetrievalFacts{
			Retrieve: &RetrievePNRType{
				Type: formats.NumericInteger_Length1To1(2),
			},
			ReservationOrProfileIdentifier: &ReservationControlInformationType{
				Reservation: &ReservationControlInformationDetailsTypeI{
					ControlNumber: formats.AlphaNumericString_Length1To20(request.PNR),
				},
			},
		},
	}
}
