package mysql

import (
	"context"
	"database/sql"
	"errors"
	gopherapi "github.com/friendsofgo/gopherapi/pkg"
	"github.com/huandu/go-sqlbuilder"
	_ "github.com/lib/pq"
	"time"
)

type gopherRepository struct {
	table string
	db    *sql.DB
}

// NewRepository instances a MySQL implementation of the gopherapi.Repository
func NewRepository(table string, db *sql.DB) gopherapi.Repository {
	return gopherRepository{table: table, db: db}
}

// CreateGopher satisfies the gopherapi.Repository interface
func (r gopherRepository) CreateGopher(ctx context.Context, g *gopherapi.Gopher) error {
	insertBuilder := sqlbuilder.NewStruct(new(sqlGopher)).InsertInto(
		r.table,
		sqlGopher{
			ID:        g.ID,
			Fixedacidity:      g.Fixedacidity,
			Volatileacidity:     g.Volatileacidity,
			Citricacid:       g.Citricacid,
			Residualsugar:      g.Residualsugar,
			Chlorides:     g.Chlorides,
			Freesulfurdioxide:       g.Freesulfurdioxide,
			Totalsulfurdioxide:      g.Totalsulfurdioxide,
			Density:     g.Density,
			PH:       g.PH,
			Sulphates:      g.Sulphates,
			Alcohol:     g.Alcohol,
			Quality :       g.Quality ,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		},
	)

	query, args := insertBuilder.Build()
	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r gopherRepository) FetchGophers(ctx context.Context) ([]gopherapi.Gopher, error) {
	sqlGopherStruct := sqlbuilder.NewStruct(new(sqlGopher))

	selectBuilder := sqlGopherStruct.SelectFrom(r.table)
	query, args := selectBuilder.Build()

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var gophers []gopherapi.Gopher
	for rows.Next() {
		sqlGopher := sqlGopher{}

		err := rows.Scan(sqlGopherStruct.Addr(&sqlGopher)...)
		if err != nil {
			return nil, err
		}

		gophers = append(gophers, gopherapi.Gopher{
				ID:        sqlGopher.ID,
			Fixedacidity:      sqlGopher.Fixedacidity,
			Volatileacidity:     sqlGopher.Volatileacidity,
			Citricacid:       sqlGopher.Citricacid,
			Residualsugar:      sqlGopher.Residualsugar,
			Chlorides:     sqlGopher.Chlorides,
			Freesulfurdioxide:       sqlGopher.Freesulfurdioxide,
			Totalsulfurdioxide:      sqlGopher.Totalsulfurdioxide,
			Density:     sqlGopher.Density,
			PH:       sqlGopher.PH,
			Sulphates:      sqlGopher.Sulphates,
			Alcohol:     sqlGopher.Alcohol,
			Quality :       sqlGopher.Quality ,
			CreatedAt: sqlGopher.CreatedAt,
			UpdatedAt: sqlGopher.UpdatedAt,
		})
	}

	return gophers, nil
}

func (r gopherRepository) DeleteGopher(ctx context.Context, ID string) error {
	deleteBuilder := sqlbuilder.NewStruct(new(sqlGopher)).DeleteFrom(r.table)
	query, args := deleteBuilder.Where(
		deleteBuilder.Equal("id", ID),
	).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}

func (r gopherRepository) UpdateGopher(ctx context.Context, ID string, g gopherapi.Gopher) error {
	updateBuilder := sqlbuilder.NewStruct(new(sqlGopher)).Update(
		r.table,
		sqlGopher{
			ID:        g.ID,
			Fixedacidity:      g.Fixedacidity,
			Volatileacidity:     g.Volatileacidity,
			Citricacid:       g.Citricacid,
			Residualsugar:      g.Residualsugar,
			Chlorides:     g.Chlorides,
			Freesulfurdioxide:       g.Freesulfurdioxide,
			Totalsulfurdioxide:      g.Totalsulfurdioxide,
			Density:     g.Density,
			PH:       g.PH,
			Sulphates:      g.Sulphates,
			Alcohol:     g.Alcohol,
			Quality :       g.Quality ,
			CreatedAt: g.CreatedAt,
			UpdatedAt: g.UpdatedAt,
		},
	)

	query, args := updateBuilder.Where(
		updateBuilder.Equal("id", ID),
	).Build()

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r gopherRepository) FetchGopherByID(ctx context.Context, ID string) (*gopherapi.Gopher, error) {
	sqlGopherStruct := sqlbuilder.NewStruct(new(sqlGopher))

	selectBuilder := sqlGopherStruct.SelectFrom(r.table)

	query, args := selectBuilder.Where(
		selectBuilder.Equal("id", ID),
	).Build()

	row := r.db.QueryRowContext(ctx, query, args...)

	sqlGopher := sqlGopher{}

	err := row.Scan(sqlGopherStruct.Addr(&sqlGopher)...)
	if err != nil {
		return nil, err
	}

	return &gopherapi.Gopher{
		ID:        sqlGopher.ID,
			Fixedacidity:      sqlGopher.Fixedacidity,
			Volatileacidity:     sqlGopher.Volatileacidity,
			Citricacid:       sqlGopher.Citricacid,
			Residualsugar:      sqlGopher.Residualsugar,
			Chlorides:     sqlGopher.Chlorides,
			Freesulfurdioxide:       sqlGopher.Freesulfurdioxide,
			Totalsulfurdioxide:      sqlGopher.Totalsulfurdioxide,
			Density:     sqlGopher.Density,
			PH:       sqlGopher.PH,
			Sulphates:      sqlGopher.Sulphates,
			Alcohol:     sqlGopher.Alcohol,
			Quality :       sqlGopher.Quality ,
			CreatedAt: sqlGopher.CreatedAt,
			UpdatedAt: sqlGopher.UpdatedAt,
	}, nil
}

type sqlGopher struct {
	ID        string     `db:"id"`
	Fixedacidity      string     `db:"Fixedacidity"`
	Volatileacidity     string     `db:"Volatileacidity"`
	Citricacid       string        `db:"Citricacid"`
	Residualsugar      string     `db:"Residualsugar"`
	Chlorides     string     `db:"Chlorides"`
	Freesulfurdioxide       string        `db:"Freesulfurdioxide"`
	Totalsulfurdioxide      string     `db:"Totalsulfurdioxide"`
	Density     string     `db:"Density"`
	PH       string        `db:"PH"`
	Sulphates      string     `db:"Sulphates"`
	Alcohol     string     `db:"Alcohol"`
	Quality       string        `db:"Quality"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
