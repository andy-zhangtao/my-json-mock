package transports

import (
	"context"
	"fmt"
	"github.com/andy-zhangtao/my-json-mock/models"
	"github.com/andy-zhangtao/my-json-mock/types"
	"net/http"
)

func RunMockDecodeRequest(c context.Context, r *http.Request) (interface{}, error) {
	// 不需要做任何处理
	mid := r.Header.Get(types.MockID)
	if mid == "" {
		return nil, fmt.Errorf("%s is required", types.MockID)
	}

	return mid, nil
}

func RunMockEncodeResponse(c context.Context, w http.ResponseWriter, res interface{}) error {
	//w.WriteHeader(http.StatusOK)
	//w.Header().Set(types.MockID, "")

	r, ok := res.(models.CommonMockResponse)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	switch r.Code {
	case types.OK:
		switch r.Result.ContentType {
		case types.Json:
			w.WriteHeader(http.StatusOK)
			w.Header().Set(types.ContentType, types.Json)
			w.Write([]byte(r.Result.Response))
			return nil
		}
	case types.FAILE:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(r.Error.Error()))
		return nil

	}
	return nil
}
