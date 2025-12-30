package ctfdgen

import (
	"log"

	"github.com/Kinirok/ctfd-acc-gen/internal/ctfd"
	"github.com/Kinirok/ctfd-acc-gen/internal/logging"
	gormodel "github.com/Kinirok/ctfd-acc-gen/internal/storage"
	"gorm.io/gorm"
)

type Config struct {
	CTFDClient ctfd.CTFdClient
	DB         *gorm.DB
}

type Generator struct {
	ctfdClient ctfd.CTFdClient
	db         *gorm.DB
	logger     *log.Logger
}

func NewGenerator(cfg Config) (*Generator, error) {
	logger := logging.Init()
	logger.Println("Checking ctfd connection")
	if cfg.CTFDClient == nil {
		return nil, ErrNoClient
	}
	logger.Println("Checking DB connection")
	if cfg.DB == nil {
		return nil, ErrNoDB
	}
	sqlDB, err := cfg.DB.DB()
	if err != nil {
		return nil, ErrNoSqlDB
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, ErrLostConnectionDB
	}
	logger.Println("Creating tables, if it doesn't exist")
	cfg.DB.AutoMigrate(&gormodel.Account{})
	cfg.DB.AutoMigrate(&gormodel.Team{})
	return &Generator{
		ctfdClient: cfg.CTFDClient,
		db:         cfg.DB,
		logger:     logger,
	}, nil
}
