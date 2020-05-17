package sabermetrics

import (
	"testing"
)

func TestLeverageIndex(t *testing.T) {
	tests := map[string]struct {
		baseState   BaseState
		score       Score
		halfInning  HalfInning
		outs        int
		expectedRes float32
		expectedErr error
	}{
		"successful top of the first": {
			baseState: BaseState{
				First:  true,
				Second: true,
				Third:  true,
			},
			score: Score{
				Away: 4,
				Home: 0,
			},
			halfInning: HalfInning{
				Inning:      1,
				TopOfInning: true,
			},
			outs:        2,
			expectedRes: 1.4,
			expectedErr: nil,
		},
		"successful bottom of the first": {
			baseState: BaseState{
				First:  false,
				Second: false,
				Third:  false,
			},
			score: Score{
				Away: 2,
				Home: 2,
			},
			halfInning: HalfInning{
				Inning:      1,
				TopOfInning: false,
			},
			outs:        2,
			expectedRes: 0.4,
			expectedErr: nil,
		},
		"successful bottom of the fifth": {
			baseState: BaseState{
				First:  true,
				Second: false,
				Third:  true,
			},
			score: Score{
				Away: 2,
				Home: 3,
			},
			halfInning: HalfInning{
				Inning:      5,
				TopOfInning: false,
			},
			outs:        0,
			expectedRes: 1.2,
			expectedErr: nil,
		},
		"successful top of the seventh": {
			baseState: BaseState{
				First:  true,
				Second: false,
				Third:  false,
			},
			score: Score{
				Away: 11,
				Home: 13,
			},
			halfInning: HalfInning{
				Inning:      7,
				TopOfInning: true,
			},
			outs:        1,
			expectedRes: 1.8,
			expectedErr: nil,
		},
		"successful bottom of the ninth": {
			baseState: BaseState{
				First:  false,
				Second: false,
				Third:  true,
			},
			score: Score{
				Away: 4,
				Home: 2,
			},
			halfInning: HalfInning{
				Inning:      9,
				TopOfInning: false,
			},
			outs:        1,
			expectedRes: 2.9,
			expectedErr: nil,
		},
		"invalid inning": {
			halfInning: HalfInning{
				Inning: 0,
			},
			outs:        2,
			expectedRes: 0.0,
			expectedErr: ErrInvalidInning,
		},
	}

	for name, tt := range tests {
		res, err := LeverageIndex(tt.baseState, tt.score, tt.halfInning, tt.outs)
		if err != tt.expectedErr {
			t.Errorf("%s: expected error: %s, received: %s", name, tt.expectedErr, err)
		}
		if res != tt.expectedRes {
			t.Errorf("%s: expected result: %v, received: %v", name, tt.expectedRes, res)
		}
	}
}
