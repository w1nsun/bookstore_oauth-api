package db

import (
	"fmt"
	"github.com/w1nsun/bookstore_oauth-api/src/clients/cassandra"
	"github.com/w1nsun/bookstore_oauth-api/src/domain/access_token"
	"github.com/w1nsun/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	_, err := cassandra.GetSession()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
