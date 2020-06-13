package cockroach

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	gopher "github.com/friendsofgo/gopherapi/pkg"
	"github.com/openzipkin/zipkin-go"
)

type gopherRepository struct {
	db     *sql.DB
	tracer *zipkin.Tracer
}

// NewRepository creates a crockoach repository with the necessary dependencies
func NewRepository(db *sql.DB, tracer *zipkin.Tracer) gopher.Repository {
	return gopherRepository{db: db, tracer: tracer}
}

func (r gopherRepository) CreateGopher(_ context.Context, g *gopher.Gopher) error {
	sqlStm := `INSERT INTO gophers (id, Fixedacidity, Volatileacidity, Citricacid, Residualsugar, Chlorides, Freesulfurdioxide, Totalsulfurdioxide, Density, PH, Sulphates, Alcohol, Quality, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW())`
	_, err := r.db.Exec(sqlStm, g.ID, g.Fixedacidity, g.Volatileacidity, g.Citricacid, g.Residualsugar, g.Chlorides, g.Freesulfurdioxide, g.Totalsulfurdioxide, g.Density, g.PH, g.Sulphates, g.Alcohol, g.Quality)
	if err != nil {
		return err
	}
	return nil
}

func (r gopherRepository) FetchGophers(ctx context.Context) ([]gopher.Gopher, error) {
	sqlStm := `SELECT id, Fixedacidity, Volatileacidity, Quality , Residualsugar, Chlorides, Freesulfurdioxide, Totalsulfurdioxide, Density, PH, Sulphates, Alcohol, Quality, created_at, updated_at FROM gophers`
	rows, err := r.db.Query(sqlStm)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var gophers []gopher.Gopher

	for rows.Next() {
		var g gopher.Gopher
		if err := rows.Scan(&g.ID, &g.Fixedacidity, &g.Volatileacidity, &g.Citricacid, &g.Residualsugar, &g.Chlorides, &g.Freesulfurdioxide, &g.Totalsulfurdioxide, &g.Density, &g.PH, &g.Sulphates, &g.Alcohol, &g.Quality , &g.CreatedAt, &g.UpdatedAt); err != nil {
			log.Println(err)
			continue
		}
		gophers = append(gophers, g)
	}
	return gophers, nil
}

func (r gopherRepository) DeleteGopher(ctx context.Context, ID string) error {
	return errors.New("method not implemented")
}

func (r gopherRepository) UpdateGopher(ctx context.Context, ID string, g gopher.Gopher) error {
	return errors.New("method not implemented")
}

func (r gopherRepository) FetchGopherByID(ctx context.Context, ID string) (*gopher.Gopher, error) {
	return nil, errors.New("method not implemented")
}
