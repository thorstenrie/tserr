// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tserr

// Error ids, error codes and error messages with their potential verbs.
var (
	errmsgCheck           = errmsg{1, 412, "check %v failed: %w"}
	errmsgNotExistent     = errmsg{2, 404, "%v does not exist"}
	errmsgOp              = errmsg{3, 422, "%v %v failed: %w"}
	errmsgNilFailed       = errmsg{4, 500, "%v returned nil, but error expected"}
	errmsgEmpty           = errmsg{5, 400, "%v cannot be empty"}
	errmsgNotEqualStr     = errmsg{6, 500, "%v does not equal %v"}
	errmsgTypeNotMatching = errmsg{7, 405, "%v does not match type %v"}
	errmsgForbidden       = errmsg{8, 403, "operation on %v forbidden"}
	errmsgReturn          = errmsg{9, 500, "%v returned %v, but %v expected"}
	errmsgHigher          = errmsg{10, 400, "value of %v is %d, but expected to be at least equal to or higher than %d"}
	errmsgEqual           = errmsg{11, 400, "value of %v is %d, but expected to be equal to %d"}
	errmsgNotSet          = errmsg{12, 404, "%v not set"}
)
