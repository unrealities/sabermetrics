package sabermetrics

// LeverageIndex (LI)
// Created by Tom Tango, Leverage Index measures the importance of a particular event by quantifying
// the extent to which win probability could change on said event, with 1.0 representing a neutral situation.
//
// bs :         BaseState{first: bool, second: bool, third: bool} representating which bases are occupied
// topOfInning: `bool` representation for if the game is currently in the top of an inning
// outs :       `int` representation of number of outs in the game currently (0,1,2)
// inning:      `int` representation of the inning number (1,2,3,4,5,6,7,8,9)
//              If inning > 9, then the leverage index of 9 is returned
//
// Returns a float64 representing the current game's leverage index
// Error occurs on invalid data
func LeverageIndex(bs BaseState, topOfInning bool, inning, outs int) (float64, error) {
	return 0.0, nil
}
