package rtconfig

type tags map[string]string

func (t tags) union(others tags) tags {
	result := t.copy()
	for k, v := range others {
		result[k] = v
	}
	return result
}

func (t tags) substract(others tags) tags {
	result := t.copy()
	for k, _ := range others {
		delete(result, k)
	}
	return result
}

func (t tags) copy() tags {
	c := tags{}
	for k, v := range t {
		c[k] = v
	}
	return c
}
