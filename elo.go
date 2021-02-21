package elo

import "math"

// MatchResult represents the result of a match (e.g. win/loss).
type MatchResult int

// Possible match results.
const (
	LeftWin MatchResult = iota
	RightWin
	Draw
)

// Option represents an option that can be used to tweak the behaviour of the
// Elo calculation.
type Option interface {
	isOption()
}

type kFactor float64

// KFactor returns an Option that adjusts the K-factor used in the Elo
// calculation.
func KFactor(k float64) Option { return kFactor(k) }

func (_ kFactor) isOption() {}

type deviation float64

// Deviation returns an Option that adjusts the Deviation used in the Elo
// calculation.
func Deviation(d float64) Option { return deviation(d) }

func (_ deviation) isOption() {}

// Calculate takes a left rating, a right rating, the result of the match, and
// some optional options. It returns the new left and right rating. K-factor
// and Deviation default to 32 and 400 respectively unless overridden by
// options.
func Calculate(lRating, rRating float64, matchResult MatchResult, options ...Option) (float64, float64) {
	k := 32.0
	d := 400.0
	for _, opt := range options {
		switch v := opt.(type) {
		case kFactor:
			k = float64(v)
		case deviation:
			d = float64(v)
		}
	}

	lExpected := (1.0 / (1.0 + math.Pow(10.0, ((rRating-lRating)/d))))
	rExpected := (1.0 / (1.0 + math.Pow(10.0, ((lRating-rRating)/d))))

	lActual := 0.5
	rActual := 0.5
	switch matchResult {
	case LeftWin:
		lActual = 1.0
		rActual = 0.0
	case RightWin:
		lActual = 0.0
		rActual = 1.0
	}

	return lRating + (k * (lActual - lExpected)),
		rRating + (k * (rActual - rExpected))
}
