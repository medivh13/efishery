package errors

const (
	UNKNOWN_ERROR                  ErrorCode = 0
	DATA_INVALID                   ErrorCode = 4_001_00_00001
	DATA_ALREADY_EXIST             ErrorCode = 4_001_00_00002
	USER_INVALID                   ErrorCode = 4_001_00_00003
	INVALID_PAYLOAD_CREATE_ORDER   ErrorCode = 4_001_00_00005
	INVALID_PAYLOAD_UPDATE_ORDER   ErrorCode = 4_001_00_00006
	INVALID_ORDER                  ErrorCode = 4_001_00_00007
	ORDER_NOT_FOUND                ErrorCode = 4_001_00_00008
	FAILED_RETRIEVE_ORDER          ErrorCode = 4_001_00_00009
	FAILED_CREATE_ORDER            ErrorCode = 4_001_00_00010
	FAILED_UPDATE_ORDER            ErrorCode = 4_001_00_00011
	FAILED_DELETE_ORDER            ErrorCode = 4_001_00_00012
	INVALID_REQUEST_RETRIEVE_ORDER ErrorCode = 4_001_00_00013
	INVALID_REQUEST_CREATE_ORDER   ErrorCode = 4_001_00_00014
	INVALID_REQUEST_UPDATE_ORDER   ErrorCode = 4_001_00_00015
	INVALID_REQUEST_DELETE_ORDER   ErrorCode = 4_001_00_00016
	INVALID_REQUEST_RETRIEVE_STORE ErrorCode = 4_001_00_00017
	FAILED_RETRIEVE_STATUS_PAGE    ErrorCode = 4_001_00_00018
	STATUS_PAGE_NOT_FOUND          ErrorCode = 4_001_00_00019
	INVALID_HEADER_X_PLATFORM      ErrorCode = 4_001_00_00020
	INVALID_HEADER_X_API_KEY       ErrorCode = 4_001_00_00021
	UNAUTHORIZED                   ErrorCode = 4_001_00_00022
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},

	INVALID_PAYLOAD_CREATE_ORDER: {
		ClientMessage: "Failed to create order.",
		SystemMessage: "Request payload for create order has an invalid form.",
		ErrorCode:     INVALID_PAYLOAD_CREATE_ORDER,
	},
	INVALID_PAYLOAD_UPDATE_ORDER: {
		ClientMessage: "Failed to update order.",
		SystemMessage: "Request payload for update order has an invalid form.",
		ErrorCode:     INVALID_PAYLOAD_UPDATE_ORDER,
	},
	INVALID_ORDER: {
		ClientMessage: "Invalid order.",
		SystemMessage: "Invalid request order.",
		ErrorCode:     INVALID_ORDER,
	},
	ORDER_NOT_FOUND: {
		ClientMessage: "Invalid order.",
		SystemMessage: "Order not found.",
		ErrorCode:     ORDER_NOT_FOUND,
	},
	FAILED_RETRIEVE_ORDER: {
		ClientMessage: "Failed to retrieve order.",
		SystemMessage: "Something wrong happened while retrieve order.",
		ErrorCode:     FAILED_RETRIEVE_ORDER,
	},
	FAILED_CREATE_ORDER: {
		ClientMessage: "Failed to create order.",
		SystemMessage: "Something wrong happened while create order.",
		ErrorCode:     FAILED_CREATE_ORDER,
	},
	FAILED_UPDATE_ORDER: {
		ClientMessage: "Failed to update order.",
		SystemMessage: "Something wrong happened while update order.",
		ErrorCode:     FAILED_UPDATE_ORDER,
	},
	FAILED_DELETE_ORDER: {
		ClientMessage: "Failed to delete order.",
		SystemMessage: "Something wrong happened while delete order.",
		ErrorCode:     FAILED_DELETE_ORDER,
	},
	INVALID_REQUEST_RETRIEVE_ORDER: {
		ClientMessage: "Failed to request order.",
		SystemMessage: "Request has an invalid query params and/or payload to retrieve order.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_ORDER,
	},
	INVALID_REQUEST_CREATE_ORDER: {
		ClientMessage: "Failed to create order.",
		SystemMessage: "Request has an invalid query params and/or payload to create order.",
		ErrorCode:     INVALID_REQUEST_CREATE_ORDER,
	},
	INVALID_REQUEST_UPDATE_ORDER: {
		ClientMessage: "Failed to update order.",
		SystemMessage: "Request has an invalid query params and/or payload to update order.",
		ErrorCode:     INVALID_REQUEST_UPDATE_ORDER,
	},
	INVALID_REQUEST_DELETE_ORDER: {
		ClientMessage: "Failed to delete order.",
		SystemMessage: "Request has an invalid query params and/or payload to delete order.",
		ErrorCode:     INVALID_REQUEST_DELETE_ORDER,
	},
	INVALID_REQUEST_RETRIEVE_STORE: {
		ClientMessage: "Failed to request store.",
		SystemMessage: "Request has an invalid query params and/or payload to retrieve store.",
		ErrorCode:     INVALID_REQUEST_RETRIEVE_STORE,
	},
	FAILED_RETRIEVE_STATUS_PAGE: {
		ClientMessage: "Failed to retrieve data",
		SystemMessage: "Something wrong happened while retrieve data",
		ErrorCode:     FAILED_RETRIEVE_ORDER,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	INVALID_HEADER_X_PLATFORM: {
		ClientMessage: "Invalid platform.",
		SystemMessage: "Invalid value of header X-Platform.",
		ErrorCode:     INVALID_HEADER_X_PLATFORM,
	},
	INVALID_HEADER_X_API_KEY: {
		ClientMessage: "Invalid Api Key.",
		SystemMessage: "Invalid value of header X-Api-Key.",
		ErrorCode:     INVALID_HEADER_X_API_KEY,
	},
	UNAUTHORIZED: {
		ClientMessage: "Unauthorized",
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	DATA_ALREADY_EXIST: {
		ClientMessage: "Data Already Exist",
		SystemMessage: "Data Already Exist",
		ErrorCode:     DATA_ALREADY_EXIST,
	},
	USER_INVALID: {
		ClientMessage: "invalid user",
		SystemMessage: "invalid user",
		ErrorCode:     USER_INVALID,
	},
}
