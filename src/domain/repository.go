package repository

import "time"

// RequestQuote -
type RequestQuote struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode"`
		} `json:"address"`
	} `json:"recipient"`
	Volumes []Volume `json:"volumes"`
}

// RequestAPI -
type RequestAPI struct {
	Shipper struct {
		RegisteredNumber string `json:"registered_number"`
		Token            string `json:"token"`
		PlatformCode     string `json:"platform_code"`
	} `json:"shipper"`
	Recipient struct {
		Type    int    `json:"type"`
		Country string `json:"country"`
		Zipcode int    `json:"zipcode"`
	} `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	SimulationType []int        `json:"simulation_type"`
	Returns        struct {
		Composition  bool `json:"composition"`
		Volumes      bool `json:"volumes"`
		AppliedRules bool `json:"applied_rules"`
	} `json:"returns"`
}

// Volume -
type Volume struct {
	Category      string  `json:"category"`
	Amount        int     `json:"amount"`
	UnitaryWeight int     `json:"unitary_weight"`
	Price         int     `json:"price"`
	UnitaryPrice  int     `json:"unitary_price"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
}

// Dispatcher -
type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          int      `json:"zipcode"`
	Volumes          []Volume `json:"volumes"`
}

// ResponseAPI -
type ResponseAPI struct {
	Dispatchers []struct {
		ID                         string  `json:"id"`
		RequestID                  string  `json:"request_id"`
		RegisteredNumberShipper    string  `json:"registered_number_shipper"`
		RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
		ZipcodeOrigin              int     `json:"zipcode_origin"`
		Offers                     []Offer `json:"offers"`
	} `json:"dispatchers"`
}

// ResponseDispatcher -
type ResponseDispatcher struct {
	ID                         string  `json:"id"`
	RequestID                  string  `json:"request_id"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int     `json:"zipcode_origin"`
	Offers                     []Offer `json:"offers"`
}

// Offer -
type Offer struct {
	Offer          int    `json:"offer"`
	TableReference string `json:"table_reference"`
	SimulationType int    `json:"simulation_type"`
	Carrier        struct {
		Name             string `json:"name"`
		RegisteredNumber string `json:"registered_number"`
		StateInscription string `json:"state_inscription"`
		Logo             string `json:"logo"`
		Reference        int    `json:"reference"`
		CompanyName      string `json:"company_name"`
	} `json:"carrier"`
	Service      string `json:"service"`
	DeliveryTime struct {
		Days          int    `json:"days"`
		EstimatedDate string `json:"estimated_date"`
	} `json:"delivery_time,omitempty"`
	Expiration time.Time `json:"expiration"`
	CostPrice  float64   `json:"cost_price"`
	FinalPrice float64   `json:"final_price"`
	Weights    struct {
		Real  int     `json:"real"`
		Cubed float64 `json:"cubed"`
		Used  float64 `json:"used"`
	} `json:"weights"`
	OriginalDeliveryTime struct {
		Days          int    `json:"days"`
		EstimatedDate string `json:"estimated_date"`
	} `json:"original_delivery_time,omitempty"`
	HomeDelivery                bool `json:"home_delivery"`
	CarrierOriginalDeliveryTime struct {
		Days          int    `json:"days"`
		EstimatedDate string `json:"estimated_date"`
	} `json:"carrier_original_delivery_time,omitempty"`
	Modal string `json:"modal"`
}

// ResponseQuote -
type ResponseQuote struct {
	Carrier []CarrierQuote `json:"carrier"`
}

// CarrierQuote -
type CarrierQuote struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline int     `json:"deadline"`
	Price    float64 `json:"price"`
}
