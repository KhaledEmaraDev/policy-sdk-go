package kubernetes

import (
	"fmt"

	cap "github.com/kubewarden/policy-sdk-go/pkg/capabilities"
)

type KubernetesHost struct {
	cap.Host
}

// ListResourcesByNamespace gets all the Kubernetes resources defined inside of
// the given namespace
// Note: cannot be used for cluster-wide resources
func (h *KubernetesHost) ListResourcesByNamespace(req ListResourcesByNamespaceRequest) ([]byte, error) {
	payload, err := req.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "kubernetes", "list_resources_by_namespace", payload)
	if err != nil {
		return []byte{}, err
	}

	return responsePayload, nil
}

// ListResources gets all the Kubernetes resources defined inside of the cluster.
// Note: this has be used for cluster-wide resources
func (h *KubernetesHost) ListResources(req ListAllResourcesRequest) ([]byte, error) {
	payload, err := req.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "kubernetes", "list_all_resources", payload)
	if err != nil {
		return []byte{}, err
	}

	return responsePayload, nil
}

// GetResource gets a specific Kubernetes resource.
func (h *KubernetesHost) GetResource(req GetResourceRequest) ([]byte, error) {
	payload, err := req.MarshalJSON()
	if err != nil {
		return []byte{}, fmt.Errorf("cannot serialize request object: %w", err)
	}

	// perform callback
	responsePayload, err := h.Client.HostCall("kubewarden", "kubernetes", "get_resource", payload)
	if err != nil {
		return []byte{}, err
	}

	return responsePayload, nil
}
