package ecode

var (
	// ErrInvalidUsername .
	ErrInvalidUsername = "username is  invalid"

	// ErrInvalidCursor - invalid cursor.
	ErrInvalidCursor = "cursor is invalid"
	// ErrInvalidLimit - invalid limit.
	ErrInvalidLimit = "limit is invalid"

	// ErrInvalidPassword - password is invalid.
	ErrInvalidPassword = "password is invalid"
	// ErrGenerateToken - generate token failed.
	ErrGenerateToken = "generate token failed"

	// ErrRegisterTokenInvalid - register token is invalid.
	// ErrRegisterTokenInvalid = "register token is invalid"

	// ErrEmailIsEmpty - email is empty.
	ErrEmailIsEmpty = "email is empty"
	// ErrUsernameIsEmpty - username is empty.
	ErrUsernameIsEmpty = "username is empty"
	// ErrPasswordIsEmpty - password is empty.
	ErrPasswordIsEmpty = "password is empty"

	// ErrParamsIsInvalid - params is invalid.
	ErrParamsIsInvalid = "email or phone is invalid"
	// ErrParamsIsEmpty - params is empty.
	ErrParamsIsEmpty = "email or phone is empty"

	// ErrCodeIsEmpty - code is empty.
	ErrCodeIsEmpty = "code is empty"
	// ErrCodeIsInvalid - code is invalid.
	ErrCodeIsInvalid = "code is invalid"
	// ErrCodeIsNotExist - code is not exist.
	ErrCodeIsNotExist = "code is not exist"
	// ErrCodeIsAlreadyUsed - code is already used.
	ErrCodeIsAlreadyUsed = "code is already used"

	// ErrCommentAlreadyExist - comment already exist
	ErrCommentAlreadyExist = "comment already exist"
	// ErrCommentNotExist - comment is not exist
	ErrCommentNotExist = "comment is not exist"

	// ErrFileAlreadyExist - file already exist
	ErrFileAlreadyExist = "file already exist"
	// ErrFileNotExist - file is not exist
	ErrFileNotExist = "file is not exist"

	// ErrProfileNotExist - profile  is not exist.
	ErrProfileNotExist = "profile is is not exist"
	// ErrProfileAlreadyExist - profile already exist.
	ErrProfileAlreadyExist = "profile is already exist"

	// ErrTaxonomyNameEmpty - taxonomy name is empty
	ErrTaxonomyNameEmpty = "taxonomy name is empty"
	// ErrTaxonomyAlreadyExist - taxonomy already exist
	ErrTaxonomyAlreadyExist = "taxonomy already exist"
	// ErrTaxonomyNotExist - taxonomy is not exist
	ErrTaxonomyNotExist = "taxonomy is not exist"

	// ErrUserEmpty - user is empty
	ErrUserEmpty = "user is empty"

	// ErrTopicNameEmpty - topic name is empty
	ErrTopicNameEmpty = "topic name is empty"
	// ErrTopicAlreadyExist - topic already exist
	ErrTopicAlreadyExist = "topic already exist"
	// ErrTopicNotExist - topic is not exist
	ErrTopicNotExist = "topic is not exist"

	// ErrInvalidOldPassword - old password is invalid
	ErrInvalidOldPassword = "old password is invalid"

	// ErrUserNotExist - user is is not exist.
	ErrUserNotExist = "user is is not exist"
	// ErrUserAlreadyExist - user is already exist.
	ErrUserAlreadyExist = "user is already exist"
)

var ecodeText = map[int]string{
	0:   "OK",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	409: "Conflict",
	500: "Internal Server Error",
}
