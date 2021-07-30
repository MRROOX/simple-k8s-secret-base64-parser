package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"encoding/base64"

	"gopkg.in/yaml.v3"
)

type Secret struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Type       string `yaml:"type"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`

	Data map[string]string `yaml:"data"`
}

func readConf(filename string) (*Secret, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Secret{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return c, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func encodeToBase64(fileName string) {
	c, err := readConf(fileName)
	checkError(err)

	s := c

	for key, value := range c.Data {

		encoded := base64.StdEncoding.EncodeToString([]byte(value))

		s.Data[key] = encoded
	}

	//fmt.Println(s)

	data, err := yaml.Marshal(s)
	checkError(err)
	err = ioutil.WriteFile("base64-"+fileName, data, 0777)
	checkError(err)

}

func decodeFromBase64(fileName string) {

	c, err := readConf(fileName)
	checkError(err)
	s := c
	for key, value := range c.Data {

		decoded, err := base64.StdEncoding.DecodeString(value)
		s.Data[key] = string(decoded)
		checkError(err)
	}

	data, err := yaml.Marshal(s)
	checkError(err)
	err = ioutil.WriteFile("decode-"+fileName, data, 0777)
	checkError(err)

}

func main() {

	file := flag.String("f", "secret.yaml", "K8S Secret File")
	procces := flag.String("p", "e", "Type of Processing => e or d | by default = e")

	flag.Parse()

	f := *file
	p := *procces
	if f != "" && p == "e" {
		encodeToBase64(f)
	}
	if f != "" && p == "d" {
		decodeFromBase64(f)
	}

}
