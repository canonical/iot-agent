package identity

import (
	"github.com/CanonicalLtd/iot-agent/config"
	"github.com/CanonicalLtd/iot-agent/snapdapi"
	"github.com/CanonicalLtd/iot-identity/domain"
	"os"
	"testing"
)

func TestService_StoreCredentials(t *testing.T) {
	settings := config.ReadParameters()
	_ = os.Remove(settings.CredentialsPath)
	_ = os.Remove("params")

	tests := []struct {
		name     string
		wantErr  bool
		snapdErr bool
	}{
		{"valid", false, false},
		{"invalid", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snapd := &snapdapi.MockClient{WithError: tt.snapdErr}
			srv := &Service{
				Settings: settings,
				Snapd:    snapd,
			}

			enroll := domain.Enrollment{}

			err := srv.storeCredentials(enroll)

			if (err != nil) != tt.wantErr {
				t.Errorf("Service.storeCredentials error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if _, err = os.Stat(settings.CredentialsPath); err != nil {
					t.Error("Service.storeCredentials error = empty enrollment")
				}
			}

			_ = os.Remove(settings.CredentialsPath)
			_ = os.Remove("params")
		})
	}
}
