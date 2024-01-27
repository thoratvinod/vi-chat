package inithandlers

import (
	"fmt"
	"log"
	"time"

	"github.com/thoratvinod/vi-chat/usermgmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DirectMessage struct {
	gorm.Model
	Sender    string
	Receiver  string
	Timestamp time.Time
	Seen      bool
}

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
	err := db.AutoMigrate(DirectMessage{})
	if err != nil {
		return fmt.Errorf("database migration failed, table=%+v", "DirectMessage")
	}
	err = db.AutoMigrate(usermgmt.User{})
	if err != nil {
		return fmt.Errorf("database migration failed, table=%+v", "User")
	}
	return nil
}
