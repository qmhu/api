// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	analysisv1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/analysis/v1alpha1"
	autoscalingv1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/autoscaling/v1alpha1"
	co2ev1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/co2e/v1alpha1"
	ensurancev1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/ensurance/v1alpha1"
	predictionv1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/prediction/v1alpha1"
	topologyv1alpha1 "github.com/gocrane/api/pkg/generated/clientset/versioned/typed/topology/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	AnalysisV1alpha1() analysisv1alpha1.AnalysisV1alpha1Interface
	AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface
	Co2eV1alpha1() co2ev1alpha1.Co2eV1alpha1Interface
	EnsuranceV1alpha1() ensurancev1alpha1.EnsuranceV1alpha1Interface
	PredictionV1alpha1() predictionv1alpha1.PredictionV1alpha1Interface
	TopologyV1alpha1() topologyv1alpha1.TopologyV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	analysisV1alpha1    *analysisv1alpha1.AnalysisV1alpha1Client
	autoscalingV1alpha1 *autoscalingv1alpha1.AutoscalingV1alpha1Client
	co2eV1alpha1        *co2ev1alpha1.Co2eV1alpha1Client
	ensuranceV1alpha1   *ensurancev1alpha1.EnsuranceV1alpha1Client
	predictionV1alpha1  *predictionv1alpha1.PredictionV1alpha1Client
	topologyV1alpha1    *topologyv1alpha1.TopologyV1alpha1Client
}

// AnalysisV1alpha1 retrieves the AnalysisV1alpha1Client
func (c *Clientset) AnalysisV1alpha1() analysisv1alpha1.AnalysisV1alpha1Interface {
	return c.analysisV1alpha1
}

// AutoscalingV1alpha1 retrieves the AutoscalingV1alpha1Client
func (c *Clientset) AutoscalingV1alpha1() autoscalingv1alpha1.AutoscalingV1alpha1Interface {
	return c.autoscalingV1alpha1
}

// Co2eV1alpha1 retrieves the Co2eV1alpha1Client
func (c *Clientset) Co2eV1alpha1() co2ev1alpha1.Co2eV1alpha1Interface {
	return c.co2eV1alpha1
}

// EnsuranceV1alpha1 retrieves the EnsuranceV1alpha1Client
func (c *Clientset) EnsuranceV1alpha1() ensurancev1alpha1.EnsuranceV1alpha1Interface {
	return c.ensuranceV1alpha1
}

// PredictionV1alpha1 retrieves the PredictionV1alpha1Client
func (c *Clientset) PredictionV1alpha1() predictionv1alpha1.PredictionV1alpha1Interface {
	return c.predictionV1alpha1
}

// TopologyV1alpha1 retrieves the TopologyV1alpha1Client
func (c *Clientset) TopologyV1alpha1() topologyv1alpha1.TopologyV1alpha1Interface {
	return c.topologyV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.analysisV1alpha1, err = analysisv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.autoscalingV1alpha1, err = autoscalingv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.co2eV1alpha1, err = co2ev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.ensuranceV1alpha1, err = ensurancev1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.predictionV1alpha1, err = predictionv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.topologyV1alpha1, err = topologyv1alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.analysisV1alpha1 = analysisv1alpha1.New(c)
	cs.autoscalingV1alpha1 = autoscalingv1alpha1.New(c)
	cs.co2eV1alpha1 = co2ev1alpha1.New(c)
	cs.ensuranceV1alpha1 = ensurancev1alpha1.New(c)
	cs.predictionV1alpha1 = predictionv1alpha1.New(c)
	cs.topologyV1alpha1 = topologyv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
