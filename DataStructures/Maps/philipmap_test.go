package maps

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestConstruction(t *testing.T) {

	var pm = NewPhilipMap[int, int](10)

	if len(pm.storage) != 10 {
		t.Fatalf("philipmap not instantiated properly")
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
