// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x3c\x6b\x73\xdb\x38\x92\xdf\xf3\x2b\x50\xbc\xb9\x93\x33\x23\xcb\xb2\xa7\x76\xef\x96\x73\xb9\x2a\x47\x76\x26\xde\xb5\x13\x9d\xe4\x64\xea\x36\x49\x6d\xc1\x24\x24\x71\x4d\x01\x1c\x00\x74\xec\x51\xe9\xbf\x5f\xe1\x41\x12\x4f\x4a\xb6\x93\xda\x6c\xed\x24\x22\x1b\x8d\x46\x77\xa3\x5f\x68\x70\xb3\x01\x39\x5a\x14\x18\x81\x04\x56\x55\x02\xb6\xdb\x17\x00\x6c\x36\xe0\x07\x58\x55\x20\x7d\x05\x46\xa7\x55\xd5\x3d\x5c\x43\x5c\x2c\x10\xe3\xf2\xcd\x55\xf3\x43\xbd\x7e\x01\x00\x00\xc9\xe9\x6f\xf3\x6b\xb4\xae\x4a\xc8\xd1\x1b\x42\xd7\x90\x7f\x44\x94\x15\x04\x27\x20\x05\xc9\xc9\xf8\x78\x7c\x38\xfe\xcb\xe1\xf8\x2f\xc9\x50\x81\x4f\x08\xce\x0b\x5e\x10\xcc\x92\x54\xa3\x90\x33\x71\x8d\x03\x24\x37\xb0\x84\x38\x43\xf4\x30\xeb\x40\xdd\xb9\xbd\x41\x15\x25\x19\x62\xec\x51\x63\x28\x5a\x16\x8c\xd3\x87\x5d\x83\x92\x0b\xcc\x11\xc5\xb0\x14\x14\x83\xe4\x0d\x4e\xd3\xf3\xdf\x6b\x58\x8a\x15\x7c\x12\x4f\x66\x68\x91\xa4\x06\x18\xd8\x0e\x41\xf2\x7f\x88\x25\xe0\x0b\xd8\x0e\x1b\x2c\x53\x5a\xdc\x41\x8e\x76\x20\x69\xa0\xc2\x38\x5e\x97\x10\xdf\xce\x51\x56\xd3\x82\x3f\xfc\x4a\x49\x5d\x09\x36\x6f\x4c\x74\x20\x05\x9f\x36\x12\x9b\x10\x80\x0d\x2b\x70\x26\x5f\xd4\xba\x34\xd2\x64\x0a\x29\x5c\x23\x8e\xa8\x1c\xda\x2f\x91\x4a\xc0\x3e\x42\x1a\x41\xf8\x66\x2d\x93\xb2\x66\x1c\x51\x43\x0d\x00\x48\xae\x1f\x2a\xa4\x08\xe7\xb4\xc0\xcb\x64\xd8\xbd\x3a\x43\x0b\x58\x97\x5c\xbe\xb5\x9f\xb3\x8c\x16\x15\x6f\x74\x2e\xd1\xaf\x3a\xae\x9d\xa1\xaa\x24\x0f\x6b\x84\xf9\x15\xbc\x2f\xd6\xf5\x3a\x30\x67\x0a\x92\x77\xf5\xfa\x06\xd1\xd0\x94\x52\x93\xc7\xb1\x49\x53\x90\x68\xbc\xa0\x42\x34\x43\x98\xc3\x25\x02\x64\x01\x34\x1b\x10\x03\x9c\x80\x5b\x84\x2a\x40\x6b\x8c\x0b\xbc\x04\x5f\x57\x45\x89\x40\x2e\xe9\x12\xcb\xec\x23\xb9\xc0\x4f\x24\xf9\x4f\xbd\x14\x2b\xb4\xdf\x8e\xe2\x73\x7c\x57\x50\x82\x05\xc9\x61\x5a\xe3\x12\xed\x11\x68\x50\x9e\xe6\x7e\xdc\x6f\x1e\x0b\xe1\x7b\x5c\x3e\x00\x58\x96\xe4\x2b\x80\x99\x58\xae\x58\x2c\x5f\x15\x0c\x08\x13\xb8\xa0\x64\x0d\x0a\xcc\x8a\x1c\x01\xbe\x42\xe0\xe3\x74\x12\xa1\xf9\x1d\x31\x5f\x9c\x0a\x84\x28\xff\x08\xcb\x1a\xa9\x4d\x2d\xb7\xef\x50\xc2\x81\x2f\xde\x22\xfe\x86\x1e\xbe\x37\x9f\x0c\x8b\xf3\x04\x36\x7d\x60\x08\xcc\xeb\x1b\x8c\x38\xd3\x88\x04\x9f\x58\x85\xb2\x62\xf1\x20\xd8\x72\x28\x79\x54\x12\x98\x83\xc6\x42\x00\x84\xf3\x8a\x14\x98\xb3\xef\xc2\xb3\x19\x2a\x11\x64\xa1\x05\x7d\x6b\x93\x31\x43\x15\x61\x05\x27\x34\x24\xa4\xe7\x4d\x36\x27\x35\xcd\x10\xc8\x48\x8e\x00\xed\xa6\xf1\x48\xb0\x4d\xf7\xb7\xa6\xe2\x7a\x85\xc0\xa5\x25\x3a\xa6\xe7\x03\x4b\x31\x21\x58\x10\xda\x6e\x8a\x00\x71\x4a\x31\x22\x64\x5d\x16\x8c\xff\xf7\xe9\x6f\xf3\x34\x3d\x9f\x9c\xa4\xa9\x02\x4e\xd3\x8b\xfc\x7f\x9e\x42\xea\xc7\xe9\x04\x30\x35\xdf\x7e\x54\xc5\xf5\xfe\xfb\x10\x57\xe9\xed\xb1\x1f\x91\x4d\x7c\x64\x51\xe7\xec\xbd\x83\xd9\xf9\xff\x7e\xb8\x98\x9d\x9f\xbd\x04\x97\x70\x7d\x93\x43\x30\xa9\x19\x27\xeb\x6b\x52\x15\x19\x78\x0b\x71\x5e\x22\x0a\xf4\x76\x00\x0d\x46\x83\xcc\xab\x02\x5f\x22\xbc\xe4\x2b\x49\xe4\xb1\xf9\xca\x31\x00\x3e\x7d\xd3\x49\x84\x73\x1d\xd3\x3e\x4e\x27\x82\x63\x4f\x65\xd8\x0e\x06\x4d\x27\x93\x8b\xb3\xd9\x37\x57\x79\x31\xb3\x40\x1c\x9e\xde\x0a\x8a\xae\x60\x55\x15\x78\x69\xea\x77\x32\x25\x94\x4f\x29\xe1\x24\x23\x8e\xe7\x59\x71\x5e\xa9\xb0\x4e\xe8\x16\xc2\x88\x1a\x70\xc9\xdb\xeb\xeb\xa9\x30\x69\x17\x98\x71\xb1\xd3\x42\xef\xe4\x5e\x47\x31\x88\x79\xd2\x71\x47\x4f\xc7\xfa\xe7\x9b\x3f\x7b\x42\x6b\x46\x9e\xf5\xac\xef\x7a\x12\x5d\x9e\x7e\x15\x9f\x6c\x3e\xbf\x74\xa7\x2a\x7b\x96\x26\xc0\x9f\x37\x15\xd8\x06\xe5\x3d\x43\x4c\x5a\x65\x4b\xe0\xc6\x96\x9b\x91\x32\xe2\x46\xe5\x9e\xb8\x38\xbd\x4a\x53\x09\x63\xac\x64\x4a\x49\x85\x28\x2f\x90\x6d\x25\x85\xdb\x63\xac\x5e\x23\x01\x3f\x25\x65\x91\x3d\x9c\x91\xac\xf6\xe2\x26\xc7\x56\x88\x54\xea\xe4\xf0\x78\x7c\x78\xfc\x9f\xc6\x24\x12\x68\xce\x21\x47\x7a\xfc\x27\xeb\x15\x70\xf0\x49\xf0\xf3\xc5\x02\x65\xd2\x19\x4b\xf7\xeb\x60\xd3\xa4\x17\x38\x2b\xaa\x26\xe3\x99\x23\x7a\x57\x64\x48\x39\xe8\x52\xda\xa3\x11\x5c\xc3\x3f\x08\x86\x5f\xd9\x28\x23\x6b\x2b\x49\x31\x17\x9a\x69\x83\xf6\x09\x24\x8c\xb3\xb4\x5b\x78\xe7\xdd\x9b\x3f\x5b\xeb\xb7\xf9\xd6\xc2\x9c\x4c\x21\x5f\x09\xe2\x8f\x32\x82\xef\xc8\xfd\x51\x62\xbf\x15\x0c\x55\x2c\xb7\x59\xe1\x32\x42\x41\x3e\xbc\x83\x6b\x25\xc6\x7c\x5d\x60\x91\x0d\x42\x4e\xa8\xc7\x92\x64\x87\x9c\xc0\xbe\xb2\x02\x9e\xbc\x04\x7f\x3d\x89\x18\x9c\x4b\x7e\x14\x3f\x1b\xfd\x54\x0f\xc0\x76\x07\xf7\xcc\x5f\x1d\xe4\xd6\xb3\xb4\x86\x86\xf7\x68\xb7\xf2\x40\x69\xfa\xa6\xc6\x8a\xaa\xbd\x94\x7c\x42\x72\xe4\x2b\xf4\xfc\xe7\xd7\x75\x76\x8b\x78\x97\x05\xff\x95\x14\x5a\x43\x0e\x93\xa1\xf8\x4b\xc9\x35\x19\x1a\x49\xb1\x24\x63\x86\x96\xd2\x92\x6f\xc1\x17\x5f\xdd\x92\xf9\xcf\x3a\xa0\x76\xb1\x2a\xa4\x54\xb9\xca\x23\x0b\x6d\x5b\xa9\x10\x79\xf1\x91\x52\xec\xa3\x85\x2c\x62\x14\x04\x8f\xfe\x28\xaa\x44\xcd\x15\x55\x46\xed\x89\x05\xb2\x02\xe7\xe8\x7e\x84\xee\x75\x6a\x62\x81\x5d\xa1\x35\xa1\x0f\xf3\xe2\x0f\xc9\xd4\xe3\x93\xff\xb2\x5f\x37\xd6\x45\x91\xfe\x2b\xe2\xa7\x5c\xe9\x86\x67\x82\x84\x66\x50\xec\x6d\xb7\x64\x56\x63\x5e\x28\x4d\xc6\x24\x47\xff\x64\xf6\x04\xd7\xc5\x1a\x91\x5a\x6a\xd8\xcf\xe3\x71\x12\xd7\x88\x70\xda\x4f\x5b\xeb\x08\x46\x91\x8c\x3f\xa3\x04\xff\x93\xdc\xec\x03\xda\x14\x07\x4c\xd0\x3d\xeb\x09\x4c\x19\xa2\x1e\xe4\x6d\x4d\x27\x86\x3d\x34\xa8\x89\x7c\x93\x08\x52\xc6\x55\x45\xc6\xf6\x19\xef\x6b\x5e\xd5\x7c\x77\x19\x8b\x68\x38\x30\xea\x5f\x5c\x07\xb7\x6f\xdd\x2a\x3c\xa2\xcb\x1f\x38\x77\x62\x18\x61\xa5\x44\xae\xa5\x94\x4d\xef\x82\x16\xce\xf5\x8d\x2f\xc4\xff\x37\x1b\x91\xd3\x49\xbc\x46\xe5\x30\x54\x6e\x6b\x6a\x86\x14\xe2\x25\x02\x3f\xdc\xca\x92\xe1\x39\xe6\x54\x1a\x59\xd6\x2c\x26\x39\xc7\xf0\xa6\x44\xf9\x66\x03\xea\xaa\x42\x54\x40\x6e\xb7\x9d\xfa\xbf\x23\x52\xf7\x83\x35\x32\xf1\x64\x8e\x4a\x65\x2c\x3f\x81\xb1\xb9\x99\x6d\x7c\x6f\x9a\x5d\xac\xec\x85\xd8\xe0\x87\xc7\x72\xdf\xe8\xad\xd3\xad\xab\x7f\x85\x4d\x09\xcb\x59\x1d\x8a\xad\xae\x23\x03\x59\x64\x18\x71\x45\x63\x5c\x27\x64\xbd\x86\x67\xa8\x2c\xd6\x05\x47\xb9\x88\x77\x12\xa3\xfe\xd3\xa6\xcc\xc7\xc3\xf1\xf0\xe4\x4f\x7f\x36\xdf\x59\xb9\x82\xaa\x01\x79\xd5\x1b\x5a\xe3\x21\x98\x4c\x3f\x80\x1a\x17\x5c\x3d\x41\x62\xff\xa0\x21\x80\x38\x07\x57\xaf\xc5\x88\xd9\xe9\x95\xf1\x26\xe9\xf4\x7b\x5f\xf6\xb4\x2a\x28\xd7\x9f\x5c\x92\xa5\x9d\xae\x06\xf4\xad\x85\x51\x1a\x36\xdc\x31\x83\xb1\x91\x63\x73\xd8\xde\x8a\x2c\x99\xfc\xaf\x02\xda\x67\x8a\xce\xac\xec\x55\xf7\x8e\xd4\xca\x8b\x45\x37\x6c\xf4\x16\xb2\x69\x2b\x0d\xad\x1b\x8e\xf6\x74\xc0\x3a\xbe\x62\x46\xc9\xd9\x50\xa3\x91\x50\x30\xb0\xdd\x9e\x4f\xe6\xd7\x90\xdd\x9e\x09\xe2\x0b\x1e\xc8\x20\x2b\x84\x73\xf6\x5e\xba\x3d\xcb\xb3\x0f\xdb\x08\x4e\xfa\x90\x2f\x81\x5c\x50\x81\x8b\xe4\xce\x9d\xc3\x00\x36\x02\x9c\xe3\xd1\x78\xbf\x28\x40\x4f\x7c\x4d\x6e\x11\xde\xe9\xe2\xa2\xee\x4d\x47\x69\x91\x88\xc1\x89\x13\xe6\x1c\x66\xb7\x72\x84\xdc\xf6\x42\x5c\x2d\x0f\x13\x3f\x76\x30\x8b\x4a\x2d\xa2\xe6\x99\x03\xea\xd4\x38\x5b\x70\xf3\xb9\x33\xa4\x8d\x4a\x34\xa8\xf8\xed\x80\x08\x8e\xef\x11\xb0\x36\xa1\xaa\xbd\x20\x2f\x54\xbd\x58\xc3\xa5\x01\x27\x7f\x86\x00\x37\x1b\xa1\xb0\x68\x24\xad\x10\xce\x47\xa7\x94\xc2\x87\xed\xd6\x0f\x57\x35\x40\x20\xb9\x00\x96\x52\xcb\x00\x68\x08\x7e\x40\xa5\x0c\x6e\xa5\x8a\xef\x46\x6f\x12\x23\x31\x6c\xb7\xc3\xcd\x06\x95\x0c\x6d\xb7\x9b\x0d\xc2\x79\x74\x4c\xb2\xd9\x34\x73\x6d\xb7\x49\x90\xb4\xf0\xf0\x2f\x3e\x2b\xc4\x7c\x62\x03\x63\x64\xd2\xac\x4a\x0d\x20\x49\xfa\xd9\xb2\xd9\x80\x3b\x61\xe5\x02\x43\xb7\x5e\x56\x14\x26\x2a\x99\x54\x75\xa7\xe0\x86\x8b\x3b\x0e\xbb\xb8\x56\xfe\x9e\x9f\x73\x11\xab\xd0\x33\x88\xfb\xe4\xb9\xb8\x63\x25\xff\x16\xe0\x74\x3a\x6d\x34\x51\x98\xca\xa8\xd2\x8a\x5d\x78\x3a\xf9\x9b\x86\x45\xf8\x4e\xff\x8e\xc0\x9e\xfe\x36\xff\xc7\xec\xfc\xd7\x8b\xf7\xef\xcc\x11\xc6\xd3\xf0\x38\x23\x36\x41\x0f\x43\xf0\x83\x12\x9a\x52\x53\x63\x29\x20\x20\x6d\xa9\x9f\x42\x39\xd4\x98\x24\x09\x01\x69\xbb\x2d\xb0\xeb\x88\xa6\x55\x0c\xf5\x97\xaf\x0d\x71\x25\xed\x3c\xd6\xde\xcb\x18\x5d\x16\xf8\xf6\x23\xa4\x2c\x4c\x9c\x47\x5b\x2f\x55\xb1\xd9\x93\xcb\xf7\xbf\xfe\xe3\xd7\xd9\xfb\x0f\xd3\x98\x53\x0f\xd5\x13\x66\xef\x27\xe7\xf3\xb9\x6f\xbd\xdc\x2c\xd6\x53\xb1\x8f\xa4\xac\xd7\x81\x74\xde\x66\x04\x1a\x5d\x91\x1a\x73\x11\x57\xea\x01\x61\x16\x28\x2f\x3d\xba\x60\xf3\x07\xc6\xd1\x3a\x22\x44\x41\xe4\xe8\x2d\x61\x7c\xbb\x4d\x37\x9b\xd1\x84\x60\x0e\x0b\x8c\x68\x50\xa9\x14\xaf\x84\xf9\x88\x20\x8b\x24\xa4\x47\x77\x8a\xd0\x23\x3f\xd1\x75\x1c\xd8\x91\xb0\x73\x92\x63\xc2\x22\x46\x08\x0b\xe5\xc4\x1d\x79\x51\x45\x8a\xbd\x01\xed\x01\xb2\xa4\xe8\x1d\x51\x61\x1c\x70\x41\x3d\x53\x9a\x9c\xdf\x73\x0a\x05\x8d\xbb\x64\x16\xd8\x83\xed\xd0\x2b\x58\x45\x04\x18\x96\x97\x18\x64\xba\x47\xad\xe5\x21\x76\x08\x0f\x59\x9d\xe6\x39\x45\x8c\x35\xe0\xcd\x3e\x08\x39\x91\x47\x6d\x8e\x67\xf0\xad\x89\x01\xc3\x5c\x7b\x3a\xde\x29\xa1\xdc\xa8\x66\xf7\x48\x64\x24\x40\x63\x1b\xc7\x55\xe2\x54\x68\x71\x4c\xdf\xe3\x2e\x45\x4c\xb1\xd9\x80\xd1\xeb\xe6\xd0\x69\xbb\x15\xb2\x0b\x5a\x0d\xa0\x6d\x56\xa7\xe7\x11\x11\x45\x54\xff\xbb\x88\x69\x4a\x8b\xbb\xa2\x44\x4b\x94\x77\xc6\xac\x7b\xe6\x11\xb8\x6f\x21\x4e\x4b\x3f\xc0\x31\x3b\xc8\x6f\xbb\x67\x54\xd4\xe9\x24\xd2\xa1\xa8\xd0\xce\x0b\x5e\x58\xec\x51\x01\xe0\x5b\xc8\x0c\x71\xbc\x70\xd9\xdf\xa5\x29\x0d\x54\x53\x8c\x94\x93\x45\x42\xd1\x10\xf3\xed\x50\x3f\x90\x25\xc8\xf4\xe4\x45\x88\xfb\x76\x8a\x77\x3e\x11\x56\x52\x17\xa5\xf7\x2b\x46\x76\xbd\x26\xad\x7a\x36\xcf\x9c\x68\xbc\xeb\xbc\x98\x10\xbc\x28\x96\x35\x75\x13\x78\x0d\xa8\x3b\x28\xde\x22\x58\xf2\xd5\xc3\x54\xf5\x51\x74\x5a\xe1\x75\x70\xf8\x16\xa9\x69\x1b\xe9\x1b\xab\x1b\x4b\x6c\xc5\x72\x29\x66\x05\x45\xf9\x44\xb8\xc0\x60\xa0\x17\xa9\x93\xec\x15\xe8\xb5\x6a\x12\xb4\x0e\xc9\x25\x81\x79\xa3\x17\x21\xfb\x12\x08\x0a\xdb\xed\xbc\x5f\x42\x63\x8e\x10\x34\xe8\x11\x07\x32\x59\xe8\x08\x1b\xbf\xb4\xed\x44\x00\x8d\x49\x6b\x97\x51\x76\x6c\xd9\x5f\xd3\x3d\x43\xe1\x1c\x5e\x38\x62\x8e\xd7\x77\x4d\xc5\x8f\x24\xbf\xc1\x9d\xe4\x17\x02\xfa\xe4\xeb\x67\xf5\x06\xc1\x8e\x31\x32\xa7\xdb\x55\x04\x0a\x76\xea\xd9\x85\xb2\x96\x95\x66\x15\xe4\x07\x5d\x78\x91\xe4\xa5\xaf\x34\xbd\xa3\xa9\xf1\xd4\x00\x6e\x66\x99\x52\xb4\x28\xee\x05\x7c\x45\x0b\xcc\x17\x20\x69\x70\xff\x3b\x4b\x6c\x9c\x6e\xc1\x65\x64\x7a\x41\xa3\xca\x22\xdb\xe9\x02\x73\x04\x1d\xd5\x44\x98\x96\x45\x91\x79\x9d\x05\xd1\x5e\x3e\x77\xa9\x3b\xd1\xca\x78\xcf\x6b\x7c\x79\x92\x48\xc2\x75\xcb\xb0\x38\xda\x16\x10\x91\x44\xec\xcd\xbc\x4e\xd1\x9a\xf1\x8e\x04\x1f\xc3\xc3\xef\xd2\xc4\xf3\x14\x0a\x65\x38\xf2\x14\xd2\x84\xa1\x54\x26\xa9\x9d\x6c\x06\x71\x4e\xd6\x0c\x1c\x14\x9c\xc0\x6e\x96\x97\x9e\x87\xee\x5d\xc8\x93\xc4\x6f\xd7\x65\x63\x25\x4b\x2d\xe0\x2b\xd7\xee\xed\xd6\x8e\x76\xef\xb5\x3c\x76\x58\xeb\xf0\xb1\x3f\x72\x71\xc6\x76\xa5\x6e\xa3\x7a\xec\x9a\x4e\x21\x37\xcb\x3e\x8b\x71\x20\x39\x7b\x37\x57\xa9\xd3\x17\xfb\x88\xff\xbb\xa8\x73\xf3\xcf\xc7\x04\x69\x11\xec\x56\xa1\x55\xaf\x3a\x71\xa6\xfb\x36\x1a\xee\xba\xc0\xef\x40\xb8\xa9\x36\x23\xd7\xed\x02\x4e\x6b\x24\xf5\x71\x64\x1a\xeb\xe7\xe9\xbb\x7b\x4a\xf0\x1d\x34\x3e\xa0\x70\xb1\x16\xbd\x67\x72\xd2\x8d\x76\x4f\x44\x34\x67\xce\x64\x74\x78\x06\x23\xde\x44\x82\xd9\xe7\x44\x5e\x0a\x07\xf6\x28\x9e\x1f\x36\xa4\x7a\x05\x06\xbb\x3d\xf1\x02\x2f\x75\x42\xed\xa4\x18\xbd\x7b\x4e\x43\xb9\x21\xa3\x2a\xd2\x9c\xeb\x93\x74\x3f\x81\x4b\x26\x45\x4e\x2f\x04\xbf\x93\xf1\x48\xfe\xef\x68\x1c\x28\x6f\x47\xaa\x32\xdd\x68\xa3\x11\x40\x77\x9c\xf9\x69\x64\x2c\x89\x4c\x2e\x2a\xb3\xb9\x88\x67\x95\x5f\x89\x7f\x43\xc9\xda\x88\x58\xad\x9d\xec\x01\x5f\x93\x18\xa8\x9d\x50\xee\x0a\x0d\x1d\x79\x06\x52\x5b\x33\xad\xfa\x58\x65\x17\xb9\xcb\x0a\xef\x20\x78\x18\xdd\x00\xa1\x63\x4d\xa5\xb4\x25\x64\xbc\xc8\xba\xbd\x5f\xe0\x65\x9a\x9a\xa6\xa0\x53\xe2\xa7\xb9\x06\x2b\xaf\xdd\x63\x77\x76\xeb\x8e\xed\x1a\xa5\x78\xe8\x77\x30\x9a\x67\x2b\xb4\x46\x20\x29\xba\x1b\x1d\x56\xf8\xad\xde\xab\xae\x8f\x50\xbf\x87\xd1\x1d\xab\x76\xdd\xc5\x42\x51\xd9\x74\xa6\xda\xe2\x37\xce\xe2\xed\x06\x56\x57\x1f\x3d\x40\x3b\x13\xb1\x36\x68\x70\x03\x74\x94\x3b\x84\xb5\x2d\xf5\x43\x73\x4d\x71\x6d\xf2\x4e\xd1\xa2\x4b\xbe\x08\x61\xf3\xd7\x19\x5c\x9b\xbf\x22\x5b\xdd\x85\xea\x60\x24\x9b\x92\xce\x28\x2c\x70\x81\x97\xaa\x53\x4b\x91\xa1\x75\x29\x49\xa5\xcb\x19\x9a\x7d\x30\x7f\x1e\x5b\xc6\xac\xc3\x63\x76\x4d\x80\xe4\x22\x2f\x91\xd1\x3c\x23\x94\xcc\x78\xa4\x52\x41\x13\x0d\x25\x8c\xfd\x9d\x60\xd4\x4c\xd9\xbd\x52\x65\x82\xc9\x0a\x65\xb7\x6e\x71\x42\x57\x10\xae\x57\x14\xb1\x15\x29\x65\x65\xe9\xc4\x56\x28\xc9\xc4\x3b\xd9\x8f\x27\x89\x50\x43\x9a\xa7\xae\x41\x49\xae\x21\x5d\x86\x7b\xac\xbc\xb2\x9d\x81\xae\x31\x68\x60\xbb\x4d\xa3\x1a\x1a\xdb\x98\x4d\xa0\xa1\x51\x11\xca\x63\xb5\x3d\x73\x46\xc8\x57\x8e\x89\xf3\xcf\x67\x1d\xfe\xab\x91\x86\x04\x2c\xe0\x0f\x78\x15\xe4\x66\x97\xee\x1a\x32\x69\x5a\x4c\xbf\xa5\xdf\xb2\x9c\xbb\x62\xe7\x28\x78\x8e\x62\xba\x0f\x3b\x5e\x72\x1a\x5f\xe5\xf8\xfd\xfd\x9b\x8d\xda\xd9\x8c\x32\xe1\xf5\x42\xf7\x27\x26\x70\xc3\xae\xeb\x76\x7e\x19\x6c\x01\x8d\x7a\x4f\xd3\x11\xec\xed\x22\x43\x5d\xbd\x16\xe7\x5c\x80\x30\xe7\x3a\x3c\x6a\xe2\x50\xe9\xe4\x91\xf9\x62\xe0\xe0\x6a\x3e\xbf\x34\x78\xd5\x38\xd9\xef\x27\x0b\x4f\x0b\xa2\xa6\xbb\x0f\xf4\xb9\x64\xf8\x15\x74\xb7\x33\xf5\xdb\xc6\x2f\x91\x76\xdf\x3d\x37\xb0\xbf\x61\xef\x1f\xfa\x76\x6d\xa0\x98\x69\x77\x11\x2b\x87\x63\xe1\x09\xb6\x57\xcb\x41\x4d\xbc\x64\x81\x1b\xaf\x42\x07\xdf\x9c\xd3\xe2\xa6\xe6\x6a\xc1\x91\xc3\xb0\x86\x98\x5d\x64\x00\x2b\xd5\x14\xee\xca\x3f\x5c\xd9\x7a\x27\x20\xce\xfe\x61\xba\x57\xef\xf9\x3b\xc8\xeb\x68\x1e\xba\xc2\xf2\x75\xe5\xd9\xfa\x73\xf9\x7a\x42\xc8\x6d\x81\xe6\xbc\xc8\x6e\x0b\x8c\x18\x6b\xe3\x07\xb1\x2a\x5b\xba\x70\x21\xeb\xa7\x0f\x89\xc5\x96\x60\x59\x79\x03\xf6\x48\x7b\x63\xc9\x94\xbe\xc3\xdb\x5a\x0b\xd0\x29\x77\xe8\x02\x70\xdb\xbb\xdb\x1e\x6a\xed\x8c\x85\xb7\xfe\x18\x07\xa0\xe3\x56\x2b\x18\x23\x19\xd8\x95\x96\x07\xba\x81\x8d\x2e\x39\xd9\x0a\x32\xa1\x04\xff\x95\xdc\x30\xbf\xdb\x55\x44\x51\xd8\xb9\x70\xb1\xeb\xba\x45\x34\x11\xde\xf3\xaa\xc5\x1e\xcd\xfb\x3d\xd7\x2c\xbc\x56\xad\x5d\x57\x2c\xbe\xcd\x05\x8b\x47\x5c\xaf\x88\x9c\x43\x9a\x96\x34\x7e\xad\x22\x6a\x65\xed\xb0\xce\xde\x2c\x5a\xbe\xee\x89\xd7\xce\x8b\x14\x7b\x5e\xa3\xe8\xbd\xf4\x12\xee\x13\xd8\xe3\xe2\x8b\xc9\xd3\x04\x65\x2c\x9d\xd5\xf8\x1a\xb2\xdb\x30\xa8\x7d\x29\x23\x08\x62\xa6\xb6\x11\x73\x7d\x4a\x71\x7b\x78\x10\x06\x01\x8a\x96\xcc\x3c\xc3\xdc\x11\xd5\x5b\x83\x21\xc5\x29\xfc\xca\x52\x81\x24\xe2\x07\x80\x6f\x39\xdb\xeb\x16\xf1\x11\xc9\x23\xd0\x9d\x66\x19\xa9\x31\xbf\xc8\x77\x60\xd4\xab\x3c\xea\xc1\xdc\x76\x7c\x4d\x2e\x3f\xcc\xaf\xcf\x67\x49\xa4\x2f\x00\x34\xe9\x44\xf0\x5d\xe8\xa9\xff\x2c\x14\x2d\x3d\x5f\xb7\xc2\xc2\x4a\x4a\xb2\x64\xe9\x84\x22\xc8\x51\xdb\x4f\x15\xf1\xdb\x36\xe8\x9c\x53\x04\xd7\xbd\xb0\xd3\x9a\x5f\x92\xe5\xf9\x1d\xc2\x9c\x05\x9b\x28\x76\xa9\x78\xa4\x8f\xa9\x51\x2e\x39\x49\xcf\xa5\x9d\xa1\xea\x19\xe9\xd3\x0a\x90\x08\x2c\x87\xf2\xfa\x71\x7a\x04\xbf\xb2\xe6\x4a\xce\x8f\xfe\x35\x1c\xf5\xe7\x5f\x28\x9d\x7f\x1d\xcb\x03\x67\x21\x9d\xb6\x18\x67\xc8\x20\x49\xe3\x8c\x73\xc3\xf4\xa8\xb7\x30\x22\x01\x3f\x0e\xd0\xce\xba\xbd\x18\x16\x73\xd8\xd1\x1b\x64\x6e\x15\xaf\x75\xfd\xbb\xab\x75\x49\x83\xcc\x68\x23\xb0\x3a\x47\x0f\x45\x18\x64\x74\x9d\x1e\x8a\xa8\xc4\x74\x6b\xde\xe5\xad\x95\x7e\x60\xc0\xf4\x5c\xcd\x6a\x48\x0d\x9e\xdb\xf7\xde\xc8\x32\x4a\x1a\x7f\x1a\x5b\x55\x28\xef\xca\x5c\xf2\xf7\xa2\x7a\x53\x94\x01\x6d\x48\x3e\x63\xbf\x98\x33\xa8\x19\x02\x8c\xd3\x22\xe3\x83\x5f\x5c\xdf\x7b\x07\x29\x80\x5f\x19\x78\x05\x28\xfa\xbd\x2e\x28\x3a\x18\xc0\xaf\xec\x90\xe5\xb7\x83\x97\x41\x60\x94\x09\x60\x8c\xbe\x8a\x61\xa3\xf3\xc9\xfc\x60\xb3\x86\xf7\x33\xc4\x69\x81\x58\x7a\x3c\xde\x86\x87\x09\xe5\x37\xc6\x4d\x4a\x52\xe7\xbf\x41\x9e\xad\x2e\xc9\x92\x1d\x84\xc7\x68\xb3\x0f\x5e\x81\x41\xc0\xba\x7b\x6b\x89\x18\x23\x3d\xbb\xdc\x0b\x02\x95\x65\x70\xcc\x16\x55\x90\x0c\x7e\x09\xf6\x4c\xf6\x20\xd6\x57\x08\x3d\xbc\xc6\x5d\x81\x28\x5a\x89\x80\x5b\x3d\x18\x82\x45\x1b\x77\x59\xfe\x01\x95\x7f\x37\x64\x07\xa9\xe2\x95\xe0\x61\xa3\xf6\x83\xd4\xa1\xb7\xab\xe0\xf5\xb4\x87\x88\xa5\x0c\xc3\x1c\x0a\xd6\x74\xd4\xb4\x83\x74\x30\x70\xa5\xeb\x35\x3d\xa1\xfb\x4a\x24\x8e\xcd\x6e\x03\xaf\xc0\x42\xef\xe4\x03\x24\x6c\xe5\x10\x64\x04\x73\x74\xcf\x5f\x7a\xfc\x91\xb3\x48\x75\x51\x6d\xf6\xe0\x15\x90\x43\x46\xfa\xf7\x88\xa2\xaa\x84\x19\x3a\x38\xfa\x8f\x7f\x3b\xf8\xfc\x39\xff\xe9\xe5\x2f\x47\xcb\x61\x87\x7f\x2d\xb4\x70\x08\x72\x94\x45\x70\x03\x40\x11\xaf\x29\x06\xea\xac\x7f\xb4\xa0\x64\x3d\x59\x41\x2a\x76\xe6\x81\x18\xe6\x29\xaf\xf8\x4f\x60\x1f\x34\x84\xaa\x36\x8c\x80\xa8\x93\xe6\x1f\x8c\x43\xca\x51\xfe\xfa\x21\x05\x03\x61\xa1\x06\xc3\x18\xa4\xad\x3f\xa9\xab\x4f\x9f\x14\x2b\x74\xc3\xc9\x97\x28\x1a\xbd\xd5\xd2\xe6\x1f\x71\x40\xe1\x9a\x53\x70\x1c\x05\x20\x77\x88\xd2\x22\x47\x2c\x8d\x2f\x4f\x21\xd2\x8d\x59\xef\xbb\x01\x9f\xfa\x06\x00\xa9\xde\x18\xae\x51\x6a\x2d\x6a\xd8\x08\x3e\xfd\x04\x06\x6c\x35\x18\x82\xc1\x61\x36\x68\x9f\x0a\x65\xed\x43\xfb\x25\xf6\x32\x38\x6a\x1b\x15\x2a\xbb\x45\x5f\xc1\x2b\x70\x05\xf9\x6a\xb4\x28\x09\xa1\x07\xf2\x9f\x54\xb6\x7b\x1c\xbc\xfc\xf1\x78\x3c\x1e\x8f\xc3\x3a\x51\x92\xe5\x81\xb5\x24\xf0\x13\x18\xa4\x03\xf0\x53\x6b\x5e\x7e\x02\x83\x23\xa1\x07\x72\x96\x57\xe2\x8d\x9c\xee\x27\x30\x58\xb3\x66\xa1\xf2\xb1\xa5\xf9\x86\x92\x23\x4a\xa3\xda\x2d\x65\xc1\x48\x89\x46\x82\x90\x01\xa2\xf4\x64\x30\x04\x62\x44\x90\x5a\xf1\x87\x21\xae\xdd\xd5\x41\x3b\xc5\x4b\xb0\x11\xce\x61\x44\x55\x7a\x74\xa0\xb4\xbc\xdd\xb8\xa3\x9c\x60\xf4\x52\x18\x11\x41\xfa\xde\x7b\xc6\x67\x78\x33\xa1\x64\xdb\x1a\x31\x06\x97\x68\x08\xb2\x9b\x1e\xcb\xc0\x64\x5c\x26\x8c\xb4\x60\xe2\xa1\x60\xd4\x81\xf0\x44\x67\x90\xa3\x83\x97\x2f\x47\x4b\xb5\x9c\x80\x1b\x02\x7b\x6f\xd9\xc6\xc5\x08\xfb\x99\xb6\xbf\xa2\xdb\xa4\x6c\xa2\x45\x05\xcf\x42\x91\xa3\xe2\x49\x44\x63\xd8\x28\xb3\xc3\xce\x96\xe1\x4f\x14\xfa\x2e\x99\xef\xc7\x06\x4d\x9d\x0a\x70\xf7\xda\xd1\x5a\x84\x29\x68\x65\x29\x62\x24\xc6\xe1\xba\x4a\x23\x62\xda\xb1\xa3\xa3\x4c\x07\x8f\x97\x13\x78\x8c\xac\x40\x4c\x5e\xc0\xe1\xb7\x5e\x6a\x9c\xd9\x52\xc0\x95\x91\x2b\x74\xdb\xe9\x26\xb6\x77\x7c\x27\x6b\x16\x78\xac\x90\x3f\x70\x96\xef\x74\xd8\xe8\x90\x63\x47\x55\x4e\x84\x14\xf3\x15\xa1\x5c\x47\x0d\xb3\xba\xa7\x42\xa7\x75\x22\x95\x40\x3b\x23\x79\x23\x82\x1f\x5d\x12\xbc\x6c\x02\x76\x96\xad\x50\x5e\xdb\x5f\x55\x99\xeb\x67\xe7\xf7\x15\x45\xac\xa9\x14\x49\xe2\xf4\x1b\xa7\x27\x49\x9d\x85\x7a\x05\x73\x19\xb3\x47\x03\xfb\x2e\x51\x89\xdc\x48\xbd\xc8\x03\x04\xeb\x63\x57\xe7\xe4\xb6\xd2\x67\x97\x9f\x9b\xfb\xc6\x9f\x93\x14\x7c\x96\x63\x75\x6d\x1c\x6c\xb7\x9f\x93\x21\xf8\x9c\x68\x63\xde\x01\xe8\xeb\x84\x12\xc0\x90\x71\xa8\x26\x1b\x10\x91\x4a\xbb\xa6\x88\xae\x0b\xc6\x42\xf9\x19\x70\x13\x34\x03\x36\x24\x35\x60\xd7\x56\xb3\xb6\x65\x51\x25\xea\xe9\x05\xbe\x23\xb7\x28\xf4\x9d\x10\x2b\x59\x03\x4f\xe4\xbb\x51\x36\x15\x93\x4a\x07\xc8\x9c\x42\xa9\xa9\x2a\x32\x8f\x96\x68\xa2\x1d\x84\x9e\x46\x1b\x13\xf7\x6d\x9c\x70\xdd\x3b\xf8\x45\x51\x75\x6b\x5d\x15\x42\xde\x42\x76\x3e\x31\xbf\x6c\x25\x89\x7a\x4f\xad\xaa\x43\xdf\xe7\x41\xbd\xa2\x4a\xcd\x0e\x11\x64\x5c\x7e\xdc\xc0\xec\x95\x78\x0a\x8e\x93\xe7\xe1\xf8\x8a\x9e\x4f\x87\xc4\xf1\x1c\x3a\x50\x7d\x98\x21\xcc\x29\x2c\x9f\x45\x0a\xaa\x9f\xbf\x1c\x58\x1d\x62\x42\xf9\xea\xd9\xf2\x81\xd5\x21\x23\xf5\xb7\x46\x24\xb9\xac\xf1\x7c\xe9\xfb\x4c\x42\xe8\x9b\x2a\x9d\x52\x8b\x37\x81\xaf\x25\x06\x3e\x2d\x30\x33\xc0\x9a\xea\x91\xd9\xfa\x65\x6d\x11\x0d\x10\xfc\xe6\x40\xd8\x91\x3c\xff\x5b\x03\xc6\x57\x1f\xbd\xfb\x21\x5e\x77\xe4\x8b\xc6\x36\xec\xe6\x9b\xf3\x91\x8c\x86\x6b\xaa\x29\x60\x5f\x46\x78\xdf\xd1\xf0\x0a\xa8\x2f\xb4\x95\xea\x15\xcb\x63\x26\x72\xd9\x17\xc0\x3b\xf4\x98\xa6\x2d\x67\x94\x2b\xce\xc7\x76\xac\x2b\x30\xe1\x2f\x7a\xb8\xc7\x7d\x11\xf1\xef\x75\xd4\x17\x3d\x37\x72\x0e\xb3\xa2\xc5\x60\xf7\xa0\xcd\x7a\xed\x5e\xfc\xe9\xaf\x3b\xdb\xa7\x80\xee\x3c\xc6\x99\xa0\xf3\x4a\x9d\x07\x39\x7e\xcf\x81\xd9\xab\xa0\x6b\xf8\xca\xf0\xa9\x5b\xe3\x01\xbb\x58\x20\x76\x4a\x18\x3e\x23\xb4\x02\x2e\xfb\x7c\xd0\xba\xf9\xe4\x83\xc5\x3f\x89\xf7\xed\xbf\x76\xd7\x5b\xf8\x4f\x90\xea\x64\x2d\x09\xcc\x6f\xda\x4e\x56\xd5\x5a\x7d\x83\x22\xc7\x7e\x91\x31\xca\x20\x20\xda\x34\x50\xb0\x37\x94\xac\x83\x3d\xb1\xbb\xb1\xcd\x5c\x5c\xbf\x15\x7c\xb5\x07\xae\xec\x64\x27\xf1\xd9\x49\x7a\x5a\xf3\x15\xa1\xc5\x1f\x28\xd8\xe5\xed\x8d\x0a\x75\x89\x18\xe7\x18\x41\xbe\xfe\x18\x40\xe3\x3c\x71\x6e\xd5\x05\x95\xd8\xf4\x60\x3b\xec\xb1\xf9\x19\x2f\xff\xeb\x58\xb6\xcd\x99\xff\x9c\xa6\xfa\x4b\x75\xda\xe8\x9c\xa1\x12\x09\x3d\x69\x1b\x44\x92\x19\xe2\xb0\xc0\x3b\x8c\x92\xfc\xc0\xf4\x84\x60\x4e\x55\xc3\x9a\xdb\x01\x9c\x5c\x43\xe7\x62\xf6\xa6\xf9\x68\x4c\xc2\xe4\x57\x0a\x84\x8d\x6d\xdb\x74\xf4\xc7\xf1\x80\x1d\x01\x68\x78\x58\x55\x26\x70\x8f\xf3\x0a\xb1\xcd\xe0\xda\xff\x07\x00\x00\xff\xff\xea\xa6\x06\x03\x24\x60\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
	}},
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

