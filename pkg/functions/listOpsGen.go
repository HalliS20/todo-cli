package functions

func Reduce[T, U any](lis []T, f func(T, U) U) U {
	switch len(lis) {
	case 0:
		panic("Too few elements to reduce")
	case 1:
		var zero U
		return f(lis[0], zero)
	default:
		return f(lis[0], Reduce(lis[1:], f))
	}
}

func Map[T, U any](lis []T, f func(T) U) []U {
	switch len(lis) {
	case 0:
		return []U{}
	case 1:
		return []U{f(lis[0])}
	default:
		return append([]U{f(lis[0])}, Map(lis[1:], f)...)
	}
}

func MapVoidFunc[T any](lis []T, f func(T)) []T {
	for _, item := range lis {
		f(item)
	}
	return lis
}

func MapVoid[T any](lis []T, f func(T)) {
	for _, item := range lis {
		f(item)
	}
}

func Zip[T, U, V any](list1 []T, list2 []U, f func(T, U) V) []V {
	if len(list1) != len(list2) {
		panic("list lengths don't match")
	}
	switch len(list2) {
	case 1:
		return []V{f(list1[0], list2[0])}
	default:
		return append([]V{f(list1[0], list2[0])}, Zip(list1[1:], list2[1:], f)...)
	}
}

func LongerList[T any](list1 []T, list2 []T) []T {
	if len(list1) > len(list2) {
		return list1
	}
	return list2
}

func LongerListFirst[T any](list1 *[]T, list2 *[]T) {
	temp := *list1
	*list1 = *list2
	*list2 = temp
}

func LPtoPL[T any](items *[]T) []*T {
	var newList []*T
	for i := range *items {
		newList = append(newList, &(*items)[i])
	}
	return newList
}

func InsertAtIndex[Type any](lisa *[]Type, index int, value Type) {
	*lisa = append(*lisa, value)
	if index >= len(*lisa) || index < 0 || len(*lisa) <= 1 {
		return
	}
	for i := len(*lisa) - 1; i > index; i-- {
		(*lisa)[i], (*lisa)[i-1] = (*lisa)[i-1], (*lisa)[i]
	}
}
