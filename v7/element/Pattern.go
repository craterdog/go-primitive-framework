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
	uti "github.com/craterdog/go-missing-utilities/v7"
	reg "regexp"
)

// CLASS INTERFACE

// Access Function

func PatternClass() PatternClassLike {
	return patternClass()
}

// Constructor Methods

func (c *patternClass_) Pattern(
	string_ string,
) PatternLike {
	var matches = patternMatcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the pattern constructor method: %s",
			string_,
		)
		panic(message)
	}
	switch matches[0] {
	case "none":
		return c.none_
	case "any":
		return c.any_
	default:
		return pattern_(matches[1]) // Strip off the trailing '?' character.
	}
}

// Constant Methods

func (c *patternClass_) None() PatternLike {
	return c.none_
}

func (c *patternClass_) Any() PatternLike {
	return c.any_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v pattern_) GetClass() PatternClassLike {
	return patternClass()
}

func (v pattern_) GetIntrinsic() string {
	return string(v)
}

// Attribute Methods

// Lexical Methods

func (v pattern_) AsString() string {
	var string_ string
	switch v {
	case `^none$`:
		string_ = `none`
	case `.*`:
		string_ = `any`
	default:
		string_ = `"` + string(v) + `"?`
	}
	return string_
}

// Matchable Methods

func (v pattern_) MatchesText(
	text string,
) bool {
	var matcher = reg.MustCompile(string(v))
	return matcher.MatchString(text)
}

func (v pattern_) GetMatches(
	text string,
) []string {
	var matcher = reg.MustCompile(string(v))
	return matcher.FindStringSubmatch(text)
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
	regex_     = "(?:\"((?:" + character_ + ")+)\"\\?)"
	character_ = "(?:(?:" + escape_ + ")|\\\\\"|[^\"" + control_ + "])"
)

var patternMatcher_ = reg.MustCompile(
	"^(?:none|(?:" + regex_ + ")|any)",
)

// Instance Structure

type pattern_ string

// Class Structure

type patternClass_ struct {
	// Declare the class constants.
	none_ PatternLike
	any_  PatternLike
}

// Class Reference

func patternClass() *patternClass_ {
	return patternClassReference_
}

var patternClassReference_ = &patternClass_{
	// Initialize the class constants.
	none_: pattern_(`^none$`),
	any_:  pattern_(`.*`),
}
