package modifying

import (
	"context"

	gopher "github.com/friendsofgo/gopherapi/pkg"
)

// Service provides modifying operations.
type Service interface {
	ModifyGopher(ctx context.Context, ID, Fixedacidity, Volatileacidity, Citricacid, Residualsugar, Chlorides, Freesulfurdioxide, Totalsulfurdioxide, Density, pH, Sulphates, Alcohol, Quality string) error
}

type service struct {
	repository gopher.Repository
}

// NewService creates a modifying service with the necessary dependencies
func NewService(repository gopher.Repository) Service {
	return &service{repository}
}

// ModifyGopher modify a gopher data
func (s *service) ModifyGopher(ctx context.Context,  ID, Fixedacidity, Volatileacidity, Citricacid, Residualsugar, Chlorides, Freesulfurdioxide, Totalsulfurdioxide, Density, pH, Sulphates, Alcohol, Quality string) error {
	g := gopher.New(ID, Fixedacidity, Volatileacidity, Citricacid,  Residualsugar, Chlorides, Freesulfurdioxide, Totalsulfurdioxide, Density, pH, Sulphates, Alcohol, Quality)
	return s.repository.UpdateGopher(ctx, ID, *g)
}
