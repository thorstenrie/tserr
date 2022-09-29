package tserr

var (
	chkFailed     = errm{412, "check %v failed: %w"}
	nilPtr        = errm{500, "nil pointer"}
	notExistent   = errm{404, "%v does not exist"}
	opFailed      = errm{422, "%v %v failed: %w"}
	errNilFailed  = errm{500, "%v returned nil, but error expected"}
	cannotBeEmpty = errm{400, "%v cannot be empty"}
	strNotEqual   = errm{500, "%v does not equal %v"}
)
