// Code generated by go-bindata.
// sources:
// assets/aws-service-operator.yaml.templ
// assets/base.go.templ
// assets/cft.go.templ
// assets/operator.go.templ
// assets/template_functions.go.templ
// assets/types.go.templ
// DO NOT EDIT!

package codegen

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _awsServiceOperatorYamlTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x0c\x76\x03\x14\x28\x2a\x39\x46\x2f\x85\xb0\xbb\xc0\xd6\x76\x53\x23\x5b\xef\xc2\x76\x7a\x4d\xc6\xd4\x58\x66\x4d\x91\x04\x49\xd9\x71\x04\xff\xf7\x82\xfa\xf2\xb7\x77\xd1\x26\x3a\x89\xa3\xc7\x79\x6f\x1e\x87\xa3\x5b\x78\xa8\x1e\xe8\x3f\xc3\xe8\x79\x0a\x83\xfe\x70\x0a\xd3\x3f\x87\x13\xf8\x63\xf8\x34\x80\xbb\xf6\x09\x6e\x61\xba\xe0\x16\xe6\x5c\x10\x70\x0b\x98\x3b\x95\x92\x24\x83\x8e\x12\x58\x71\x84\x2f\xb8\xb6\xa1\xd2\x3e\xa2\x0c\x34\xdf\xbe\x04\xb7\x30\x9c\xc3\x46\xe5\x3f\x25\x20\xf8\x92\xc0\x2d\x08\xd8\x02\x65\x4a\x80\x72\xe3\x16\x5c\xa6\x80\x33\x95\x3b\x70\x2d\x41\x86\x4b\x02\x4a\xb8\xb3\xe0\x54\xb9\x23\x72\x94\x69\x11\xdc\xd6\x02\x64\x19\xd4\xcb\xb4\xc3\x54\x42\x29\xc9\x0e\x5a\x4b\xce\x42\xc2\x0d\x31\xa7\xcc\x26\x0a\x02\xd4\xfc\x6f\x32\x96\x2b\x19\xc3\xaa\x1b\x2c\xb9\x4c\x62\x78\xe2\xd6\x05\xdc\x51\x66\xe3\x20\x84\x2a\x36\xc2\x8c\xac\x46\x46\x01\xc0\xd1\x26\x80\x8c\x1c\x26\xe8\x30\x0e\x00\x00\x24\x66\x14\x83\x2f\xd5\x92\x59\x71\x46\x6d\xc9\x41\x51\x98\xb2\xaa\x77\x5c\x26\xf4\xf5\x17\x78\x47\x82\x32\x92\x0e\xe2\x7b\x88\x86\x9e\x71\xbb\x0d\xc2\x03\x02\xd4\x9c\xbe\x3a\x92\x7e\x65\xa3\xe5\x6f\x36\xe2\xaa\xb3\xea\xce\xc8\xa1\xa7\xae\xd4\xf5\x72\xeb\x54\x36\x26\xab\x72\xc3\xa8\x4f\x73\x2e\xb9\xe3\x4a\x5e\xd0\x56\x14\x0d\x71\x34\xd1\xc4\xa2\x66\x63\xf4\x22\x72\x83\x62\xbb\x8d\x8e\x95\x47\xb8\xb6\x01\x80\xd5\xc4\xaa\x44\xa9\x51\xb9\x8e\xe1\x02\xae\x22\xb2\x15\xb4\x11\x79\xcc\xfa\x91\xcb\x64\xbb\xad\x21\x82\x5b\xf7\xf1\x0a\xac\x3c\x93\x0a\xaa\x4b\x91\x6f\xa8\xa2\xc6\xdb\x85\x32\x6e\xb4\xaf\xa7\x28\x42\x38\x73\x10\xfe\x14\x2e\xa4\x9c\xf8\x1c\x65\x4d\x6d\xda\xb0\x11\x10\xf9\xdc\x6d\xd8\xe7\xa6\xbd\xc2\x2c\x97\x69\x2e\xd0\xbc\x55\xaf\x65\x4a\x5f\x3b\xa2\x89\xff\x5e\x63\x57\xbb\x2e\x44\xa1\x17\xd8\x0d\x8a\xa2\xe2\x6e\xda\xb6\x27\x72\xeb\xc8\x8c\x95\x38\x6e\x5c\x33\x43\x16\x61\xee\x16\xca\xf0\x6f\xe8\x9b\xe5\xb4\xb9\xde\xdc\xd7\x00\x26\x17\x95\xbf\x65\xf7\x7e\xf0\xdd\x51\xdb\x1d\xc2\xcd\x4d\xf9\x62\xea\x12\xda\xb8\x25\x66\xc8\xd9\x7a\xa5\x55\xd2\xbc\x32\x25\xe7\x3c\xcd\x50\xdb\x16\x59\x12\x36\x4b\x5a\x91\xac\xf7\xad\xc8\xcc\xda\x84\x29\xb9\xfa\x4d\x34\xed\x12\xc2\x1a\x1d\x5b\x34\x89\x0d\xa1\xa3\x7a\x91\x90\xa0\x76\x91\xeb\x64\xf7\x45\xd7\x5b\xce\xd4\x72\xee\x3a\x9e\xaf\x8e\x95\x77\xb2\x09\x27\xed\x9d\xfc\x7e\xba\xcf\xc8\xbb\x78\x1f\x4f\xe4\xdd\xfc\x7c\x73\x2a\xc4\x07\xdb\xde\x99\x54\xb9\x1e\x19\x53\xb9\x74\xff\x6b\xee\xb5\x13\xc1\x4f\xd0\x0b\x98\x73\x3d\xfb\x3b\x97\x09\x97\xe9\x0f\x6e\x5d\x25\x68\x4c\xf3\x0a\xd9\x38\x7a\x85\x25\xd8\x8d\xb4\xc3\xeb\xf5\x0a\x8f\xcd\x67\xff\x10\x73\xf5\x2d\xb9\xe0\xf1\xf7\xb3\xb1\x4f\x5a\xa8\x8d\x9f\x20\x47\xf6\xa1\xd6\xf6\xbf\x39\xf5\x3a\xfb\xfe\x2f\xc2\x90\x16\x9c\xa1\x8d\xa1\x5b\xae\xcb\xdf\x33\x3a\x6a\xa6\xf0\x21\x71\x69\xbe\x94\xca\x95\x56\xdb\x5d\x10\x80\x63\x16\x61\x86\xdf\x94\xc4\xb5\x8d\x98\xca\x3a\xfe\xc8\xae\xa8\x2c\x7f\x28\x38\x23\x71\x90\x06\xb5\xbe\xba\x67\xa7\xbc\x5c\x1d\x9c\xcc\xe8\xba\x2d\xfe\x61\x4a\x3a\xe4\x92\xcc\x1e\x69\xf8\x9a\x9f\x75\x81\x19\xa6\x15\xaa\x06\x35\x98\xce\xb9\x8d\xf1\xea\x7d\xf4\x3e\xea\x86\xe5\xc8\xff\xf5\x38\xcd\x4b\x2e\xc4\x8b\x12\x9c\x6d\x62\x78\x14\x6b\xdc\xd8\x7d\x0b\x4c\x7a\x60\x49\x33\x32\xc8\x1c\x05\xc3\x90\x55\xad\x1d\xfa\x02\xee\xef\x7a\x4f\x9f\x26\xd3\xc1\xf8\xf3\xe8\xf1\xaf\xc1\xc3\x09\xd6\x50\xca\x95\xbc\xbf\x1b\x0f\x3e\x0c\x9f\x47\xa7\xdf\xb1\x32\x31\xe4\xc9\xfd\xdd\x63\xaf\xf7\xfc\x69\x34\xfd\x3c\xec\x3f\xfc\x1b\x00\x00\xff\xff\x8d\xfe\x5b\xa0\x5b\x0a\x00\x00")

func awsServiceOperatorYamlTemplBytes() ([]byte, error) {
	return bindataRead(
		_awsServiceOperatorYamlTempl,
		"aws-service-operator.yaml.templ",
	)
}

