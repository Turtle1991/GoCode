package qsort

/*func QuickSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i] > in[j] {
				in[i], in[j] = in[j], in[i]
			}
		} // end for j = ...
	} // end for i = ...
}*/

func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= temp && i <= p {
			i++
		}

		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
