package bubblesort

func BubbleSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		flag := true

		for j := 0; j < len(in)-1-i; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
				flag = false
			}
		} // end for j = ...

		if flag == true {
			break
		}

	} // end for i = ...
}
