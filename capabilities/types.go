package capabilities

// Note: all the types defined inside of this file will get the easyjson
// interface contract be generated by `make generate-easyjson`

type KeylessInfo struct {
	// Issuer is identifier of the OIDC provider. E.g: https://github.com/login/oauth
	Issuer string `json:"issuer"`
	// Subject contains the information of the user used to authenticate against
	// the OIDC provider. E.g: mail@example.com
	Subject string `json:"subject"`
}

type VerificationResponse struct {
	// informs if the image was verified or not
	IsTrusted bool `json:"is_trusted"`
	// digest of the verified image
	Digest string `json:"digest"`
}

// sigstorePubKeysVerify represents the WaPC JSON contract, used for marshalling
// and unmarshalling payloads to wapc host calls
type sigstorePubKeysVerify struct {
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// List of PEM encoded keys that must have been used to sign the OCI object
	PubKeys []string `json:"pub_keys"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

type sigstorePubKeysVerifyRequest struct {
	SigstorePubKeysVerify sigstorePubKeysVerify `json:"SigstorePubKeyVerify"`
}

// sigstoreKeylessVerify represents the WaPC JSON contract, used for marshalling
// and unmarshalling payloads to wapc host calls
type sigstoreKeylessVerify struct {
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// List of PEM encoded keys that must have been used to sign the OCI object
	Keyless []KeylessInfo `json:"keyless"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

type sigstoreKeylessVerifyRequest struct {
	SigstoreKeylessVerify sigstoreKeylessVerify `json:"SigstoreKeylessVerify"`
}

type KeylessPrefixInfo struct {
	// Issuer is identifier of the OIDC provider. E.g: https://github.com/login/oauth
	Issuer string `json:"issuer"`
	// Valid prefix of the Subject field in the signature used to authenticate
	// against the OIDC provider. It forms a valid URL on its own, and will get
	// sanitized by appending `/` to protect against typosquatting
	UrlPrefix string `json:"url_prefix"`
}

// sigstorePubKeysVerifyV2 represents the WaPC JSON contract, used for marshalling
// and unmarshalling payloads to wapc host calls
type sigstorePubKeysVerifyV2 struct {
	Type SigstorePubKeyVerifyType `json:"type"`
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// List of PEM encoded keys that must have been used to sign the OCI object
	PubKeys []string `json:"pub_keys"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

// sigstoreKeylessVerifyExactV2 represents the WaPC JSON contract, used for marshalling
// and unmarshalling payloads to wapc host calls
type sigstoreKeylessVerifyExactV2 struct {
	Type SigstoreKeylessVerifyType `json:"type"`
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// List of PEM encoded keys that must have been used to sign the OCI object
	Keyless []KeylessInfo `json:"keyless"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

// sigstoreKeylessVerifyV2 represents the WaPC JSON contract, used for marshalling
// and unmarshalling payloads to wapc host calls
type sigstoreKeylessPrefixVerifyV2 struct {
	Type SigstoreKeylessPrefixVerifyType `json:"type"`
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// List of keyless signatures that must be found
	KeylessPrefix []KeylessPrefixInfo `json:"keyless_prefix"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

type sigstoreGithubActionsVerifyV2 struct {
	Type SigstoreGithubActionsVerifyType `json:"type"`
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// owner of the repository. E.g: octocat
	Owner string `json:"owner"`
	// Optional - Repo of the GH Action workflow that signed the artifact. E.g: example-repo
	Repo string `json:"repo"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

type sigstoreCertificateVerifyV2 struct {
	Type SigstoreCertificateVerifyType `json:"type"`
	// String pointing to the object (e.g.: `registry.testing.lan/busybox:1.0.0`)
	Image string `json:"image"`
	// PEM encoded certificate used to verify the signature
	Certificate RuneString `json:"certificate"`
	// Optional - the certificate chain that is used to verify the provided
	// certificate. When not specified, the certificate is assumed to be trusted
	CertificateChain []RuneString `json:"certificate_chain"`
	// Require the  signature layer to have a Rekor bundle.
	// Having a Rekor bundle allows further checks to be performed,
	// like ensuring the signature has been produced during the validity
	// time frame of the certificate.
	//
	// It is recommended to set this value to `true` to have a more secure
	// verification process.
	RequireRekorBundle bool `json:"require_rekor_bundle"`
	// Annotations that must have been provided by all signers when they signed
	// the OCI artifact. Optional
	Annotations map[string]string `json:"annotations"`
}

// We don't need to expose that to consumers of the library
// This is a glorified wrapper needed to unmarshal a string
// inside of TinyGo. As of release 0.29.0, unmarshal a simple
// string causes a runtime panic
type OciManifestResponse struct {
	// digest of the image
	Digest string `json:"digest"`
}

// We don't need to expose that to consumers of the library
// This is a glorified wrapper needed to unmarshal a list
// of string inside of TinyGo. As of release 0.29.0, unmarshal a simple
// list of string causes a runtime panic
type LookupHostResponse struct {
	// List of IP addresses associated with the host
	Ips []string `json:"ips"`
}

// Set of parameters used by the `list_resources_by_namespace` function
type ListResourcesByNamespaceRequest struct {
	// apiVersion of the resource (v1 for core group, groupName/groupVersions for other).
	APIVersion string `json:"api_version"`
	// Singular PascalCase name of the resource
	Kind string `json:"kind"`
	// Namespace scoping the search
	Namespace string `json:"namespace"`
	// A selector to restrict the list of returned objects by their labels.
	// Defaults to everything if omitted
	LabelSelector string `json:"label_selector,omitempty"`
	// A selector to restrict the list of returned objects by their fields.
	// Defaults to everything if omitted
	FieldSelector string `json:"field_selector,omitempty"`
}

// Set of parameters used by the `list_all_resources` function
type ListAllResourcesRequest struct {
	// apiVersion of the resource (v1 for core group, groupName/groupVersions for other).
	APIVersion string `json:"api_version"`
	// Singular PascalCase name of the resource
	Kind string `json:"kind"`
	// A selector to restrict the list of returned objects by their labels.
	// Defaults to everything if omitted
	LabelSelector string `json:"label_selector,omitempty"`
	// A selector to restrict the list of returned objects by their fields.
	// Defaults to everything if omitted
	FieldSelector string `json:"field_selector,omitempty"`
}

// Set of parameters used by the `get_resource` function
type GetResourceRequest struct {
	APIVersion string `json:"api_version"`
	// Singular PascalCase name of the resource
	Kind string `json:"kind"`
	// Namespace scoping the search
	Namespace string `json:"namespace"`
	// The name of the resource
	Name string `json:"name"`
	// Disable caching of results obtained from Kubernetes API Server
	// By default query results are cached for 5 seconds, that might cause
	// stale data to be returned.
	// However, making too many requests against the Kubernetes API Server
	// might cause issues to the cluster
	DisableCache bool `json:"disable_cache"`
}

// A x509 certificate
type Certificate struct {
	// Which encoding is used by the certificate
	Encoding CertificateEncoding `json:"encoding"`
	// Actual certificate
	Data RuneString `json:"data"`
}

// CertificateVerificationRequest holds information about a certificate and
// a chain to validate it with.
type CertificateVerificationRequest struct {
	/// PEM-encoded certificate
	Cert Certificate `json:"cert"`
	// list of PEM-encoded certs, ordered by trust usage (intermediates first, root last)
	// If empty, certificate is assumed trusted
	CertChain []Certificate `json:"cert_chain"`
	// RFC 3339 time format string, to check expiration against. If None,
	// certificate is assumed never expired
	NotAfter string `json:"not_after"`
}

type CertificateVerificationResponse struct {
	Trusted bool `json:"trusted"`
	// empty when trusted is true
	Reason string `json:"reason,omitempty"`
}
