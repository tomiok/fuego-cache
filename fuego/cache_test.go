package cache

import (
	"testing"
	"time"
)

func defaultConfigs() FuegoConfig {
	return FuegoConfig{
		DiskPersistence: false,
		FileLocation:    "",
		WebPort:         "9919",
		Mode:            "http",
	}
}

func Test_cacheConstructor(t *testing.T) {
	c := NewCache(defaultConfigs())
	if c == nil {
		t.Fail()
	}
}

func Test_SetAndGetOne(t *testing.T) {
	fuegoCache := NewCache(defaultConfigs())

	res, err := fuegoCache.SetOne(1, "1")

	if err != nil {
		t.Fail()
	}

	if res != "ok" {
		t.Fatal("cannot add")
	}

	value, err := fuegoCache.GetOne(1)

	if err != nil {
		t.Fail()
	}

	if value != "1" {
		t.Fatalf("cannot read %s", value)
	}
}

func Test_DeleteOne(t *testing.T) {
	fuegoCache := NewCache(defaultConfigs())

	res, err := fuegoCache.SetOne(1, "hell there")

	if res != "ok" {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}

	count := fuegoCache.Count()

	if count != 1 {
		t.Fail()
	}

	res = fuegoCache.DeleteOne(2) //should be nil the response since 2 is not a key inserted

	if res != "nil" {
		t.Fail()
	}

	res = fuegoCache.DeleteOne(1)

	if res != "ok" {
		t.Fail()
	}

	count = fuegoCache.Count()

	if count != 0 {
		t.Fail()
	}
}

func Test_expiredEntry(t *testing.T) {
	ttlInSeconds := 3
	fuegoCache := NewCache(defaultConfigs())

	ok, err := fuegoCache.SetOne(1, "hello there", ttlInSeconds)

	if err != nil {
		t.Fatalf("test failed %s", err.Error())
	}

	if ok != "ok" {
		t.Fatalf("error setting value %s", ok)
	}

	val, err := fuegoCache.GetOne(1)

	if ok != "ok" {
		t.Fatalf("error setting value %s", ok)
	}

	if val != "hello there" {
		t.Fatal("entry shold be stored")
	}

	time.Sleep(time.Second * 4) // 4 seconds

	ok, err = fuegoCache.GetOne(1)

	if err == nil {
		t.Fail()
	}

	if err != nil {
		if err.Error() != "key expired" {
			t.Fatalf("key should be expired")
		}
	}
}
