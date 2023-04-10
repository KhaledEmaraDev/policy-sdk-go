package verify_v1

import (
	oci "github.com/kubewarden/policy-sdk-go/pkg/capabilities/oci"
)

type VerifyV1Host struct {
	oci.OciHost
}

// VerifyPubKeys verifies sigstore signatures of an image using public keys
// Arguments
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
// * pubKeys: list of PEM encoded keys that must have been used to sign the OCI object
// * annotations: annotations that must have been provided by all signers when they signed the OCI artifact
func (h *VerifyV1Host) VerifyPubKeys(image string, pubKeys []string, annotations map[string]string) (oci.VerificationResponse, error) {
	requestObj := sigstorePubKeysVerifyRequest{
		SigstorePubKeysVerify: sigstorePubKeysVerify{
			Image:       image,
			PubKeys:     pubKeys,
			Annotations: annotations,
		},
	}

	return h.Verify(requestObj, oci.V1)
}

// VerifyKeyless verifies sigstore signatures of an image using keyless signing
// Arguments
// * image: image to be verified (e.g.: `registry.testing.lan/busybox:1.0.0`)
// * keyless: list of KeylessInfo pairs, containing Issuer and Subject info from OIDC providers
// * annotations: annotations that must have been provided by all signers when they signed the OCI artifact
func (h *VerifyV1Host) VerifyKeyless(image string, keyless []oci.KeylessInfo, annotations map[string]string) (oci.VerificationResponse, error) {
	requestObj := sigstoreKeylessVerifyRequest{
		SigstoreKeylessVerify: sigstoreKeylessVerify{
			Image:       image,
			Keyless:     keyless,
			Annotations: annotations,
		},
	}

	return h.Verify(requestObj, oci.V1)
}
