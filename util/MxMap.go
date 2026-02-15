package util

import "sync"

/*
 * MxMap: Thread-safe Hash Map
 *
 * Copyright (C) 2023 Louis Trevino, Torino Consulting, Ltd.
 */

type MxMap[K, V comparable] struct {
	Mx  *sync.Mutex
	Map map[K]V
}

func NewMutexMap[K, V comparable]() MxMap[K, V] {
	mutexMap := MxMap[K, V]{
		Mx:  &sync.Mutex{},
		Map: make(map[K]V),
	}
	return mutexMap
}

func (mxMap *MxMap[K, V]) Get(key K) (V, bool) {
	mxMap.Mx.Lock()
	defer mxMap.Mx.Unlock()
	val, pr := mxMap.Map[key]
	return val, pr
}

func (mxMap *MxMap[K, V]) Put(key K, val V) {
	mxMap.Mx.Lock()
	defer mxMap.Mx.Unlock()
	mxMap.Map[key] = val
}
