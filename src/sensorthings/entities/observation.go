package entities

import (
	"encoding/json"

	"errors"
	"fmt"
	"time"

	gostErrors "github.com/geodan/gost/src/errors"
)

// Observation in SensorThings represents a single Sensor reading of an ObservedProperty. A physical device, a Sensor, sends
// Observations to a specified Datastream. An Observation requires a FeatureOfInterest entity, if none is provided in the request,
// the Location of the Thing associated with the Datastream, will be assigned to the new Observation as the FeaturOfInterest.
type Observation struct {
	BaseEntity
	PhenomenonTime       string                 `json:"phenomenonTime,omitempty"`
	Result               interface{}            `json:"result,omitempty"`
	ResultTime           string                 `json:"resultTime,omitempty"`
	ResultQuality        string                 `json:"resultQuality,omitempty"`
	ValidTime            string                 `json:"validTime,omitempty"`
	Parameters           map[string]interface{} `json:"parameters,omitempty"`
	NavDatastream        string                 `json:"Datastream@iot.navigationLink,omitempty"`
	NavFeatureOfInterest string                 `json:"FeatureOfInterest@iot.navigationLink,omitempty"`
	Datastream           *Datastream            `json:"Datastream,omitempty"`
	FeatureOfInterest    *FeatureOfInterest     `json:"FeatureOfInterest,omitempty"`
}

// GetEntityType returns the EntityType for Observation
func (o Observation) GetEntityType() EntityType {
	return EntityTypeObservation
}

// GetPropertyNames returns the available properties for an Observation
func (o *Observation) GetPropertyNames() []string {
	return []string{"id", "phenomenonTime", "result", "resultTime", "resultQuality", "validTime", "parameters"}
}

// ParseEntity tries to parse the given json byte array into the current entity
func (o *Observation) ParseEntity(data []byte) error {
	observation := &o
	err := json.Unmarshal(data, observation)
	if err != nil {
		return gostErrors.NewBadRequestError(errors.New("Unable to parse Observation"))
	}

	return nil
}

// ContainsMandatoryParams checks if all mandatory params for Observation are available before posting.
func (o *Observation) ContainsMandatoryParams() (bool, []error) {
	// When a SensorThings service receives a POST Observations without phenomenonTime, the service SHALL
	// assign the current server time to the value of the phenomenonTime.
	var errors []error

	if len(o.PhenomenonTime) == 0 {
		o.PhenomenonTime = time.Now().UTC().Format(time.RFC3339Nano)
	} else {
		if t, err := time.Parse(time.RFC3339Nano, o.PhenomenonTime); err != nil {
			errors = append(errors, gostErrors.NewBadRequestError(fmt.Errorf("Invalid phenomenonTime: %v", err.Error())))
		} else {
			o.PhenomenonTime = t.UTC().Format("2006-01-02T15:04:05.000Z")
		}
	}

	// From spec: "When a SensorThings service receives a POST Observations without resultTime, the service SHALL assign a
	// null value to the resultTime."
	// Implementation: omit resultTime in database when null (see also https://github.com/Geodan/gost/issues/68)
	if len(o.ResultTime) != 0 {
		if t, err := time.Parse(time.RFC3339Nano, o.ResultTime); err != nil {
			errors = append(errors, gostErrors.NewBadRequestError(fmt.Errorf("Invalid resultTime: %v", err.Error())))
		} else {
			o.ResultTime = t.UTC().Format("2006-01-02T15:04:05.000Z")
		}
	}

	CheckMandatoryParam(&errors, o.PhenomenonTime, o.GetEntityType(), "phenomenonTime")
	CheckMandatoryParam(&errors, o.Result, o.GetEntityType(), "result")
	CheckMandatoryParam(&errors, o.Datastream, o.GetEntityType(), "Datastream")

	if len(errors) != 0 {
		return false, errors
	}

	return true, nil
}

// SetAllLinks sets the self link and relational links
func (o *Observation) SetAllLinks(externalURL string) {
	o.SetSelfLink(externalURL)
	o.SetLinks(externalURL)

	if o.Datastream != nil {
		o.Datastream.SetAllLinks(externalURL)
	}

	if o.FeatureOfInterest != nil {
		o.FeatureOfInterest.SetAllLinks(externalURL)
	}
}

// SetSelfLink sets the self link for the entity
func (o *Observation) SetSelfLink(externalURL string) {
	o.NavSelf = CreateEntitySelfLink(externalURL, EntityLinkObservations.ToString(), o.ID)
}

// SetLinks sets the entity specific navigation links, empty string if linked(expanded) data is not nil
func (o *Observation) SetLinks(externalURL string) {
	o.NavDatastream = CreateEntityLink(o.Datastream == nil, externalURL, EntityLinkObservations.ToString(), EntityTypeDatastream.ToString(), o.ID)
	o.NavFeatureOfInterest = CreateEntityLink(o.FeatureOfInterest == nil, externalURL, EntityLinkObservations.ToString(), EntityTypeFeatureOfInterest.ToString(), o.ID)
}

// MarshalPostgresJSON marshalls an observation entity for saving into PostgreSQL
func (o Observation) MarshalPostgresJSON() ([]byte, error) {
	return json.Marshal(&struct {
		PhenomenonTime string                 `json:"phenomenonTime,omitempty"`
		Result         interface{}            `json:"result,omitempty"`
		ResultTime     string                 `json:"resultTime,omitempty"`
		ResultQuality  string                 `json:"resultQuality,omitempty"`
		ValidTime      string                 `json:"validTime,omitempty"`
		Parameters     map[string]interface{} `json:"parameters,omitempty"`
	}{
		PhenomenonTime: o.PhenomenonTime,
		Result:         o.Result,
		ResultTime:     o.ResultTime,
		ResultQuality:  o.ResultQuality,
		ValidTime:      o.ValidTime,
		Parameters:     o.Parameters,
	})
}

// GetSupportedEncoding returns the supported encoding tye for this entity
func (o Observation) GetSupportedEncoding() map[int]EncodingType {
	return map[int]EncodingType{}
}
