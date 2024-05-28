package pokecache

import (
	"testing"
	"time"
)

const interval = 5 * time.Second

func TestCreateCache(t *testing.T) {
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(interval)

	cases := []struct {
		inputURL string
		inputVal []byte
	}{
		{
			inputURL: "pageURL1",
			inputVal: []byte("home"),
		},
		{
			inputURL: "pageURL2",
			inputVal: []byte("location-area"),
		},
		{
			inputURL: "",
			inputVal: []byte("gymz"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputURL, cas.inputVal)
		actual, ok := cache.Get(cas.inputURL)
		if !ok {
			t.Errorf("%s not found", cas.inputURL)
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s doesn't match %s",
				string(actual),
				string(cas.inputVal),
			)
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.Add("pageURL", []byte("testdata"))

	_, ok := cache.Get("pageURL")
	if !ok {
		t.Errorf("pageURL not found")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("pageURL")
	if ok {
		t.Errorf("expected to not find pageURL")
	}

}
