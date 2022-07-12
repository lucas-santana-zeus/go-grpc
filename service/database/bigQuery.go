package database

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

var (
	ctx context.Context
)

func QueryConnection(bqQuery *bigquery.Query) (*bigquery.RowIterator, error) {

	bqQuery.Location = "US"

	job, err := bqQuery.Run(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	status, err := job.Wait(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if err := status.Err(); err != nil {
		return nil, err
	}

	it, err := job.Read(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return it, nil

}
