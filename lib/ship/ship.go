package ship

import ()

type ShipList interface {
	LoadInfo(string) error
	GetImage(string) error
}
