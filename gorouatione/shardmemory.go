package main

import "fmt"

type shareMap struct {
	m map[string]interface{}
	c chan command

}

type command struct {
	action int
	key string
	value interface{}
	result chan<- interface{}
}

const (
	set = iota
	get
	remove
	cont
)

func (sm shardMap) Set(k string, v interface{}) {
	sm.c <- command{action: set, key: k, value: v}
}

func (sm shardMap) Get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.c <- command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

func (sm sharedMap) Remove(k string) {
	sm.c <- command{action: remove, key:k}
}

