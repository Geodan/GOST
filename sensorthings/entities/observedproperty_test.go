package entities

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonObservedProperty = `{
  "name": "ObservedPropertyUp Tempomatic 2000",
  "description": "http://schema.org/description",
  "definition": "Calibration date:  Jan 1, 2014"
}`

var jsonObservedPropertyError = `{
  "name": "ObservedPropertyUp Tempomatic 2000",
}`

func TestMissingMandatoryParametersObservedProperty(t *testing.T) {
	//arrange
	op := &ObservedProperty{}

	//act
	_, err := op.ContainsMandatoryParams()

	//assert
	assert.NotNil(t, err, "ObservedProperty mandatory param description not filled in should have returned error")
	if len(err) > 0 {
		assert.Contains(t, fmt.Sprintf("%v", err[0]), "name")
	}
}

func TestObservedPropertyGetPropertyNames(t *testing.T) {
	// arrange
	op := &ObservedProperty{}

	// act
	propertynames := op.GetPropertyNames()

	// assert
	assert.True(t, propertynames[0] == "id")
}


func TestMandatoryParametersExistObservedProperty(t *testing.T) {
	//arrange
	op := &ObservedProperty{
		Name:        "test",
		Definition:  "test",
		Description: "test",
	}

	//act
	_, err := op.ContainsMandatoryParams()

	//assert
	assert.Nil(t, err, "All mandatory params are filled in should not have returned an error")
}

func TestParseEntityResultOkObservedProperty(t *testing.T) {
	//arrange
	op := &ObservedProperty{}

	//act
	err := op.ParseEntity([]byte(jsonObservedProperty))

	//assert
	assert.Equal(t, err, nil, "Unable to parse json into ObservedProperty")
}

func TestParseEntityResultNotOkObservedProperty(t *testing.T) {
	//arrange
	op := &ObservedProperty{}

	//act
	err := op.ParseEntity([]byte(jsonObservedPropertyError))

	//assert
	assert.NotEqual(t, err, nil, "ObservedProperty parse from json should have failed")
}

func TestSetLinksSObservedProperty(t *testing.T) {
	//arrange
	op := &ObservedProperty{}
	op.ID = id
	ds1 := &Datastream{}
	ds2 := &Datastream{}
	datastreams := []*Datastream{ds1, ds2}
	op.Datastreams = datastreams

	//act
	op.SetAllLinks(externalURL)

	//assert
	assert.Equal(t, op.NavSelf, fmt.Sprintf("%s/v1.0/%s(%s)", externalURL, EntityLinkObservedProperties.ToString(), id), "ObservedProperty navself incorrect")
	// assert.Equal(t, op.NavDatastreams, fmt.Sprintf("%s/v1.0/%s(%s)/%s", externalURL, EntityLinkObservedProperties.ToString(), id, EntityLinkDatastreams.ToString()), "ObservedProperty NavDatastreams incorrect")
}

func TestGetSupportedEncodingObservedProperty(t *testing.T) {
	//arrange
	op := &ObservedProperty{}

	// act
	supportedEncoding := op.GetSupportedEncoding()

	//assert
	assert.Equal(t, 0, len(supportedEncoding), "ObservedProperty should not support any encoding")
}
