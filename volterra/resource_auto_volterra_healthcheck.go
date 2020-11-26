//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_healthcheck "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/healthcheck"
)

// resourceVolterraHealthcheck is implementation of Volterra's Healthcheck resources
func resourceVolterraHealthcheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraHealthcheckCreate,
		Read:   resourceVolterraHealthcheckRead,
		Update: resourceVolterraHealthcheckUpdate,
		Delete: resourceVolterraHealthcheckDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"http_health_check": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"headers": {
							Type:     schema.TypeMap,
							Optional: true,
						},

						"host_header": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"use_origin_server_name": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"path": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"request_headers_to_remove": {

							Type: schema.TypeList,

							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"use_http2": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},

			"tcp_health_check": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"expected_response": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"send_payload": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"healthy_threshold": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"interval": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"jitter_percent": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},

			"unhealthy_threshold": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

// resourceVolterraHealthcheckCreate creates Healthcheck resource
func resourceVolterraHealthcheckCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_healthcheck.CreateSpecType{}
	createReq := &ves_io_schema_healthcheck.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	healthCheckTypeFound := false

	if v, ok := d.GetOk("http_health_check"); ok && !healthCheckTypeFound {

		healthCheckTypeFound = true
		healthCheckInt := &ves_io_schema_healthcheck.CreateSpecType_HttpHealthCheck{}
		healthCheckInt.HttpHealthCheck = &ves_io_schema_healthcheck.HttpHealthCheck{}
		createSpec.HealthCheck = healthCheckInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["headers"]; ok && !isIntfNil(v) {

				ms := map[string]string{}
				for k, v := range v.(map[string]interface{}) {
					ms[k] = v.(string)
				}
				healthCheckInt.HttpHealthCheck.Headers = ms
			}

			hostHeaderChoiceTypeFound := false

			if v, ok := cs["host_header"]; ok && !isIntfNil(v) && !hostHeaderChoiceTypeFound {

				hostHeaderChoiceTypeFound = true
				hostHeaderChoiceInt := &ves_io_schema_healthcheck.HttpHealthCheck_HostHeader{}

				healthCheckInt.HttpHealthCheck.HostHeaderChoice = hostHeaderChoiceInt

				hostHeaderChoiceInt.HostHeader = v.(string)

			}

			if v, ok := cs["use_origin_server_name"]; ok && !isIntfNil(v) && !hostHeaderChoiceTypeFound {

				hostHeaderChoiceTypeFound = true

				if v.(bool) {
					hostHeaderChoiceInt := &ves_io_schema_healthcheck.HttpHealthCheck_UseOriginServerName{}
					hostHeaderChoiceInt.UseOriginServerName = &ves_io_schema.Empty{}
					healthCheckInt.HttpHealthCheck.HostHeaderChoice = hostHeaderChoiceInt
				}

			}

			if v, ok := cs["path"]; ok && !isIntfNil(v) {

				healthCheckInt.HttpHealthCheck.Path = v.(string)
			}

			if v, ok := cs["request_headers_to_remove"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				healthCheckInt.HttpHealthCheck.RequestHeadersToRemove = ls

			}

			if v, ok := cs["use_http2"]; ok && !isIntfNil(v) {

				healthCheckInt.HttpHealthCheck.UseHttp2 = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("tcp_health_check"); ok && !healthCheckTypeFound {

		healthCheckTypeFound = true
		healthCheckInt := &ves_io_schema_healthcheck.CreateSpecType_TcpHealthCheck{}
		healthCheckInt.TcpHealthCheck = &ves_io_schema_healthcheck.TcpHealthCheck{}
		createSpec.HealthCheck = healthCheckInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expected_response"]; ok && !isIntfNil(v) {

				healthCheckInt.TcpHealthCheck.ExpectedResponse = v.(string)
			}

			if v, ok := cs["send_payload"]; ok && !isIntfNil(v) {

				healthCheckInt.TcpHealthCheck.SendPayload = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("healthy_threshold"); ok && !isIntfNil(v) {

		createSpec.HealthyThreshold =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("interval"); ok && !isIntfNil(v) {

		createSpec.Interval =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("jitter_percent"); ok && !isIntfNil(v) {

		createSpec.JitterPercent =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("timeout"); ok && !isIntfNil(v) {

		createSpec.Timeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("unhealthy_threshold"); ok && !isIntfNil(v) {

		createSpec.UnhealthyThreshold =
			uint32(v.(int))
	}

	log.Printf("[DEBUG] Creating Volterra Healthcheck object with struct: %+v", createReq)

	createHealthcheckResp, err := client.CreateObject(context.Background(), ves_io_schema_healthcheck.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Healthcheck: %s", err)
	}
	d.SetId(createHealthcheckResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraHealthcheckRead(d, meta)
}

func resourceVolterraHealthcheckRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_healthcheck.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Healthcheck %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Healthcheck %q: %s", d.Id(), err)
	}
	return setHealthcheckFields(client, d, resp)
}

func setHealthcheckFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraHealthcheckUpdate updates Healthcheck resource
func resourceVolterraHealthcheckUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_healthcheck.ReplaceSpecType{}
	updateReq := &ves_io_schema_healthcheck.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	healthCheckTypeFound := false

	if v, ok := d.GetOk("http_health_check"); ok && !healthCheckTypeFound {

		healthCheckTypeFound = true
		healthCheckInt := &ves_io_schema_healthcheck.ReplaceSpecType_HttpHealthCheck{}
		healthCheckInt.HttpHealthCheck = &ves_io_schema_healthcheck.HttpHealthCheck{}
		updateSpec.HealthCheck = healthCheckInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["headers"]; ok && !isIntfNil(v) {

				ms := map[string]string{}
				for k, v := range v.(map[string]interface{}) {
					ms[k] = v.(string)
				}
				healthCheckInt.HttpHealthCheck.Headers = ms
			}

			hostHeaderChoiceTypeFound := false

			if v, ok := cs["host_header"]; ok && !isIntfNil(v) && !hostHeaderChoiceTypeFound {

				hostHeaderChoiceTypeFound = true
				hostHeaderChoiceInt := &ves_io_schema_healthcheck.HttpHealthCheck_HostHeader{}

				healthCheckInt.HttpHealthCheck.HostHeaderChoice = hostHeaderChoiceInt

				hostHeaderChoiceInt.HostHeader = v.(string)

			}

			if v, ok := cs["use_origin_server_name"]; ok && !isIntfNil(v) && !hostHeaderChoiceTypeFound {

				hostHeaderChoiceTypeFound = true

				if v.(bool) {
					hostHeaderChoiceInt := &ves_io_schema_healthcheck.HttpHealthCheck_UseOriginServerName{}
					hostHeaderChoiceInt.UseOriginServerName = &ves_io_schema.Empty{}
					healthCheckInt.HttpHealthCheck.HostHeaderChoice = hostHeaderChoiceInt
				}

			}

			if v, ok := cs["path"]; ok && !isIntfNil(v) {

				healthCheckInt.HttpHealthCheck.Path = v.(string)
			}

			if v, ok := cs["request_headers_to_remove"]; ok && !isIntfNil(v) {

				ls := make([]string, len(v.([]interface{})))
				for i, v := range v.([]interface{}) {
					ls[i] = v.(string)
				}
				healthCheckInt.HttpHealthCheck.RequestHeadersToRemove = ls

			}

			if v, ok := cs["use_http2"]; ok && !isIntfNil(v) {

				healthCheckInt.HttpHealthCheck.UseHttp2 = v.(bool)
			}

		}

	}

	if v, ok := d.GetOk("tcp_health_check"); ok && !healthCheckTypeFound {

		healthCheckTypeFound = true
		healthCheckInt := &ves_io_schema_healthcheck.ReplaceSpecType_TcpHealthCheck{}
		healthCheckInt.TcpHealthCheck = &ves_io_schema_healthcheck.TcpHealthCheck{}
		updateSpec.HealthCheck = healthCheckInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["expected_response"]; ok && !isIntfNil(v) {

				healthCheckInt.TcpHealthCheck.ExpectedResponse = v.(string)
			}

			if v, ok := cs["send_payload"]; ok && !isIntfNil(v) {

				healthCheckInt.TcpHealthCheck.SendPayload = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("healthy_threshold"); ok && !isIntfNil(v) {

		updateSpec.HealthyThreshold =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("interval"); ok && !isIntfNil(v) {

		updateSpec.Interval =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("jitter_percent"); ok && !isIntfNil(v) {

		updateSpec.JitterPercent =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("timeout"); ok && !isIntfNil(v) {

		updateSpec.Timeout =
			uint32(v.(int))
	}

	if v, ok := d.GetOk("unhealthy_threshold"); ok && !isIntfNil(v) {

		updateSpec.UnhealthyThreshold =
			uint32(v.(int))
	}

	log.Printf("[DEBUG] Updating Volterra Healthcheck obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_healthcheck.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Healthcheck: %s", err)
	}

	return resourceVolterraHealthcheckRead(d, meta)
}

func resourceVolterraHealthcheckDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_healthcheck.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Healthcheck %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Healthcheck before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Healthcheck obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_healthcheck.ObjectType, namespace, name)
}
