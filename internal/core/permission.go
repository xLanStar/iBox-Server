package core

type Permission uint8

const (
	NO_PERMISSION Permission = 0
	READ_ONLY     Permission = 1
	ALL           Permission = 7
)
