package app

import (
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
	"testing"
)

func TestService_WelcomeEveryone(t *testing.T) {
	type fields struct {
		namesLister namesLister
		helloer     party.Helloer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PartyService{
				namesLister: tt.fields.namesLister,
				helloer:     tt.fields.helloer,
			}
			if err := s.WelcomeEveryone(); (err != nil) != tt.wantErr {
				t.Errorf("WelcomeEveryone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
