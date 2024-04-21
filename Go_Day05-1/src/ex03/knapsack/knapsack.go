package knapsack

type Present struct {
	Value, Size int
}

type PresentCell struct {
	MaxVal   int
	Presents []Present
}

func GrabPresents(arr []Present, capacity int) []Present {
	if capacity < 1 {
		return []Present{}
	}
	table := make([][]PresentCell, len(arr)+1)
	for i := range table {
		table[i] = make([]PresentCell, capacity+1)
	}

	for row_number := 1; row_number < len(arr)+1; row_number++ {
		pr := arr[row_number-1]
		for cap := 1; cap < capacity+1; cap++ {
			if pr.Size <= cap {
				cur_val := pr.Value + table[row_number-1][cap-pr.Size].MaxVal

				prev_val := table[row_number-1][cap].MaxVal

				if cur_val > prev_val {
					table[row_number][cap].MaxVal = cur_val
					table[row_number][cap].Presents = make([]Present, len(table[row_number-1][cap-pr.Size].Presents))
					copy(table[row_number][cap].Presents, table[row_number-1][cap-pr.Size].Presents)
					table[row_number][cap].Presents = append(table[row_number][cap].Presents, pr)
				} else {
					table[row_number][cap] = table[row_number-1][cap]
				}
			} else {
				table[row_number][cap] = table[row_number-1][cap]
			}
		}
	}

	return table[len(table)-1][capacity-1].Presents
}
