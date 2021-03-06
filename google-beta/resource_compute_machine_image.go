// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceComputeMachineImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeMachineImageCreate,
		Read:   resourceComputeMachineImageRead,
		Update: resourceComputeMachineImageUpdate,
		Delete: resourceComputeMachineImageDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeMachineImageImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the resource.`,
			},
			"source_instance": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The source instance used to create the machine image. You can provide this as a partial or full URL to the resource.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A text description of the resource.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeMachineImageCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	nameProp, err := expandComputeMachineImageName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeMachineImageDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceInstanceProp, err := expandComputeMachineImageSourceInstance(d.Get("source_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_instance"); !isEmptyValue(reflect.ValueOf(sourceInstanceProp)) && (ok || !reflect.DeepEqual(v, sourceInstanceProp)) {
		obj["sourceInstance"] = sourceInstanceProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new MachineImage: %#v", obj)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequestWithTimeout(config, "POST", project, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating MachineImage: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating MachineImage %q: %#v", d.Id(), res)

	return resourceComputeMachineImageRead(d, meta)
}

func resourceComputeMachineImageRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	res, err := sendRequest(config, "GET", project, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeMachineImage %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}

	if err := d.Set("name", flattenComputeMachineImageName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("description", flattenComputeMachineImageDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("source_instance", flattenComputeMachineImageSourceInstance(res["sourceInstance"], d, config)); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading MachineImage: %s", err)
	}

	return nil
}

func resourceComputeMachineImageUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandComputeMachineImageName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeMachineImageDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceInstanceProp, err := expandComputeMachineImageSourceInstance(d.Get("source_instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("source_instance"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, sourceInstanceProp)) {
		obj["sourceInstance"] = sourceInstanceProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating MachineImage %q: %#v", d.Id(), obj)
	_, err = sendRequestWithTimeout(config, "PUT", project, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating MachineImage %q: %s", d.Id(), err)
	}

	return resourceComputeMachineImageRead(d, meta)
}

func resourceComputeMachineImageDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting MachineImage %q", d.Id())

	res, err := sendRequestWithTimeout(config, "DELETE", project, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "MachineImage")
	}

	log.Printf("[DEBUG] Finished deleting MachineImage %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeMachineImageImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/machineImages/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/machineImages/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeMachineImageName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeMachineImageDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeMachineImageSourceInstance(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeMachineImageName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeMachineImageSourceInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseZonalFieldValue("instances", v.(string), "project", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for source_instance: %s", err)
	}
	return f.RelativeLink(), nil
}
