package db

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"testDelivery/authorizationService/internal/database/user"
	"testDelivery/authorizationService/pkg/config"
)

var Module = fx.Provide(Setup)

type Params struct {
	fx.In
	*logrus.Logger
	*config.Tuner
}

type DbInter interface {
	Close() error
	GetDB() *gorm.DB
}

type gormDB struct {
	db  *gorm.DB
	log *logrus.Logger
}

func Setup(param Params) DbInter {
	var err error
	dbrUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		param.Tuner.DB.Host,
		param.Tuner.DB.Port,
		param.Tuner.DB.User,
		param.Tuner.DB.Password,
		param.Tuner.DB.Database, param.Tuner.DB.SSlMode,
	)
	log.Println(dbrUri)
	db, err := gorm.Open(postgres.Open(dbrUri), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		param.Logger.Fatal("db.Setup err:", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	// AutoMigrate
	if err := autoMigrate(db); err != nil {

		param.Logger.Fatal("create model migrate err: ", err)

	}
	param.Logger.Println("DB successfully connected! ")

	return &gormDB{
		db:  db,
		log: param.Logger,
	}

}

func (d gormDB) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (d gormDB) GetDB() *gorm.DB {
	return d.db
}

func autoMigrate(db *gorm.DB) error {

	if err := addUserEnum(db); err != nil {
		return err
	}

	for _, model := range []interface{}{
		(*user.UserEntity)(nil),
	} {
		dbSilent := db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
		if err := dbSilent.AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}

func addUserEnum(db *gorm.DB) error {
	roles := []interface{}{
		user.Admin, user.Courier, user.Customer,
	}
	return db.Exec(fmt.Sprintf(`
		DO
		$$
		BEGIN
			IF NOT EXISTS (SELECT * FROM pg_type typ
				INNER JOIN pg_namespace nsp ON nsp.oid = typ.typnamespace
				WHERE nsp.nspname = current_schema() AND typ.typname = 'enum_user_role') THEN
				CREATE TYPE enum_user_role AS ENUM('%s', '%s', '%s');
			END IF;
		END;
		$$
		LANGUAGE plpgsql;
	`, roles...)).Error
}
