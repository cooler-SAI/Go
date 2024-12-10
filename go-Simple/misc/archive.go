package misc

func ChangeValue(point *int) {
	*point = 500

}

func ChangeArray(array *[5]int, value int) {
	for i := 0; i < len(*array); i++ {
		(*array)[i] = value
	}

}

func ChangeArrayMulti(array *[5]int, value int) {
	changedValue := 1000
	for i := 0; i < len(*array); i++ {
		(*array)[i] = value + changedValue
	}

}
