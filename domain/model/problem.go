package model

// Problem ...
type Problem struct {
	ID         int
	Name       string
	Charset    string
	CanCompile bool
	Scores     []bool
}

// Problems ...
type Problems []Problem
