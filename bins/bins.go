package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(private bool, name string) *Bin {
	createdAt := time.Now()
	id := time.Now().String()
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
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
