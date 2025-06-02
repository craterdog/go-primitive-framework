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
	mat "math"
	reg "regexp"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func DurationClass() DurationClassLike {
	return durationClass()
}

// Constructor Methods

func (c *durationClass_) Duration(
	milliseconds int,
) DurationLike {
	return duration_(milliseconds)
}

func (c *durationClass_) DurationFromString(
	string_ string,
) DurationLike {
	var matches = durationMatcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the duration constructor method: %s",
			string_,
		)
		panic(message)
	}
	return duration_(c.durationFromMatches(matches))
}

// Constant Methods

func (c *durationClass_) Minimum() DurationLike {
	return c.minimum_
}

func (c *durationClass_) Maximum() DurationLike {
	return c.maximum_
}

func (c *durationClass_) MillisecondsPerSecond() int {
	return c.millisecondsPerSecond_
}

func (c *durationClass_) MillisecondsPerMinute() int {
	return c.millisecondsPerMinute_
}

func (c *durationClass_) MillisecondsPerHour() int {
	return c.millisecondsPerHour_
}

func (c *durationClass_) MillisecondsPerDay() int {
	return c.millisecondsPerDay_
}

func (c *durationClass_) MillisecondsPerWeek() int {
	return c.millisecondsPerWeek_
}

func (c *durationClass_) MillisecondsPerMonth() int {
	return c.millisecondsPerMonth_
}

func (c *durationClass_) MillisecondsPerYear() int {
	return c.millisecondsPerYear_
}

func (c *durationClass_) DaysPerMonth() float64 {
	return c.daysPerMonth_
}

func (c *durationClass_) DaysPerYear() float64 {
	return c.daysPerYear_
}

func (c *durationClass_) WeeksPerMonth() float64 {
	return c.weeksPerMonth_
}

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v duration_) GetClass() DurationClassLike {
	return durationClass()
}

func (v duration_) GetIntrinsic() int {
	return int(v)
}

// Attribute Methods

// Discrete Methods

func (v duration_) AsBoolean() bool {
	return v != 0
}

func (v duration_) AsInteger() int {
	return int(v)
}

// Lexical Methods

func (v duration_) AsString() string {
	var builder sts.Builder
	builder.WriteString("~")
	if v.IsNegative() {
		builder.WriteString("-")
	}
	builder.WriteString("P")
	var float = mat.Abs(v.AsWeeks())
	var weeks = int(float)
	if float64(weeks) == float {
		// It is an exact number of weeks.
		builder.WriteString(stc.FormatInt(int64(weeks), 10))
		builder.WriteString("W")
		return builder.String()
	}
	var years = v.GetYears()
	if years > 0 {
		builder.WriteString(stc.FormatInt(int64(years), 10))
		builder.WriteString("Y")
	}
	var months = v.GetMonths()
	if months > 0 {
		builder.WriteString(stc.FormatInt(int64(months), 10))
		builder.WriteString("M")
	}
	var days = v.GetDays()
	if days > 0 {
		builder.WriteString(stc.FormatInt(int64(days), 10))
		builder.WriteString("D")
	}
	var hours = v.GetHours()
	var minutes = v.GetMinutes()
	var seconds = v.GetSeconds()
	var milliseconds = v.GetMilliseconds()
	if hours+minutes+seconds+milliseconds == 0 {
		// There is no time part of the duration.
		return builder.String()
	}
	builder.WriteString("T")
	if hours > 0 {
		builder.WriteString(stc.FormatInt(int64(hours), 10))
		builder.WriteString("H")
	}
	if minutes > 0 {
		builder.WriteString(stc.FormatInt(int64(minutes), 10))
		builder.WriteString("M")
	}
	if seconds+milliseconds > 0 {
		builder.WriteString(stc.FormatInt(int64(seconds), 10))
		if milliseconds > 0 {
			builder.WriteString(".")
			builder.WriteString(stc.FormatInt(int64(milliseconds), 10))
		}
		builder.WriteString("S")
	}
	return builder.String()
}

// Polarized Methods

func (v duration_) IsNegative() bool {
	return v < 0
}

// Temporal Methods

func (v duration_) AsMilliseconds() float64 {
	return float64(v)
}

func (v duration_) AsSeconds() float64 {
	return float64(v) / float64(durationClass().millisecondsPerSecond_)
}

func (v duration_) AsMinutes() float64 {
	return float64(v) / float64(durationClass().millisecondsPerMinute_)
}

func (v duration_) AsHours() float64 {
	return float64(v) / float64(durationClass().millisecondsPerHour_)
}

func (v duration_) AsDays() float64 {
	return float64(v) / float64(durationClass().millisecondsPerDay_)
}

func (v duration_) AsWeeks() float64 {
	return float64(v) / float64(durationClass().millisecondsPerWeek_)
}

func (v duration_) AsMonths() float64 {
	return float64(v) / float64(durationClass().millisecondsPerMonth_)
}

func (v duration_) AsYears() float64 {
	return float64(v) / float64(durationClass().millisecondsPerYear_)
}

// Factored Methods

func (v duration_) GetMilliseconds() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off everything but the milliseconds.
	milliseconds = milliseconds % durationClass().millisecondsPerSecond_
	return milliseconds
}

