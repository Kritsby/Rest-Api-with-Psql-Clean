package postgreSQL

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"os"
	"time"
)

type DsnConfig struct {
	Name     string
	Password string
	Host     string
	Port     string
	Database string
}

func InitPostgres() (*pgxpool.Pool, error) {
	pool, err := NewPsqlConnection(context.Background(), DsnConfig{
		Name:     viper.GetString("Name"),
		Password: viper.GetString("Password"),
		Host:     viper.GetString("Host"),
		Port:     viper.GetString("Port"),
		Database: viper.GetString("Database"),
	})
	if err != nil {
		return nil, fmt.Errorf("error conection: %w", err)
	}

	return pool, err
}

func NewPsqlConnection(ctx context.Context, dsnConfig DsnConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dsnConfig.Name,
		dsnConfig.Password,
		dsnConfig.Host,
		dsnConfig.Port,
		dsnConfig.Database,
	)
	pool, err := doWithTry(ctx, dsn, 5, 5)
	if err != nil {
		return nil, fmt.Errorf("can't connect to DB: %w", err)
	}

	return pool, nil
}

func doWithTry(ctx context.Context, dsn string, try int, second time.Duration) (conn *pgxpool.Pool, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for try > 0 {
		conn, err = pgxpool.Connect(ctx, dsn)
		if err == nil {
			return conn, err
		}
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		time.Sleep(second * time.Second)
		try--
	}
	return nil, err
}
