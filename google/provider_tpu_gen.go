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

import "github.com/hashicorp/terraform/helper/schema"

// If the base path has changed as a result of your PR, make sure to update
// the provider_reference page!
var TpuDefaultBasePath = "https://tpu.googleapis.com/v1/"
var TpuCustomEndpointEntryKey = "tpu_custom_endpoint"
var TpuCustomEndpointEntry = &schema.Schema{
	Type:         schema.TypeString,
	Optional:     true,
	ValidateFunc: validateCustomEndpoint,
	DefaultFunc: schema.MultiEnvDefaultFunc([]string{
		"GOOGLE_TPU_CUSTOM_ENDPOINT",
	}, TpuDefaultBasePath),
}

var GeneratedTpuResourcesMap = map[string]*schema.Resource{
	"google_tpu_node": resourceTpuNode(),
}
