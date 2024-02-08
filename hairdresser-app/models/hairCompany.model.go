package models

type HairCompany struct {
	Email   string             `json:"email,omitempty" validate:"required"`
	Name    string             `json:"name,omitempty" validate:"required"`
	SIREN   string             `json:"siren,omitempty" validate:"required"`
	Address Address            `json:"address,omitempty" validate:"required"`
	Status  string             `json:"status,omitempty" Validate:"required"` //status can be oppened or closed 
}

type Address struct {
	Line1      string `json:"line1,omitempty" validate:"required"`
	Line2      string `json:"line2,omitempty"`
	PostalCode string `json:"postalCode,omitempty" validate:"required"`
	City       string `json:"city,omitempty" validate:"required"`
	Country    string `json:"country,omitempty" validate:"required"`
}
