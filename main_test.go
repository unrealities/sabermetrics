package sabermetrics

import "testing"

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
			expectedRes: LeverageIndices[23][0],
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
			expectedRes: LeverageIndices[16][9],
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
			expectedRes: LeverageIndices[5][82],
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
			expectedRes: LeverageIndices[9][110],
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
			expectedRes: LeverageIndices[11][151],
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
		"invalid outs": {
			halfInning: HalfInning{
				Inning: 1,
			},
			outs:        7,
			expectedRes: 0.0,
			expectedErr: ErrInvalidOuts,
		},
		"invalid: end of the game (bottom of ninth)": {
			halfInning: HalfInning{
				Inning: 9,
			},
			score: Score{
				Away: 0,
				Home: 1,
			},
			outs:        3,
			expectedRes: 0.0,
			expectedErr: ErrGameOver,
		},
		"invalid: end of the game (top of ninth)": {
			halfInning: HalfInning{
				Inning:      9,
				TopOfInning: true,
			},
			score: Score{
				Away: 0,
				Home: 1,
			},
			outs:        3,
			expectedRes: 0.0,
			expectedErr: ErrGameOver,
		},
		"advance inning if outs == 3": {
			halfInning: HalfInning{
				Inning: 1,
			},
			outs:        3,
			expectedRes: LeverageIndices[0][18], // Top of the second. No runnes. No score. No outs.
			expectedErr: nil,
		},
		"if inning > 9, set to 9": {
			halfInning: HalfInning{
				Inning: 19,
			},
			outs:        0,
			expectedRes: LeverageIndices[0][153], // Bottom of the ninth. No runners. No score. No outs.
			expectedErr: nil,
		},
		"if run differential is greater than 4, use 4": {
			halfInning: HalfInning{
				Inning: 1,
			},
			score: Score{
				Away: 0,
				Home: 15,
			},
			outs:        0,
			expectedRes: LeverageIndices[0][13], // Bottom of the first. No Runners. RunDiff = 4. No outs.
			expectedErr: nil,
		},
		"if run differential is less than -4, use -4": {
			halfInning: HalfInning{
				Inning: 1,
			},
			score: Score{
				Away: 11,
				Home: 4,
			},
			outs:        0,
			expectedRes: LeverageIndices[0][5], // Bottom of the first. No Runners. RunDiff = -4. No outs.
			expectedErr: nil,
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