func (v duration_) GetSeconds() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass().millisecondsPerDay_)

	// Strip off the hours.
	milliseconds = milliseconds - (v.GetHours() * durationClass().millisecondsPerHour_)

	// Strip off the minutes.
	milliseconds = milliseconds - (v.GetMinutes() * durationClass().millisecondsPerMinute_)

	// Convert to seconds.
	var seconds = milliseconds / durationClass().millisecondsPerSecond_
	return seconds
}

func (v duration_) GetMinutes() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass().millisecondsPerDay_)

	// Strip off the hours.
	milliseconds = milliseconds - (v.GetHours() * durationClass().millisecondsPerHour_)

	// Convert to minutes.
	var minutes = milliseconds / durationClass().millisecondsPerMinute_
	return minutes
}

func (v duration_) GetHours() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass().millisecondsPerDay_)

	// Convert to hours.
	var hours = milliseconds / durationClass().millisecondsPerHour_
	return hours
}

func (v duration_) GetDays() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass().millisecondsPerMonth_)

	// Convert to days.
	var days = milliseconds / durationClass().millisecondsPerDay_
	return days
}

func (v duration_) GetWeeks() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Convert to weeks.
	var weeks = milliseconds / durationClass().millisecondsPerWeek_
	return weeks
}

func (v duration_) GetMonths() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass().millisecondsPerYear_)

	// Convert to months.
	var months = milliseconds / durationClass().millisecondsPerMonth_
	return months
}

func (v duration_) GetYears() int {
	// Retrieve the total number of milliseconds.
	var milliseconds = durationClass().magnitude(int(v))

	// Convert to years.
	var years = milliseconds / durationClass().millisecondsPerYear_
	return years
}

// PROTECTED INTERFACE

// Private Methods

func (c *durationClass_) durationFromMatches(matches []string) int {
	var milliseconds = 0.0
	var sign = 1.0
	var isTime = false
	for _, match := range matches[1:] {
		if match != "" {
			var stype = match[len(match)-1:] // Strip off the time span.
			var tspan = match[:len(match)-1] // Strip off the span type.
			var float, _ = stc.ParseFloat(tspan, 64)
			switch stype {
			case "-":
				sign = -1.0
			case "W":
				milliseconds += float * float64(durationClass().millisecondsPerWeek_)
			case "Y":
				milliseconds += float * float64(durationClass().millisecondsPerYear_)
			case "M":
				if isTime {
					milliseconds += float * float64(durationClass().millisecondsPerMinute_)
				} else {
					milliseconds += float * float64(durationClass().millisecondsPerMonth_)
				}
			case "D":
				milliseconds += float * float64(durationClass().millisecondsPerDay_)
			case "T":
				isTime = true
			case "H":
				milliseconds += float * float64(durationClass().millisecondsPerHour_)
			case "S":
				milliseconds += float * float64(durationClass().millisecondsPerSecond_)
			}
		}
	}
	return int(sign * milliseconds)
}

func (c *durationClass_) magnitude(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	days_     = "(?:(?:" + timespan_ + ")D)"
	hours_    = "(?:(?:" + timespan_ + ")H)"
	minutes_  = "(?:(?:" + timespan_ + ")M)"
	months_   = "(?:(?:" + timespan_ + ")M)"
	seconds_  = "(?:(?:" + timespan_ + ")S)"
	timespan_ = "(?:0|((?:" + ordinal_ + ")(?:" + fraction_ + ")?))"
	weeks_    = "(?:(?:" + timespan_ + ")W)"
	years_    = "(?:(?:" + timespan_ + ")Y)"
)

var durationMatcher_ = reg.MustCompile(
	"^(?:~(?:" + sign_ + ")?P((?:" + weeks_ + ")|((?:" + years_ + ")?(?:" +
		months_ + ")?(?:" + days_ + ")?(T(?:" + hours_ + ")?(?:" + minutes_ +
		")?(?:" + seconds_ + ")?)?)))",
)

// Instance Structure

type duration_ int

// Class Structure

type durationClass_ struct {
	// Declare the class constants.
	minimum_               DurationLike
	maximum_               DurationLike
	millisecondsPerSecond_ int
	millisecondsPerMinute_ int
	millisecondsPerHour_   int
	millisecondsPerDay_    int
	millisecondsPerWeek_   int
	millisecondsPerMonth_  int
	millisecondsPerYear_   int
	daysPerMonth_          float64
	daysPerYear_           float64
	weeksPerMonth_         float64
}

// Class Reference

func durationClass() *durationClass_ {
	return durationClassReference_
}

var durationClassReference_ = &durationClass_{
	// Initialize the class constants.
	minimum_: duration_(0),
	maximum_: duration_(mat.MaxInt),

	// These are locked to the Earth's daily revolutions.
	millisecondsPerSecond_: 1000,
	millisecondsPerMinute_: 60000,
	millisecondsPerHour_:   3600000,
	millisecondsPerDay_:    86400000,
	millisecondsPerWeek_:   604800000,

	// These are locked to the Earth's yearly orbit around the sun.
	millisecondsPerMonth_: 2629746000, // An average but exact value.
	millisecondsPerYear_:  31556952000,

	// Tying the two together is where things get messy.
	daysPerMonth_:  30.436875,
	daysPerYear_:   365.2425,
	weeksPerMonth_: 4.348125,
}
