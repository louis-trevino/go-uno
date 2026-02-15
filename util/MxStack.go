package util

import "sync"

/*
 * MxStack: Thread-safe Stack
 *
 * Copyright (C) 2023 Louis Trevino, Torino Consulting, Ltd.
 */

type MxStack[T any] struct {
	Stack []T
	Mx    *sync.Mutex
}

func NewMxStack[T any]() *MxStack[T] {
	stack := make([]T, 0)
	return &MxStack[T]{Stack: stack, Mx: &sync.Mutex{}}
}

func (ms *MxStack[T]) Push(val T) {
	ms.Mx.Lock()
	defer ms.Mx.Unlock()
	ms.Stack = append(ms.Stack, val)
}

func (ms *MxStack[T]) Peek() (val T, pr bool) {
	last := len(ms.Stack) - 1
	if last < 0 {
		var nada T
		return nada, false
	}
	return ms.Stack[last], true
}

func (ms *MxStack[T]) Pop() (val T, pr bool) {
	ms.Mx.Lock()
	defer ms.Mx.Unlock()
	last := len(ms.Stack) - 1
	if last < 0 {
		var nada T
		return nada, false
	}
	elem := ms.Stack[last]
	ms.Stack = ms.Stack[:last]
	return elem, true
}
