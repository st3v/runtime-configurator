package rtconfig

import (
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type RuntimeConfig struct {
	Releases releases `yaml:"releases"`
	Addons   addons   `yaml:"addons"`
	Tags     tags     `yaml:"tags"`
}

func (r RuntimeConfig) Add(other RuntimeConfig) RuntimeConfig {
	return RuntimeConfig{
		Releases: r.Releases.union(other.Releases),
		Addons:   r.Addons.union(other.Addons),
		Tags:     r.Tags.union(other.Tags),
	}
}

func (r RuntimeConfig) Remove(other RuntimeConfig) RuntimeConfig {
	return RuntimeConfig{
		Releases: r.Releases.substract(other.Releases),
		Addons:   r.Addons.substract(other.Addons),
		Tags:     r.Tags.substract(other.Tags),
	}
}

func (r RuntimeConfig) Write(w io.Writer) error {
	data, err := yaml.Marshal(&r)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func Read(r io.Reader) (RuntimeConfig, error) {
	rc := RuntimeConfig{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return rc, err
	}

	return rc, yaml.Unmarshal(data, &rc)
}

func FromFile(path string) (RuntimeConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return RuntimeConfig{}, err
	}
	defer file.Close()

	return Read(file)
}
