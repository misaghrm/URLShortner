package repository

import (
	"database/sql"
	"github.com/misaghrm/urlshortener/domain/entity"
	"github.com/misaghrm/urlshortener/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"time"
)

var (
	connectionString string
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
}

func InsertNewURL(Model *url.URL) int64 {
	domain := &entity.BaseDomain{
		BaseEntity: entity.BaseEntity{
			Id:              util.ID.Generate().Int64(),
			CreationTimeUtc: time.Now().UTC(),
		},
		Scheme: Model.Scheme,
		Opaque: Model.Opaque,
		Host:   Model.Host,
	}
	path := &entity.Path{
		BaseEntity: entity.BaseEntity{
			Id:              util.ID.Generate().Int64(),
			CreationTimeUtc: time.Now().UTC(),
		},
		Path:         sql.NullString{Model.Path, true},
		Query:        sql.NullString{Model.RawQuery, true},
		Fragment:     sql.NullString{Model.RawFragment, true},
		BaseDomainId: (domain.Id),
	}
	Db.Begin()
	Db.Model(&entity.BaseDomain{}).Where(&entity.BaseDomain{
		Scheme: Model.Scheme,
		Opaque: Model.Opaque,
		Host:   Model.Host,
	}).FirstOrCreate(domain)
	Db.Model(&entity.Path{}).Where(&entity.Path{
		Path:         sql.NullString{Model.Path, true},
		Query:        sql.NullString{Model.RawQuery, true},
		Fragment:     sql.NullString{Model.RawFragment, true},
		BaseDomainId: domain.Id,
	}).FirstOrCreate(path)
	Db.Commit()
	return path.Id
}

func Find(id int64) (url string) {
	path := entity.Path{}
	Db.Model(&entity.Path{}).Where(`"Id" = ?`, id).First(&path)
	Db.Model(&entity.BaseDomain{}).Where(`"Id" = ?`, path.BaseDomainId).First(&path.BaseDomain)
	url = path.BaseDomain.Scheme + path.BaseDomain.Opaque + path.BaseDomain.Host + path.Path.String + path.Query.String + path.Fragment.String

	return url
}
