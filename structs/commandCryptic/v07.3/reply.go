package CommandCryptic_v07_3 // hsfres073
import "github.com/tmconsulting/amadeus-golang-sdk/structs/commandCryptic"

//import "encoding/xml"

type Response struct {
	// XMLName xml.Name `xml:"http://xml.amadeus.com/HSFRES_07_3_1A Command_CrypticReply"`
	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails,omitempty"`
	LongTextString       *LongTextString       `xml:"longTextString"`
}

type MessageActionDetails struct {
	MessageFunctionDetails *MessageFunctionDetails `xml:"messageFunctionDetails"`
	ResponseType           string                  `xml:"responseType,omitempty"` // Format limitations: an..3
}

type LongTextString struct {
	TextStringDetails string `xml:"textStringDetails"` // Format limitations: an..9999
}

func (r *Response) ToCommon() *commandCryptic.Response {
	return &commandCryptic.Response{
		MessageActionDetails: &commandCryptic.MessageActionDetails{
			MessageFunctionDetails: &commandCryptic.MessageFunctionDetails{
				BusinessFunction:          r.MessageActionDetails.MessageFunctionDetails.BusinessFunction,
				MessageFunction:           r.MessageActionDetails.MessageFunctionDetails.MessageFunction,
				AdditionalMessageFunction: r.MessageActionDetails.MessageFunctionDetails.AdditionalMessageFunction,
			},
			ResponseType: r.MessageActionDetails.ResponseType,
		},
		LongTextString: &commandCryptic.LongTextString{
			TextStringDetails: r.LongTextString.TextStringDetails,
		},
	}
}
