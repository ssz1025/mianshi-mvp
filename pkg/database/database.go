package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/d60-Lab/gin-template/internal/model"
	"github.com/d60-Lab/gin-template/pkg/config"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	var logLevel logger.LogLevel
	if cfg.Server.Mode == "release" {
		logLevel = logger.Error
	} else {
		logLevel = logger.Info
	}

	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logLevel),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	var db *gorm.DB
	var err error

	switch cfg.Database.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Database.Database), gormConfig)
		if err != nil {
			return nil, err
		}
	default:
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Shanghai",
			cfg.Database.Username,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Database,
		)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			return nil, err
		}
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)
		if err := db.Exec("SET search_path TO public").Error; err != nil {
			return nil, fmt.Errorf("failed to set search_path: %w", err)
		}
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.QuestionBank{},
		&model.QuestionBankQuestion{},
		&model.QuestionTag{},
		&model.QuestionTagRelation{},
		&model.Question{},
		&model.UserQuestionRecord{},
		&model.UserFavorite{},
		&model.PracticeRoute{},
		&model.RoutePhase{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
