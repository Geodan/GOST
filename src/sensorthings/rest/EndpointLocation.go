package rest

import (
	"fmt"

	"github.com/geodan/gost/src/sensorthings/models"
	"github.com/geodan/gost/src/sensorthings/odata"
)

func createLocationsEndpoint(externalURL string) *Endpoint {
	return &Endpoint{
		Name:       "Locations",
		OutputInfo: true,
		URL:        fmt.Sprintf("%s/%s/%s", externalURL, models.APIPrefix, fmt.Sprintf("%v", "Locations")),
		SupportedQueryOptions: []odata.QueryOptionType{
			odata.QueryOptionTop, odata.QueryOptionSkip, odata.QueryOptionOrderBy, odata.QueryOptionCount, odata.QueryOptionResultFormat,
			odata.QueryOptionExpand, odata.QueryOptionSelect, odata.QueryOptionFilter,
		},
		SupportedExpandParams: []string{
			"Things",
			"HistoricalLocations",
		},
		SupportedSelectParams: []string{
			"description",
			"encodingType",
			"location",
			"Things",
			"HistoricalLocations",
		},
		Operations: []models.EndpointOperation{
			{models.HTTPOperationGet, "/v1.0/locations", HandleGetLocations},
			{models.HTTPOperationGet, "/v1.0/locations{id}", HandleGetLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/things", HandleGetThingsByLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/historicallocations", HandleGetHistoricalLocationsByLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/things/{params}", HandleGetThingsByLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/historicallocations/{params}", HandleGetHistoricalLocationsByLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/historicallocations/{params}/$value", HandleGetHistoricalLocationsByLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/{params}", HandleGetLocation},
			{models.HTTPOperationGet, "/v1.0/locations{id}/{params}/$value", HandleGetLocation},
			{models.HTTPOperationGet, "/v1.0/locations/{params}", HandleGetLocations},

			{models.HTTPOperationPost, "/v1.0/locations", HandlePostLocation},
			{models.HTTPOperationPost, "/v1.0/things{id}/locations", HandlePostLocationByThing},
			{models.HTTPOperationDelete, "/v1.0/locations{id}", HandleDeleteLocation},
			{models.HTTPOperationPatch, "/v1.0/locations{id}", HandlePatchLocation},
			{models.HTTPOperationPut, "/v1.0/locations{id}", HandlePutLocation},

			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations", HandleGetLocations},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}", HandleGetLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/things", HandleGetThingsByLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/historicallocations", HandleGetHistoricalLocationsByLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/things/{params}", HandleGetThingsByLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/historicallocations/{params}", HandleGetHistoricalLocationsByLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/historicallocations/{params}/$value", HandleGetHistoricalLocationsByLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/{params}", HandleGetLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations{id}/{params}/$value", HandleGetLocation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/locations/{params}", HandleGetLocations},

			{models.HTTPOperationPost, "/v1.0/{c:.*}/locations", HandlePostLocation},
			{models.HTTPOperationPost, "/v1.0/{c:.*}/things{id}/locations", HandlePostLocationByThing},
			{models.HTTPOperationDelete, "/v1.0/{c:.*}/locations{id}", HandleDeleteLocation},
			{models.HTTPOperationPatch, "/v1.0/{c:.*}/locations{id}", HandlePatchLocation},
			{models.HTTPOperationPut, "/v1.0/{c:.*}/locations{id}", HandlePutLocation},
		},
	}
}
