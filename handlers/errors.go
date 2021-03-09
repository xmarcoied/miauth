package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/go-chi/render"
	"github.com/xmarcoied/miauth/pkg"
)

func errDetailsMsg(r *http.Request, httpStatusCode int, details string) string {
	q := r.URL.String()
	if qun, e := url.QueryUnescape(q); e == nil {
		q = qun
	}

	srcFileInfo := ""
	if pc, file, line, ok := runtime.Caller(2); ok {
		fnameElems := strings.Split(file, "/")
		funcNameElems := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		srcFileInfo = fmt.Sprintf(
			" [caused by %s:%d %s]", strings.Join(fnameElems[len(fnameElems)-3:], "/"),
			line, funcNameElems[len(funcNameElems)-1])
	}

	remoteIP := r.RemoteAddr
	if pos := strings.Index(remoteIP, ":"); pos >= 0 {
		remoteIP = remoteIP[:pos]
	}
	return fmt.Sprintf("%d - %s - %s - %s%s",
		httpStatusCode, details, remoteIP, q, srcFileInfo)
}

// RenderJSONError sends a json error
func RenderJSONError(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error) {
	// Transforming the error to the pkg.Error entity
	var appError *pkg.Error
	if _, ok := err.(*pkg.Error); !ok {
		appError = &pkg.Error{
			Code: pkg.ErrInternal,
			Msg:  err.Error(),
		}
	} else {
		appError = err.(*pkg.Error)
	}

	if appError.Msg == "" {
		appError.Msg = http.StatusText(httpStatusCode)
	}

	// structing logger
	logger := pkg.GetLogRequest(r)
	logger.Error(errDetailsMsg(r, httpStatusCode, appError.Msg))
	render.Status(r, httpStatusCode)
	render.JSON(w, r, pkg.JSON{"error": err})
}
