package constant

const (
	OK      = "OK"
	SUCCESS = "success"
	FAILED  = "failed"

	SUCCESS_LOGIN         = "successfully logged in"
	SUCCESS_REGISTER_USER = "User registered successfully"

	FAILED_GET_CATS              = "failed to get list of cats"
	FAILED_CAT_NOT_FOUND         = "failed to get list of cats"
	FAILED_CAT_GENDER_IDENTIC    = "failed to match, gender of both cats are same"
	FAILED_CAT_MATCHED           = "failed to match, some of cat already matched"
	FAILED_CAT_USER_IDENCTIC     = "failed to match, user of both cats are same"
	FAILED_CAT_USER_UNAUTHORIZED = "failed to match, you're not the owner of the cat"
	FAILED_LOGIN                 = "failed to login, email or password is wrong"

	EMAIL_REGISTERED = "email already registered"

	SUCCESS_ADD_CAT    = "successfully add cat"
	SUCCESS_UPDATE_CAT = "successfully update cat"

	FAILED_GET_USER_CAT = "id is not found"
	HAS_REQUESTED_MATCH = "your cat has match requests"
)
