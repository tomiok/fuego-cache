package cache

import (
	"strconv"
	"testing"
)

var _cache = NewCache(FuegoConfig{
	DiskPersistence: false,
})

func Test_bulkGet(t *testing.T) {
	populate(_cache)
	keys := []interface{}{"test1", "test2"}
	r := _cache.BulkGet(keys)

	for _, value := range r {
		if !(value.Value == "1" || value.Value == "2") {
			t.Fail()
		}
	}

	tearDown(_cache)
}

func Test_bulkSet(t *testing.T) {
	_cache.BulkSet([]entry{{
		key: 1,
		object: fuegoValue{
			value: "1",
		},
	}, {
		key: 2,
		object: fuegoValue{
			value: "2",
		},
	}, {
		key: 3,
		object: fuegoValue{
			value: "3",
		},
	}})
	count := _cache.Count()
	if count != 3 {
		t.Error("the count is " + strconv.Itoa(count))
		t.Fail()
	}
	tearDown(_cache)
}

func Test_bulkDelete(t *testing.T) {
	populate(_cache)
	count := _cache.Count()
	if count != 4 {
		t.Fail()
	}
	keys := []interface{}{"test1", "test2"}
	_cache.BulkDelete(keys)

	count = _cache.Count()
	if count != 2 {
		t.Fail()
	}
	tearDown(_cache)
}

func populate(c *cache) {
	_, _ = c.SetOne("test1", "1")
	_, _ = c.SetOne("test2", "2")
	_, _ = c.SetOne("test3", "3")
	_, _ = c.SetOne("test4", "4")
}

func tearDown(c *cache) {
	c.Clear()
}
