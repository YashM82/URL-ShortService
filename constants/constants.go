package constants

import "net/http"

// Response represents HTTP response
type Response struct {
	HTTP_STATUS int
	MESSAGE     string
}

var StatusCodes = struct {
	OK                    int  
	CREATED               int
	BAD_REQUEST           int
	UNAUTHORIZED          int
	FORBIDDEN             int
	NOT_FOUND             int
	INTERNAL_SERVER_ERROR int
}{
	OK:                    http.StatusOK,
	CREATED:               http.StatusCreated,
	BAD_REQUEST:           http.StatusBadRequest,
	UNAUTHORIZED:          http.StatusUnauthorized,
	FORBIDDEN:             http.StatusForbidden,
	NOT_FOUND:             http.StatusNotFound,
	INTERNAL_SERVER_ERROR: http.StatusInternalServerError,
}

var ResponseConstants = struct {
	General struct {
		OK                    Response
		CREATED               Response
		BAD_REQUEST           Response
		UNAUTHORIZED          Response
		FORBIDDEN             Response
		NOT_FOUND             Response
		INTERNAL_SERVER_ERROR Response
	}
}{
	General: struct {
		OK                    Response
		CREATED               Response
		BAD_REQUEST           Response
		UNAUTHORIZED          Response
		FORBIDDEN             Response
		NOT_FOUND             Response
		INTERNAL_SERVER_ERROR Response
	}{
		OK:                    Response{HTTP_STATUS: StatusCodes.OK, MESSAGE: "Ok"},
		CREATED:               Response{HTTP_STATUS: StatusCodes.CREATED, MESSAGE: "Created"},
		BAD_REQUEST:           Response{HTTP_STATUS: StatusCodes.BAD_REQUEST, MESSAGE: "Bad Request"},
		UNAUTHORIZED:          Response{HTTP_STATUS: StatusCodes.UNAUTHORIZED, MESSAGE: "Unauthorized"},
		FORBIDDEN:             Response{HTTP_STATUS: StatusCodes.FORBIDDEN, MESSAGE: "Forbidden"},
		NOT_FOUND:             Response{HTTP_STATUS: StatusCodes.NOT_FOUND, MESSAGE: "Not Found"},
		INTERNAL_SERVER_ERROR: Response{HTTP_STATUS: StatusCodes.INTERNAL_SERVER_ERROR, MESSAGE: "Internal Server Error"},
	},
}
