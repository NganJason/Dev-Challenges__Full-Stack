package wrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)
type Processor func(ctx context.Context, req, resp interface{}) error

func WrapProcessor(
	proc Processor,
	req, resp interface{},
	cookie *http.Cookie,
) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		newReq := reflect.New(reflect.TypeOf(req).Elem()).Interface()
		newResp := reflect.New(reflect.TypeOf(resp).Elem()).Interface()

		err := json.NewDecoder(r.Body).Decode(&newReq)
		if err != nil {
			if err == io.EOF {
				newReq = nil
			} else {
				writeToClient(w, newResp, err)
				return
			}
		}

		ctx := clog.ContextWithTraceID(
			r.Context(),
			strconv.Itoa(int(time.Now().Unix())),
		)

		c, _ := r.Cookie(string(cookies.GetCookieKey()))
		if c != nil {
			ctx = cookies.AddCookieToCtx(ctx, c)
		}

		err = proc(ctx, newReq, newResp)

		if err == nil && cookie != nil {
			http.SetCookie(w, cookie)
		}

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
