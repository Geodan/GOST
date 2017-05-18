package postgis

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/geodan/gost/sensorthings/entities"
)

func TestDatastreamParamFactory(t *testing.T) {
	// arrange
	values := map[string]interface{}{
		"datastream_id":4,
		"datastream_name": "name",
		"datastream_description": "desc",
	}

	// act
	entity, err := datastreamParamFactory(values)
	entitytype := entity.GetEntityType()

	// assert
	assert.True(t, entity != nil)
	// entities..
	assert.True(t, err == nil)
	assert.True(t,entity.GetID() == 4)
	assert.True(t,entitytype == entities.EntityTypeDatastream)
}

