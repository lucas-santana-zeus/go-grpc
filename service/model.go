package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"go-grpc/commons/models"
	block "go-grpc/commons/pb"
)

const (
	PRECIPITATION_DATA_TYPE_ID = "6"
	TEMPERATURE_DATA_TYPE_ID   = "1"
)

func (server *Server) GetBlockById(context context.Context, req *block.RequestID) (*block.ResponseBlock, error) {
	bqQuery := server.BQClient.Query(
		"select (a_temp.data_timestamp) as data_timestamp, (a_temp.created_timestamp) as created_timestamp," +
			"a_temp.data_inst as temp_inst, a_temp.data_min as temp_min, a_temp.data_max as temp_max," +
			"a_prec.data_inst as prec_inst, a_prec.data_min as prec_min, a_prec.data_max as prec_max" +
			"from `athena-dsv.athena.pixel` a_temp " +
			"inner join (" +
			"select data_timestamp, created_timestamp, data_inst, data_min, data_max, source_id" +
			"from `athena-dsv.athena.pixel`" +
			`where date(created_timestamp) = "2022-06-28"` +
			"and datatype_id = " + PRECIPITATION_DATA_TYPE_ID +
			"and source_type_id = 3" +
			"and source_subtype_id = 16" +
			"and source_id = @source_id" +
			") as a_prec" +
			"on a_temp.source_id = a_prec.source_id" +
			`where date(a_temp.created_timestamp) = "2022-06-28"` +
			"and a_temp.datatype_id = " + TEMPERATURE_DATA_TYPE_ID +
			"and a_temp.source_type_id = 3" +
			"and a_temp.source_subtype_id = 16" +
			"and a_temp.source_id = @source_id" +
			"order by a_temp.created_timestamp desc, a_prec.created_timestamp desc" +
			"limit 1")

	bqQuery.Parameters = []bigquery.QueryParameter{
		{
			Name:  "source_id",
			Value: req.GetId(),
		},
	}

	//todo: call bq connection function passing bqQuery as parameter and finish construction of resBlock

	resBlock := models.TransformBlockDAOIntoResponse(blockDAO)
	return resBlock, nil
}