func awsServiceOperatorYamlTempl() (*asset, error) {
	bytes, err := awsServiceOperatorYamlTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-service-operator.yaml.templ", size: 2651, mode: os.FileMode(436), modTime: time.Unix(1548187260, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _baseGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x3f\xcf\x9b\x30\x10\x87\xe7\xf8\x53\x9c\xa2\x57\x15\x44\x60\x76\xd4\x4e\xed\xd2\xa1\xa9\xda\x77\xe8\x50\x75\x30\xce\xc5\x41\x09\x36\xb5\x8f\x86\x0a\xf9\xbb\x57\x36\xd0\x90\xe1\x15\xc9\x74\xf1\xff\xdf\xf3\xe4\x68\x85\x3c\x0b\x85\x50\x09\x87\x8c\xd5\x4d\x6b\x2c\x41\xc2\x00\xb6\xd2\x68\xc2\x9e\xb6\xe1\xb7\xaa\xe9\xd4\x55\x5c\x9a\xa6\x10\x57\x77\x11\x95\x0b\x35\x77\x68\xff\xd4\x12\x73\xd3\xa2\x15\x64\x6c\xd1\x9e\x55\x21\x8d\x3e\xd6\xea\xe9\x63\xbf\x3b\xec\xb0\x11\x5a\x28\xb4\xe1\xf0\x30\xe4\x60\x85\x56\x08\x2f\xb5\x3e\x60\x9f\xc1\x0b\x5e\xb0\x41\x4d\x50\x7e\x00\xfe\x99\xb0\x71\xde\x3f\xfb\xca\x3c\x70\xc5\x30\xcc\xf7\xf1\xd7\x16\x25\xff\x8e\xce\x74\x56\x22\xdf\x8b\x06\xbd\x9f\x23\xa0\x3e\x78\xcf\x52\xc6\xe8\x6f\x3b\x5a\x02\x47\xb6\x93\x04\x03\x03\x18\x59\xa7\xc2\x3f\xc6\xc2\x00\x22\xcb\x97\x91\x05\x76\x4b\x32\xfe\x6d\xb1\xf4\x0c\xe5\x4a\x5c\xd8\xad\x6c\xe0\x5f\x27\xf2\x3b\x2e\xcf\xd8\xb1\xd3\x12\xf6\x78\x4d\xde\xc0\xc9\x1e\xe7\xc9\x58\x0a\xbb\xa8\x28\xb8\xb1\x48\x9d\xd5\xf0\x2e\x4c\x84\xf1\x7c\x7d\x39\xd5\x2c\xce\x2d\xaf\x2e\xef\x46\xe3\xfa\xc3\x82\x56\x15\x95\x6b\x1b\xf8\x1e\xaf\xb3\xa5\x64\xca\x78\x97\x28\xbd\x45\x1a\xf5\x01\xdc\x14\x26\xd5\xc8\x9e\xc2\x0f\x41\xf2\x94\x48\xea\x61\xfa\x84\x82\xc9\x50\x33\xd0\xa2\x41\xd7\x0a\x19\x9b\xa8\xd6\x2a\x8d\xa6\x1e\x66\xac\x8f\x10\xfa\x3c\xfe\x3b\x73\x78\xf7\x73\xbb\xd6\xca\xbf\xe2\x2b\x1b\x65\xa0\xe2\x6b\x0e\x5e\x49\x58\xfa\x4f\xb0\x48\x9c\x46\xda\xcd\x0d\x9e\x6d\xde\xe7\x92\x7a\xfe\xc9\x68\x4c\x52\xe6\xff\x05\x00\x00\xff\xff\x59\xfb\x37\x1f\x4a\x04\x00\x00")

func baseGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_baseGoTempl,
		"base.go.templ",
	)
}

