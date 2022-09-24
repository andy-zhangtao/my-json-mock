package endpoints

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/andy-zhangtao/my-json-mock/models"
	"github.com/andy-zhangtao/my-json-mock/services"
	"github.com/andy-zhangtao/my-json-mock/types"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

func MakeRunMockEndpoint(ts services.IMockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		mid, ok := request.(string)
		if !ok {
			return models.CommonMockResponse{
				Code: types.FAILE,
			}, errors.New("Wrong Mock ID")
		}

		mock, err := ts.FindSpecifyMockWithMId(mid)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, types.NotFoundMockError{}
			}
			return nil, err
		}

		// get content-type
		var params models.HttpMockParamsRequest
		err = json.Unmarshal([]byte(mock.Params), &params)
		if err != nil {
			return nil, err
		}

		return models.CommonMockResponse{
			Code: types.OK,
			Result: models.MockResponse{
				ContentType: params.ContentType,
				Response:    mock.Mock,
			},
		}, nil
	}
}
