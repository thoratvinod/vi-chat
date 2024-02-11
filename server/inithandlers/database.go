package inithandlers

import (
	"fmt"
	"log"

	"github.com/thoratvinod/vi-chat/server/pkg/dm"
	"github.com/thoratvinod/vi-chat/server/pkg/group"
	"github.com/thoratvinod/vi-chat/server/pkg/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	host := "localhost"
	port := 5432
	dbName := "vi-chat"
	dbUser := "postgres"
	password := ""
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, dbUser, dbName, password)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("database connection failed, err=%+v", err.Error())
		return nil, err
	}
	err = migrate(DB)
	if err != nil {
		return nil, fmt.Errorf("database migration failed, err=%+v", err.Error())
	}
	return DB, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		user.User{},
		dm.DirectMessage{},
		group.Group{},
		group.GroupMessage{},
		group.GroupMessageStatus{},
	)
	if err != nil {
		return fmt.Errorf("database migration failed, err=%+v", err.Error())
	}
	return nil
}
