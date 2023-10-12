package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cacheClient := NewCache(interval)
	if cacheClient.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cacheClient := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cacheClient.Add(cas.inputKey, cas.inputVal)
		actual, ok := cacheClient.Get(cas.inputKey)
		if !ok {
			t.Errorf("%v not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inputVal))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cacheClient := NewCache(interval)

	key1 := "key1"
	val1 := []byte("val1")
	cacheClient.Add(key1, val1)

	time.Sleep(interval + time.Millisecond)

	_, ok := cacheClient.Get(key1)
	if ok {
		t.Errorf("%v: should have reap %v", key1, ok)
	}
}
