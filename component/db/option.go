package db

type Callbacker interface{}

type HandlerFunc func(Callbacker)
