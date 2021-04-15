//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package http_loadbalancer

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	google_api "google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create ApiepLBCustomAPI GRPC Client satisfying server.CustomClient
type ApiepLBCustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient ApiepLBCustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *ApiepLBCustomAPIGrpcClient) doRPCGetSwaggerSpec(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SwaggerSpecReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.http_loadbalancer.SwaggerSpecReq", yamlReq)
	}
	rsp, err := c.grpcClient.GetSwaggerSpec(ctx, req, opts...)
	return rsp, err
}

func (c *ApiepLBCustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewApiepLBCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &ApiepLBCustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewApiepLBCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["GetSwaggerSpec"] = ccl.doRPCGetSwaggerSpec

	ccl.rpcFns = rpcFns

	return ccl
}

// Create ApiepLBCustomAPI REST Client satisfying server.CustomClient
type ApiepLBCustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *ApiepLBCustomAPIRestClient) doRPCGetSwaggerSpec(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SwaggerSpecReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.http_loadbalancer.SwaggerSpecReq: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP POST request for custom API")
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &google_api.HttpBody{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		// server strips HTTP Body proto message and sends only data, re-build it here
		pbRsp.ContentType = rsp.Header.Get("Content-Type")
		pbRsp.Data = body

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *ApiepLBCustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewApiepLBCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &ApiepLBCustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["GetSwaggerSpec"] = ccl.doRPCGetSwaggerSpec

	ccl.rpcFns = rpcFns

	return ccl
}

// Create ApiepLBCustomAPIInprocClient

// INPROC Client (satisfying ApiepLBCustomAPIClient interface)
type ApiepLBCustomAPIInprocClient struct {
	svc svcfw.Service
}

func (c *ApiepLBCustomAPIInprocClient) GetSwaggerSpec(ctx context.Context, in *SwaggerSpecReq, opts ...grpc.CallOption) (*google_api.HttpBody, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI")
	cah, ok := ah.(ApiepLBCustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *ApiepLBCustomAPISrv", ah)
	}

	var (
		rsp *google_api.HttpBody
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.http_loadbalancer.SwaggerSpecReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'ApiepLBCustomAPI.GetSwaggerSpec' operation on 'http_loadbalancer'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI.GetSwaggerSpec"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetSwaggerSpec(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "google.api.HttpBody", rsp)...)

	return rsp, nil
}

func NewApiepLBCustomAPIInprocClient(svc svcfw.Service) ApiepLBCustomAPIClient {
	return &ApiepLBCustomAPIInprocClient{svc: svc}
}

// RegisterGwApiepLBCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwApiepLBCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterApiepLBCustomAPIHandlerClient(ctx, mux, NewApiepLBCustomAPIInprocClient(s))
}

var ApiepLBCustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "HTTP loadbalancer",
        "description": "HTTP Loadbalancer view defines a required parameters that can be used in CRUD, to create and manage HTTP loadbalancer.\nIt can be used to create HTTP loadbalancer and HTTPS loadbalancer.\n\nView will create following child objects.\n\n* Virtual-host\n* routes\n* clusters\n* endpoints\n* advertise policy",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": null,
    "paths": {
        "/public/namespaces/{namespace}/http_loadbalancers/{name}/api_endpoints/swagger_spec": {
            "get": {
                "summary": "Get Swagger Spec for App Type",
                "description": "Get the corresponding Swagger spec for the given Http Loadbalancer",
                "operationId": "ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI.GetSwaggerSpec",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/apiHttpBody"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "name",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    }
                ],
                "tags": [
                    "ApiepLBCustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-http_loadbalancer-ApiepLBCustomAPI-GetSwaggerSpec"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI.GetSwaggerSpec"
            },
            "x-displayname": "HTTP loadbalancer Custom API",
            "x-ves-proto-service": "ves.io.schema.views.http_loadbalancer.ApiepLBCustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "apiHttpBody": {
            "type": "object",
            "description": "Message that represents an arbitrary HTTP body. It should only be used for\npayload formats that can't be represented as JSON, such as raw binary or\nan HTML page.\n\n\nThis message can be used both in streaming and non-streaming API methods in\nthe request as well as the response.\n\nIt can be used as a top-level request field, which is convenient if one\nwants to extract parameters from either the URL or HTTP template into the\nrequest fields and also want access to the raw HTTP body.\n\nExample:\n\n    message GetResourceRequest {\n      // A unique request id.\n      string request_id = 1;\n\n      // The raw HTTP body is bound to this field.\n      google.api.HttpBody http_body = 2;\n    }\n\n    service ResourceService {\n      rpc GetResource(GetResourceRequest) returns (google.api.HttpBody);\n      rpc UpdateResource(google.api.HttpBody) returns\n      (google.protobuf.Empty);\n    }\n\nExample with streaming methods:\n\n    service CaldavService {\n      rpc GetCalendar(stream google.api.HttpBody)\n        returns (stream google.api.HttpBody);\n      rpc UpdateCalendar(stream google.api.HttpBody)\n        returns (stream google.api.HttpBody);\n    }\n\nUse of this type only changes how the request and response bodies are\nhandled, all other features will continue to work unchanged.",
            "properties": {
                "content_type": {
                    "type": "string",
                    "description": "The HTTP Content-Type header value specifying the content type of the body."
                },
                "data": {
                    "type": "string",
                    "description": "The HTTP request/response body as raw binary.",
                    "format": "byte"
                },
                "extensions": {
                    "type": "array",
                    "description": "Application specific response metadata. Must be set in the first response\nfor streaming APIs.",
                    "items": {
                        "$ref": "#/definitions/protobufAny"
                    }
                }
            }
        },
        "protobufAny": {
            "type": "object",
            "description": "-Any- contains an arbitrary serialized protocol buffer message along with a\nURL that describes the type of the serialized message.\n\nProtobuf library provides support to pack/unpack Any values in the form\nof utility functions or additional generated methods of the Any type.\n\nExample 1: Pack and unpack a message in C++.\n\n    Foo foo = ...;\n    Any any;\n    any.PackFrom(foo);\n    ...\n    if (any.UnpackTo(\u0026foo)) {\n      ...\n    }\n\nExample 2: Pack and unpack a message in Java.\n\n    Foo foo = ...;\n    Any any = Any.pack(foo);\n    ...\n    if (any.is(Foo.class)) {\n      foo = any.unpack(Foo.class);\n    }\n\n Example 3: Pack and unpack a message in Python.\n\n    foo = Foo(...)\n    any = Any()\n    any.Pack(foo)\n    ...\n    if any.Is(Foo.DESCRIPTOR):\n      any.Unpack(foo)\n      ...\n\n Example 4: Pack and unpack a message in Go\n\n     foo := \u0026pb.Foo{...}\n     any, err := ptypes.MarshalAny(foo)\n     ...\n     foo := \u0026pb.Foo{}\n     if err := ptypes.UnmarshalAny(any, foo); err != nil {\n       ...\n     }\n\nThe pack methods provided by protobuf library will by default use\n'type.googleapis.com/full.type.name' as the type URL and the unpack\nmethods only use the fully qualified type name after the last '/'\nin the type URL, for example \"foo.bar.com/x/y.z\" will yield type\nname \"y.z\".\n\n\nJSON\n====\nThe JSON representation of an -Any- value uses the regular\nrepresentation of the deserialized, embedded message, with an\nadditional field -@type- which contains the type URL. Example:\n\n    package google.profile;\n    message Person {\n      string first_name = 1;\n      string last_name = 2;\n    }\n\n    {\n      \"@type\": \"type.googleapis.com/google.profile.Person\",\n      \"firstName\": \u003cstring\u003e,\n      \"lastName\": \u003cstring\u003e\n    }\n\nIf the embedded message type is well-known and has a custom JSON\nrepresentation, that representation will be embedded adding a field\n-value- which holds the custom JSON in addition to the -@type-\nfield. Example (for message [google.protobuf.Duration][]):\n\n    {\n      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",\n      \"value\": \"1.212s\"\n    }",
            "properties": {
                "type_url": {
                    "type": "string",
                    "description": "A URL/resource name that uniquely identifies the type of the serialized\nprotocol buffer message. This string must contain at least\none \"/\" character. The last segment of the URL's path must represent\nthe fully qualified name of the type (as in\n-path/google.protobuf.Duration-). The name should be in a canonical form\n(e.g., leading \".\" is not accepted).\n\nIn practice, teams usually precompile into the binary all types that they\nexpect it to use in the context of Any. However, for URLs which use the\nscheme -http-, -https-, or no scheme, one can optionally set up a type\nserver that maps type URLs to message definitions as follows:\n\n* If no scheme is provided, -https- is assumed.\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\n  value in binary format, or produce an error.\n* Applications are allowed to cache lookup results based on the\n  URL, or have them precompiled into a binary to avoid any\n  lookup. Therefore, binary compatibility needs to be preserved\n  on changes to types. (Use versioned type names to manage\n  breaking changes.)\n\nNote: this functionality is not currently available in the official\nprotobuf release, and it is not used for type URLs beginning with\ntype.googleapis.com.\n\nSchemes other than -http-, -https- (or the empty scheme) might be\nused with implementation specific semantics."
                },
                "value": {
                    "type": "string",
                    "description": "Must be a valid serialized protocol buffer of the above specified type.",
                    "format": "byte"
                }
            }
        }
    },
    "x-displayname": "Configure HTTP Loadbalancer",
    "x-ves-proto-file": "ves.io/schema/views/http_loadbalancer/public_apiep_customapi.proto"
}`
