package driver

import (
	"context"
	"database/sql"
	"github.com/MikhailLipanin/how2amuse/pkg/models"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL creates database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// testDB tries to ping the database
func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// GetRegions gets all regions
func (m *DB) GetRegions() ([]models.Region, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select * from region
		`

	rows, err := m.SQL.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var ret []models.Region

	for rows.Next() {
		var region models.Region
		err := rows.Scan(
			&region.Country,
			&region.ID,
			&region.CityCount,
			&region.Area,
			&region.Name,
			&region.ImgHref,
		)
		if err != nil {
			return ret, err
		}
		ret = append(ret, region)
	}
	if err = rows.Err(); err != nil {
		return ret, err
	}

	return ret, nil
}

// GetCities gets all cities
func (m *DB) GetCities() ([]models.City, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select * from city
		`

	rows, err := m.SQL.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var ret []models.City

	for rows.Next() {
		var city models.City
		err := rows.Scan(
			&city.ID,
			&city.RegionID,
			&city.Name,
			&city.Population,
			&city.TimeZone,
			&city.ImgHref,
		)
		if err != nil {
			return ret, err
		}
		ret = append(ret, city)
	}
	if err = rows.Err(); err != nil {
		return ret, err
	}

	return ret, nil
}

// GetPlaces gets all cities
func (m *DB) GetPlaces() ([]models.Place, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select * from place
		`

	rows, err := m.SQL.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var ret []models.Place

	for rows.Next() {
		var place models.Place
		err := rows.Scan(
			&place.ID,
			&place.CityID,
			&place.Name,
			&place.Description,
			&place.ImgHref,
		)
		if err != nil {
			return ret, err
		}
		ret = append(ret, place)
	}
	if err = rows.Err(); err != nil {
		return ret, err
	}

	return ret, nil
}
