package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	ca_parksconfig "ca_parks/config"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/config"

	"ca_parks/core/module/park"

	"ca_parks/core/module/recreation_area"

	"ca_parks/core/module/user"

	"ca_parks/core/module/feature"

	"ca_parks/core/module/park_has_feature"

	"ca_parks/core/module/event"

	"ca_parks/core/repository"
	ca_parksdb "ca_parks/core/repository/gen"
)

type Implementation struct {
	repository ca_parksdb.Queries

	park park.Module

	recreation_area recreation_area.Module

	user user.Module

	feature feature.Module

	park_has_feature park_has_feature.Module

	event event.Module
}

func New(ctx context.Context, provider config.Provider) (*Implementation, error) {

	var dbs ca_parksconfig.DBs
	if err := provider.Get("db").Populate(&dbs); err != nil {
		return nil, err
	}

	if len(dbs) == 0 {
		return nil, errors.New("db configuration not found")
	}

	dbconfig := dbs[0]
	db, err := sql.Open(dbconfig.Driver, dbconfig.Path())
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}
	defer db.Close()
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(0)
	repository := repository.New(*db)

	return &Implementation{
		repository: *repository,
	}, nil
}

func (i Implementation) Park() park.Module {
	if i.park == nil {
		i.park = park.New(&i.repository)
	}
	return i.park
}

func (i Implementation) RecreationArea() recreation_area.Module {
	if i.recreation_area == nil {
		i.recreation_area = recreation_area.New(&i.repository)
	}
	return i.recreation_area
}

func (i Implementation) User() user.Module {
	if i.user == nil {
		i.user = user.New(&i.repository)
	}
	return i.user
}

func (i Implementation) Feature() feature.Module {
	if i.feature == nil {
		i.feature = feature.New(&i.repository)
	}
	return i.feature
}

func (i Implementation) ParkHasFeature() park_has_feature.Module {
	if i.park_has_feature == nil {
		i.park_has_feature = park_has_feature.New(&i.repository)
	}
	return i.park_has_feature
}

func (i Implementation) Event() event.Module {
	if i.event == nil {
		i.event = event.New(&i.repository)
	}
	return i.event
}
