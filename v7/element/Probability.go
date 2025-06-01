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
	ran "crypto/rand"
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v7"
	big "math/big"
	reg "regexp"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func ProbabilityClass() ProbabilityClassLike {
	return probabilityClass()
}

// Constructor Methods

func (c *probabilityClass_) Probability(
	float float64,
) ProbabilityLike {
	var probability ProbabilityLike
	switch {
	case float < 0.0:
		probability = probability_(0.0)
	case float > 1.0:
		probability = probability_(1.0)
	default:
		probability = probability_(float)
	}
	return probability
}

func (c *probabilityClass_) ProbabilityFromBoolean(
	boolean bool,
) ProbabilityLike {
	var probability ProbabilityLike
	switch boolean {
	case true:
		probability = probability_(1)
	case false:
		probability = probability_(0)
	}
	return probability
}

func (c *probabilityClass_) ProbabilityFromString(
	string_ string,
) ProbabilityLike {
	var matches = probabilityMatcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the probability constructor method: %s",
			string_,
		)
		panic(message)
	}
	var float, _ = stc.ParseFloat(matches[1], 64) // Strip off the leading 'p'.
	return probability_(float)
}

// Constant Methods

func (c *probabilityClass_) Minimum() ProbabilityLike {
	return c.minimum_
}

func (c *probabilityClass_) Maximum() ProbabilityLike {
	return c.maximum_
}

// Function Methods

func (c *probabilityClass_) Random() ProbabilityLike {
	return probability_(c.randomProbability())
}

// INSTANCE INTERFACE

// Principal Methods

func (v probability_) GetClass() ProbabilityClassLike {
	return probabilityClass()
}

func (v probability_) GetIntrinsic() float64 {
	return float64(v)
}

// Attribute Methods

// Continuous Methods

func (v probability_) AsFloat() float64 {
	return float64(v)
}

func (v probability_) IsZero() bool {
	return v == 0
}

func (v probability_) IsInfinite() bool {
	return false
}

func (v probability_) IsUndefined() bool {
	return false
}

func (v probability_) HasMagnitude() bool {
	return !v.IsZero()
}

// Discrete Methods

func (v probability_) AsBoolean() bool {
	if v == 0.5 {
		return probabilityClass().randomBoolean()
	}
	return v > 0.5
}

func (v probability_) AsInteger() int {
	if v.AsBoolean() {
		return 1
	}
	return 0
}

// Lexical Methods

func (v probability_) AsString() string {
	return "p" + numberClass().stringFromFloat(float64(v))
}

// PROTECTED INTERFACE

// Private Methods

const maximum_ = 1 << 53

func (c *probabilityClass_) randomBoolean() bool {
	var random, err = ran.Int(ran.Reader, big.NewInt(int64(2)))
	if err != nil {
		panic(fmt.Sprintf("The random number generator gave the following error: %v", err))
	}
	return int(random.Int64()) > 0
}

func (c *probabilityClass_) randomInteger(max int) int {
	var random, err = ran.Int(ran.Reader, big.NewInt(int64(max+1)))
	if err != nil {
		panic(fmt.Sprintf("The random number generator gave the following error: %v", err))
	}
	return int(random.Int64())
}

func (c *probabilityClass_) randomProbability() float64 {
	return float64(c.randomInteger(maximum_)) / float64(maximum_)
}

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
var probabilityMatcher_ = reg.MustCompile(
	"^(?:p(0(?:" + fraction_ + ")?|1))",
)

// Instance Structure

type probability_ float64

// Class Structure

type probabilityClass_ struct {
	// Declare the class constants.
	minimum_ ProbabilityLike
	maximum_ ProbabilityLike
}

// Class Reference

func probabilityClass() *probabilityClass_ {
	return probabilityClassReference_
}

var probabilityClassReference_ = &probabilityClass_{
	// Initialize the class constants.
	// minimum_: constantValue,
	// maximum_: constantValue,
}
