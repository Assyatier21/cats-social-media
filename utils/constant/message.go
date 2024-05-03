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
	FAILED_GET_MATCH_CATS        = "failed to get list of match cats"
	FAILED_REQUEST_MATCH_CATS    = "failed to match cat, already have a pending request"
	FAILED_LOGIN                 = "failed to login, email or password is wrong"

	EMAIL_REGISTERED = "This email is already registered. Please choose another one"

	SUCCESS_ADD_CAT    = "The cat record has been added successfully"
	SUCCESS_UPDATE_CAT = "The cat record has been updated successfully"
	SUCCESS_DELETE_CAT = "The cat record was deleted successfully"

	FAILED_GET_USER_CAT = "UserID is not found"
	HAS_REQUESTED_MATCH = "The corresponding cat is matched or already has match request(s)"

	FAILED_MATCH_ID_INVALID = "match_id is no longer valid"
	FAILED_GET_MATCH_ID     = "matchId is not found"
	FAILED_MATCH_NOT_VALID  = "matchId is no longer valid"
	FAILED_CAN_NOT_APPROVE  = "you can not appove this match"

	SUCCESS_APPROVE_MATCH = "successfully matches the cat match request"
)
