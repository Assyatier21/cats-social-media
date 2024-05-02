package constant

const (
	OK      = "OK"
	SUCCESS = "success"
	FAILED  = "failed"

	SUCCESS_LOGIN         = "successfully logged in"
	SUCCESS_REGISTER_USER = "User registered successfully"

	FAILED_GET_CATS              = "failed to get list of cats"
	FAILED_WRONG_AGE_FORMAT      = "invalid age input format"
	FAILED_CAT_NOT_FOUND         = "failed to get cat, data not found"
	FAILED_MATCH_CAT_NOT_FOUND   = "failed to get match cat, data not found"
	FAILED_CAT_GENDER_IDENTIC    = "gender of both cats are same"
	FAILED_CAT_MATCHED           = "some of cat already matched"
	FAILED_CAT_USER_IDENCTIC     = "user of both cats are same"
	FAILED_CAT_USER_UNAUTHORIZED = "you are not the owner of the cat"
	FAILED_LOGIN                 = "failed to login, email or password is wrong"

	EMAIL_REGISTERED = "email already registered"

	SUCCESS_ADD_CAT    = "successfully add cat"
	SUCCESS_UPDATE_CAT = "successfully update cat"

	FAILED_GET_USER_CAT = "id is not found"
	HAS_REQUESTED_MATCH = "your cat has match requests"

	FAILED_MATCH_ID_INVALID = "match_id is no longer valid"
	FAILED_GET_MATCH_ID    = "matchId is not found"
	FAILED_MATCH_NOT_VALID = "matchId is no longer valid"
	FAILED_CAN_NOT_APPROVE = "you can not appove this match"

	SUCCESS_APPROVE_MATCH = "successfully matches the cat match request"
)
