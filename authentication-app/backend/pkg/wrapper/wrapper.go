package wrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)
type Processor func(ctx context.Context, req, resp interface{}) error

func WrapProcessor(
	proc Processor,
	req, resp interface{},
) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		newReq := reflect.New(reflect.TypeOf(req).Elem()).Interface()
		newResp := reflect.New(reflect.TypeOf(resp).Elem()).Interface()

		err := json.NewDecoder(r.Body).Decode(&newReq)
		if err != nil {
			writeToClient(w, newResp, err)
			return
		}

		ctx := ContextWithTraceID(r.Context(), time.Now().String())

		err = proc(ctx, newReq, newResp)
		writeToClient(w, newResp, err)
	}
}

func writeToClient(w http.ResponseWriter, resp interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		code := cerr.Code(err)
		w.WriteHeader(code)

		setDebugMessage(resp, err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
	}

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func setDebugMessage(resp interface{}, msg string) {
	if resp == nil {
		return
	}

	debugMsgField := "DebugMsg"
	structField, found := reflect.TypeOf(resp).Elem().FieldByName(debugMsgField)
	if !found {
		fmt.Println("not found")
		return
	}

	fieldType := structField.Type
	if fieldType.Kind() != reflect.Ptr || fieldType.Elem().Kind() != reflect.String {
		return
	}

	requiredField := reflect.ValueOf(resp).Elem().FieldByName(debugMsgField)

	if requiredField.CanSet() {
		var finalMsg string

		elem := requiredField.Elem()
		if elem.IsValid() && len(elem.String()) != 0 {
			finalMsg = elem.String() + ": " + msg
		} else {
			finalMsg = msg
		}

		requiredField.Set(reflect.ValueOf(&finalMsg))
	}
}

type contextKey string

const (
	traceIDKey = contextKey("trace_id")
)

func ContextWithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}
