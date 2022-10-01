package tserr

var (
	errmsgCheck           = errmsg{1, 412, "check %v failed: %w"}
	errmsgNotExistent     = errmsg{2, 404, "%v does not exist"}
	errmsgOp              = errmsg{3, 422, "%v %v failed: %w"}
	errmsgNilFailed       = errmsg{4, 500, "%v returned nil, but error expected"}
	errmsgEmpty           = errmsg{5, 400, "%v cannot be empty"}
	errmsgNotEqualStr     = errmsg{6, 500, "%v does not equal %v"}
	errmsgTypeNotMatching = errmsg{7, 405, "%v does not match type %v"}
)
