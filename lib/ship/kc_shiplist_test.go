package ship

import (
	"testing"
)

func TestShipGetImage(t *testing.T) {
	ship := KCShip{
		ID:   541,
		Name: "test",
	}

	if err := ship.GetImage(""); err != nil {
		t.Error(err)
	}
}

func TestLoadInfo(t *testing.T) {
	var kc KC

	if err := kc.LoadInfo("./testdata"); err != nil {
		t.Fatal(err)
	} else {
		t.Log(kc.Data.KCShips)
	}
}

func TestKCGetImage(t *testing.T) {
	var kc KC

	if err := kc.LoadInfo("./testdata"); err != nil {
		t.Fatal(err)
	} else {
		kc.GetImage("./test")
	}
}
