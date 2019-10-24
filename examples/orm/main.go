package main

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	bigquery "github.com/solcates/go-sql-bigquery"
	_ "github.com/solcates/go-sql-bigquery/dialects/bigquery"
	"os"
	"time"
)

func main() {
	var err error
	var db *gorm.DB
	logrus.SetLevel(logrus.DebugLevel)

	// Get the Connection String from the Environment Variable of BIGQUERY_CONNECTION_STRING
	uri := os.Getenv(bigquery.ConnectionStringEnvKey)
	if db, err = gorm.Open("bigquery", uri); err != nil {
		logrus.Fatal(err)
	}
	db.LogMode(true)
	db.DropTable("dataset1.animals")
	db.AutoMigrate(&Animal{})

	// Add an animal
	django := &Animal{
		Name: "Django",
		Size: 1,
		Born: time.Now(),
	}
	err = db.Save(django).Error
	if err != nil {
		logrus.Fatal(err)
	}

}

type Animal struct {
	Name string
	Size int64
	Born time.Time
}