package basic

// Bitmap
// 位图实现
type Bitmap struct {
	data []int64
}

func NewBitmap(maxValue int64) *Bitmap {
	return &Bitmap{data: make([]int64, maxValue/(1<<6))}
}

func (b *Bitmap) Add(value int64) {
	var flag int64 = 1
	b.data[value/(1<<6)] |= flag << (value & 63)
}

func (b *Bitmap) Pop(value int64) {
	var flag int64 = 1
	b.data[value/(1<<6)] &= ^(flag << (value & 63))
}

func (b *Bitmap) Contains(value int64) bool {
	var flag int64 = 1
	return b.data[value/(1<<6)]&(flag<<(value&63)) != 0
}