func baseGoTempl() (*asset, error) {
	bytes, err := baseGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "base.go.templ", size: 1098, mode: os.FileMode(436), modTime: time.Unix(1544578164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _cftGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6d\x6f\x9b\xc8\x13\x7f\x0d\x9f\x62\xfe\x56\xd5\x3f\x44\x0e\xbe\xbe\xb5\x2e\x27\xf5\x9a\x3e\x44\xad\x92\x2a\x76\xdb\x17\x51\xd4\xae\x61\x8c\xb7\xc1\x2c\xc7\x2e\xf1\xf9\x2c\xbe\xfb\x69\x1f\x30\x0b\x36\x76\xdc\xf6\xda\xeb\xc9\xa8\x95\x61\x99\xf9\xcd\xe3\xce\xec\x84\xc1\x00\x7e\xd3\x17\x9c\x5f\xc1\xe5\xd5\x18\x9e\x9f\x5f\x8c\x61\xfc\xea\x62\x04\x2f\x2e\xde\x3c\x87\x5f\xd7\x97\x3b\x18\xc0\x78\x46\x39\x4c\x69\x82\x40\x39\x90\x42\xb0\x18\x53\xcc\x89\xc0\x08\xee\x29\x81\x4f\x64\xc1\x4f\x59\x26\x57\x58\x7e\x1a\xb2\x08\x63\x4c\x21\xcb\x59\x88\x9c\x7f\x92\x00\x17\x53\x58\xb2\xe2\xff\x11\x24\xf4\x0e\x41\xcc\x10\xc2\x19\x49\x63\x04\x92\x2e\xc5\x8c\xa6\x31\x90\x09\x2b\x04\x88\xb5\xa0\x39\xb9\x43\xc0\x88\x0a\x0e\x82\x29\x8e\x40\xe0\x3c\x4b\x24\x9a\xd6\x24\x55\xab\xd9\x5d\x3c\x30\x12\x07\x84\x73\x14\x1c\x22\x9a\x63\x28\x58\xbe\x0c\x5c\x37\x23\xe1\x1d\x89\x11\x56\xab\x60\x94\x61\x18\x5c\x23\x67\x45\x1e\x62\x70\x49\xe6\x58\x96\xae\x4b\xe7\x19\xcb\x05\x78\xae\xd3\xc3\x3c\x67\x39\xef\xb9\x4e\x2f\xa6\x62\x56\x4c\x82\x90\xcd\x07\x64\xc1\xe5\xff\x53\x1e\xdd\x9d\xc6\x4c\xde\xee\x26\xe0\x98\xdf\xd3\x10\x07\x61\xc2\x8a\x68\xca\xf2\x39\x11\x94\xa5\x3d\xd7\x21\x0b\xfe\xfe\x09\x49\xb2\x19\x79\x02\x2d\xfe\x84\x4c\x0c\x86\x66\x5e\xfb\x72\x20\xad\x23\x19\xe5\x83\xf6\x9b\x40\x8a\xbd\x37\x78\x9b\x1a\xed\x46\x0c\x59\x3a\xa5\xf1\xa1\x5c\x33\x4c\x32\x94\xfe\xf1\x5d\x19\x84\x4b\x5c\x40\x95\x06\x1c\x08\xa4\xb8\x00\x36\xf9\x8c\xa1\x70\xa7\x45\x1a\xca\xf7\x9e\x16\x04\xfa\x27\x78\xa6\x7e\xfa\x5d\xb1\x80\x13\xcb\x45\x41\x45\xf4\x9a\xa6\x51\x59\xf6\x41\xb0\x8c\x86\x4f\xaf\x2f\x81\x8b\x9c\xa6\xb1\x0f\x27\xcf\x1a\x1e\x86\x95\xeb\xe4\x28\x8a\x3c\x85\xc7\xcd\x37\x2b\xd7\x71\x9a\x68\xc3\x2e\x15\xfa\xae\xe3\x68\x65\x87\x8e\xbc\xf4\x7d\xdf\x05\x80\xb5\x02\x43\xd0\x57\xf5\xdc\x77\x9d\xd2\x2d\x95\x4b\x5a\x1a\x45\x38\xa5\x29\x72\x95\xa6\x5d\x36\x87\x53\xc1\x5d\xb1\xcc\xb0\xcd\xcc\x45\x5e\x84\x42\x5a\x65\x9c\x58\x5d\x0d\x67\xba\x2d\xcb\x76\xf9\xd0\xad\x75\x36\x50\xda\x95\x46\xf9\x91\x20\xe1\x9d\xd4\x0a\xb4\x1b\xb5\xde\xa9\x5c\x60\x53\x75\xcf\x25\x05\x4c\x08\xc7\x08\x98\xde\x7d\xad\x8d\xaf\x54\x52\xd1\xf7\x78\x3b\x40\x7e\x2d\xc1\xf3\x8d\x68\x2b\x68\x26\xbb\x82\x9a\x88\x07\x95\xa9\x49\xc1\x05\xe6\x72\xb1\x0f\xbd\x0e\x4f\xf6\xfa\xc0\x5b\x06\x07\x9a\x63\xeb\x32\xcf\x48\x88\xbe\x31\xfd\x25\x8a\xab\x42\x64\x85\xe0\xc6\x76\xcb\x5c\x66\x5e\x4c\x73\x36\x57\xcb\xe7\xc8\xc3\x9c\x4e\x50\x29\xca\x21\x24\x49\xd2\x6d\x72\x8d\xec\xf9\xe0\xcd\x49\x76\xa3\x0d\xbf\xd5\x3f\x7d\x50\x35\xc7\x97\x7e\xa8\x04\x0d\xcf\x60\x83\x6e\x55\xba\x0e\x47\xae\x5e\xae\xdd\xf2\xf4\xc3\x68\x84\x9c\x53\x96\xba\x0e\xbf\x0f\xe5\xbb\x66\xd1\x09\xe4\x16\x94\x6c\xbe\xeb\x3a\xca\x9a\x8b\xb4\x12\xd1\xa2\x6c\x1a\xa5\xc8\xe4\xb6\x59\x07\x43\x26\x3d\x59\xc8\xe8\x48\x75\x3c\x3b\x4c\xbe\xaf\x76\x40\x65\x80\x32\x49\xe9\x79\x1f\xb6\x60\xbd\xc7\x96\x12\xbe\xeb\xd0\xa9\xa2\xfd\xdf\x19\xa4\x34\x91\x2e\xa8\x72\x21\xa5\x89\x82\x91\xb8\x8e\x2c\x35\x4c\x00\x2f\x72\x04\x3a\xd5\x1d\x82\x72\xc0\x7b\xd9\x62\x18\xe7\x74\x92\xa0\xc2\x4a\x30\xf5\xb4\x0e\x5a\x3b\xee\x4b\xe8\x27\xdb\x80\x59\xce\x95\x73\x7a\x29\xd3\x61\xae\x02\x8f\x11\x2c\xa8\x98\x81\x98\x11\x61\x12\x40\xee\x80\x9e\xaf\x4d\x9c\xb2\x1c\x3e\xf6\x65\x4e\x48\x0b\x73\xd5\xc1\x1a\x22\x6f\x7e\xb9\x0d\xaa\x54\x92\x72\x4d\x50\x6f\x4e\x58\x21\xcc\x8b\xd7\xb8\xbc\x85\x33\xb0\x56\xde\x93\xa4\x40\x2d\xc0\xe8\x69\xd8\xb4\x13\x4c\x6d\xc9\x91\x08\xed\x48\x58\xd0\x24\x81\x50\x2d\x58\x99\x6a\x14\x47\xe0\x45\x96\x25\x14\x23\xc8\x48\x4e\xe6\xbc\x3b\x3b\x2d\x4c\x99\x9e\x5a\x2a\x9c\xb4\x72\xc3\xa2\xba\xb2\x42\x5c\x67\xee\x57\x66\x66\x38\x55\x8d\x5d\x1a\x33\x3c\x5b\xd7\x81\x97\x28\x94\xba\x2f\x2a\x9e\xb1\x21\x5a\xd7\x85\xc3\x6a\x41\xf5\x38\x5e\x66\x58\x96\xc1\x76\xec\x8e\x82\xf1\x60\x5e\x53\x55\xf6\xed\x36\xcb\xa3\x87\x6f\x35\xa7\x12\xf8\xee\xfa\xcd\xd0\xa6\xac\xfd\xa8\xc8\x2e\x99\xa0\x53\x1a\x2a\x89\x4f\xaf\x2f\xf9\x10\x6e\x6e\x4f\x4c\x39\x71\x1d\xc7\x69\xc8\xa8\x1a\x83\xe2\x2c\xfb\x55\x32\x6a\xaf\xaa\xa6\x60\x85\x46\xab\xff\x56\x26\x97\xd7\xbb\xb6\x88\xba\x6a\xb0\x5f\x63\xbd\xc7\x5c\xe6\xc5\x5e\x38\x43\xb7\x0d\xb1\x45\xe2\xbb\x4e\x5a\xb9\xbe\x13\x76\x1d\x9c\x2e\x15\x4d\xe4\x9c\xb0\xee\x35\x9d\x60\x56\x3f\x52\x70\x9b\x6d\xca\x77\x57\xab\x53\x53\x1f\x1e\x09\x96\x49\x28\x2d\xf2\x77\x16\x2d\x83\x51\x38\xc3\x39\x09\xde\xe6\xb2\x75\x0a\x8a\xbc\x2c\x65\x23\x3f\x95\xf5\x6d\xcb\xaa\xc1\x49\x91\xcb\x03\xb7\x84\x6a\xd0\xd4\xac\x26\x33\xe8\x5f\x18\x69\x5e\x29\x3b\x78\x8d\xcb\xb2\x5c\xad\x0c\xbf\x7e\x94\x94\x12\xa9\x57\xaf\x57\x69\x55\x96\x3d\x8d\x88\x09\xc7\x87\xc3\xc8\x7f\x65\x19\x5c\x4d\x3e\x07\xab\xd5\xa3\xc6\x7e\x31\x00\x23\x75\xa2\x51\x7c\x2a\x02\xf5\xa3\x62\xed\xd5\x82\xd3\x7d\xea\xab\x8a\xb9\xee\x34\x55\x90\x6a\xf3\xbd\x3d\x3a\xf7\xd7\x3c\xe7\x44\x90\xd5\xd5\xe4\xf3\x50\xa5\xc5\xa3\xe6\x99\x53\x9f\xb1\x86\x50\x57\x9c\x57\x9a\x6d\xb8\xe6\x97\x85\xcc\x2f\x77\x35\x33\xab\x2d\xca\x6d\x05\xb0\x43\xb7\xce\x94\xdb\x0c\x93\xa2\xef\xf5\xad\x73\x93\xdc\xca\x74\xba\xdc\x65\xbb\xf2\x9b\xef\xb7\xdc\x6c\x87\x5a\x25\x92\x87\x7f\xa8\xb4\x55\x01\x84\x9e\x5d\x07\x7a\x7e\x3b\x34\xbb\x74\x56\x18\x7b\x15\xde\xf0\xbc\x2e\x1a\x7e\x43\xb7\x07\x65\xb9\xc9\x47\xa7\x2d\x7b\x67\x4e\x5b\x4c\x07\x25\xf1\xfe\xac\x3d\x28\x4d\x7f\x70\x5e\x7e\xe3\x98\x6e\xb8\x61\x5b\xd6\xa9\xdb\xfa\xce\x75\xd4\x81\x05\x05\xe6\xaa\x6b\xde\xdc\xb6\x8f\x22\x6f\xab\xf7\xf2\x40\x6c\x11\x9f\x01\xc9\x32\x4c\x23\xaf\x5e\xeb\x83\x9d\xb6\xfe\xc3\xc9\xeb\xb6\xb2\x8f\x23\xb5\xfa\xc6\x3e\xda\xf0\x07\x35\x87\xbd\x8a\xed\xa8\x16\xdd\x75\xe2\x10\x50\xbf\x2b\xd8\xd6\x19\x29\x18\xa1\x58\xc7\x96\x5b\x50\x7e\xeb\x18\x32\x26\xf1\x66\x6e\x8e\x49\xfc\xa5\xe7\x90\xfd\x78\x5f\x76\x10\xe9\xc6\xfd\x92\x93\x48\x37\xda\x43\x8e\x22\xae\x23\x48\xdc\xb1\xa1\xc6\x44\xcd\x96\x8a\x60\x1d\x49\xf9\xd4\xdc\x3e\x63\x12\xfb\xbb\xa9\x6a\x7f\x76\x10\xda\xae\xe9\x20\x69\x5a\xec\x6f\xa6\xc8\x98\xc4\x5c\xd1\xfa\xae\x0b\x76\x01\x03\x3d\x6e\xda\xf3\x4c\x6b\xd6\xd4\x45\xcf\xcc\x52\xef\xb2\xa8\x39\x4b\x15\x6a\x41\x0d\x4f\xf8\x27\xe5\x82\xa6\xb1\x1e\xaa\xba\x67\x27\x0b\xc3\xd3\xec\xd1\xae\x3f\xbe\x74\x8f\x57\x16\xd0\xcf\x33\x5e\x19\x8b\x0f\x99\xaa\x0e\x66\x79\xe0\x30\x65\xf9\xef\x38\x4c\x1d\x87\xa9\xe3\x30\xf5\x6f\x18\xa6\xcc\x66\xff\x0e\x47\xd5\xe3\x04\xf5\x4f\x4c\x50\xf0\x5f\x9a\xa0\xbe\x5f\x32\x1e\xe7\xa6\xe3\xdc\x74\x9c\x9b\x8e\x73\xd3\x4f\x36\x37\xd9\xb3\xcc\xae\xb9\xe9\x1c\x13\x6c\xcc\x4d\x91\x5a\xa8\xbf\x41\x75\x8f\x4b\x16\xab\xe7\x83\xf7\x4d\x27\x9c\x7d\x9f\x36\xd7\x92\xf5\x7c\x50\x6e\xf8\xc7\xfe\xe6\x6c\x0f\x08\xd2\x5b\x1f\x6d\x47\xd9\x56\xec\x72\xd4\x07\x42\xc5\xbb\x54\xd0\x44\x91\x6a\xae\xe8\x50\x97\x6d\x05\xf9\xee\xce\xfb\xea\xef\xc2\x60\xf9\x6f\x9b\x49\xcf\xd8\x3c\x93\xbf\x9d\x0e\xfd\x3b\x00\x00\xff\xff\xd3\x02\x8f\xa7\xac\x24\x00\x00")

func cftGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_cftGoTempl,
		"cft.go.templ",
	)
}

