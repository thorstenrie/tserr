package tserr

var (
	chkFailed     = errm{412, "check %v failed: %w"}
	nilPtr        = errm{500, "nil pointer"}
	opFailed      = errm{422, "%v %v failed: %w"}
	testNilFailed = errm{500, "%v returned nil, but error expected"}
)

