package endpoints

import (
	"context"
	"encoding/json"
	"github.com/andy-zhangtao/my-json-mock/models"
	"github.com/andy-zhangtao/my-json-mock/services"
	"github.com/andy-zhangtao/my-json-mock/types"
	"github.com/andy-zhangtao/my-json-mock/types/db"
	"github.com/andy-zhangtao/my-json-mock/utils"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func MakeFindAllMockEndpoint(ts services.IMockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		all, err := ts.FindAllMock()
		if err != nil {
			return nil, err
		}

		return models.CommonResponse{
			Code:   types.OK,
			Result: all,
		}, nil
	}
}

// MakeFindSpecifyMockWithIdEndpoint
func MakeFindSpecifyMockWithIdEndpoint(ts services.IMockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		id, ok := request.(int)
		if !ok {
			return nil, errors.New(types.ConvertRequestError)
		}

		t, err := ts.FindSpecifyMockWithId(id)
		if err != nil {
			return nil, err
		}

		return models.CommonResponse{
			Code:   types.OK,
			Result: t,
		}, nil
	}
}

// MakeAddNewMockEndpoint
func MakeAddNewMockEndpoint(ts services.IMockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(db.MockRequest)
		if !ok {
			return nil, errors.New(types.ConvertRequestError)
		}

		err = addNewMockPreCheck(&r)
		if err != nil {
			return nil, err
		}

		data, _ := json.Marshal(r)
		r.Mid = utils.GenerateID(string(data))

		err = ts.AddNewMock(r)
		if err != nil {
			return nil, err
		}

		return models.CommonResponse{
			Code:   types.OK,
			Result: types.HttpCommonSucces,
		}, nil
	}
}

// MakeUpdateMockEndpoint
func MakeUpdateMockEndpoint(ts services.IMockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(db.MockRequest)
		if !ok {
			return nil, errors.New(types.ConvertRequestError)
		}

		err = ts.UpdateMock(r)
		if err != nil {
			return nil, err
		}

		return models.CommonResponse{
			Code:   types.OK,
			Result: types.HttpCommonSucces,
		}, nil
	}
}

// MakeDeleteMockEndpoint
func MakeDeleteMockEndpoint(ts services.IMockService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(db.MockRequest)
		if !ok {
			return nil, errors.New(types.ConvertRequestError)
		}

		err = ts.DeleteMock(r.Id)
		if err != nil {
			return nil, err
		}

		return models.CommonResponse{
			Code:   types.OK,
			Result: types.HttpCommonSucces,
		}, nil
	}
}

func addNewMockPreCheck(req *db.MockRequest) error {
	req.User = strings.TrimSpace(req.User)
	req.Method = strings.ToUpper(strings.TrimSpace(req.Method))
	req.Name = strings.TrimSpace(req.Name)
	req.Mock = strings.TrimSpace(req.Mock)
	req.Path = strings.TrimSpace(req.Path)
	req.Params = strings.TrimSpace(req.Params)
	req.Remark = strings.TrimSpace(req.Remark)

	if req.User == "" {
		return types.EmptyUserError{}
	}

	if req.Method == "" {
		req.Method = http.MethodGet
	}

	if req.Name == "" {
		return types.EmptyNameError{}
	}

	if req.Mock == "" {
		return types.EmptyMockContentError{}
	}

	if req.Path == "" {
		return types.EmptyPathError{}
	}
	return nil
}
