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
	tim "time"
)

// CLASS INTERFACE

// Access Function

func MomentClass() MomentClassLike {
	return momentClass()
}

// Constructor Methods

func (c *momentClass_) Moment(
	milliseconds int,
) MomentLike {
	return moment_(milliseconds)
}

func (c *momentClass_) MomentFromString(
	string_ string,
) MomentLike {
	var matches = c.matcher_.FindStringSubmatch(string_)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the moment constructor method: %s",
			string_,
		)
		panic(message)
	}
	return moment_(c.momentFromMatches(matches))
}

// Constant Methods

func (c *momentClass_) Minimum() MomentLike {
	return c.minimum_
}

func (c *momentClass_) Maximum() MomentLike {
	return c.maximum_
}

func (c *momentClass_) Epoch() MomentLike {
	return c.epoch_
}

// Function Methods

func (c *momentClass_) Now() MomentLike {
	var now = tim.Now().UTC().UnixMilli()
	return moment_(now)
}

func (c *momentClass_) Earlier(
	moment MomentLike,
	duration DurationLike,
) MomentLike {
	return moment_(moment.AsInteger() - duration.AsInteger())
}

func (c *momentClass_) Later(
	moment MomentLike,
	duration DurationLike,
) MomentLike {
	return moment_(moment.AsInteger() + duration.AsInteger())
}

func (c *momentClass_) Duration(
	first MomentLike,
	second MomentLike,
) DurationLike {
	return duration_(second.AsInteger() - first.AsInteger())
}

// INSTANCE INTERFACE

// Principal Methods

func (v moment_) GetClass() MomentClassLike {
	return momentClass()
}

func (v moment_) GetIntrinsic() int {
	return int(v)
}

// Attribute Methods

// Discrete Methods

func (v moment_) AsBoolean() bool {
	// There is no beginning of time (EPOCH is arbitrary).
	return true
}

func (v moment_) AsInteger() int {
	return int(v)
}

// Lexical Methods

func (v moment_) AsString() string {
	var builder sts.Builder
	var year = v.GetYears()
	var month = v.GetMonths()
	var day = v.GetDays()
	var hour = v.GetHours()
	var minute = v.GetMinutes()
	var second = v.GetSeconds()
	var millisecond = v.GetMilliseconds()
	builder.WriteString("<")
	builder.WriteString(stc.FormatInt(int64(year), 10))
	if month > 1 || day > 1 || hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
		builder.WriteString("-")
		builder.WriteString(momentClass().formatOrdinal(month, 2))
		if day > 1 || hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
			builder.WriteString("-")
			builder.WriteString(momentClass().formatOrdinal(day, 2))
			if hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
				builder.WriteString("T")
				builder.WriteString(momentClass().formatOrdinal(hour, 2))
				if minute > 0 || second > 0 || millisecond > 0 {
					builder.WriteString(":")
					builder.WriteString(momentClass().formatOrdinal(minute, 2))
					if second > 0 || millisecond > 0 {
						builder.WriteString(":")
						builder.WriteString(momentClass().formatOrdinal(second, 2))
						if millisecond > 0 {
							builder.WriteString(".")
							builder.WriteString(momentClass().formatOrdinal(millisecond, 3))
						}
					}
				}
			}
		}
	}
	builder.WriteString(">")
	return builder.String()
}

// Temporal Methods

func (v moment_) AsMilliseconds() float64 {
	return float64(v)
}

func (v moment_) AsSeconds() float64 {
	return float64(v) / float64(durationClass().millisecondsPerSecond_)
}

func (v moment_) AsMinutes() float64 {
	return float64(v) / float64(durationClass().millisecondsPerMinute_)
}

func (v moment_) AsHours() float64 {
	return float64(v) / float64(durationClass().millisecondsPerHour_)
}

func (v moment_) AsDays() float64 {
	return float64(v) / float64(durationClass().millisecondsPerDay_)
}

func (v moment_) AsWeeks() float64 {
	return float64(v) / float64(durationClass().millisecondsPerWeek_)
}

func (v moment_) AsMonths() float64 {
	return float64(v) / float64(durationClass().millisecondsPerMonth_)
}

func (v moment_) AsYears() float64 {
	return float64(v) / float64(durationClass().millisecondsPerYear_)
}

// Factored Methods

func (v moment_) GetMilliseconds() int {
	var time = v.asTime()
	var milliseconds = time.Nanosecond() / 1e6
	return int(milliseconds)
}

func (v moment_) GetSeconds() int {
	var time = v.asTime()
	var seconds = time.Second()
	return int(seconds)
}

func (v moment_) GetMinutes() int {
	var time = v.asTime()
	var minutes = time.Minute()
	return int(minutes)
}

func (v moment_) GetHours() int {
	var time = v.asTime()
	var hours = time.Hour()
	return int(hours)
}

func (v moment_) GetDays() int {
	var time = v.asTime()
	var days = time.Day()
	return int(days)
}

