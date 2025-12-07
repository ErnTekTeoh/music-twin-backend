package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
	"music-twin-backend/common"
	"music-twin-backend/module"
	"music-twin-backend/proto/pb"
	"net/http"
	"runtime/debug"
)

type HttpProcessor struct {
	Processor func(ctx context.Context, request, response interface{}) (errorCode int32)
	Request   interface{}
	Response  interface{}
}

func enableHttpCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// TODO check if need auth
func HttpProcessorWrapper(processor HttpProcessor) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				panicErr := fmt.Errorf("%s", r)
				common.LogWithError(context.Background(), "Panic error in HTTP Processor Wrapper")
				common.LogWithError(context.Background(), panicErr.Error())
				common.LogWithError(context.Background(), string(debug.Stack()))
				return
			}
		}()

		request := common.CloneEmpty(processor.Request)
		response := common.CloneEmpty(processor.Response)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Accept, Content-Type, Authorization, Member-Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			return
		}
		enableHttpCors(&w)
		_ = json.NewDecoder(r.Body).Decode(request)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}
		ctx := context.Background()
		token := r.Header.Get("Authorization")

		userId := module.GetUserIdFromToken(ctx, token)
		requestMeta := &pb.RequestMeta{
			UserId: proto.Int32(userId),
		}

		common.SetFieldValueToStruct(request, "RequestMeta", requestMeta)
		errCode := processor.Processor(ctx, request, response)
		if errCode != int32(pb.Constant_ERROR_CODE_SUCCESS) {
			// TODO for now nothing, for future monitoring purposes
		}
		json.NewEncoder(w).Encode(response)
	}

}
