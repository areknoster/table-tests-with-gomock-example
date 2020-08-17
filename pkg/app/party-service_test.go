package app

import (
	"fmt"
	mock_party "github.com/areknoster/table-driven-tests-gomock/mocks/pkg/party"
	"github.com/areknoster/table-driven-tests-gomock/pkg/party"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPartyService_GreetVisitors_NotNiceReturnsError(t *testing.T) {
	// initialize gomock controller
	ctrl := gomock.NewController(t)
	// if not all expectations set on the controller are fulfilled at the end of function, the test will fail!
	defer ctrl.Finish()
	// init structure which implements namesLister interface
	mockedVisitorLister := mock_party.NewMockVisitorLister(ctrl)
	// mockedVisitorLister called once with names.NiceVisitor argument would return []string{"Peter"}, nil
	mockedVisitorLister.EXPECT().ListVisitors(party.NiceVisitor).Return([]party.Visitor{{"Peter", "TheSmart"}}, nil)
	// mockedVisitorLister called once with names.NotNiceVisitor argument would return nil and error
	mockedVisitorLister.EXPECT().ListVisitors(party.NotNiceVisitor).Return(nil, fmt.Errorf("dummyErr"))
	// mockedVisitorLister implements namesLister interface, so it can be assigned in PartyService
	sp := &PartyService{
		visitorLister: mockedVisitorLister,
	}
	gotErr := sp.GreetVisitors(false)
	if gotErr == nil {
		t.Errorf("did not get an error")
	}
}

func TestPartyService_GreetVisitors(t *testing.T) {
	type fields struct {
		namesLister *mock_party.MockVisitorLister
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
			name: "namesLister.ListVisitors(names.NiceVisitor) returns error, error expected",
			prepare: func(f *fields) {
				f.namesLister.EXPECT().ListVisitors(party.NiceVisitor).Return(nil, fmt.Errorf("dummyErr"))
			},
			args:    args{justNice: true},
			wantErr: true,
		},
		{
			name: "namesLister.ListVisitors(names.NotNiceVisitor) returns error, error expected",
			prepare: func(f *fields) {
				// if given calls do not happen in expected order, the test would fail!
				gomock.InOrder(
					f.namesLister.EXPECT().ListVisitors(party.NiceVisitor).Return([]string{"Peter"}, nil),
					f.namesLister.EXPECT().ListVisitors(party.NotNiceVisitor).Return(nil, fmt.Errorf("dummyErr")),
				)
			},
			args:    args{justNice: false},
			wantErr: true,
		},
		{
			name: " name of nice person, 1 name of not-nice person. greeter should be called with a nice person first, then with not-nice person as an argument",
			prepare: func(f *fields) {
				nice := []string{"Peter"}
				notNice := []string{"Buka"}
				gomock.InOrder(
					f.namesLister.EXPECT().ListVisitors(party.NiceVisitor).Return(nice, nil),
					f.namesLister.EXPECT().ListVisitors(party.NotNiceVisitor).Return(notNice, nil),
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
				namesLister: mock_party.NewMockVisitorLister(ctrl),
				greeter:     mock_party.NewMockGreeter(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			s := &PartyService{
				visitorLister: f.namesLister,
				greeter:       f.greeter,
			}
			if err := s.GreetVisitors(tt.args.justNice); (err != nil) != tt.wantErr {
				t.Errorf("GreetVisitors() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