func (v moment_) GetWeeks() int {
	var time = v.asTime()
	var _, weeks = time.ISOWeek()
	return int(weeks)
}

func (v moment_) GetMonths() int {
	var time = v.asTime()
	var months = time.Month()
	return int(months)
}

func (v moment_) GetYears() int {
	var time = v.asTime()
	var years = time.Year()
	return int(years)
}

// PROTECTED INTERFACE

func (v moment_) String() string {
	return v.AsString()
}

// Private Methods

func (c *momentClass_) formatOrdinal(ordinal int, digits int) string {
	return fmt.Sprintf("%0"+stc.Itoa(digits)+"d", ordinal)
}

// This list contains the supported ISO 8601 date-time formats delimited by
// angle brackets. Note: the Go templates in this list must contain their exact
// numeric values. If you are curious why this is, check out this posting:
//   - https://medium.com/@simplyianm/how-go-solves-date-and-time-formatting-8a932117c41c
//
// A good mnemonic to use to remember the pattern is:
//
//	  1    2    3     4      5     6      7
//	month day hour minute second year timezone
//
//	"January 2nd, 2006 at 3:04:05pm in the MST timezone"
//
// Anyway, it is what it is, but this hides it from the rest of the code.
var hackedIsoFormats_ = [...]string{
	"<2006-01-02T15:04:05.000>",
	"<2006-01-02T15:04:05>",
	"<2006-01-02T15:04>",
	"<2006-01-02T15>",
	"<2006-01-02>",
	"<2006-01>",
	"<2006>",
	"<-2006>",
	"<-2006-01>",
	"<-2006-01-02>",
	"<-2006-01-02T15>",
	"<-2006-01-02T15:04>",
	"<-2006-01-02T15:04:05>",
	"<-2006-01-02T15:04:05.000>",
}

// The Go time.Parse() function can only handle years in the range [0000..9999]
// even though the time.Time.Format() method will correctly print any size year
// positive or negative. The Go team has labeled this issue as "unfortunate" and
// will not fix it.
//
// To support the entire ISO 8601 calendar (and beyond) as shown here:
//
//	https://en.wikipedia.org/wiki/Holocene_calendar#Conversion
//
// So we must resort to some hacking with this private function...
func (c *momentClass_) momentFromMatches(matches []string) int {
	// First, we replace the year with year zero.
	var yearString = matches[2]
	var patched = sts.Replace(matches[0], yearString, "0000", 1)

	// Next, we attempt to parse the patched moment using our Go based formats.
	for _, format := range hackedIsoFormats_ {
		var date, err = tim.Parse(format, patched) // Parsed in UTC.
		if err == nil {

			// We found a match, now we add back in the correct year.
			var year, _ = stc.ParseInt(yearString, 10, 64)
			date = date.AddDate(int(year), 0, 0)
			if sts.HasPrefix(format, "<-") {

				// We change the positive date to a negative one.
				date = date.AddDate(-2*date.Year(), 0, 0)
			}

			// And return the correct date as milliseconds.
			var milliseconds = int(date.UnixMilli())
			return milliseconds
		}
	}

	// This panic will only happen if the scanner regular expressions are out
	// of sync with the ISO 8601 standard formats. The moment has already been
	// succussfully scanned by the the scanner.
	panic(fmt.Sprintf("The moment does not match a known format: %v", matches[0]))
}

func (v moment_) asTime() tim.Time {
	return tim.UnixMilli(int64(v)).UTC()
}

// NOTE:
// These private constants are used to define the private regular expression
// matcher that is used to match legal string patterns for this intrinsic type.
// Unfortunately there is no way to make them private to this class since they
// must be TRUE Go constants to be used in this way.  We append an underscore to
// each name to lessen the chance of a name collision with other private Go
// class constants in this package.
const (
	day_    = "(?:([0-2][1-9])|(3[0-1]))"
	hour_   = "(?:([0-1][0-9])|(2[0-3]))"
	minute_ = "(?:[0-5][0-9])"
	month_  = "(?:(0[1-9])|(1[0-2]))"
	second_ = "(?:([0-5][0-9])|(6[0-1]))"
	year_   = "(?:0|(?:" + ordinal_ + "))"
)

// Instance Structure

type moment_ int

// Class Structure

type momentClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
	minimum_ MomentLike
	maximum_ MomentLike
	epoch_   MomentLike
}

// Class Reference

func momentClass() *momentClass_ {
	return momentClassReference_
}

var momentClassReference_ = &momentClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^(?:<(?:" + sign_ + ")?(?:" + year_ + ")(-(?:" + month_ + ")(-(?:" +
			day_ + ")(T(?:" + hour_ + ")(:(?:" + minute_ + ")(:(?:" + second_ +
			")(?:" + fraction_ + ")?)?)?)?)?)?>)",
	),
	minimum_: moment_(mat.MinInt),
	maximum_: moment_(mat.MaxInt),
	epoch_:   moment_(0),
}
