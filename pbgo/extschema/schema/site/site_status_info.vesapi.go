//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package site

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

// Create CustomSiteStatusAPI GRPC Client satisfying server.CustomClient
type CustomSiteStatusAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomSiteStatusAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomSiteStatusAPIGrpcClient) doRPCSiteStatusMetrics(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SiteStatusMetricsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.SiteStatusMetricsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SiteStatusMetrics(ctx, req, opts...)
	return rsp, err
}

func (c *CustomSiteStatusAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomSiteStatusAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomSiteStatusAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomSiteStatusAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["SiteStatusMetrics"] = ccl.doRPCSiteStatusMetrics

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomSiteStatusAPI REST Client satisfying server.CustomClient
type CustomSiteStatusAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomSiteStatusAPIRestClient) doRPCSiteStatusMetrics(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SiteStatusMetricsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.SiteStatusMetricsRequest: %s", yamlReq, err)
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
		q.Add("field_selector", fmt.Sprintf("%v", req.FieldSelector))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("site", fmt.Sprintf("%v", req.Site))
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
	pbRsp := &SiteStatusMetricsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.site.SiteStatusMetricsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomSiteStatusAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomSiteStatusAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomSiteStatusAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["SiteStatusMetrics"] = ccl.doRPCSiteStatusMetrics

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomSiteStatusAPIInprocClient

// INPROC Client (satisfying CustomSiteStatusAPIClient interface)
type CustomSiteStatusAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomSiteStatusAPIInprocClient) SiteStatusMetrics(ctx context.Context, in *SiteStatusMetricsRequest, opts ...grpc.CallOption) (*SiteStatusMetricsResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.site.CustomSiteStatusAPI")
	cah, ok := ah.(CustomSiteStatusAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomSiteStatusAPISrv", ah)
	}

	var (
		rsp *SiteStatusMetricsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.site.SiteStatusMetricsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomSiteStatusAPI.SiteStatusMetrics' operation on 'site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.site.CustomSiteStatusAPI.SiteStatusMetrics"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SiteStatusMetrics(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.site.SiteStatusMetricsResponse", rsp)...)

	return rsp, nil
}

func NewCustomSiteStatusAPIInprocClient(svc svcfw.Service) CustomSiteStatusAPIClient {
	return &CustomSiteStatusAPIInprocClient{svc: svc}
}

// RegisterGwCustomSiteStatusAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomSiteStatusAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomSiteStatusAPIHandlerClient(ctx, mux, NewCustomSiteStatusAPIInprocClient(s))
}

var CustomSiteStatusAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Site Status Info",
        "description": "APIs to get status of a site or individual nodes in a multi-node site.\nSome of the status metrics for a site are listed below:\n- Cpu, Memory and Disk usage\n- In/Out throughput\n- Number of active flows and flow setup rate\n- Number active system and vK8s pods\n- Throughput to Regional Edges (REs)",
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
        "/public/namespaces/{namespace}/site/{site}/status/metrics": {
            "post": {
                "summary": "Site Status Metrics",
                "description": "Get status metrics for a site",
                "operationId": "ves.io.schema.site.CustomSiteStatusAPI.SiteStatusMetrics",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/siteSiteStatusMetricsResponse"
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
                        "name": "site",
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/siteSiteStatusMetricsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomSiteStatusAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-site-CustomSiteStatusAPI-SiteStatusMetrics"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.CustomSiteStatusAPI.SiteStatusMetrics"
            },
            "x-displayname": "Site Status API",
            "x-ves-proto-service": "ves.io.schema.site.CustomSiteStatusAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
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
        "siteSiteStatusMetricsData": {
            "type": "object",
            "description": "Site Status Data contains name of the field and the corresponding data",
            "title": "SiteStatusMetricsData",
            "x-displayname": "Site Status Metrics Data",
            "x-ves-proto-message": "ves.io.schema.site.SiteStatusMetricsData",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " List of metric data",
                    "title": "Data",
                    "items": {
                        "$ref": "#/definitions/siteSiteStatusMetricsFieldData"
                    },
                    "x-displayname": "Data"
                },
                "field": {
                    "description": " Site Status Metrics Field",
                    "title": "Field",
                    "$ref": "#/definitions/siteSiteStatusMetricsField",
                    "x-displayname": "Field"
                }
            }
        },
        "siteSiteStatusMetricsField": {
            "type": "string",
            "description": "Lists the fields that can queried in the site status metrics request\n",
            "title": "SiteStatusMetricsField",
            "enum": [
                "SITE_ACTIVE_FLOW_COUNT",
                "SITE_FLOW_SETUP_RATE",
                "SITE_IN_THROUGHPUT",
                "SITE_OUT_THROUGHPUT",
                "SITE_SYSTEM_ACTIVE_POD_COUNT",
                "SITE_VK8S_ACTIVE_POD_COUNT",
                "SITE_RE_IN_THROUGHPUT",
                "SITE_RE_OUT_THROUGHPUT",
                "SITE_SITE_IN_THROUGHPUT",
                "SITE_SITE_OUT_THROUGHPUT",
                "SITE_NODE_CPU_USAGE",
                "SITE_NODE_MEMORY_USAGE",
                "SITE_NODE_DISK_USAGE",
                "SITE_NODE_ACTIVE_FLOW_COUNT",
                "SITE_NODE_FLOW_SETUP_RATE",
                "SITE_NODE_IN_THROUGHPUT",
                "SITE_NODE_OUT_THROUGHPUT",
                "SITE_NODE_IF_IN_THROUGHPUT",
                "SITE_NODE_IF_OUT_THROUGHPUT"
            ],
            "default": "SITE_ACTIVE_FLOW_COUNT",
            "x-displayname": "Site Status Metrics Field",
            "x-ves-proto-enum": "ves.io.schema.site.SiteStatusMetricsField"
        },
        "siteSiteStatusMetricsFieldData": {
            "type": "object",
            "description": "Site Status Field Data contains key/value pair for a field",
            "title": "SiteStatusMetricsFieldData",
            "x-displayname": "Site Status Metrics Field Data",
            "x-ves-proto-message": "ves.io.schema.site.SiteStatusMetricsFieldData",
            "properties": {
                "key": {
                    "type": "object",
                    "description": " Key contains name/value pair.\n For SITE_RE_IN_THROUGHPUT, the key will be \"name\": \"\u003cRE-name\u003e\"\n For SITE_NODE_IF_IN_THROUGHPUT, the key will be \"name\": \"\u003cInterface-name\u003e\"",
                    "title": "Key",
                    "x-displayname": "Key"
                },
                "value": {
                    "type": "array",
                    "description": " List of metric values",
                    "title": "Value",
                    "items": {
                        "$ref": "#/definitions/schemaMetricValue"
                    },
                    "x-displayname": "Value"
                }
            }
        },
        "siteSiteStatusMetricsRequest": {
            "type": "object",
            "description": "Request to get status metrics for a site",
            "title": "Site Status Metrics Request",
            "x-displayname": "Site Status Metrics Request",
            "x-ves-proto-message": "ves.io.schema.site.SiteStatusMetricsRequest",
            "properties": {
                "end_time": {
                    "type": "string",
                    "description": " end time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the end_time will be evaluated to start_time+10m\n           If start_time is not specified, then the end_time will be evaluated to \u003ccurrent time\u003e\n\nExample: - \"1570007981\"-",
                    "title": "End time",
                    "x-displayname": "End Time",
                    "x-ves-example": "1570007981"
                },
                "field_selector": {
                    "type": "array",
                    "description": " List of fields to be returned in the response\nRequired: YES",
                    "title": "Field Selector",
                    "items": {
                        "$ref": "#/definitions/siteSiteStatusMetricsField"
                    },
                    "x-displayname": "Field Selector",
                    "x-ves-required": "true"
                },
                "namespace": {
                    "type": "string",
                    "description": " Only system namespace is valid for this request\n\nExample: - \"system\"-\nRequired: YES",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "system",
                    "x-ves-required": "true"
                },
                "site": {
                    "type": "string",
                    "description": " Name of the site\n\nExample: - \"ce01\"-\nRequired: YES",
                    "title": "Site",
                    "x-displayname": "Site",
                    "x-ves-example": "ce01",
                    "x-ves-required": "true"
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
                    "description": " step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.\n The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn \u003c= end_time.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\n Optional: If not specified, then step size is evaluated to \u003cend_time - start_time\u003e\n\nExample: - \"15m\"-",
                    "title": "Step",
                    "x-displayname": "Step",
                    "x-ves-example": "15m"
                }
            }
        },
        "siteSiteStatusMetricsResponse": {
            "type": "object",
            "description": "Response for the Site Status Metrics Request",
            "title": "Site Status Metrics Response",
            "x-displayname": "Site Status Metrics Response",
            "x-ves-proto-message": "ves.io.schema.site.SiteStatusMetricsResponse",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " Data contains time-series data for the status fields in the request",
                    "title": "Data",
                    "items": {
                        "$ref": "#/definitions/siteSiteStatusMetricsData"
                    },
                    "x-displayname": "Data"
                }
            }
        }
    },
    "x-displayname": "Site",
    "x-ves-proto-file": "ves.io/schema/site/site_status_info.proto"
}`
