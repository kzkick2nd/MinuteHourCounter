package conveyor

type QueneBuckets struct {
	maxItems int
	quene    []int
	totalSum int
}

func NewBuckets(maxItems int) *QueneBuckets {
	return &QueneBuckets{maxItems: maxItems}
}

func (q *QueneBuckets) AddToBack(count int) {
	q.totalSum += count
}

func (q *QueneBuckets) Shift(shifted int) {}

func (q *QueneBuckets) TotalSum() int {
	return q.totalSum
}
