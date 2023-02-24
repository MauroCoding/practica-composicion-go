package items

type item struct {
	name  string
	id    byte
	value float64
}

type items []item

type Items = items

func (i *item) GetValue() float64 {
	return i.value
}

func NewItem(name string, id byte, value float64) item {
	return item{
		name, id, value,
	}
}

func NewItemSlice(el ...item) items {
	var itemSlice []item
	for _, v := range el {
		itemSlice = append(itemSlice, v)
	}

	return itemSlice
}
