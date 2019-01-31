package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
)

type ConfigReader interface {
	Read(interface{}) error
}

type SystemPaaSTAConfigFileReader struct {
	Basedir  string
	Filename string
}

func ParseContent(r io.Reader, t interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, t)
	return err
}

func (c SystemPaaSTAConfigFileReader) FileNameForConfig() string {
	return path.Join(c.Basedir, c.Filename)
}

func (c SystemPaaSTAConfigFileReader) Read(t interface{}) error {
	r, err := os.Open(c.FileNameForConfig())
	defer r.Close()
	if err != nil {
		return err
	}
	return ParseContent(r, t)
}