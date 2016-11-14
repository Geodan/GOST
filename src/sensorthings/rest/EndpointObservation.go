package rest

import (
	"fmt"

	"github.com/geodan/gost/src/sensorthings/models"
	"github.com/geodan/gost/src/sensorthings/odata"
)

func createObservationsEndpoint(externalURL string) *Endpoint {
	return &Endpoint{
		Name:       "Observations",
		OutputInfo: true,
		URL:        fmt.Sprintf("%s/%s/%s", externalURL, models.APIPrefix, fmt.Sprintf("%v", "Observations")),
		SupportedQueryOptions: []odata.QueryOptionType{
			odata.QueryOptionTop, odata.QueryOptionSkip, odata.QueryOptionOrderBy, odata.QueryOptionCount, odata.QueryOptionResultFormat,
			odata.QueryOptionExpand, odata.QueryOptionSelect, odata.QueryOptionFilter,
		},
		SupportedExpandParams: []string{
			"Datastream",
			"FeatureOfInterest",
		},
		SupportedSelectParams: []string{
			"description",
			"encodingType",
			"feature",
			"Observations",
		},
		Operations: []models.EndpointOperation{
			{models.HTTPOperationGet, "/v1.0/observations", HandleGetObservations},
			{models.HTTPOperationGet, "/v1.0/observations{id}", HandleGetObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/datastream", HandleGetDatastreamByObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/featureofinterest", HandleGetFeatureOfInterestByObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/datastream/{params}", HandleGetDatastreamByObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/datastream/{params}/$value", HandleGetDatastreamByObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/featureofinterest/{params}", HandleGetFeatureOfInterestByObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/featureofinterest/{params}/$value", HandleGetFeatureOfInterestByObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/{params}", HandleGetObservation},
			{models.HTTPOperationGet, "/v1.0/observations{id}/{params}/$value", HandleGetObservation},
			{models.HTTPOperationGet, "/v1.0/observations/{params}", HandleGetObservations},

			{models.HTTPOperationPost, "/v1.0/observations", HandlePostObservation},
			{models.HTTPOperationPost, "/v1.0/datastreams{id}/observations", HandlePostObservationByDatastream},
			{models.HTTPOperationDelete, "/v1.0/observations{id}", HandleDeleteObservation},
			{models.HTTPOperationPatch, "/v1.0/observations{id}", HandlePatchObservation},
			{models.HTTPOperationPut, "/v1.0/observations{id}", HandlePutObservation},

			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations", HandleGetObservations},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}", HandleGetObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/datastream", HandleGetDatastreamByObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/featureOfInterest", HandleGetFeatureOfInterestByObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/datastream/{params}", HandleGetDatastreamByObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/datastream/{params}/$value", HandleGetDatastreamByObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/featureofinterest/{params}", HandleGetFeatureOfInterestByObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/featureofinterest/{params}/$value", HandleGetFeatureOfInterestByObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/{params}", HandleGetObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations{id}/{params}/$value", HandleGetObservation},
			{models.HTTPOperationGet, "/v1.0/{c:.*}/observations/{params}", HandleGetObservations},

			{models.HTTPOperationPost, "/v1.0/{c:.*}/observations", HandlePostObservation},
			{models.HTTPOperationPost, "/v1.0/{c:.*}/datastreams{id}/observations", HandlePostObservationByDatastream},
			{models.HTTPOperationDelete, "/v1.0/{c:.*}/observations{id}", HandleDeleteObservation},
			{models.HTTPOperationPatch, "/v1.0/{c:.*}/observations{id}", HandlePatchObservation},
			{models.HTTPOperationPut, "/v1.0/{c:.*}/observations{id}", HandlePutObservation},
		},
	}
}
