package collection

type CollectionBase struct {
	value  interface{}
	length int
}

func (c CollectionBase) Value() interface{} {
	return nil
}

func (c CollectionBase) Length() int {
	return 0
}

func (c CollectionBase) Json() string {
	return ""
}

func (c CollectionBase) Join(delimiter string) string {
	return ""
}

func (c CollectionBase) Min() float64 {
	return 0
}

func (c CollectionBase) Max() float64 {
	return 0
}

func (c CollectionBase) Contains(value interface{}) bool {
	return false
}

func (c CollectionBase) Unique() Collection {
	return nil
}

func (c CollectionBase) DelKey(key int) Collection {
	return nil
}

func (c CollectionBase) DelValue(value interface{}) Collection {
	return nil
}

func (c CollectionBase) Pluck(key string) Collection {
	return nil
}

func (c CollectionBase) DelKeyValue(key string, value interface{}) Collection {
	return nil
}

func (c CollectionBase) Filter(callback FilterCallback) Collection {
	return nil
}

func (c CollectionBase) GroupBy(key string) map[string]interface{} {
	return map[string]interface{}{}
}
