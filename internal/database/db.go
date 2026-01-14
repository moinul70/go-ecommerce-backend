package database

import (
	"ecommerce/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Database struct {
    *sqlx.DB
    Driver string
}

func BuildDSN(cfg config.DbConfig) string {
    switch cfg.Driver {

    case "mysql":
        return fmt.Sprintf(
            "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
            cfg.User,
            cfg.Password,
            cfg.Host,
            cfg.Port,
            cfg.DBName,
        )

     case "postgres", 
     "postgresql":
        sslmode := cfg.Sslmode
        if sslmode == "" {
            sslmode = "disable"
        }
        return fmt.Sprintf(
            "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
            cfg.Host,
            cfg.Port,
            cfg.User,
            cfg.Password,
            cfg.DBName,
            cfg.Sslmode,
        )

    default:
        panic("unsupported driver")
    }
}

func DbConnect(cfg config.DbConfig) (*Database,error){
    dsn := BuildDSN(cfg)

    db, err := sqlx.Open(cfg.Driver, dsn)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(cfg.MaxOpen)
    db.SetMaxIdleConns(cfg.MaxIdle)
    db.SetConnMaxLifetime(time.Duration(cfg.LifeTime) * time.Second)

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &Database{
        DB:     db,
        Driver: cfg.Driver,
    }, nil
}