package capabilities

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_capabilities "github.com/kubewarden/policy-sdk-go/mock"
	"github.com/mailru/easyjson"
)

func TestPubKeyTrusted(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_capabilities.NewMockWapcClient(ctrl)

	image := "image"
	pubKeys := []string{"key"}
	annotations := map[string]string{}
	requestObj := sigstorePubKeysVerifyV2{
		Image:       image,
		PubKeys:     pubKeys,
		Annotations: annotations,
	}
	payload, err := easyjson.Marshal(requestObj)
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
		HostCall("kubewarden", "oci", V2.String(), payload).
		Return(verificationPayload, nil)

	host := Host{
		Client: m,
	}

	host.VerifyPubKeysImageV2(image, pubKeys, annotations)
}
