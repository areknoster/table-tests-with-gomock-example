package app

import (
	"fmt"
	mock_app "github.com/areknoster/table-driven-tests-gomock/mocks/pkg/app"
	mock_party "github.com/areknoster/table-driven-tests-gomock/mocks/pkg/party"
	"github.com/areknoster/table-driven-tests-gomock/pkg/names"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPartyService_GreetVisitors(t *testing.T) {
	nice := []string{"Peter"}
	notNice := []string{"Buka"}
	type fields struct {
		namesLister *mock_app.MocknamesLister
		greeter     *mock_party.MockGreeter
	}
	type args struct {
		justNice bool
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		wantErr bool
	}{
		{
			name: "namesLister.ListNames(names.Nice) returns error, error expected",
			prepare: func(f *fields) {
				f.namesLister.EXPECT().ListNames(names.Nice).Return(nil, fmt.Errorf("dummyErr"))
			},
			args:    args{justNice: true},
			wantErr: true,
		},
		{
			name: "namesLister.ListNames(names.NotNice) returns error, error expected",
			prepare: func(f *fields) {
				// if given calls do not happen in expected order, the test would fail!
				gomock.InOrder(
					f.namesLister.EXPECT().ListNames(names.Nice).Return(nice, nil),
					f.namesLister.EXPECT().ListNames(names.NotNice).Return(nil, fmt.Errorf("dummyErr")),
				)
			},
			args:    args{justNice: false},
			wantErr: true,
		},
		{
			name: " name of nice person, 1 name of not-nice person. greeter should be called with a nice person first, then with not-nice person as an argument",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.namesLister.EXPECT().ListNames(names.Nice).Return(nice, nil),
					f.namesLister.EXPECT().ListNames(names.NotNice).Return(notNice, nil),
					f.greeter.EXPECT().Hello(nice[0]),
					f.greeter.EXPECT().Hello(notNice[0]),
				)
			},
			args:    args{justNice: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				namesLister: mock_app.NewMocknamesLister(ctrl),
				greeter:     mock_party.NewMockGreeter(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := &PartyService{
				namesLister: f.namesLister,
				greeter:     f.greeter,
			}
			if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
				t.Errorf("GreetVisitors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
