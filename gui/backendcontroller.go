package main

type BackendController[T any, O any] interface {
	Run(options ...O) error
	Cancel() bool
	GetData() T
	LockData()
	UnlockData()
	SetState(key string, value string)
	GetState(key string) string
	SetSetting(key string, value string)
	GetSetting(key string) string
}
