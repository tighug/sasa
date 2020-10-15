package model

// Problem ...
type Problem struct {
	ID         int
	Name       string
	CanCompile bool
	Charset    string
	Scores     []bool
}
