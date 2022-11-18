package medium

type stockSpannerPoint struct {
	price       int
	growingDays int
}

type StockSpanner struct {
	points []*stockSpannerPoint
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (s *StockSpanner) Next(price int) int {
	days := 1
	for i := len(s.points) - 1; i >= 0; {
		if s.points[i].price > price {
			break
		}
		days += s.points[i].growingDays
		i = i - s.points[i].growingDays
	}
	s.points = append(s.points, &stockSpannerPoint{price: price, growingDays: days})
	return days
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
