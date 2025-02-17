/*
 * 3gpp-traffic-influence
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package TrafficInfluence

// APIClient manages communication with the 3gpp-traffic-influence API v1.0.1
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	TrafficInfluenceAPISCSASLevelGETOperationApi           *TrafficInfluenceAPISCSASLevelGETOperationApiService
	TrafficInfluenceAPISubscriptionLevelDELETEOperationApi *TrafficInfluenceAPISubscriptionLevelDELETEOperationApiService
	TrafficInfluenceAPISubscriptionLevelGETOperationApi    *TrafficInfluenceAPISubscriptionLevelGETOperationApiService
	TrafficInfluenceAPISubscriptionLevelPATCHOperationApi  *TrafficInfluenceAPISubscriptionLevelPATCHOperationApiService
	TrafficInfluenceAPISubscriptionLevelPOSTOperationApi   *TrafficInfluenceAPISubscriptionLevelPOSTOperationApiService
	TrafficInfluenceAPISubscriptionLevelPUTOperationApi    *TrafficInfluenceAPISubscriptionLevelPUTOperationApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.TrafficInfluenceAPISCSASLevelGETOperationApi = (*TrafficInfluenceAPISCSASLevelGETOperationApiService)(&c.common)
	c.TrafficInfluenceAPISubscriptionLevelDELETEOperationApi = (*TrafficInfluenceAPISubscriptionLevelDELETEOperationApiService)(&c.common)
	c.TrafficInfluenceAPISubscriptionLevelGETOperationApi = (*TrafficInfluenceAPISubscriptionLevelGETOperationApiService)(&c.common)
	c.TrafficInfluenceAPISubscriptionLevelPATCHOperationApi = (*TrafficInfluenceAPISubscriptionLevelPATCHOperationApiService)(&c.common)
	c.TrafficInfluenceAPISubscriptionLevelPOSTOperationApi = (*TrafficInfluenceAPISubscriptionLevelPOSTOperationApiService)(&c.common)
	c.TrafficInfluenceAPISubscriptionLevelPUTOperationApi = (*TrafficInfluenceAPISubscriptionLevelPUTOperationApiService)(&c.common)

	return c
}
