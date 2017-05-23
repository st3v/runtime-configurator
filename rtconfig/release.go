package rtconfig

type release struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type releases []release

func (r releases) union(rels releases) releases {
	for _, rel := range rels {
		if i := r.indexOf(rel); i >= 0 {
			r[i] = rel
			continue
		}

		r = append(r, rel)
	}
	return r
}

func (r releases) substract(rels releases) releases {
	for _, rel := range rels {
		if i := r.indexOf(rel); i >= 0 {
			r[i] = r[len(r)-1]
			r = r[:len(r)-1]
		}
	}
	return r
}

func (r releases) indexOf(rel release) int {
	for i, v := range r {
		if v.Name == rel.Name {
			return i
		}
	}
	return -1
}
