package main

type BackendController[T any, O any] interface {
	Run(options ...O) error
	Cancel() bool
	GetData() T
	LockData()
	UnlockData()
	SetSetting(key string, value string)
	GetSetting(key string) string
}
