package collection

type collectionBase struct {
	value  interface{}
	length int
}

func (c collectionBase) Value() interface{} {
	return nil
}

func (c collectionBase) Length() int {
	return 0
}

func (c collectionBase) Json() string {
	return ""
}

func (c collectionBase) Join(delimiter string) string {
	return ""
}

func (c collectionBase) Min() float64 {
	return 0
}

func (c collectionBase) Max() float64 {
	return 0
}

func (c collectionBase) Contains(value interface{}) bool {
	return false
}

func (c collectionBase) Unique() collection {
	return nil
}

func (c collectionBase) DelKey(key int) collection {
	return nil
}

func (c collectionBase) DelValue(value interface{}) collection {
	return nil
}

func (c collectionBase) Pluck(key string) collection {
	return nil
}

func (c collectionBase) DelKeyValue(key string, value interface{}) collection {
	return nil
}

func (c collectionBase) Filter(callback filterCallback) collection {
	return nil
}

func (c collectionBase) GroupBy(key string) map[string]interface{} {
	return map[string]interface{}{}
}
