package service

import (
	"frete-rapido/src/config"
	mongodb "frete-rapido/src/db"
	domain "frete-rapido/src/domain"
	"reflect"
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		request domain.RequestQuote
	}
	tests := []struct {
		name     string
		args     args
		wantArgs []string
	}{
		{
			name: "Test Zipcode and Volumes",
			args: args{
				request: domain.RequestQuote{
					Recipient: struct {
						Address struct {
							Zipcode string "json:\"zipcode\""
						} "json:\"address\""
					}{Address: struct {
						Zipcode string "json:\"zipcode\""
					}{Zipcode: ""}},
					Volumes: []domain.Volume{},
				},
			},
			wantArgs: []string{"Zipcode is required", "Volumes is required"},
		},
		{
			name: "Test Volumes",
			args: args{
				request: domain.RequestQuote{
					Recipient: struct {
						Address struct {
							Zipcode string "json:\"zipcode\""
						} "json:\"address\""
					}{Address: struct {
						Zipcode string "json:\"zipcode\""
					}{Zipcode: "12345678"}},
					Volumes: []domain.Volume{
						{
							Category:      "",
							Amount:        0,
							UnitaryWeight: 0,
							Price:         0,
							Sku:           "",
							Height:        0,
							Width:         0,
							Length:        0,
						},
					},
				},
			},
			wantArgs: []string{"Category is required for Volume 0", "Amount is required for Volume 0", "UnitaryWeight is required for Volume 0", "Price is required for Volume 0", "Sku is required for Volume 0", "Height is required for Volume 0", "Width is required for Volume 0", "Length is required for Volume 0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArgs := Check(tt.args.request); !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("Check() = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	type args struct {
		responseAPI domain.ResponseAPI
	}
	tests := []struct {
		name              string
		args              args
		wantResponseQuote domain.ResponseQuote
	}{
		{
			name: "Test Format",
			args: args{
				responseAPI: domain.ResponseAPI{
					Dispatchers: []domain.DispatcherAPI{
						{
							Offers: []domain.Offer{
								{
									Modal:      "Test",
									FinalPrice: 9.9,
								},
							},
						},
					},
				},
			},
			wantResponseQuote: domain.ResponseQuote{
				Carrier: []domain.CarrierQuote{
					{
						Name:     "",
						Service:  "Test",
						Deadline: 0,
						Price:    9.9,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResponseQuote := Format(tt.args.responseAPI); !reflect.DeepEqual(gotResponseQuote, tt.wantResponseQuote) {
				t.Errorf("Format() = %v, want %v", gotResponseQuote, tt.wantResponseQuote)
			}
		})
	}
}

func TestPrepare(t *testing.T) {
	type args struct {
		quotes []mongodb.QuoteBD
	}
	tests := []struct {
		name                string
		args                args
		wantResponseMetrics domain.ResponseMetrics
		wantErr             bool
	}{
		{
			name: "Test Prepare",
			args: args{
				quotes: []mongodb.QuoteBD{
					{
						Carriers: []mongodb.Carrier{
							{
								Name:     "Test",
								Service:  "Test",
								Deadline: 1,
								Price:    9.9,
							},
						},
					},
				},
			},
			// wantResponseMetrics: domain.ResponseMetrics{
			// 	Metrics: []domain.Metric{
			// 		{
			// 			ResultsPerCarrier:    map[string]int{"Test": 1},
			// 			TotalPricePerCarrier: map[string]float64{"Test": 9.9},
			// 			AvgPricePerCarrier:   map[string]float64{"Test": 9.9},
			// 			CheapestFreight:      map[string]float64{"Test": 9.9},
			// 		},
			// 	},
			// },
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Prepare(tt.args.quotes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResponseMetrics, tt.wantResponseMetrics) {
			// 	t.Errorf("Prepare() = %v, want %v", gotResponseMetrics, tt.wantResponseMetrics)
			// }
		})
	}
}

func TestBuild(t *testing.T) {
	type args struct {
		requestQuote domain.RequestQuote
	}
	tests := []struct {
		name           string
		args           args
		wantRequestAPI domain.RequestAPI
	}{
		{
			name: "Test Build",
			args: args{
				requestQuote: domain.RequestQuote{
					Recipient: struct {
						Address struct {
							Zipcode string "json:\"zipcode\""
						} "json:\"address\""
					}{Address: struct {
						Zipcode string "json:\"zipcode\""
					}{Zipcode: "12345678"}},
					Volumes: []domain.Volume{
						{
							Category:      "Test",
							Amount:        1,
							UnitaryWeight: 1,
							Price:         1,
							Sku:           "Test",
							Height:        1,
							Width:         1,
							Length:        1,
						},
					},
				},
			},
			wantRequestAPI: domain.RequestAPI{
				Shipper: struct {
					RegisteredNumber string "json:\"registered_number\""
					Token            string "json:\"token\""
					PlatformCode     string "json:\"platform_code\""
				}{RegisteredNumber: config.RegisteredNumber, Token: config.Token, PlatformCode: config.PlatformCode},
				Recipient: struct {
					Type    int    "json:\"type\""
					Country string "json:\"country\""
					Zipcode int    "json:\"zipcode\""
				}{Type: 0, Country: "BRA", Zipcode: 12345678},
				Dispatchers: []domain.Dispatcher{
					{
						RegisteredNumber: config.RegisteredNumber,
						Zipcode:          12345678,
						Volumes: []domain.Volume{
							{
								Category:      "Test",
								Amount:        1,
								UnitaryWeight: 1,
								UnitaryPrice:  1,
								Price:         1,
								Sku:           "Test",
								Height:        1,
								Width:         1,
								Length:        1,
							},
						},
					},
				},
				SimulationType: []int{0},
				Returns: struct {
					Composition  bool "json:\"composition\""
					Volumes      bool "json:\"volumes\""
					AppliedRules bool "json:\"applied_rules\""
				}{Composition: false, Volumes: false, AppliedRules: false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRequestAPI := Build(tt.args.requestQuote); !reflect.DeepEqual(gotRequestAPI, tt.wantRequestAPI) {
				t.Errorf("Build() = %v, want %v", gotRequestAPI, tt.wantRequestAPI)
			}
		})
	}
}
