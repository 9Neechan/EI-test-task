package api

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// mapGRPCToHTTPStatus maps a gRPC status code to an HTTP status code
func mapGRPCToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.Internal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
