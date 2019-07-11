package ironic

import (
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/baremetal/noauth"
	noauthintrospection "github.com/gophercloud/gophercloud/openstack/baremetalintrospection/noauth"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"net/http"
	"time"
)

// Clients stores the client connection information for Ironic and Inspector
type Clients struct {
	ironic    *gophercloud.ServiceClient
	inspector *gophercloud.ServiceClient

	ironicUp    bool
	inspectorUp bool
	timeout     time.Duration
}

// GetIronicClient returns the API client for Ironic, optionally retrying to reach the API if timeout is set.
func (c *Clients) GetIronicClient() (*gophercloud.ServiceClient, error) {
	if c.ironicUp || c.timeout == 0 {
		return c.ironic, nil
	}

	err := waitForAPI(c.timeout, c.ironic)
	if err != nil {
		return nil, err

	}
	c.ironicUp = true
	return c.ironic, nil
}

// GetInspectorClient returns the API client for Ironic, optionally retrying to reach the API if timeout is set.
func (c *Clients) GetInspectorClient() (*gophercloud.ServiceClient, error) {
	if c.inspector == nil {
		return nil, fmt.Errorf("no inspector endpoint was specified")
	} else if c.inspectorUp || c.timeout == 0 {
		return c.inspector, nil
	}

	err := waitForAPI(c.timeout, c.inspector)
	if err != nil {
		return nil, err
	}
	c.inspectorUp = true
	return c.inspector, nil
}

// Provider Ironic
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IRONIC_ENDPOINT", ""),
				Description: descriptions["url"],
			},
			"inspector": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("IRONIC_INSPECTOR_ENDPOINT", ""),
				Description: descriptions["inspector"],
			},
			"microversion": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("IRONIC_MICROVERSION", "1.52"),
				Description: descriptions["microversion"],
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: descriptions["timeout"],
				Default:     0,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ironic_node_v1":       resourceNodeV1(),
			"ironic_port_v1":       resourcePortV1(),
			"ironic_allocation_v1": resourceAllocationV1(),
			"ironic_deployment":    resourceDeployment(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ironic_introspection": dataSourceIronicIntrospection(),
		},
		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"url":          "The authentication endpoint for Ironic",
		"inspector":    "The endpoint for Ironic inspector",
		"microversion": "The microversion to use for Ironic",
		"timeout":      "Wait the specified number of seconds for the API to become available",
	}
}

// Creates a noauth Ironic client
func configureProvider(schema *schema.ResourceData) (interface{}, error) {
	var clients Clients

	url := schema.Get("url").(string)
	if url == "" {
		return nil, fmt.Errorf("url is required for ironic provider")
	}
	log.Printf("[DEBUG] Ironic endpoint is %s", url)

	ironic, err := noauth.NewBareMetalNoAuth(noauth.EndpointOpts{
		IronicEndpoint: url,
	})
	if err != nil {
		return nil, err
	}
	ironic.Microversion = schema.Get("microversion").(string)
	clients.ironic = ironic

	inspectorURL := schema.Get("inspector").(string)
	if inspectorURL != "" {
		log.Printf("[DEBUG] Inspector endpoint is %s", inspectorURL)
		inspector, err := noauthintrospection.NewBareMetalIntrospectionNoAuth(noauthintrospection.EndpointOpts{
			IronicInspectorEndpoint: inspectorURL,
		})
		if err != nil {
			return nil, fmt.Errorf("could not configure inspector endpoint: %s", err.Error())
		}
		clients.inspector = inspector
	}

	clients.timeout = time.Duration(schema.Get("timeout").(int)) * time.Second

	return &clients, err
}

// Retries a gophercloud API until a timeout is reached.
func waitForAPI(duration time.Duration, client *gophercloud.ServiceClient) error {
	success := make(chan bool, 1)
	timeout := make(chan bool, 1)
	defer close(success)
	defer close(timeout)

	go func() {
		for tries := 0; ; tries++ {
			// Return if we've hit the timeout
			select {
			case <-timeout:
				return
			default:
			}

			r, _ := client.HTTPClient.Get(client.Endpoint)
			if r.StatusCode == http.StatusOK {
				success <- true
				return
			}

			log.Printf("[DEBUG] Retrying API, attempt %d failed", tries)
			time.Sleep(5 * time.Second)
		}
	}()

	// Block until success or timeout is reached
	select {
	case _, ok := <-success:
		if !ok {
			return fmt.Errorf("could not contact API: channel unexpectedly closed")
		}
	case <-time.After(duration):
		timeout <- true
		return fmt.Errorf("could not contact API: timeout reached")
	}

	return nil
}
