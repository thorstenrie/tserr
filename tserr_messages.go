// Copyright (c) 2023 thorstenrie.
// All Rights Reserved. Use is governed with GNU Affero General Public License v3.0
// that can be found in the LICENSE file.
package tserr

// Import Go standard package http
import "net/http" // http

// Error ids, error codes and error messages with their potential verbs.
var (
	errmsgCheck           = errmsg{1, http.StatusPreconditionFailed, "check %v failed: %w"}
	errmsgNotExistent     = errmsg{2, http.StatusNotFound, "%v does not exist"}
	errmsgOp              = errmsg{3, http.StatusUnprocessableEntity, "%v %v failed: %w"}
	errmsgNilFailed       = errmsg{4, http.StatusInternalServerError, "%v returned nil, but error expected"}
	errmsgEmpty           = errmsg{5, http.StatusBadRequest, "%v cannot be empty"}
	errmsgEqualStr        = errmsg{6, http.StatusInternalServerError, "value of %v is %v, but expected to be equal to %v"}
	errmsgTypeNotMatching = errmsg{7, http.StatusMethodNotAllowed, "%v does not match type %v"}
	errmsgForbidden       = errmsg{8, http.StatusForbidden, "operation on %v forbidden"}
	errmsgReturn          = errmsg{9, http.StatusInternalServerError, "%v returned %v, but %v expected"}
	errmsgHigher          = errmsg{10, http.StatusInternalServerError, "value of %v is %d, but expected to be at least equal to or higher than %d"}
	errmsgEqual           = errmsg{11, http.StatusInternalServerError, "value of %v is %d, but expected to be equal to %d"}
	errmsgLower           = errmsg{12, http.StatusInternalServerError, "value of %v is %d, but expected to be lower than %d"}
	errmsgNotSet          = errmsg{14, http.StatusNotFound, "%v not set"}
	errmsgNotAvailable    = errmsg{15, http.StatusServiceUnavailable, "%v not available: %w"}
	errmsgEqualf          = errmsg{16, http.StatusInternalServerError, "value of %v is %f, but expected to be equal to %f"}
	errmsgNonPrintable    = errmsg{17, http.StatusBadRequest, "%v contains non-printable runes, but only printable runes are allowed"}
	errmsgNotEqual        = errmsg{18, http.StatusInternalServerError, "variable %v equals variable %v, but not allowed to equal"}
	errmsgDuplicate       = errmsg{19, http.StatusForbidden, "%v is a duplicate and already exists"}
	errmsgLocked          = errmsg{20, http.StatusLocked, "%v is locked"}
)
