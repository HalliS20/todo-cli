package functions

type MapPart int

const (
	KEY MapPart = iota
	VALUE
)

func SetMapI[T comparable, U any](mapObj map[T]U, item T) {
	var zeroI U
	mapObj[item] = zeroI
}

func IncMapVal[T comparable](mapObj map[T]int, key T) {
	if _, exists := mapObj[key]; exists {
		mapObj[key] += 1
	}
}

func IncMapValByKey(mapObj map[int]int, key int) {
	if _, exists := mapObj[key]; exists {
		mapObj[key] += key
	}
}

func PopMapItem[T comparable, U any](mapObj map[T]U) (T, U) {
	var k T
	var v U
	for k, v = range mapObj {
		break
	}
	delete(mapObj, k)
	return k, v
}

func GetMapItems[T comparable, U any, R any](mapObj map[T]U, part MapPart) []R {
	if len(mapObj) == 0 {
		return []R{}
	}
	k, v := PopMapItem(mapObj)
	switch part {
	case KEY:
		key := any(k).(R)
		return append([]R{key}, GetMapItems[T, U, R](mapObj, part)...)
	case VALUE:
		val := any(v).(R)
		return append([]R{val}, GetMapItems[T, U, R](mapObj, part)...)
	}
	return []R{}
}
