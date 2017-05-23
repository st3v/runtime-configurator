package rtconfig

type addon struct {
	Name       string                 `yaml:"name"`
	Jobs       jobs                   `yaml:"jobs,omitempty"`
	Properties map[string]interface{} `yaml:"properties,omitempty"`
	Include    map[string]interface{} `yaml:"include,omitempty"`
}

type jobs []map[string]interface{}

type addons []addon

func (a addons) union(others addons) addons {
	for _, other := range others {
		if i := a.indexOf(other); i >= 0 {
			a[i] = other
			continue
		}

		a = append(a, other)
	}
	return a
}

func (a addons) substract(others addons) addons {
	for _, other := range others {
		if i := a.indexOf(other); i >= 0 {
			a[i] = a[len(a)-1]
			a = a[:len(a)-1]
		}
	}
	return a
}

func (a addons) indexOf(item addon) int {
	for i, v := range a {
		if v.Name == item.Name {
			return i
		}
	}
	return -1
}
