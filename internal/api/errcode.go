package api

import "net/http"

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

type errCode struct {
	status int
	extra  int32
}

func newErrCode(statusCode int, extraCode int32) errCode {
	if extraCode == 0 {
		extraCode = int32(statusCode)
	}
	return errCode{status: statusCode, extra: extraCode}
}

var (
	codeInternal  = newErrCode(http.StatusInternalServerError, 0)
	codeForbidden = newErrCode(http.StatusForbidden, 0)
	codeNotFound  = newErrCode(http.StatusNotFound, 0)
)
