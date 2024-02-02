package utils

type ErrorResponse struct {
	Message CodeError `json:"message"`
	Detail  string    `json:"detail"`
}

type CodeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	LOGIC_CRASH            = CodeError{Code: "1000", Message: "Logic Crush"}
	UNAUTHORIZED           = CodeError{Code: "1001", Message: "User not allowed"}
	ID_NOT_FOUND           = CodeError{Code: "1002", Message: "ID must be provided"}
	ID_NOT_VALID           = CodeError{Code: "1003", Message: "ID not valid"}
	WRITER_NOT_FOUND       = CodeError{Code: "1004", Message: "Writer must be provided"}
	SCHEMA_NOT_MATCH       = CodeError{Code: "1005", Message: "Schema does not match"}
	UNSUPPORTED_MEDIA_TYPE = CodeError{Code: "415", Message: "This type of file is not supported"}
	WARNING_FILE           = CodeError{Code: "1006", Message: "File could have potential virus"}
	NOT_FOUND              = CodeError{Code: "404", Message: "Not found"}
	OPERATION_FAIL         = CodeError{Code: "417", Message: "Operation fail"}
	REQUIRED_FIELD         = CodeError{Code: "418", Message: "Required param"}
	IS_NOT_MEMBER          = CodeError{Code: "1007", Message: "Is not a member of"}
	BROKER_CONNECTION_FAIL = CodeError{Code: "1008", Message: "Broker not responding"}
	BROKER_SEND_FAIL       = CodeError{Code: "1009", Message: "Broker attempt to send failed"}
	BROKER_CHANNEL_ERROR   = CodeError{Code: "1010", Message: "Broker channel has an error"}
	DUPLICATE_ID           = CodeError{Code: "1011", Message: "ID is present in the repository"}
	DB_CONNECTION_FAIL     = CodeError{Code: "1012", Message: "Database not responding"}
	DB_COLLECTION_ERROR    = CodeError{Code: "1013", Message: "Table is not present in the database"}
	DB_GET_FAIL            = CodeError{Code: "1014", Message: "Database attempt get"}
	DB_ID_NOT_FOUND        = CodeError{Code: "1015", Message: "Database attempt get by identifier"}
	DB_CREATE_FAIL         = CodeError{Code: "1016", Message: "Database attempt create"}
	DB_UPDATE_FAIL         = CodeError{Code: "1017", Message: "Database attempt update"}
	DB_DELETE_FAIL         = CodeError{Code: "1018", Message: "Database attempt delete"}
)
