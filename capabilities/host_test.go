package capabilities

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_capabilities "github.com/kubewarden/policy-sdk-go/mock"
	"github.com/mailru/easyjson"
)

type v2VerifyTestCase struct {
	request            easyjson.Marshaler
	checkIsTrustedFunc func(host Host, request easyjson.Marshaler) (bool, error)
}

func TestV2Verify(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_capabilities.NewMockWapcClient(ctrl)

	for description, testCase := range map[string]v2VerifyTestCase{
		"PubKeysImage": {
			request: sigstorePubKeysVerifyV2{
				Image:       "image",
				PubKeys:     []string{"key"},
				Annotations: map[string]string{},
			},
			checkIsTrustedFunc: CheckPubKeysImageTrusted,
		},
		"KeylessExactMatch": {
			request: sigstoreKeylessVerifyExactV2{
				Image: "image",
				Keyless: []KeylessInfo{
					{Issuer: "https://github.com/login/oauth", Subject: "mail@example.com"},
				},
				Annotations: map[string]string{},
			},
			checkIsTrustedFunc: CheckKeylessExactMatchTrusted,
		},
		"KeylessPrefixMatch": {
			request: sigstoreKeylessPrefixVerifyV2{
				Image: "image",
				KeylessPrefix: []KeylessPrefixInfo{
					{Issuer: "https://github.com/login/oauth", UrlPrefix: "https://example.com"},
				},
				Annotations: map[string]string{},
			},
			checkIsTrustedFunc: CheckKeylessPrefixMatchTrusted,
		},
		"KeylessGithubActions": {
			request: sigstoreGithubActionsVerifyV2{
				Image:       "image",
				Owner:       "owner",
				Repo:        "repo",
				Annotations: map[string]string{},
			},
			checkIsTrustedFunc: CheckKeylessGithubActionsTrusted,
		},
		"Certificate": {
			request: sigstoreCertificateVerifyV2{
				Image:              "image",
				Certificate:        []rune("certificate"),
				CertificateChain:   make([][]rune, 0),
				RequireRekorBundle: true,
				Annotations:        map[string]string{},
			},
			checkIsTrustedFunc: CheckCertificateTrusted,
		},
	} {
		t.Run(description, func(t *testing.T) {
			requestPayload, err := easyjson.Marshal(testCase.request)
			if err != nil {
				t.Fatalf("cannot serialize request object: %v", err)
			}

			verificationResponse := VerificationResponse{
				IsTrusted: true,
				Digest:    "",
			}
			verificationPayload, err := easyjson.Marshal(verificationResponse)
			if err != nil {
				t.Fatalf("cannot serialize response object: %v", err)
			}

			m.
				EXPECT().
				HostCall("kubewarden", "oci", V2.String(), requestPayload).
				Return(verificationPayload, nil).
				Times(1)

			host := Host{
				Client: m,
			}

			res, err := testCase.checkIsTrustedFunc(host, testCase.request)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !res {
				t.Fatalf("expected trusted image, got untrusted")
			}
		})
	}
}

func CheckPubKeysImageTrusted(host Host, request easyjson.Marshaler) (bool, error) {
	requestPubKeys := request.(sigstorePubKeysVerifyV2)
	res, err := host.VerifyPubKeysImageV2(requestPubKeys.Image, requestPubKeys.PubKeys, requestPubKeys.Annotations)
	if err != nil {
		return false, err
	}
	return res.IsTrusted, nil
}

func CheckKeylessExactMatchTrusted(host Host, request easyjson.Marshaler) (bool, error) {
	requestKeylessExactMatch := request.(sigstoreKeylessVerifyExactV2)
	res, err := host.VerifyKeylessExactMatchV2(requestKeylessExactMatch.Image, requestKeylessExactMatch.Keyless, requestKeylessExactMatch.Annotations)
	if err != nil {
		return false, err
	}
	return res.IsTrusted, nil
}

func CheckKeylessPrefixMatchTrusted(host Host, request easyjson.Marshaler) (bool, error) {
	requestKeylessPrefixMatch := request.(sigstoreKeylessPrefixVerifyV2)
	res, err := host.VerifyKeylessPrefixMatchV2(requestKeylessPrefixMatch.Image, requestKeylessPrefixMatch.KeylessPrefix, requestKeylessPrefixMatch.Annotations)
	if err != nil {
		return false, err
	}
	return res.IsTrusted, nil
}

func CheckKeylessGithubActionsTrusted(host Host, request easyjson.Marshaler) (bool, error) {
	requestKeylessGithubActions := request.(sigstoreGithubActionsVerifyV2)
	res, err := host.VerifyKeylessGithubActionsV2(requestKeylessGithubActions.Image, requestKeylessGithubActions.Owner, requestKeylessGithubActions.Repo, requestKeylessGithubActions.Annotations)
	if err != nil {
		return false, err
	}
	return res.IsTrusted, nil
}

func CheckCertificateTrusted(host Host, request easyjson.Marshaler) (bool, error) {
	requestCertificate := request.(sigstoreCertificateVerifyV2)

	chain := make([]string, len(requestCertificate.CertificateChain))
	for i, c := range requestCertificate.CertificateChain {
		chain[i] = string(c)
	}

	res, err := host.VerifyCertificateV2(requestCertificate.Image, string(requestCertificate.Certificate), chain, requestCertificate.RequireRekorBundle, requestCertificate.Annotations)
	if err != nil {
		return false, err
	}
	return res.IsTrusted, nil
}
