package gopher

import (
	"context"
	"time"
)

// Gopher defines the properties of a gopher to be listed
type Gopher struct {
	ID        string     `json:"ID"`
	Fixedacidity      string     `json:"Fixedacidity,omitempty"`
	Volatileacidity     string     `json:"Volatileacidity,omitempty"`
	Citricacid       string      `json:"Citricacid,omitempty"`
	Residualsugar      string     `json:"Residualsugar,omitempty"`
	Chlorides     string     `json:"Chlorides,omitempty"`
	Freesulfurdioxide       string      `json:"Freesulfurdioxide,omitempty"`
	Totalsulfurdioxide      string     `json:"Totalsulfurdioxide,omitempty"`
	Density     string     `json:"Density,omitempty"`
	PH       string      `json:"PH,omitempty"`
	Sulphates      string     `json:"Sulphates,omitempty"`
	Alcohol     string     `json:"Alcohol,omitempty"`
	Quality        string      `json:"Quality,omitempty"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

// New creates a gopher
func New(ID, Fixedacidity, Volatileacidity, Citricacid, Residualsugar, Chlorides, Freesulfurdioxide, Totalsulfurdioxide, Density, PH, Sulphates, Alcohol, Quality string) *Gopher {
	return &Gopher{
		ID:    ID,
		Fixedacidity:    Fixedacidity,
		Volatileacidity:  Volatileacidity,
		Citricacid: Citricacid,
		Residualsugar:    Residualsugar,
		Chlorides:  Chlorides,
		Freesulfurdioxide: Freesulfurdioxide,
		Totalsulfurdioxide:    Totalsulfurdioxide,
		Density:  Density,
		PH: PH,
		Sulphates:    Sulphates,
		Alcohol:  Alcohol,
		Quality : Quality,
	}
}

//Repository provides access to the gopher storage
type Repository interface {
	// CreateGopher saves a given gopher
	CreateGopher(ctx context.Context, gopher *Gopher) error
	// FetchGophers return all gophers saved in storage
	FetchGophers(ctx context.Context) ([]Gopher, error)
	// DeleteGopher remove gopher with given ID
	DeleteGopher(ctx context.Context, ID string) error
	// UpdateGopher modify gopher with given ID and given new data
	UpdateGopher(ctx context.Context, ID string, gopher Gopher) error
	// FetchGopherByID returns the gopher with given ID
	FetchGopherByID(ctx context.Context, ID string) (*Gopher, error)
}
