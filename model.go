package sabermetrics

// BaseState represents the current offensive base state
// Each base is given a `bool` to tell if it is occupied or not
type BaseState struct {
	First  bool
	Second bool
	Third  bool
}
