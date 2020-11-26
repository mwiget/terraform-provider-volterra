//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package service_policy

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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

// Create CustomDataAPI GRPC Client satisfying server.CustomClient
type CustomDataAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomDataAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomDataAPIGrpcClient) doRPCServicePolicyHits(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ServicePolicyHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.service_policy.ServicePolicyHitsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.ServicePolicyHits(ctx, req, opts...)
	return rsp, err
}

func (c *CustomDataAPIGrpcClient) doRPCServicePolicyHitsLatency(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ServicePolicyHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.service_policy.ServicePolicyHitsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.ServicePolicyHitsLatency(ctx, req, opts...)
	return rsp, err
}

func (c *CustomDataAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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
	return rsp, nil
}

func NewCustomDataAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomDataAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomDataAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["ServicePolicyHits"] = ccl.doRPCServicePolicyHits

	rpcFns["ServicePolicyHitsLatency"] = ccl.doRPCServicePolicyHitsLatency

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomDataAPI REST Client satisfying server.CustomClient
type CustomDataAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomDataAPIRestClient) doRPCServicePolicyHits(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ServicePolicyHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.service_policy.ServicePolicyHitsRequest: %s", yamlReq, err)
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
		q.Add("end_time", fmt.Sprintf("%v", req.EndTime))
		q.Add("group_by", fmt.Sprintf("%v", req.GroupBy))
		q.Add("label_filter", fmt.Sprintf("%v", req.LabelFilter))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("start_time", fmt.Sprintf("%v", req.StartTime))
		q.Add("step", fmt.Sprintf("%v", req.Step))

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
	pbRsp := &ServicePolicyHitsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.service_policy.ServicePolicyHitsResponse", body)
	}
	return pbRsp, nil
}

func (c *CustomDataAPIRestClient) doRPCServicePolicyHitsLatency(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ServicePolicyHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.service_policy.ServicePolicyHitsRequest: %s", yamlReq, err)
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
		q.Add("end_time", fmt.Sprintf("%v", req.EndTime))
		q.Add("group_by", fmt.Sprintf("%v", req.GroupBy))
		q.Add("label_filter", fmt.Sprintf("%v", req.LabelFilter))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("start_time", fmt.Sprintf("%v", req.StartTime))
		q.Add("step", fmt.Sprintf("%v", req.Step))

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
	pbRsp := &ServicePolicyHitsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.service_policy.ServicePolicyHitsResponse", body)
	}
	return pbRsp, nil
}

func (c *CustomDataAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomDataAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["ServicePolicyHits"] = ccl.doRPCServicePolicyHits

	rpcFns["ServicePolicyHitsLatency"] = ccl.doRPCServicePolicyHitsLatency

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomDataAPIInprocClient

// INPROC Client (satisfying CustomDataAPIClient interface)
type CustomDataAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomDataAPIInprocClient) ServicePolicyHits(ctx context.Context, in *ServicePolicyHitsRequest, opts ...grpc.CallOption) (*ServicePolicyHitsResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.service_policy.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPISrv", ah)
	}

	var (
		rsp *ServicePolicyHitsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.service_policy.ServicePolicyHitsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.ServicePolicyHits' operation on 'service_policy'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	rsp, err = cah.ServicePolicyHits(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.service_policy.ServicePolicyHitsResponse", rsp)...)

	return rsp, nil
}
func (c *CustomDataAPIInprocClient) ServicePolicyHitsLatency(ctx context.Context, in *ServicePolicyHitsRequest, opts ...grpc.CallOption) (*ServicePolicyHitsResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.service_policy.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPISrv", ah)
	}

	var (
		rsp *ServicePolicyHitsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.service_policy.ServicePolicyHitsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.ServicePolicyHitsLatency' operation on 'service_policy'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	rsp, err = cah.ServicePolicyHitsLatency(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.service_policy.ServicePolicyHitsResponse", rsp)...)

	return rsp, nil
}

func NewCustomDataAPIInprocClient(svc svcfw.Service) CustomDataAPIClient {
	return &CustomDataAPIInprocClient{svc: svc}
}

// RegisterGwCustomDataAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomDataAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomDataAPIHandlerClient(ctx, mux, NewCustomDataAPIInprocClient(s))
}

var CustomDataAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Service Policy",
        "description": "Monitoring APIs for Service Policy",
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
        "/public/namespaces/{namespace}/service_policy/hits": {
            "post": {
                "summary": "Service Policy Hits",
                "description": "Get the counter for Service Policy hits for a given namespace.",
                "operationId": "ves.io.schema.service_policy.CustomDataAPI.ServicePolicyHits",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service_policyServicePolicyHitsResponse"
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
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service_policyServicePolicyHitsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-service_policy-CustomDataAPI-ServicePolicyHits"
                },
                "x-ves-proto-rpc": "ves.io.schema.service_policy.CustomDataAPI.ServicePolicyHits"
            },
            "x-displayname": "Service Policy View Monitoring APIs",
            "x-ves-proto-service": "ves.io.schema.service_policy.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/service_policy/latency": {
            "post": {
                "summary": "Service Policy Latency",
                "description": "Get the average latency for Service policy evaluation.",
                "operationId": "ves.io.schema.service_policy.CustomDataAPI.ServicePolicyHitsLatency",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/service_policyServicePolicyHitsResponse"
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
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service_policyServicePolicyHitsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-service_policy-CustomDataAPI-ServicePolicyHitsLatency"
                },
                "x-ves-proto-rpc": "ves.io.schema.service_policy.CustomDataAPI.ServicePolicyHitsLatency"
            },
            "x-displayname": "Service Policy View Monitoring APIs",
            "x-ves-proto-service": "ves.io.schema.service_policy.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "schemaMetricLabelOp": {
            "type": "string",
            "description": "The operator to use when filtering metrics based on label values.\n",
            "title": "Metric Label Operator",
            "enum": [
                "EQ",
                "NEQ"
            ],
            "default": "EQ",
            "x-displayname": "Metric Label Operator",
            "x-ves-proto-enum": "ves.io.schema.MetricLabelOp"
        },
        "schemaMetricValue": {
            "type": "object",
            "description": "Metric data contains timestamp and the value.",
            "title": "Metric Value",
            "x-displayname": "Metric Value",
            "x-ves-proto-message": "ves.io.schema.MetricValue",
            "properties": {
                "timestamp": {
                    "type": "number",
                    "description": " timestamp\n\nExample: - \"1570007981\"-",
                    "title": "Timestamp",
                    "format": "double",
                    "x-displayname": "Timestamp",
                    "x-ves-example": "1570007981"
                },
                "value": {
                    "type": "string",
                    "description": "\n\nExample: - \"15\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "15"
                }
            }
        },
        "service_policyServicePolicyHits": {
            "type": "object",
            "description": "ServicePolicyHits contains the timeseries data of Service policy hits",
            "title": "Service Policy Hits",
            "x-displayname": "Service Policy Hits",
            "x-ves-proto-message": "ves.io.schema.service_policy.ServicePolicyHits",
            "properties": {
                "id": {
                    "description": " ID identifies the unique combination of group_by label values in the response",
                    "title": "ID",
                    "$ref": "#/definitions/service_policyServicePolicyHitsId",
                    "x-displayname": "ID"
                },
                "metric": {
                    "type": "array",
                    "description": " x-unit: \"count\"\n List of metric values",
                    "title": "Metric",
                    "items": {
                        "$ref": "#/definitions/schemaMetricValue"
                    },
                    "x-displayname": "Metric"
                }
            }
        },
        "service_policyServicePolicyHitsId": {
            "type": "object",
            "description": "ServicePolicyHitsId uniquely identifies an entry in the response to Service policy hits request.\nService policy hits counter is aggregated based on the labels specified in the group_by field in the request.\nTherefore, only the feields that corresponds to the MetricLabel in the group_by field will have non-empty\nvalue in the response.",
            "title": "Service Policy Hits ID",
            "x-displayname": "Service Policy Hits ID",
            "x-ves-proto-message": "ves.io.schema.service_policy.ServicePolicyHitsId",
            "properties": {
                "action": {
                    "type": "string",
                    "description": " Action associated with the policy rule\n\nExample: - \"allow\"-",
                    "title": "Action",
                    "x-displayname": "Action",
                    "x-ves-example": "allow"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace in which the policy rule was hit\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "policy": {
                    "type": "string",
                    "description": " Policy name\n\nExample: - \"policy1\"-",
                    "title": "Policy",
                    "x-displayname": "Policy",
                    "x-ves-example": "policy1"
                },
                "policy_rule": {
                    "type": "string",
                    "description": " Policy Rule name\n\nExample: - \"rule1\"-",
                    "title": "Policy Rule",
                    "x-displayname": "Policy Rule",
                    "x-ves-example": "rule1"
                },
                "policy_set": {
                    "type": "string",
                    "description": " Policy Set name\n\nExample: - \"policy-set1\"-",
                    "title": "Policy Set",
                    "x-displayname": "Policy Set",
                    "x-ves-example": "policy-set1"
                },
                "site": {
                    "type": "string",
                    "description": " Site name\n\nExample: - \"ce1\"-",
                    "title": "Site",
                    "x-displayname": "Site",
                    "x-ves-example": "ce1"
                },
                "virtual_host": {
                    "type": "string",
                    "description": " Virtual Host name\n\nExample: - \"productpage\"-",
                    "title": "Virtual Host",
                    "x-displayname": "Virtual Host",
                    "x-ves-example": "productpage"
                }
            }
        },
        "service_policyServicePolicyHitsRequest": {
            "type": "object",
            "description": "Request to get the Service policy hits counter.",
            "title": "Service Policy Hits Request",
            "x-displayname": "Service Policy Hits Request",
            "x-ves-proto-message": "ves.io.schema.service_policy.ServicePolicyHitsRequest",
            "properties": {
                "end_time": {
                    "type": "string",
                    "description": " end time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the end_time will be evaluated to start_time+10m\n           If start_time is not specified, then the end_time will be evaluated to \u003ccurrent time\u003e\n\nExample: - \"1570007981\"-",
                    "title": "End time",
                    "x-displayname": "End Time",
                    "x-ves-example": "1570007981"
                },
                "group_by": {
                    "type": "array",
                    "description": " Aggregate data by one of more labels specified in -group_by-.\n\n Optional: If not specified, then the rule hits are aggregated/grouped by -POLICY-.",
                    "title": "Group by",
                    "items": {
                        "$ref": "#/definitions/service_policyServicePolicyMetricLabel"
                    },
                    "x-displayname": "Group By"
                },
                "label_filter": {
                    "type": "array",
                    "description": " List of label filter expressions of the form \"label\" -Op- \"value\".\n Response will only contain data that matches all the conditions specified in the -label_filter-.\n\n Optional: If not specified, then the metrics will be filtered only based on the -namespace- in the request.",
                    "title": "Label Filter",
                    "items": {
                        "$ref": "#/definitions/service_policyServicePolicyMetricLabelFilter"
                    },
                    "x-displayname": "Label Filter"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace is used to scope Service policy hits for the given namespace.\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "start_time": {
                    "type": "string",
                    "description": " start time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the start_time will be evaluated to end_time-10m\n           If end_time is not specified, then the start_time will be evaluated to \u003ccurrent time\u003e-10m\n\nExample: - \"1570007981\"-",
                    "title": "Start time",
                    "x-displayname": "Start Time",
                    "x-ves-example": "1570007981"
                },
                "step": {
                    "type": "string",
                    "description": " step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.\n The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn \u003c= end_time.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\n Optional: If not specified, then step size is evaluated to \u003cend_time - start_time\u003e\n\nExample: - \"5m\"-",
                    "title": "Step",
                    "x-displayname": "Step",
                    "x-ves-example": "5m"
                }
            }
        },
        "service_policyServicePolicyHitsResponse": {
            "type": "object",
            "description": "Number of Service policy rule hits for each unique combination of group_by labels in the request.",
            "title": "Service Policy Hits Response",
            "x-displayname": "Service Policy Hits Response",
            "x-ves-proto-message": "ves.io.schema.service_policy.ServicePolicyHitsResponse",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " List of Service policy hits data",
                    "title": "Service policy Hits",
                    "items": {
                        "$ref": "#/definitions/service_policyServicePolicyHits"
                    },
                    "x-displayname": "Service Policy Hits"
                }
            }
        },
        "service_policyServicePolicyMetricLabel": {
            "type": "string",
            "description": "Service policy hits can be sliced and diced based on one or more labels listed below.\n",
            "title": "Service Policy Metric Labels",
            "enum": [
                "NAMESPACE",
                "POLICY",
                "POLICY_RULE",
                "POLICY_SET",
                "ACTION",
                "SITE",
                "VIRTUAL_HOST"
            ],
            "default": "NAMESPACE",
            "x-displayname": "Service Policy Metric Labels",
            "x-ves-proto-enum": "ves.io.schema.service_policy.ServicePolicyMetricLabel"
        },
        "service_policyServicePolicyMetricLabelFilter": {
            "type": "object",
            "description": "Label filter can be specified to filter metrics based on label match",
            "title": "Service Policy Metric Label Filter",
            "x-displayname": "Service Policy Metric Label Filter",
            "x-ves-proto-message": "ves.io.schema.service_policy.ServicePolicyMetricLabelFilter",
            "properties": {
                "label": {
                    "description": " Label associated with Service policy hits",
                    "title": "Label",
                    "$ref": "#/definitions/service_policyServicePolicyMetricLabel",
                    "x-displayname": "Label"
                },
                "op": {
                    "description": " Operator to evaluate the label ",
                    "title": "Operator",
                    "$ref": "#/definitions/schemaMetricLabelOp",
                    "x-displayname": "Operator"
                },
                "value": {
                    "type": "string",
                    "description": " Value to be compared with\n\nExample: - \"policy1\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "policy1"
                }
            }
        }
    },
    "x-displayname": "Service Policy",
    "x-ves-proto-file": "ves.io/schema/service_policy/public_custom_data_api.proto"
}`
