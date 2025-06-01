/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element_test

import (
	ele "github.com/craterdog/go-primitive-framework/v7/element"
	ass "github.com/stretchr/testify/assert"
	mat "math"
	tes "testing"
)

var AngleClass = ele.AngleClass()

func TestZeroAngles(t *tes.T) {
	var v = AngleClass.Angle(0)
	ass.Equal(t, 0.0, v.AsFloat())

	v = AngleClass.AngleFromString("~0")
	ass.Equal(t, "~0", v.AsString())
}

func TestPositiveAngles(t *tes.T) {
	var v = AngleClass.Angle(mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = AngleClass.AngleFromString("~π")
	ass.Equal(t, "~π", v.AsString())
}

func TestNegativeAngles(t *tes.T) {
	var v = AngleClass.Angle(-mat.Pi)
	ass.Equal(t, mat.Pi, v.AsFloat())

	v = AngleClass.Angle(-mat.Pi / 2.0)
	ass.Equal(t, 1.5*mat.Pi, v.AsFloat())
}

func TestAnglesLibrary(t *tes.T) {
	var v0 = AngleClass.Zero()
	var v1 = AngleClass.Angle(mat.Pi * 0.25)
	var v2 = AngleClass.Angle(mat.Pi * 0.5)
	var v3 = AngleClass.Angle(mat.Pi * 0.75)
	var v4 = AngleClass.Pi()
	var v5 = AngleClass.Angle(mat.Pi * 1.25)
	var v6 = AngleClass.Angle(mat.Pi * 1.5)
	var v7 = AngleClass.Angle(mat.Pi * 1.75)
	var v8 = AngleClass.Tau()

	ass.Equal(t, v4, AngleClass.Inverse(v0))
	ass.Equal(t, v5, AngleClass.Inverse(v1))
	ass.Equal(t, v6, AngleClass.Inverse(v2))
	ass.Equal(t, v7, AngleClass.Inverse(v3))
	ass.Equal(t, v0, AngleClass.Inverse(v4))
	ass.Equal(t, v4, AngleClass.Inverse(v8))

	ass.Equal(t, v1, AngleClass.Sum(v0, v1))
	ass.Equal(t, v0, AngleClass.Difference(v1, v1))
	ass.Equal(t, v3, AngleClass.Sum(v1, v2))
	ass.Equal(t, v1, AngleClass.Difference(v3, v2))
	ass.Equal(t, v5, AngleClass.Sum(v2, v3))
	ass.Equal(t, v2, AngleClass.Difference(v5, v3))
	ass.Equal(t, v7, AngleClass.Sum(v3, v4))
	ass.Equal(t, v3, AngleClass.Difference(v7, v4))
	ass.Equal(t, v1, AngleClass.Sum(v8, v1))
	ass.Equal(t, v0, AngleClass.Difference(v8, v8))

	ass.Equal(t, v3, AngleClass.Scaled(v1, 3.0))
	ass.Equal(t, v0, AngleClass.Scaled(v4, 2.0))
	ass.Equal(t, v4, AngleClass.Scaled(v4, -1.0))
	ass.Equal(t, v0, AngleClass.Scaled(v8, 1.0))

	ass.Equal(t, v0, AngleClass.ArcCosine(AngleClass.Cosine(v0)))
	ass.Equal(t, v1, AngleClass.ArcCosine(AngleClass.Cosine(v1)))
	ass.Equal(t, v2, AngleClass.ArcCosine(AngleClass.Cosine(v2)))
	ass.Equal(t, v3, AngleClass.ArcCosine(AngleClass.Cosine(v3)))
	ass.Equal(t, v4, AngleClass.ArcCosine(AngleClass.Cosine(v4)))
	ass.Equal(t, v0, AngleClass.ArcCosine(AngleClass.Cosine(v8)))

	ass.Equal(t, v0, AngleClass.ArcSine(AngleClass.Sine(v0)))
	ass.Equal(t, v1, AngleClass.ArcSine(AngleClass.Sine(v1)))
	ass.Equal(t, v2, AngleClass.ArcSine(AngleClass.Sine(v2)))
	ass.Equal(t, v6, AngleClass.ArcSine(AngleClass.Sine(v6)))
	ass.Equal(t, v7, AngleClass.ArcSine(AngleClass.Sine(v7)))
	ass.Equal(t, v0, AngleClass.ArcSine(AngleClass.Sine(v8)))

	ass.Equal(t, v0, AngleClass.ArcTangent(AngleClass.Cosine(v0), AngleClass.Sine(v0)))
	ass.Equal(t, v1, AngleClass.ArcTangent(AngleClass.Cosine(v1), AngleClass.Sine(v1)))
	ass.Equal(t, v2, AngleClass.ArcTangent(AngleClass.Cosine(v2), AngleClass.Sine(v2)))
	ass.Equal(t, v3, AngleClass.ArcTangent(AngleClass.Cosine(v3), AngleClass.Sine(v3)))
	ass.Equal(t, v4, AngleClass.ArcTangent(AngleClass.Cosine(v4), AngleClass.Sine(v4)))
	ass.Equal(t, v5, AngleClass.ArcTangent(AngleClass.Cosine(v5), AngleClass.Sine(v5)))
	ass.Equal(t, v0, AngleClass.ArcTangent(AngleClass.Cosine(v8), AngleClass.Sine(v8)))
}
