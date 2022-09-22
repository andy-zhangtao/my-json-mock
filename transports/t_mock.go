package transports

import (
	"context"
	"encoding/json"
	"github.com/andy-zhangtao/my-json-mock/models"
	"github.com/andy-zhangtao/my-json-mock/utils"
	"net/http"
)

// FindAllMockDecodeRequest
func FindAllMockDecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	//values := r.URL.Query()
	//var mr models.HttpMockRequest
	//err := utils.Bind(r, &mr)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}

// FindAllMockEncodeResponse
func FindAllMockEncodeResponse(c context.Context, w http.ResponseWriter, res interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

// FindSpecifyMockWithIdDecodeRequest
func FindSpecifyMockWithIdDecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	var mr models.HttpMockRequest
	err := utils.Bind(r, &mr)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

// FindSpecifyMockWithIdEncodeResponse
func FindSpecifyMockWithIdEncodeResponse(c context.Context, w http.ResponseWriter, res interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

// AddNewMockDecodeRequest
func AddNewMockDecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	var mr models.HttpMockRequest
	err := utils.Bind(r, &mr)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

// AddNewMockEncodeResponse
func AddNewMockEncodeResponse(c context.Context, w http.ResponseWriter, res interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

// UpdateMockDecodeRequest
func UpdateMockDecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	var mr models.HttpMockRequest
	err := utils.Bind(r, &mr)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

// UpdateMockEncodeResponse
func UpdateMockEncodeResponse(c context.Context, w http.ResponseWriter, res interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}

// DeleteMockDecodeRequest
func DeleteMockDecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	var mr models.HttpMockRequest
	err := utils.Bind(r, &mr)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

// DeleteMockEncodeResponse
func DeleteMockEncodeResponse(c context.Context, w http.ResponseWriter, res interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(res)
}
