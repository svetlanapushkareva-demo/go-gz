package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func newBin(private bool, name string) *Bin {
	createdAt := time.Now()
	id := time.Now().String()
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

type BinList []*Bin

func NewBinList() BinList {
	bl := make(BinList, 0)
	return bl
}
func (bl *BinList) Add(private bool, name string) {

	*bl = append(*bl, newBin(private, name))
}
