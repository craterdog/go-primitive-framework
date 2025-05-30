/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element

import (
	fmt "fmt"
	reg "regexp"
)

// CLASS INTERFACE

// Access Function

func ResourceClass() ResourceClassLike {
	return resourceClass()
}

// Constructor Methods

func (c *resourceClass_) Resource(
	string_ string,
) ResourceLike {
	if !resourceMatcher_.MatchString(string_) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the resource constructor method: %s",
			string_,
		)
		panic(message)
	}
	return resource_(string_)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v resource_) GetClass() ResourceClassLike {
	return resourceClass()
}

func (v resource_) GetIntrinsic() string {
	return string(v)
}

// Attribute Methods

// Lexical Methods

func (v resource_) AsString() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

// Segmented Methods

func (v resource_) GetScheme() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

func (v resource_) GetAuthority() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

func (v resource_) GetPath() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

func (v resource_) GetQuery() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

func (v resource_) GetFragment() string {
	var result_ string
	// TBD - Add the method implementation.
	return result_
}

// PROTECTED INTERFACE

// Private Methods

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	scheme_       = "(?:(?:" + alpha_ + ")((?:" + alphanumeric_ + ")|\\+|-|\\.)*)"
	authority_    = "(?:[^/" + control_ + "]+)"
	path_         = "(?:[^\\?#>" + control_ + "]*)"
	query_        = "(?:[^#>" + control_ + "]*)"
	fragment_     = "(?:[^>" + control_ + "]*)"
	alpha_        = "(?:[A-Za-z])"
	alphanumeric_ = "(?:(?:" + alpha_ + ")|(?:" + base10_ + "))"
)

var resourceMatcher_ = reg.MustCompile(
	"^(?:<(?:" + scheme_ + "):(//(?:" + authority_ + "))?/(?:" + path_ +
		")(\\?(?:" + query_ + "))?(#(?:" + fragment_ + "))?>)",
)

// Instance Structure

type resource_ string

// Class Structure

type resourceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func resourceClass() *resourceClass_ {
	return resourceClassReference_
}

var resourceClassReference_ = &resourceClass_{
	// Initialize the class constants.
}
