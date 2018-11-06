package conveyor

type QueneBuckets struct {
	maxItems int
	quene    []int
	totalSum int
}

func NewBuckets(maxItems int) *QueneBuckets {
	return &QueneBuckets{maxItems: maxItems}
}

func (q *QueneBuckets) Shift(shifted int) {
	if shifted > q.maxItems {
		q.quene = []int{}
		q.totalSum = 0
		return
	}

	for shifted > 0 {
		q.quene = append(q.quene, 0)
		shifted--
	}

	for len(q.quene) > q.maxItems {
		q.totalSum -= q.quene[0]
		q.quene = q.quene[1:]
	}
}

func (q *QueneBuckets) AddToBack(count int) {
	if len(q.quene) == 0 {
		q.Shift(1)
	}
	q.quene[len(q.quene)-1] += count
	q.totalSum += count
}

func (q *QueneBuckets) TotalSum() int {
	return q.totalSum
}
