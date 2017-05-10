package postgis

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/geodan/gost/src/sensorthings/entities"
)

func TestFoiParamFactory(t *testing.T) {
	// arrange
	values := map[string]interface{}{
		"featureofinterest_id":4,
		"featureofinterest_name": "name",
		"featureofinterest_description": "desc",
	}
	// todo: encodingtype + feature

	// act
	entity, err := featureOfInterestParamFactory(values)
	entitytype := entity.GetEntityType()

	// assert
	assert.True(t, entity != nil)
	// entities..
	assert.True(t, err == nil)
	assert.True(t,entity.GetID() == 4)
	assert.True(t,entitytype == entities.EntityTypeFeatureOfInterest)
}

