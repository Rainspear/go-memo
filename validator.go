package main

type Enum interface {
	IsValid() bool
}

func (l Level) IsValid() bool {
	switch l {
	case LevelEssential, LevelImportant, LevelCritical, LevelMajor, LevelMinor:
		return true
	default:
		return false
	}
}

func (s Status) IsValid() bool {
	switch s {
	case StatusSuccess, StatusFailure, StatusSkipped, StatusUntouch:
		return true
	default:
		return false
	}
}
