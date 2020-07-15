package e

const (
	VALID                                = "valid"
	OK                                   = "ok"
	TYPE_ERROR                           = "type error"
	UPDATE_TOKEN_FAIL                    = "update token fail and token out of date"
	TOKEN_OUT_OF_DATE                    = "token out of date"
	USERNAME_NOT_EXIST_OR_PASSWORD_WRONG = "username or password wrong"
	PHONE_NOT_EXIST                      = "phone not exist"
	INFO_ERROR                           = "info not exist or wrong"
	TOKEN_ERROR                          = "token error"
	MYSQL_ERROR                          = "mysql error"
	IMAGE_FORMAT_OR_SIZE_ERROR           = "file format or size error"
	UPLOAD_CHECK_FILE_ERROR              = "check file error"
	UPLOAD_SAVE_FILE_ERROR               = "save file error"
	ROUTING_NOT_EXIST                    = "router not exist"
	APP_ERROR                            = "platform error"
	OPERATING_ERROR                      = "operating error"
	NOT_LOGIN                            = "websocket user not login"
	SERVER_ERROR                         = "server error"
)

var msgFlags = map[int]string{
	Valid:                           VALID,
	Ok:                              OK,
	TypeError:                       TYPE_ERROR,
	UpdateTokenFail:                 UPDATE_TOKEN_FAIL,
	TokenOutOfDate:                  TOKEN_OUT_OF_DATE,
	UserNameNotExistOrPasswordWrong: USERNAME_NOT_EXIST_OR_PASSWORD_WRONG,
	PhoneNotExist:                   PHONE_NOT_EXIST,
	InfoError:                       INFO_ERROR,
	TokenError:                      TOKEN_ERROR,
	MysqlError:                      MYSQL_ERROR,
	ImageFormatOrSizeError:          IMAGE_FORMAT_OR_SIZE_ERROR,
	UploadCheckFileError:            UPLOAD_CHECK_FILE_ERROR,
	UploadSaveFileError:             UPLOAD_SAVE_FILE_ERROR,
	RoutingNotExist:                 ROUTING_NOT_EXIST,
	AppError:                        APP_ERROR,
	OperatingError:                  OPERATING_ERROR,
	NotLogin:                        NOT_LOGIN,
	ServerError:                     SERVER_ERROR,
}

// 获取错误代码对应的中文意思
func GetMsg(code int) string {
	msg, ok := msgFlags[code]
	if ok {
		return msg
	}

	return msgFlags[ERROR]
}
