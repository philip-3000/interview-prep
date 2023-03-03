package maps

import (
	"testing"
	"go.uber.org/goleak"
)

func TestConstruction(t *testing.T) {

	var pm = NewPhilipMap[int, int](10)

	if len(pm.storage) != 10 {
		t.Fatalf("philipmap not instantiated properly")
	}
}

func TestItemsIteration(t *testing.T) {
	var pm = NewPhilipMap[int, int](10)
	
	// test for goroutine leaks
	defer goleak.VerifyNone(t)

	for i := 1; i <= 50; i++ {
		pm.Put(i, i)
	}

	var builtInMap = map[int]int{}
	for kvp := range pm.Items() {
		builtInMap[kvp.Key] = kvp.Value
	}

	for i := 1; i <= 50; i++ {
		_, ok := builtInMap[i]
		if !ok {
			t.Fatalf("key was not in dictionary: %v", i)
		}
	}

	// what if we don't read all...shouldn't get a leak
	var ctr = 0
	for kvp := range pm.Items() {
		ctr += 1
		kvp.Value = 0
		if ctr == 10 {
			break
		}
	}


}

func TestDelete(t *testing.T) {
	var pm = NewPhilipMap[int, int](2)

	for i := 0; i < 10; i += 1 {
		pm.Put(i, i)
	}

	var previousLength = pm.Length()

	pm.Delete(8)

	ok, _ := pm.Get(8)

	if ok {
		t.Fatalf("key '8' should have been deleted.")
	}

	if previousLength != (pm.Length() + 1) {
		t.Fatalf("Length not correct! Expected: %v, Actual: %v", previousLength-1, pm.Length())
	}

}

func TestPutGet(t *testing.T) {

	var pm = NewPhilipMap[string, string](2)
	//t.Logf("Calling Put for key and value")
	pm.Put("key", "value")
	if pm.Length() != 1 {
		t.Fatalf("Put() length did not increase by 1")
	}

	pm.Put("key2", "value2")
	if pm.Length() != 2 {
		t.Fatalf("Put() length did not increase by 1")
	}

	pm.Put("key", "updated value")
	if pm.Length() != 2 {
		t.Fatalf("Put() length incorrect")
	}

	// try to pull them back
	var ok, val = pm.Get("key")
	if !ok {
		t.Fatalf("Get() couldn't find value for key 'key'\n")
	}
	if val != "updated value" {
		t.Fatalf("Get() retrieved incorrect value: %v", val)
	}

}
