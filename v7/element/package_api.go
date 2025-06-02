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

/*
Package "element" provides a framework of aspects and class definitions for a
rich set of primitive data types that are elemental.  All primitive types are
immutable and—for better performance—are implemented as extensions to existing
Go primitive types.

For detailed documentation on this package refer to the wiki:
  - https://github.com/craterdog/go-primitive-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package element

import (
	uri "net/url"
)

// TYPE DECLARATIONS

/*
Units is a constrained type representing the possible units for an angle.
*/
type Units uint8

const (
	Degrees Units = iota
	Radians
	Gradians
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AngleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
angle-like concrete class.
*/
type AngleClassLike interface {
	// Constructor Methods
	Angle(
		radians float64,
	) AngleLike
	AngleFromString(
		string_ string,
	) AngleLike

	// Constant Methods
	Minimum() AngleLike
	Maximum() AngleLike
	Zero() AngleLike
	Pi() AngleLike
	Tau() AngleLike

	// Function Methods
	Inverse(
		angle AngleLike,
	) AngleLike
	Sum(
		first AngleLike,
		second AngleLike,
	) AngleLike
	Difference(
		first AngleLike,
		second AngleLike,
	) AngleLike
	Scaled(
		angle AngleLike,
		factor float64,
	) AngleLike
	Complement(
		angle AngleLike,
	) AngleLike
	Supplement(
		angle AngleLike,
	) AngleLike
	Conjugate(
		angle AngleLike,
	) AngleLike
	Cosine(
		angle AngleLike,
	) float64
	ArcCosine(
		x float64,
	) AngleLike
	Sine(
		angle AngleLike,
	) float64
	ArcSine(
		y float64,
	) AngleLike
	Tangent(
		angle AngleLike,
	) float64
	ArcTangent(
		x float64,
		y float64,
	) AngleLike
}

/*
BooleanClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
boolean-like concrete class.
*/
type BooleanClassLike interface {
	// Constructor Methods
	Boolean(
		boolean bool,
	) BooleanLike
	BooleanFromString(
		string_ string,
	) BooleanLike

	// Constant Methods
	False() BooleanLike
	True() BooleanLike

	// Function Methods
	Not(
		boolean BooleanLike,
	) BooleanLike
	And(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
	San(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
	Ior(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
	Xor(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
}

/*
CitationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
citation-like concrete class.
*/
type CitationClassLike interface {
	// Constructor Methods
	Citation(
		string_ string,
	) CitationLike
}

/*
DurationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
duration-like concrete class.
*/
type DurationClassLike interface {
	// Constructor Methods
	Duration(
		milliseconds int,
	) DurationLike
	DurationFromString(
		string_ string,
	) DurationLike

	// Constant Methods
	Minimum() DurationLike
	Maximum() DurationLike
	MillisecondsPerSecond() int
	MillisecondsPerMinute() int
	MillisecondsPerHour() int
	MillisecondsPerDay() int
	MillisecondsPerWeek() int
	MillisecondsPerMonth() int
	MillisecondsPerYear() int
	DaysPerMonth() float64
	DaysPerYear() float64
	WeeksPerMonth() float64
}

/*
GlyphClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
glyph-like concrete class.
*/
type GlyphClassLike interface {
	// Constructor Methods
	Glyph(
		rune_ rune,
	) GlyphLike
	GlyphFromInteger(
		integer int,
	) GlyphLike
	GlyphFromString(
		string_ string,
	) GlyphLike

	// Constant Methods
	Minimum() GlyphLike
	Maximum() GlyphLike

	// Function Methods
	ToLowercase(
		glyph GlyphLike,
	) GlyphLike
	ToUppercase(
		glyph GlyphLike,
	) GlyphLike
}

/*
MomentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
moment-like concrete class.
*/
type MomentClassLike interface {
	// Constructor Methods
	Moment(
		milliseconds int,
	) MomentLike
	MomentFromString(
		string_ string,
	) MomentLike

	// Constant Methods
	Minimum() MomentLike
	Maximum() MomentLike
	Epoch() MomentLike

	// Function Methods
	Now() MomentLike
	Earlier(
		moment MomentLike,
		duration DurationLike,
	) MomentLike
	Later(
		moment MomentLike,
		duration DurationLike,
	) MomentLike
	Duration(
		first MomentLike,
		second MomentLike,
	) DurationLike
}

/*
NumberClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
number-like concrete class.
*/
type NumberClassLike interface {
	// Constructor Methods
	Number(
		complex_ complex128,
	) NumberLike
	NumberFromPolar(
		magnitude float64,
		phase float64,
	) NumberLike
	NumberFromRectangular(
		real_ float64,
		imaginary float64,
	) NumberLike
	NumberFromString(
		string_ string,
	) NumberLike

	// Constant Methods
	Minimum() NumberLike
	Maximum() NumberLike
	Zero() NumberLike
	One() NumberLike
	I() NumberLike
	E() NumberLike
	Pi() NumberLike
	Phi() NumberLike
	Tau() NumberLike
	Infinity() NumberLike
	Undefined() NumberLike

	// Function Methods
	Inverse(
		number NumberLike,
	) NumberLike
	Sum(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Difference(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Scaled(
		number NumberLike,
		factor float64,
	) NumberLike
	Reciprocal(
		number NumberLike,
	) NumberLike
	Conjugate(
		number NumberLike,
	) NumberLike
	Product(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Quotient(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Remainder(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Power(
		base NumberLike,
		exponent NumberLike,
	) NumberLike
	Logarithm(
		base NumberLike,
		number NumberLike,
	) NumberLike
}

/*
PatternClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
pattern-like concrete class.
*/
type PatternClassLike interface {
	// Constructor Methods
	Pattern(
		string_ string,
	) PatternLike

	// Constant Methods
	None() PatternLike
	Any() PatternLike
}

/*
PercentageClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
percentage-like concrete class.
*/
type PercentageClassLike interface {
	// Constructor Methods
	Percentage(
		float float64,
	) PercentageLike
	PercentageFromInteger(
		integer int,
	) PercentageLike
	PercentageFromString(
		string_ string,
	) PercentageLike
}

/*
ProbabilityClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
probability-like concrete class.
*/
type ProbabilityClassLike interface {
	// Constructor Methods
	Probability(
		float float64,
	) ProbabilityLike
	ProbabilityFromBoolean(
		boolean bool,
	) ProbabilityLike
	ProbabilityFromString(
		string_ string,
	) ProbabilityLike

	// Constant Methods
	Minimum() ProbabilityLike
	Maximum() ProbabilityLike

	// Function Methods
	Random() ProbabilityLike
}

/*
ResourceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
resource-like concrete class.
*/
type ResourceClassLike interface {
	// Constructor Methods
	Resource(
		string_ string,
	) ResourceLike
	ResourceFromUri(
		url *uri.URL,
	) ResourceLike
}

// INSTANCE DECLARATIONS

/*
AngleLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete angle-like class.
*/
type AngleLike interface {
	// Principal Methods
	GetClass() AngleClassLike
	GetIntrinsic() float64

	// Aspect Interfaces
	Angular
	Continuous
	Lexical
}

/*
BooleanLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a boolean-like elemental class.
*/
type BooleanLike interface {
	// Principal Methods
	GetClass() BooleanClassLike
	GetIntrinsic() bool

	// Aspect Interfaces
	Discrete
	Lexical
}

/*
CitationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a citation-like elemental class.
*/
type CitationLike interface {
	// Principal Methods
	GetClass() CitationClassLike
	GetIntrinsic() string

	// Aspect Interfaces
	Lexical
	Named
	Versioned
}

/*
DurationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a duration-like elemental class.
*/
type DurationLike interface {
	// Principal Methods
	GetClass() DurationClassLike
	GetIntrinsic() int

	// Aspect Interfaces
	Discrete
	Lexical
	Polarized
	Temporal
	Factored
}

/*
GlyphLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a glyph-like elemental class.
*/
type GlyphLike interface {
	// Principal Methods
	GetClass() GlyphClassLike
	GetIntrinsic() rune

	// Aspect Interfaces
	Discrete
	Lexical
}

/*
MomentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a moment-like elemental class.
*/
type MomentLike interface {
	// Principal Methods
	GetClass() MomentClassLike
	GetIntrinsic() int

	// Aspect Interfaces
	Discrete
	Lexical
	Temporal
	Factored
}

/*
NumberLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a number-like elemental class.
*/
type NumberLike interface {
	// Principal Methods
	GetClass() NumberClassLike
	GetIntrinsic() complex128

	// Aspect Interfaces
	Complex
	Continuous
	Lexical
	Polarized
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a pattern-like elemental class.
*/
type PatternLike interface {
	// Principal Methods
	GetClass() PatternClassLike
	GetIntrinsic() string

	// Aspect Interfaces
	Lexical
	Matchable
}

/*
PercentageLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a percentage-like elemental class.
*/
type PercentageLike interface {
	// Principal Methods
	GetClass() PercentageClassLike
	GetIntrinsic() float64

	// Aspect Interfaces
	Continuous
	Discrete
	Lexical
	Polarized
}

/*
ProbabilityLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a probability-like elemental class.
*/
type ProbabilityLike interface {
	// Principal Methods
	GetClass() ProbabilityClassLike
	GetIntrinsic() float64

	// Aspect Interfaces
	Continuous
	Discrete
	Lexical
}

/*
ResourceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a resource-like elemental class.
*/
type ResourceLike interface {
	// Principal Methods
	GetClass() ResourceClassLike
	GetIntrinsic() string
	AsUri() *uri.URL

	// Aspect Interfaces
	Lexical
	Segmented
}

// ASPECT DECLARATIONS

/*
Angular is an aspect interface that declares a set of method signatures that
must be supported by each instance of an angular concrete class.
*/
type Angular interface {
	InUnits(
		units Units,
	) float64
	GetParts() (
		x float64,
		y float64,
	)
}

/*
Complex is an aspect interface that defines a set of method signatures
that must be supported by each instance of a complex elemental class.
*/
type Complex interface {
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() float64
}

/*
Continuous is an aspect interface that defines a set of method signatures
that must be supported by each instance of a continuous elemental class.
*/
type Continuous interface {
	AsFloat() float64
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
	HasMagnitude() bool
}

/*
Discrete is an aspect interface that defines a set of method signatures
that must be supported by each instance of a discrete elemental class.
*/
type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

/*
Factored is an aspect interface that defines a set of method signatures
that must be supported by each instance of a factored elemental class.
*/
type Factored interface {
	GetMilliseconds() int
	GetSeconds() int
	GetMinutes() int
	GetHours() int
	GetDays() int
	GetWeeks() int
	GetMonths() int
	GetYears() int
}

/*
Lexical is an aspect interface that defines a set of method signatures
that must be supported by each instance of a lexical elemental class.
*/
type Lexical interface {
	AsString() string
}

/*
Matchable is an aspect interface that defines a set of method signatures
that must be supported by each instance of a matchable elemental class.
*/
type Matchable interface {
	MatchesText(
		text string,
	) bool
	GetMatches(
		text string,
	) []string
}

/*
Named is an aspect interface that defines a set of method signatures
that must be supported by each instance of a named elemental class.
*/
type Named interface {
	GetName() string
}

/*
Polarized is an aspect interface that defines a set of method signatures
that must be supported by each instance of a polarized elemental class.
*/
type Polarized interface {
	IsNegative() bool
}

/*
Segmented is an aspect interface that defines a set of method signatures
that must be supported by each instance of a segmented elemental class.
*/
type Segmented interface {
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

/*
Temporal is an aspect interface that defines a set of method signatures
that must be supported by each instance of a temporal elemental class.
*/
type Temporal interface {
	AsMilliseconds() float64
	AsSeconds() float64
	AsMinutes() float64
	AsHours() float64
	AsDays() float64
	AsWeeks() float64
	AsMonths() float64
	AsYears() float64
}

/*
Versioned is an aspect interface that defines a set of method signatures
that must be supported by each instance of a versioned elemental class.
*/
type Versioned interface {
	GetVersion() string
}
