package handler

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/private/protocol/xml/xmlutil"
)

// UnmarshalXML unmarshals a response body for the XML protocol.
func UnmarshalXML(v interface{}, r *http.Response) error {

	defer r.Body.Close()

	// Workaround for now, until I get another Idea on how to deal with
	// empty responses
	if v == nil {
		return nil
	}

	decoder := xml.NewDecoder(r.Body)
	err := xmlutil.UnmarshalXML(v, decoder, "")

	if err != nil {
		return errors.New("SerializationError" + "failed decoding EC2 Query response" + fmt.Sprint(err))
	}

	return nil
}
