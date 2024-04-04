package concurrency

// select statements

func BubbleSort(slc []int) []int {

	for i := 0; i < len(slc); i++ {
		for j := i + 1; j < len(slc); j++ {
			if slc[i] > slc[j] {
				slc[j], slc[i] = slc[i], slc[j]
			}

		}

	}
	return slc
}

func InnerLoop(i int, slc []int) {
	for j := i + 1; j < len(slc); j++ {
		if slc[i] > slc[j] {
			slc[j], slc[i] = slc[i], slc[j]
		}
	}
}
func OuterLoop(slc []int, ch chan int, closech chan struct{}) {

	for i := 0; i < len(slc); i++ {
		ch <- i
		//slc = innerLoop(i, slc)
	}

	close(closech)

}
