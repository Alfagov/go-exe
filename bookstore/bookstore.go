package bookstore

// Cost returns the cost of the book basket
// given the number of books in each basket.
func Cost(chkOut []int) int {

	// If no books, no cost.
	if len(chkOut) == 0 {
		return 0
	}
	// If only one book, no discount.
	if len(chkOut) == 1 {
		return 800
	}

	// Create map and slice to track number of books in each basket.
	cOutMap := map[int]int{}
	cOutSlice := make([]int, 0)

	for _, value := range chkOut {

		// Increase the number of books in the basket.
		cOutMap[value]++

		// If the number of books in the basket is greater than the number of
		// books in the slice then add a new basket to the slice initialized at 0.
		if len(cOutSlice) < cOutMap[value] {
			cOutSlice = append(cOutSlice, 0)
		}

		// If current basket is full, add book to next basket.
		if cOutSlice[cOutMap[value]-1] == 4 && len(cOutSlice) > cOutMap[value] {

			var isBookAdded bool
			for i := cOutMap[value]; i < len(cOutSlice); i++ {
				if cOutSlice[i] < 4 {
					cOutSlice[i]++
					isBookAdded = true
					break
				}
			}
			if !isBookAdded {
				cOutSlice[cOutMap[value]-1]++
			}
			continue
		}
		cOutSlice[cOutMap[value]-1]++
	}

	return calculateCost(cOutSlice)
}

func calculateCost(pSlice []int) int {
	cost := 0
	for _, value := range pSlice {
		switch value {
		case 1:
			cost += 1 * 800
		case 2:
			//2 books with 5% discount
			cost += 2 * 760
		case 3:
			//3 books with 10% discount
			cost += 3 * 720
		case 4:
			// 4 books with 20% discount
			cost += 4 * 640
		default:
			//Any other amount with 25% discount
			cost += value * 600
		}
	}

	return cost
}
