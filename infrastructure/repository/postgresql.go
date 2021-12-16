package repository

import (
	"github.com/misaghrm/urlshortener/domain/entity"
	"github.com/misaghrm/urlshortener/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	connectionString string = "host=127.0.0.1 user=postgres password= dbname=postgres port=5432 sslmode=disable"
	Db               *gorm.DB
	err              error
)

func init() {
	Db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Error,
	}), CreateBatchSize: 500})
	if err != nil {
		log.Panicln(err)
		return
	}
	err = Db.AutoMigrate(entity.BaseDomain{})
	if err != nil {
		log.Panicln(err)
		return
	}
}

func InsertNewURL(Model string) int64 {
	domain := entity.BaseDomain{
		Id:              util.ID.Generate().Int64(),
		CreationTimeUtc: time.Now().UTC(),
		Url:             Model,
	}
	Db.
		Model(&entity.BaseDomain{}).
		Where(`"Url" = ?`, Model).
		FirstOrCreate(&domain)
	return domain.Id
}

func Find(id int64) (url string) {
	Url := entity.BaseDomain{}
	err = Db.Model(&entity.BaseDomain{}).
		Where(`"Id" = ?`, id).
		First(&Url).
		Error
	if err != nil {
		log.Println("error of finding path:", err)
	}
	Url.Counter += 1
	err = Db.
		Save(&Url).
		Error
	if err != nil {
		log.Println("error of saving path:", err)
	}
	return Url.Url
}