func cftGoTempl() (*asset, error) {
	bytes, err := cftGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cft.go.templ", size: 9388, mode: os.FileMode(436), modTime: time.Unix(1544578164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _operatorGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\xdb\xb6\xb2\x7f\x26\x3f\x05\xca\xdb\xa6\x64\xca\x90\xcd\xdb\x1d\xcf\xf5\xed\xb8\xb6\xda\x7a\xea\xd8\x3e\xb6\xd3\x3c\x64\x32\x29\x4c\xae\x24\xc4\x24\xc0\x02\x90\x15\x57\xd5\x77\x3f\x03\x2c\x40\x91\x32\x25\x4b\x4d\xa6\xe7\xcc\x39\xf5\x8b\xf8\x07\x58\xec\xee\xef\xb7\x8b\x5d\xd0\x79\x4e\xfe\x1f\xff\xc8\xc9\x05\x39\xbf\xb8\x21\xa3\x93\xd3\x1b\x72\xf3\xd3\xe9\x35\xf9\xe1\xf4\x6c\x44\xfe\xaf\xfd\x0b\xf3\x9c\xdc\x4c\x99\x22\x63\x56\x01\x61\x8a\xd0\x99\x16\x13\xe0\x20\xa9\x86\x92\xdc\x33\x4a\x7e\xa5\x73\xf5\x42\x34\xe6\x89\x90\xc4\xbf\xfb\xd5\xcc\x3c\x1d\x93\x07\x31\xfb\xba\x24\x15\xbb\x03\xa2\xa7\x40\x8a\x29\xe5\x13\x20\x94\x3f\xe8\x29\xe3\x13\x42\x6f\xc5\x4c\x13\xdd\xae\x50\xd3\x3b\x20\x50\x32\xad\x88\x16\x76\x46\xa6\xa1\x6e\x2a\x23\x0d\x55\xe0\xf6\x69\x73\x37\xc9\x0b\x51\xc2\x04\x78\x4e\x95\x02\xad\x48\xc9\x24\x14\x5a\xc8\x87\x2c\x0c\x1b\x5a\xdc\xd1\x09\x90\xc5\x22\xbb\x6e\xa0\xc8\xae\x40\x89\x99\x2c\x20\x3b\xa7\x35\x2c\x97\x61\xc8\xea\x46\x48\x4d\xe2\x30\x88\x0a\xc1\x35\x7c\xd4\x51\x18\x2c\x16\x2f\x08\x1b\x13\x9c\x72\xaa\x8e\x67\x4a\x8b\x9a\xfd\x0e\xe5\x72\x19\x06\xa4\x06\x4d\xef\x5f\x92\xe8\xee\x7f\x55\xc6\x44\x4e\x1b\x56\xd3\x62\xca\x38\xc8\x87\xdc\xa8\x43\x1b\xa6\x72\x33\x28\xbf\x7f\x19\x85\x41\x34\x61\x7a\x3a\xbb\xcd\x0a\x51\xe7\x74\xae\x2a\x7a\xab\xcc\xef\x0b\x05\xf2\x9e\x15\xd0\x7a\xcc\xce\x9d\x42\xd5\x80\x54\x66\x9a\x84\x71\x05\x85\x57\x07\xb8\x59\x7c\x4f\x69\x85\xe0\x63\x36\x89\x42\x42\xf6\x9a\xf6\xdb\x0c\x66\x50\x53\x4e\x27\x20\xcd\xe4\x9e\x3b\xfe\x61\x5e\x2e\x97\x21\x21\x85\x90\xd0\xf7\x43\x6e\x1e\x59\xab\xfb\x2b\x32\xca\x0b\x51\x19\x91\xb9\xd2\xb2\xa0\x0a\xec\x08\xa5\x25\xe3\x13\x63\x2c\x9d\xab\xa2\x62\xc0\xf5\x7e\x8a\xe2\x1c\xf7\xa3\x40\xe7\xf7\x20\x15\x13\x1c\xca\x5c\x3f\x34\x50\xe6\xeb\xb3\x32\x3a\x57\xf9\xfd\x4b\x5a\x35\x53\xba\x37\x36\xd6\x2d\x3d\x38\xf6\x9b\xef\x6f\xcc\xba\xce\x67\xa8\xf9\x8b\x89\xc8\xb5\x10\x95\xca\x0b\x5a\x4c\x21\x0a\xad\x43\x7e\x71\x6a\xee\xe7\x12\x4b\xbe\xa7\xcc\xee\x01\xea\xd9\x4d\x35\x13\x5c\x65\x97\x18\x31\x16\xe1\xc8\x87\xcd\xa6\x31\x3d\x6f\x24\xa1\x09\xce\x0b\x9f\x00\x24\x34\x12\x94\x01\x86\x50\x62\x62\x4b\x8a\xaa\x02\x49\xc4\xed\x07\x28\x34\x19\x8b\xf6\x52\x69\x21\x81\x14\x76\x0d\x22\x5d\x84\xaa\xd0\x40\xb8\x12\xa7\xb4\x9c\x15\x9a\x2c\xc2\x00\x59\x4d\xf0\x0f\x6f\xb2\x63\xfb\x13\x06\x5a\x34\xac\x38\xba\x3a\xb7\xef\x90\x5f\x61\x60\x81\x7b\x85\x7c\x26\xcf\xbb\xec\x46\x36\xbb\x57\xe1\xd2\x1a\x70\x0e\xf3\x76\xd1\x42\x02\xd5\xd0\xd5\xde\xa8\x3d\xa7\xba\xb0\x49\x6b\xab\xfe\x6e\x72\x19\x8e\x67\xbc\xe8\x4a\x8d\x9d\x01\x3d\xd5\x53\xb2\xa3\x96\x09\x79\xde\xaa\xb7\x58\x83\xd2\xc7\x26\x5a\x5c\x68\x59\x91\x83\x43\x14\x9c\x9d\xc3\xdc\x2d\x9c\xfa\x95\x8f\xde\x5c\x1f\xfb\xd0\x49\xc9\xcb\x6f\x93\x95\xff\x52\xf2\xbe\x9d\x6a\xe4\x64\x57\x30\x61\x4a\x83\x8c\xa3\x0d\xa9\x34\x4a\xfa\x8e\xce\x8e\xca\x32\x5e\x89\xeb\xd9\xf3\x13\xe5\x65\x05\xf2\x87\x19\x2f\x62\xab\xf3\xeb\xa6\xa4\x1a\x64\x92\xf4\x93\x9d\x04\x3d\x93\x9c\x3c\xf3\x06\x2f\xc2\xc0\xa1\x7f\xd0\x83\x3f\x0d\x83\x0d\x8e\x68\x0d\x3a\x20\xad\x2e\x6e\xb0\x0b\xe1\x9e\xd2\x07\x3d\x14\xd2\x30\x58\x3a\x52\x5c\x6b\x2a\xf5\x1b\x83\x3b\xa2\x0f\xca\x32\x81\x71\xa5\x29\x37\x60\x8b\x31\xb9\x40\x36\x5c\x0f\xb3\x81\xf2\x92\xd0\x42\x2b\x22\xec\xc6\x55\x23\x2d\xe2\x62\x05\x67\xd2\x59\x25\x2e\xf4\x47\xe2\x36\x24\xc3\x10\xf3\x9b\x12\x4e\x6b\x50\x0d\x2d\xc0\x51\x3b\x31\x0c\xf0\x4b\x38\xa7\x2a\x83\x9b\xcd\x23\x2d\x42\xa3\x7b\xe0\xba\xe3\x73\x65\xfc\x78\x54\x96\xe6\xda\x3a\xb2\xc8\x04\x3f\x2a\x4b\xe3\x19\x44\x02\xdf\x98\xc7\x78\x6f\xde\x9c\x40\x05\xdd\x37\x78\x6f\x7d\x14\x06\x26\xd3\x98\x85\xdb\x8c\x63\xe8\xf6\x98\x2a\x97\xd5\x4c\xd2\x6a\xb9\x8c\x3a\xb6\xa4\x64\xdd\x82\x94\x14\xd9\x00\x45\xb3\xab\xd1\xf5\x0d\xde\xc5\x86\x29\x66\xad\x0c\xbd\xf5\xac\x93\x2f\x33\xbf\xea\xcf\xcc\x20\xbc\x58\xa6\xa4\xd0\x1f\xb3\x13\xc1\xc1\x4c\x5b\x86\xe1\x20\x55\xf2\x9c\x74\xa9\x48\xe6\xac\xaa\x88\xa6\xae\x5c\xa9\x41\x29\x3a\x31\xa8\x4b\x51\xdb\x27\x96\x28\x16\xd5\x46\x8a\x02\x94\xea\xa0\xda\x15\xb4\x21\xda\x6b\x35\x59\x0b\xf2\x57\xb8\xc4\xf7\xa2\x7c\x48\x08\x48\x89\xf1\x5d\x89\xc9\x04\x5d\xeb\x04\x9c\xd9\x07\x61\x70\x4f\xa5\xf5\xe1\x63\x56\x84\x01\x1b\x1b\xf9\x99\xd5\x80\xde\x56\x60\x04\x05\x66\x18\x39\xb4\x2f\x3c\x1e\x26\x72\xdd\x1b\x14\x80\xaf\xcf\xfd\x7d\x18\x2c\x09\x54\x0a\xe7\xe3\x5e\x75\x6d\x52\x85\xcd\x0d\xed\x9e\x6d\xb0\xfe\x41\x48\xb4\xcc\x99\x8b\x58\xd9\xcb\x24\x0c\x5a\x8e\xaa\xd4\x58\x66\xad\xf1\xd2\x5a\xb8\x90\x1b\x98\x4c\xe2\x28\x4a\xb2\x33\xa6\x74\x8c\xc5\x96\xbd\xbe\x68\xec\xee\xb3\x58\x1a\x89\x6c\x6c\x25\x7d\x71\x48\x38\xab\xac\x7e\xce\x55\xd9\x1b\xa6\xa7\x23\xe3\xbe\x18\xa4\x4c\x32\xbc\x8c\xd0\xa1\x13\xd0\xda\x24\xee\xcd\xc4\x34\xb2\x7d\xde\x01\x29\xc3\x20\x30\x39\xc2\xc4\xfa\xfb\x15\x51\x8d\x01\xd2\xd6\xaf\xad\x61\xd9\xa9\x86\x5a\xa1\x22\x6c\xdc\x3e\xcf\xae\x35\xd5\x33\x65\x7e\x8a\xbb\xd3\x13\x72\x88\x1e\xbe\xa4\x52\x41\xe9\x10\x7f\x1b\xe1\xdb\x32\x7a\x87\xf3\x3d\x54\xb2\x9b\x61\xdb\x17\x1e\xa9\xde\x5b\x07\x17\x6a\xbb\xc4\x98\x64\x63\x4b\x0d\xe3\xa3\x28\x22\xcf\x9e\x75\x88\x82\x8f\xcc\x62\x94\x73\xa1\x71\x5b\x37\x56\xd5\xb4\x79\x8b\x2c\x7a\x87\x3f\x56\x21\xa7\xe0\x49\x74\xb0\x55\xfb\x74\x35\xd6\xe8\xb4\x65\xb4\x7d\xed\xc6\x7b\x0c\x6e\x1e\x9a\x0d\x53\x7a\x23\xec\xac\x25\x32\x60\xcb\x58\xf4\x7b\xf4\xce\x78\x3c\xba\xba\x38\x3b\xfb\xfe\xe8\xf8\xe7\xf7\xc7\x17\xaf\x2e\xcf\x46\x37\x23\xb4\x3d\x10\xb7\x1f\x5a\x42\x96\x36\x9d\x59\xed\xda\x7d\x72\x2d\xbe\xd2\xad\xd6\x27\x0e\xfa\x75\x5a\xf6\xd9\x64\x55\x0f\x7c\x90\x40\x21\x64\x69\xf6\x4a\x44\x01\x4a\x9b\xaa\xc7\xb1\x55\xac\x03\x4d\xea\x8a\xed\xcc\xbe\x37\x7e\x78\x43\x25\x67\x7c\x92\x12\x57\x50\x67\x37\xe2\x98\xd6\x50\xc5\xae\xa8\xce\x6e\xc4\x99\x98\x83\x8c\x77\xf0\x51\x92\x0c\x5a\xd6\x1f\x76\x05\x54\x09\x8e\x66\xba\xac\xb0\x07\x02\x27\x23\xe3\xf6\x6d\xfe\x9f\xd9\x74\x89\x93\xfe\x1c\x00\x3b\x18\xb1\xd3\xa0\xae\xa5\xff\x25\x80\x06\xc6\xc6\x43\xc2\x78\x21\xa1\x06\xae\xaf\x44\x55\xdd\xd2\xe2\xee\x58\xcc\xb8\xde\x84\xc6\x1e\xfe\xe9\x6c\x23\x7f\x83\xde\x01\xfd\x5c\xc8\x9a\x56\xff\xa2\x20\x0e\x71\x9b\x70\x86\x71\x56\x85\xcb\xb0\x53\x7d\x0f\xd4\xa7\xb6\x4a\x34\x66\x11\xc6\x35\xc8\x31\x2d\x60\xb1\xb4\x45\xa8\xdd\x3c\xc4\xed\x87\x2c\x7e\xbe\xb9\x14\x4b\xb2\x13\x80\xe6\x58\x34\x0f\x71\x12\xb6\xad\x0b\x17\x7a\xb8\x13\x3d\x2a\x4b\xdb\x85\xb2\x31\x51\x7e\x1f\xed\x9b\x63\x53\x4b\x44\xfe\xf8\x63\xeb\x80\x81\xdc\x63\x8b\xde\xb1\x36\x4a\xdb\xbe\x28\xf3\xdc\xb3\xb5\xa7\xef\x14\x12\x3b\x4e\xcc\x74\x33\xd3\xab\xc2\x65\xac\xb3\x63\xdb\xd9\xe1\x5e\x81\x83\xd6\x19\xe3\xba\x93\xac\x57\xb4\x0d\x96\x26\x63\x5f\x9b\xd8\x76\x71\xb0\x38\xc1\x9a\x88\x7c\xfd\x95\xfa\x3a\x4a\x89\xb2\xf7\x89\x5b\x02\xd1\xb3\x37\xcb\x70\x68\xcd\x53\x3e\x16\xe3\x38\xa2\x65\x09\xe5\x56\xd1\x64\xce\xf4\x94\x28\x57\xa8\xf4\x16\x4b\x5d\x65\x19\x3f\x47\x67\xb8\x72\xa6\x4c\x92\x6d\x4b\xde\x33\x98\x13\xaa\xc9\x54\xeb\x46\x1d\xe4\x79\x21\xb8\x12\x15\x64\x74\xae\x32\x5a\xd3\xdf\x05\xb7\x47\x1a\x45\x25\x66\xe5\xd8\x04\x82\x81\x3d\x9f\x8a\x1a\xbe\xfb\x9f\xdc\xea\x91\x97\xa0\x29\xab\xbe\x43\xa5\xca\xc3\xaf\x54\xb4\x45\x95\x30\x08\xde\x23\x4c\xeb\x89\x65\x05\xaf\x37\x67\x55\x34\x6d\x14\x98\x92\xe8\xf8\x6a\x74\x74\x33\x7a\x7f\x7a\xfe\xfe\xf2\xea\xe2\xc7\xab\xd1\xf5\x75\x94\x92\x28\xda\x50\x7f\xee\x82\xb6\x07\xdb\xea\x67\xc0\x56\x18\xbe\x18\x90\x06\x43\xec\x4d\x2b\x65\xdb\xd7\x0d\xe7\x2f\x36\x32\xba\x66\xf5\xfa\xe6\x4d\xb1\x8b\x0d\x49\x2c\xaa\xf2\xc2\x24\x26\x0e\xf3\x8b\x81\x48\x16\xc2\x86\xb2\x1d\xb4\x7b\x34\x07\xdc\x4e\x43\x99\xfb\x24\x01\x9b\x00\xf6\x0b\x5e\xe3\x69\xcc\x44\x5c\x24\xd6\x67\x4f\xe7\x12\xb4\xdd\xa7\x13\x77\xca\x9a\x9d\x2a\x0b\xf6\xb1\xa8\x1b\x53\xfc\xc5\x62\x83\x26\x29\x19\xd3\x4a\x41\x62\x8a\xe8\x2f\xdc\xb9\xac\x35\x63\xf4\xdb\x8c\x56\x66\x9a\xb7\xd1\x24\xf5\xe5\x32\x35\x36\xf5\x1f\x25\xdb\xb2\x8e\x10\x3b\xa5\x9d\xd7\x9e\xd4\xc5\x9d\x33\xfd\xb3\x24\x9e\x96\x8b\x4f\x67\x07\x0e\x73\xd2\x50\x49\x6b\x45\xbe\xfa\xe6\xde\x76\xbf\xa2\x2a\xcd\x75\x64\x6d\xc6\xe8\xe2\xc2\x98\xb4\x6f\x86\xc2\x90\xdd\x25\x47\x39\x0d\x5c\x8a\x6a\x97\xfd\x0f\xc9\x51\xc2\xdb\xe3\x2e\x9e\xcc\x52\xaf\x2f\x4f\xfe\xf2\x2c\x85\x1f\x05\x9e\x48\x53\x3e\xe8\xd6\x98\x6e\xa8\xbb\x43\xba\xc2\xf3\xa5\xbf\xac\xda\xc0\xe5\x8c\x35\xbb\x14\x06\x41\x27\x2c\x4f\x3a\x9d\x63\x12\x0e\xf8\x7c\xaf\x68\xb4\x7d\xe8\x1e\x65\x80\xaf\x4c\xb1\x98\xdb\x40\x73\x6c\x6e\xb7\xc7\x56\x57\xe6\x0e\xe0\x7a\x7f\x6d\xdb\x86\xc2\x6d\x1f\xd2\x2c\xe8\x5b\x3b\x8e\xf5\x33\x33\x7b\x9a\x81\x71\x30\x74\x16\xfa\xf4\x99\xd9\x27\x1c\x61\xf9\xd3\x96\x5d\x0f\xb0\x56\x5d\x52\xf6\x23\xe8\x18\x7b\x1a\x77\x98\xf5\x23\xf4\xce\xb2\x1e\x13\xe6\xb3\x9d\x64\x75\x9b\x16\xac\xf5\x71\x94\x89\x0a\x7b\x88\xe5\x67\x75\xb7\xf2\xee\x20\x27\xbc\x0b\x4e\xe7\xe8\x69\xc3\x88\x6f\x5e\x86\xe1\x2a\xcd\xed\xee\x28\x57\xa0\x74\xa5\xff\x59\xf7\xb4\x99\xcb\x0b\x1b\x70\x47\xbf\xf3\x41\x3e\x0e\x34\xa4\x7b\xb1\x30\x6d\x0b\xe8\xce\xbd\xa9\x65\xfc\xad\xb4\x1d\x58\xcb\xd9\x6d\x09\x2c\x45\x46\x27\x7f\x53\x7a\x80\xd2\x9c\x55\xe9\x27\xf1\x7a\xb8\xe2\x74\x70\xed\x32\x18\x7b\x69\x1b\x0d\xe6\x62\x78\x4a\x7b\xee\xeb\x79\x81\x25\xef\xa6\x02\x54\xf5\xcb\x4d\xbb\x79\x74\x76\x23\xc7\xc0\xee\x42\x7e\xa7\xc7\xb2\x40\xf5\xea\x45\x03\x0a\x3e\x8e\x3f\xfd\xc8\xdc\x2d\x10\xf9\xa3\x03\x9b\xe9\xf1\x10\xfc\xcb\x46\x8a\xc6\x2c\x8a\xb8\x5d\xb8\x0a\xa5\x98\x42\x4d\xb3\x4b\x29\x1a\x90\x9a\x81\x22\xf6\x73\x9b\xdb\x12\xec\x9c\xec\x06\xea\xa6\xa2\xda\xff\x6b\x45\xb0\x58\xe0\xf3\x9f\xe1\xc1\xf0\xdf\x32\xda\xfb\x6a\x35\x36\x8e\xfc\x38\xf7\x0c\xec\x17\x25\x3f\xf0\x84\x6a\xba\xb8\xb8\xfd\x70\xb0\xe6\xa8\x63\xf7\xe5\x10\xdd\xb8\xec\x7e\x97\xb0\x90\x39\xbd\xbd\xec\x6b\xfb\xc1\xd9\x6a\x42\x0e\x49\x4f\x33\xff\x05\xd1\xed\x90\x7b\x88\x71\x5e\x7c\xfb\xc8\x02\x3b\x20\x7a\xb7\xf6\x6d\xb2\x73\xbd\xfc\xf7\xce\xa9\xfd\x78\xdc\x8f\xe2\x68\x93\x7a\xe0\xc5\x51\x59\x32\x93\x45\x68\xe5\x63\x4d\x0d\xb2\x7e\x2f\x3a\x9b\x1a\xc8\xdb\x60\x16\xe9\x9a\xe0\xf9\xdc\xd9\x0e\xfa\xa4\xe9\x6e\x0e\x8f\x3f\x11\x7c\xd2\xde\xf0\x77\xf6\xff\x4c\xd9\x7f\x4b\x8a\x74\xe9\xd1\x45\xcd\x2e\xc5\xfa\xe0\x0a\xab\xf9\x6f\x28\xd3\xaf\xb9\x66\x95\x15\x81\xd2\x4a\xdc\x5d\x7a\xec\xc1\xc9\x9e\x39\xdb\xb9\xbd\xce\x22\x45\xb6\xb5\x35\xc4\x38\xaf\x43\x8e\xbf\x90\x03\x9d\x63\x33\x64\x81\x3f\x4f\xdb\x95\x07\x8f\x0a\x30\xf7\x3d\x75\x78\xc7\x36\xdd\xc4\xaa\x7d\x18\xf0\x5f\x76\x8d\xff\xe8\x64\xb6\x16\xe5\x2f\x0f\x0e\xc9\x5b\xff\xf1\x72\xb9\x58\xb8\xf4\xd9\xd9\xad\xba\x5f\x71\x77\x12\xdd\x36\x8f\xf0\xdb\x6a\xb6\x3d\xd7\x21\xd1\xe8\xa3\x06\xc9\xd1\x43\x11\x0e\xfd\x52\xae\x75\x57\x9d\x4d\x0c\x4f\x8e\xbb\x93\xdc\x3a\x71\xa7\xd1\x1c\x3c\xa5\x8c\xba\x82\xbb\x02\xec\xde\x17\x2d\x16\x31\xe3\x25\x7c\xec\x28\x78\x29\xa4\x56\xe4\xdb\xc4\x5e\x60\xfc\xb4\x4e\x3a\x24\xb4\x69\x80\x97\xb1\x7f\x92\x92\xc7\x8a\xf7\xda\xb9\xfe\xe5\xae\xa8\xb4\x12\xb7\x8e\x3b\x24\x5e\x0d\x8f\xd7\xd3\xd0\x23\x89\x5f\xd1\xc6\x2c\x83\xce\xab\xcd\xcd\x67\x82\xbf\x27\x7e\x08\x55\x53\x65\x6c\xfc\x62\xde\x59\x8e\x69\xa8\xcd\xb8\xd5\x7c\x3b\xd3\x6e\xef\x06\x53\xf3\x1e\xab\x8a\xe8\x80\xb4\x0f\x7e\xa1\xd5\xcc\x22\xbb\x5e\x04\xec\xc4\xaf\x56\xf9\x27\x59\x35\x6c\x57\xe2\xff\xc9\xce\x7a\xb4\xe5\xca\xea\xd9\x2e\x6c\xd9\x1d\xbd\xad\x24\xe9\x8c\xf4\xdb\x9f\xd1\xa0\x25\xca\xee\x55\x51\x2f\x7b\xad\xd5\x45\x3b\x64\x2b\x7f\x92\xf9\xe8\x4b\xd9\x3f\x03\x00\x00\xff\xff\x14\x2d\xec\x35\x46\x2d\x00\x00")

func operatorGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_operatorGoTempl,
		"operator.go.templ",
	)
}

func operatorGoTempl() (*asset, error) {
	bytes, err := operatorGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "operator.go.templ", size: 11590, mode: os.FileMode(436), modTime: time.Unix(1548177925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _template_functionsGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x8f\xdb\x36\x10\x3d\x8b\xbf\x62\x2a\x04\xa8\x05\x78\x45\xec\xad\x08\x9a\x1e\xda\x6c\x12\x23\xc1\x6e\x91\x35\x90\x63\x43\x4b\x63\x69\x60\x8a\x24\xc8\x91\x1d\x43\xd0\x7f\x2f\x48\x7d\x04\xd8\x22\xdd\x8d\x2e\x94\x46\x9c\xf7\xe6\x3d\xce\x50\x4a\xf8\x63\x7a\xe0\xed\x03\xdc\x3f\xec\xe1\xee\xed\x6e\x0f\xfb\x0f\xbb\x47\x78\xb7\xfb\x74\x07\xbf\xaf\x8f\x90\x12\xf6\x2d\x05\x38\x92\x46\xa0\x00\xaa\x67\xdb\xa0\x41\xaf\x18\x6b\x38\x93\x82\xaf\xea\x12\x6e\xac\x8b\x11\xeb\x6f\x2a\x5b\x63\x83\x06\x9c\xb7\x15\x86\xf0\x35\x02\xec\x8e\x70\xb5\xfd\xaf\x35\x68\x3a\x21\x70\x8b\x50\xb5\xca\x34\x08\xca\x5c\xb9\x25\xd3\x80\x3a\xd8\x9e\x81\x57\xa2\x4e\x9d\x10\xb0\x26\x0e\xc0\x36\x65\x94\x8c\x9d\xd3\x11\x6d\xaa\xc4\xa4\xa8\x3b\x35\x72\x66\x94\x2a\x04\xe4\x00\x35\x79\xac\xd8\xfa\x6b\x29\x84\x53\xd5\x49\x35\x08\x2d\x6a\x87\x3e\x08\x41\x9d\xb3\x9e\x61\x23\x32\x75\x09\x95\x26\x34\x0c\x79\x43\xdc\xf6\x87\xb2\xb2\x9d\x54\x97\xa0\xd5\x21\xc4\xf5\x26\xa0\x3f\x53\x85\xab\x34\x99\xc8\x52\xce\xbc\x04\x64\x79\x46\x1f\xc8\x1a\xac\x25\x5f\x1d\xd6\xf2\x69\x56\xa9\x2e\x41\x9e\x6f\x95\x76\xad\xba\xcd\x45\xf6\x73\x6c\xd6\x1c\xa9\xc9\x45\xd6\x21\xab\xf3\x2d\xe4\xa7\xdf\x42\x49\x56\x2a\x47\x9d\xaa\x5a\x32\xe8\xaf\x69\xa3\x72\x14\x64\xdc\x24\xcf\xb7\xb9\x28\x44\xf4\xe9\x1e\x2f\xf0\x85\xb4\x06\x8f\xdc\x7b\xb3\x98\x10\x0d\x3d\x20\xb8\x68\x57\x0d\x64\x66\x7f\x93\xbd\x8a\x31\x88\x63\x6f\xaa\x98\xbc\x29\xe0\xc3\x9c\x32\x88\x6c\x06\x99\x23\x83\xc8\x00\x3e\xf6\x07\xf4\x06\x19\xc3\x67\x0c\xb6\xf7\x15\xde\xab\x0e\x5f\xff\x20\xbe\x15\x59\x36\x0c\x37\xe0\xd3\xc9\xbf\x22\x53\xe3\xb7\x2d\xbc\x42\x8d\x5d\x3c\x85\xd7\x6f\xa0\xdc\x31\x76\x01\xc6\x51\x64\xd9\x7b\xe4\x61\x58\x7e\x96\x8f\x0e\xab\xf2\x23\x99\x7a\x1c\xff\xbc\x4e\x24\xcf\x6c\x58\xd8\x30\x86\x44\x36\x8a\x31\x79\xb2\x08\xaa\xf1\x48\x06\x03\x28\xad\x93\xfa\x29\x0e\x51\x3a\x93\x35\x01\xb8\x55\x0c\xca\x23\xe0\x37\x67\xa3\x51\xff\xb1\x29\x1e\xf7\x8a\x17\xd8\xf7\x15\xc3\x20\x7e\xe4\x4a\x82\xde\x04\xf6\x64\x9a\x02\x96\x17\xf1\x72\x47\x9e\xd1\x3b\xe1\x4f\xfd\x52\xfe\x95\x96\x2d\x4c\x2c\xcb\x5a\xc0\x86\x0c\xa3\x3f\xaa\x0a\x87\x71\x0b\xe8\xbd\xf5\x73\x09\x93\x4d\xa3\x10\x2f\xad\x47\xca\xe7\x4e\x00\x2e\xb1\xf7\x8e\x64\xea\x64\x9c\x9f\xcd\x80\xc3\x15\x8c\xea\x70\x6a\xb3\x67\x30\x66\x41\xf0\x44\x57\xcc\x5f\xc5\xc5\x8f\xe0\x54\x85\xff\x2b\x33\xb6\xb0\xb6\x4d\x83\x3e\xea\x98\xf1\x3e\xa5\x80\xc8\xa6\x69\x7e\x44\xde\xc2\x3f\xf1\xf7\x7a\x37\x94\xf7\x78\x79\x67\xfd\xc4\xbb\xb8\xfb\xf9\xee\x71\x3f\x45\x8a\x38\x16\x93\xac\xc4\x93\xa0\x17\xac\xf2\xa9\xb0\xbf\x75\xef\x95\x8e\xb2\xc6\x71\xb3\x56\x5d\x94\xef\x91\xd3\xe7\x16\xa6\x29\x8f\x81\x07\x97\xda\x70\x18\x0b\x91\xd1\x31\x61\xff\xf2\x06\x0c\xe9\xa8\x63\x16\x52\x7e\x21\x6e\xef\xa2\xba\x0d\x7a\x5f\x94\xd3\x6b\x9e\xf4\x42\x83\xcc\xf1\x62\x0d\xb1\x97\xad\xa3\x2a\x2f\x44\xb6\x0c\x71\x9e\xa7\x72\xe3\x58\xac\x83\xfd\x5d\x88\x21\x2d\x46\xf1\xbd\x2b\xfe\x0d\x00\x00\xff\xff\x2d\x5b\xdf\x7e\x2e\x06\x00\x00")

func template_functionsGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_template_functionsGoTempl,
		"template_functions.go.templ",
	)
}

