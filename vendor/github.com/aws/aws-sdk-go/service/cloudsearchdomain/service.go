// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudsearchdomain

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

// You use the AmazonCloudSearch2013 API to upload documents to a search domain
// and search those documents.
//
// The endpoints for submitting UploadDocuments, Search, and Suggest requests
// are domain-specific. To get the endpoints for your domain, use the Amazon
// CloudSearch configuration service DescribeDomains action. The domain endpoints
// are also displayed on the domain dashboard in the Amazon CloudSearch console.
// You submit suggest requests to the search endpoint.
//
// For more information, see the Amazon CloudSearch Developer Guide (http://docs.aws.amazon.com/cloudsearch/latest/developerguide).
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CloudSearchDomain struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "cloudsearchdomain" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName         // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CloudSearchDomain client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudSearchDomain client from just a session.
//     svc := cloudsearchdomain.New(mySession)
//
//     // Create a CloudSearchDomain client with additional configuration
//     svc := cloudsearchdomain.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudSearchDomain {
	var c client.Config
	if v, ok := p.(client.ConfigNoResolveEndpointProvider); ok {
		c = v.ClientConfigNoResolveEndpoint(cfgs...)
	} else {
		c = p.ClientConfig(EndpointsID, cfgs...)
	}
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *CloudSearchDomain {
	if len(signingName) == 0 {
		signingName = "cloudsearch"
	}
	svc := &CloudSearchDomain{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2013-01-01",
				JSONVersion:   "1.1",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restjson.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CloudSearchDomain operation and runs any
// custom request initialization.
func (c *CloudSearchDomain) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
