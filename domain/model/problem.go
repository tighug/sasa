package model

// Problem ...
type Problem struct {
	ID         int
	Name       string
	Charset    string
	CanCompile bool
	Score      int
}

// Problems ...
type Problems []Problem