func template_functionsGoTempl() (*asset, error) {
	bytes, err := template_functionsGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template_functions.go.templ", size: 1582, mode: os.FileMode(436), modTime: time.Unix(1544578164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesGoTempl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x4d\x6f\xe3\x36\x10\x3d\x4b\xbf\x62\x60\x04\x85\xd5\x26\x36\x7c\x5b\x18\xd8\xc3\x6e\x16\x2d\xb6\x69\x36\x45\x9c\xf6\xb2\x58\x20\x0c\x35\x92\xb9\x96\x48\x96\xa4\xd2\x1a\x86\xfe\x7b\xc1\x0f\xc9\xa6\x25\x6d\xd3\x43\x0e\xab\x93\x34\xf3\xe6\x91\x7c\xf3\x21\x4a\x42\x77\xa4\x44\x78\x5e\x91\x4a\x6e\xc9\x2a\x4d\x59\x2d\x85\x32\x30\x4f\x93\x1a\x0d\x79\x5e\xc1\x6c\xf7\x46\x2f\x98\x58\x12\xc9\x6a\x42\xb7\x8c\xa3\xda\x2f\xe5\xae\xb4\x06\xbd\xb4\xa0\xe5\xf3\x6a\x96\x26\x93\x38\xd5\x70\xc3\x6a\x9c\xa5\x59\x9a\x2e\x97\xf0\x53\x89\x9c\x56\x0c\xb9\x89\xbf\xd6\x5c\x6c\x0c\x31\x8d\x76\xe6\xdd\x1b\xbd\xce\x11\x25\x15\x72\x7f\x55\x22\x5f\x33\x6e\x50\x15\x84\xa2\x7e\xfb\x1f\x0b\x2d\xee\x9e\xbe\x22\x35\x6e\xb1\xc3\x61\xb1\x91\x48\x17\x37\x8c\xe7\x6d\x0b\x39\x16\x8c\xa3\x06\xb3\x45\x78\x22\x1a\x41\xa1\x16\x8d\xa2\x98\x9a\xbd\xc4\x73\xb4\x36\xaa\xa1\x06\x0e\x9d\x14\x8b\x87\xbd\xc4\x5b\x34\x04\x00\x1e\xbf\x6a\xc1\xd7\xb3\x4b\xc6\x2b\xc6\x71\xf6\xd8\x63\xfc\xe2\x0e\x15\x30\xd6\x91\x13\x43\x2c\xa8\x5b\xc1\x32\xb5\x2d\x00\x9c\xad\x39\xf0\x07\x8e\x53\xfb\xaf\x9b\xbb\x4f\x6d\x6b\xd9\xbc\x5e\x70\xf2\xc4\x6c\x03\x7f\x60\xd3\xce\x6e\x19\xee\x1a\x23\x1b\x33\xcd\x30\xf0\x07\x06\xe1\xec\x96\xe1\x5d\x9e\x33\xc3\x04\x27\xd5\x7d\xd0\x52\x9f\x71\x8c\x21\x02\x0d\x19\xba\x66\x8f\x69\x9b\xa6\x87\xc3\x15\xb0\x02\xe6\x25\xc2\xbc\x42\x0e\x9e\xef\xbd\xc8\xf7\x8b\x0d\xdd\x62\x4d\x16\x37\xb8\xbf\x25\x52\x32\x5e\x66\xb0\xca\xda\xd6\x85\x28\xc2\x4b\x84\x0b\xc6\x73\xfc\xe7\x12\x2e\xb0\xc2\x1a\xb9\x81\xf5\xdb\x6f\x12\xb4\xad\x2f\x95\x8b\x38\x13\x9e\xe6\xac\x6a\x4e\xcc\x5d\xed\x40\x21\xd4\x79\x74\x57\x4f\x53\x94\xc7\xd2\xea\x4e\x2a\x14\xcc\xf1\xaf\xb0\x77\x98\xd9\xb0\x59\x16\x99\x3e\xd8\x1a\xca\xec\x51\x93\xeb\x4a\x34\xf9\xcf\x42\xd5\xc4\xaa\xf7\x80\xb5\xac\x88\xc1\x4f\xa4\x46\xcb\xcc\x78\xd9\x09\x4c\x27\x81\x36\x77\xd3\x34\x5a\x12\xfa\x72\x2e\x87\xb6\x84\xf7\xa2\xaa\x9e\x08\xdd\x5d\x8b\x86\x1b\x60\xdc\x74\xb1\xea\xd4\xe1\x1b\xe1\x0a\xd0\x29\x65\x4b\xae\xcf\x9c\x54\x42\xda\x74\x75\xa9\xb3\x87\xed\x14\xe2\xc1\xef\x9a\x00\x66\x9d\xfc\xee\x30\x59\x20\xf2\x80\x8d\x53\xf7\x06\xf7\x6d\xdb\xdb\x6c\xd0\x5d\xe1\xbb\x13\xa2\xb4\x1c\x7b\xcc\x23\x5d\xdc\xd8\x1e\xfd\xbb\xad\x34\xff\xe6\x6c\x95\xc6\xc9\x74\x47\xbd\x7c\x4c\xf9\x77\x9f\xbd\x70\x74\x27\xc2\x70\xca\x86\x99\x71\xda\x35\x7e\x5c\x9c\x77\xcc\x58\xc3\x8c\x10\xf5\xc2\x8d\x56\x8a\x0f\xf0\xd0\xae\xb5\x7f\x57\x42\xa2\x32\x0c\x35\xbc\x42\x5d\xc4\xc5\x30\x3c\x7f\x98\xba\x5b\x51\xe5\xfe\xf4\x7e\xdc\x82\x28\xdc\x97\x4b\x5b\xd1\x25\x02\x4c\xc8\xc4\x98\x00\x81\xe9\x44\x80\x6e\x4c\x46\x93\x3d\x4e\xb4\x8a\x20\x6e\xbf\x71\xd4\x3d\x12\x2d\xf8\x37\xa3\x3c\xa4\x3f\x2b\x2b\x82\xcc\x7f\x68\x8c\xab\xce\xc9\xbb\x31\x84\xee\x3e\x7e\x80\xe8\x89\xe9\xb5\x87\xbc\x40\xbd\xb1\xbf\xc5\x51\xca\xe3\x0f\xa3\x2f\x26\x3d\xa6\xdc\x18\xcb\x60\xe6\xfa\x80\x11\xe8\x62\x83\xea\x99\x51\xd4\x76\xfc\x74\xef\x49\xf2\xf9\x8b\x3f\x54\xd2\x1d\x2a\x78\xe2\x59\xf1\x12\x76\xaa\xd0\x04\x72\xf7\x9a\x8c\x91\x3b\xcf\xff\xe5\xbe\x16\xbc\x60\xe5\x2d\x91\x8e\xfe\xf8\x35\xe0\xa7\xbd\x2b\x5e\xc2\xe7\xe4\x75\xee\x5f\xbf\x31\x1d\xcf\x85\xca\x1a\x88\x31\x8a\x3d\x35\xc6\x8f\x05\xff\x93\x8d\xfa\xd0\xe6\x77\x2c\xc9\x8e\x6e\xfa\x92\x36\x79\x45\xb3\x71\x53\x17\xb4\x8f\x06\xeb\xd3\x1b\xd3\xe7\x2f\x67\x9b\x09\x31\xcc\xe2\xfc\x5d\xa5\x68\x38\x05\xc6\x99\x99\x67\x76\x1b\x95\xa0\xa4\x72\x83\x08\xdf\x37\xac\xca\x51\x2d\xee\xb1\x64\xda\xa0\x9a\x93\x3c\x8f\xe9\xec\x66\x75\xd6\xb3\x8c\xfb\xe7\xda\xb1\xc1\x8f\x9d\xc0\x9e\x3d\x03\x54\xca\x0e\xd2\x34\xf1\x00\x5b\x10\x37\x5c\xfc\xcd\x7d\x94\x47\xfd\xa2\x44\x23\xff\x44\xa5\x99\xe0\x97\x69\x92\xfc\x70\x76\xe1\x6c\x47\x8c\x56\x20\xe7\xc8\x7a\xcd\xde\xe5\xf9\x83\x38\xe5\x0a\x9b\xba\x84\xe1\x32\x59\x9a\x28\x34\x8d\xe2\xc0\x59\x95\xb6\xe9\xbf\x01\x00\x00\xff\xff\x10\x11\xf5\x38\x65\x0c\x00\x00")

func typesGoTemplBytes() ([]byte, error) {
	return bindataRead(
		_typesGoTempl,
		"types.go.templ",
	)
}

func typesGoTempl() (*asset, error) {
	bytes, err := typesGoTemplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types.go.templ", size: 3173, mode: os.FileMode(436), modTime: time.Unix(1544578164, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"aws-service-operator.yaml.templ": awsServiceOperatorYamlTempl,
	"base.go.templ":                   baseGoTempl,
	"cft.go.templ":                    cftGoTempl,
	"operator.go.templ":               operatorGoTempl,
	"template_functions.go.templ":     template_functionsGoTempl,
	"types.go.templ":                  typesGoTempl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"aws-service-operator.yaml.templ": &bintree{awsServiceOperatorYamlTempl, map[string]*bintree{}},
	"base.go.templ":                   &bintree{baseGoTempl, map[string]*bintree{}},
	"cft.go.templ":                    &bintree{cftGoTempl, map[string]*bintree{}},
	"operator.go.templ":               &bintree{operatorGoTempl, map[string]*bintree{}},
	"template_functions.go.templ":     &bintree{template_functionsGoTempl, map[string]*bintree{}},
	"types.go.templ":                  &bintree{typesGoTempl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
