// Code generated by go-bindata.
// sources:
// bundled_scripts/idrac.py
// bundled_scripts/idrac.toml
// bundled_scripts/ilo.py
// bundled_scripts/ilo.toml
// DO NOT EDIT!

package main

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

var _idracPy = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x3b\x7b\x53\xe2\xc8\xb7\xff\xf3\x29\xfa\x66\x6a\x2b\x50\xab\x01\x7c\xcc\xce\x70\x87\xd9\x42\x88\x0e\xb5\x82\x16\xe0\xee\x6c\x79\xad\x54\x08\x8d\x66\xcd\x83\x4d\x07\x91\x6b\xf9\xdd\x7f\xe7\x74\xe7\xd1\x79\xa9\x38\xbf\x61\xaa\x4c\xd2\x7d\x5e\x7d\x5e\x7d\x4e\x27\xf3\xe1\x7f\x9a\x6b\x16\x34\xe7\xb6\xd7\xa4\xde\x03\x59\x6d\xc3\x3b\xdf\xab\xd5\x6c\x77\xe5\x07\x21\xf9\x87\xc1\x43\x74\xef\xb3\xf8\x8e\x6d\x93\xdb\xf5\xda\x5e\xd4\x96\x81\xef\x12\xcb\x5f\x6d\x49\x34\xba\xa0\x74\x85\xcf\x62\xe6\xd1\x75\x34\x1a\x06\x94\xc6\xd3\xba\x43\x5d\xea\x85\x33\x1c\x32\x19\xd1\x67\x09\xbf\x80\xfe\xbb\xa6\x2c\x64\x02\x31\x7e\xd2\x56\xa6\x75\x6f\xde\x52\xa6\xad\x03\xc7\xb1\xe7\x87\x1a\x7d\xb4\xe8\x2a\xb4\x7d\x8f\xc5\x34\x87\x1e\xa3\xd6\x3a\xa0\x13\x81\xf3\x97\x19\x78\xb6\x77\x5b\xab\x4d\xfb\xdf\xf4\x51\x8f\x74\x89\x72\x17\x86\xab\x4e\xb3\xc9\xac\x3b\xea\x9a\x4c\x5b\xb8\xe1\x52\xf3\x83\xdb\xe6\x66\x4e\xdd\xe6\x86\x59\xb6\xdb\x6c\x37\xe1\xef\xbe\x80\x68\x1e\x28\xb5\xef\xa3\xf3\xf1\xd4\x98\x22\xf6\x53\x84\xbe\xd9\x6c\xb4\xcd\x21\x47\x3c\x68\xb5\x0e\x9b\xad\xe3\x26\xf3\xcd\xd5\x3e\x28\x8f\x3a\xfe\x8a\x3e\xc7\x58\x7f\x4d\xf5\xb1\x8c\x18\xf3\x05\x6d\x20\x82\x60\xcd\x90\xc8\x51\xb3\xf5\x19\x94\xbf\x76\x69\x60\xe2\x9a\x24\x12\xa3\x5e\x29\x8d\xbc\xec\xae\xe9\x81\xec\xfc\xaa\x3d\xb2\x45\x42\x60\xdc\x36\x4e\x7a\x53\xbd\x94\x04\x75\x1c\xcd\xf2\xdd\xea\xe5\x37\x7f\x61\x40\xe8\x03\x39\xa3\x1e\x0d\x6c\x8b\x70\xf2\xfb\x56\x60\x2e\x43\xba\x20\xb8\x06\xe2\x52\xc6\xc0\x2e\xb5\xe9\x45\xef\xd2\xd0\xc7\x57\x23\x21\xb2\x31\xd3\x47\x97\xe7\xbd\x19\x32\x56\x55\xf5\xcb\xef\xb0\x66\xf2\x40\x03\x06\x8b\xeb\x2a\x6d\xad\xa5\xfc\xfe\xb5\xf6\x85\x75\xf4\x48\x67\xe8\x21\x1e\xeb\xb0\xae\xf2\x46\x25\x2b\x11\xc6\x86\x99\xdd\xbc\x5d\x4b\xf5\xfb\xa9\x69\x2e\x16\x01\x48\x0b\x3e\x91\x22\xc3\x7a\x0a\xe8\xaf\xa8\x36\x45\xa6\x45\xdc\xd7\x4d\xab\x7c\xad\x11\x02\x2b\xff\x46\xcd\x05\x0d\xf0\x01\x1e\x61\x15\x9d\x9e\x85\xd3\x84\x75\xdc\x35\x0b\xaf\x3c\x98\x64\xa1\xe9\x2d\xba\x4a\x18\xac\xa9\xf2\x75\x67\x46\x4d\x3d\xba\xa7\x5f\x9a\x29\x7d\x89\xe1\xcc\xaf\x64\xf6\x04\x6b\x05\xab\x62\x88\x1a\x10\x70\xcf\x82\xc2\xcc\x4f\xb0\x61\xbe\x33\xa1\xcc\x5f\x07\x16\xbd\x9a\x0c\xab\x09\x05\x11\x10\x27\x91\xc3\x92\x64\x19\x09\x37\x1a\x0e\x2a\x29\x61\x9a\xe9\x3c\xe1\xdf\x48\x9a\x04\x45\x22\x33\xa1\x2b\x67\x1b\x8b\x19\xeb\x55\x98\xfd\x4d\x0a\x94\x9d\xa4\x19\xf8\x0e\x6d\x9a\x9e\xef\x6d\x5d\x7f\xcd\x22\x25\x46\xc4\x04\xcb\x66\x81\x67\xa4\x9a\x29\x75\xa8\x15\xfa\xc1\x94\x86\x92\x2c\xf2\x0c\x19\x9b\x2e\xed\x2a\x86\x01\xf1\xe6\xc1\x2d\x83\x04\x87\xfa\x62\xd1\x7c\xa2\xaf\x18\x21\x65\x59\xc2\xe0\x4b\x53\x76\x28\xf0\xae\x13\x7f\xb1\x4d\x24\xa2\x5e\x27\xf1\x85\x9c\x38\x17\x90\x42\x5d\xfb\xff\xa9\x2e\xf9\x4d\x0e\x64\x64\x3e\x46\xd9\x9a\xa1\x5f\x3c\x1a\x34\x7a\x4a\x44\x94\x21\x12\x29\xf3\x4c\x51\x44\x21\x15\xde\xc5\x71\xff\xb5\x06\xd9\xa1\x56\x1b\xf5\xfa\xc6\xe5\x44\x3f\x1d\x7e\x37\x4e\xce\x7b\xfd\x3f\xce\x87\xd3\x19\x64\x8e\x6b\x4e\x4d\x3d\x6e\xc1\xbf\x23\x75\x8f\xa8\x87\x87\xc7\xad\x8f\xa7\x78\xd7\x6a\x7d\xfe\xf4\xdb\x47\x71\x87\xbf\xe4\xae\x8f\x77\x07\xad\xa3\xf6\xf1\x21\xde\xb5\x8f\x3e\xb7\x0f\x60\x56\x50\x6a\x81\x99\x05\xa5\x53\xfd\x14\x7e\x1c\xa2\x77\xfa\xf9\x40\xe0\x1f\xb4\x3e\x89\xbb\x81\xde\x1b\x1c\xf4\x05\x5c\x6f\x70\x34\x00\xfc\x9b\xda\x54\x9f\x0c\x7b\xe7\x25\x02\x8e\x7d\x8f\x02\x28\x82\x8f\xfd\x90\xf4\x1e\x4c\xdb\x31\xe7\x0e\xc5\x81\xef\x8f\xfc\x1f\xde\xee\xe3\x0f\x6f\xae\xaf\xbc\x7b\xcf\xdf\x78\x37\xa9\xf8\x62\x09\x42\x48\xa4\x31\x5d\x51\xcb\x5e\xda\x74\x81\x20\x7f\xff\xd1\x6e\xf5\x07\x5c\xd6\x83\xc3\xa3\xe3\x8f\xbf\x7d\xfa\xdc\x12\xcc\x3c\xce\x03\x02\xf9\x84\x92\x53\xdb\x71\x20\x2b\x9f\x6c\xc9\x85\xa6\x6b\x23\x8d\xcb\x5c\x1b\xe8\x7f\x0e\xfb\xba\x31\x1c\x9f\x5e\xc8\x39\xf9\x89\xb3\x52\x5c\x1f\x36\x01\x03\xfd\x4f\xe9\x10\x45\x11\x02\x28\x34\xbc\xa3\x81\x47\x43\x06\x83\xd7\x37\xd1\xa0\x4b\x5d\x3f\xd8\xca\x23\x4b\x7b\x1e\x50\xc3\xba\x33\x3d\x0f\x88\x58\x66\xb0\xc8\x20\xac\x02\xdf\x82\x50\xf1\x83\xcc\xe8\xc2\x66\xf7\x99\x01\x06\x3b\x8b\x09\x32\xac\xdd\x39\x0d\x64\x29\x96\x76\xe0\x6e\x4c\xe0\x10\x6d\x1b\xf2\xdc\xdc\xf6\x59\x6e\xfc\xb9\xa6\xcf\xbe\xe9\x93\xb1\x3e\x2b\x5b\xa7\x69\xc9\xe8\xe5\xcb\x66\x2b\x4a\x17\x38\xb0\x16\xe6\x21\x62\xe0\x65\x71\x9e\x6b\x97\x93\x8b\xbe\x3e\x9d\x5e\x4c\xde\xac\xdf\x98\x11\xf7\x1b\x31\x64\xf9\x90\x56\x92\xa1\xe7\xda\x48\x1f\x5d\x4c\xfe\x7e\x3b\x45\x08\xe1\x2c\xc1\x2c\x8f\xe7\xda\xe9\xf0\x64\xa2\x1b\xfd\x6f\xbd\xf1\x58\x3f\x37\xfa\xbd\xc9\xa0\x48\xfc\x03\xc9\x2f\x72\x4f\xe8\x80\x40\x1e\x26\x1b\xd0\x08\x4c\x91\xb5\xb7\x66\x30\x54\xb7\x07\x93\x5e\x9f\x2c\x7c\xca\x3c\x35\x24\x60\xed\x07\x7b\x41\x49\x78\x67\x43\x41\xe6\x2d\x7d\xb2\xa5\x61\x43\x7b\xdd\x94\xef\xb1\x05\x88\x92\xa8\x7f\x30\x9c\xfe\xf1\x23\x7a\xaa\x72\x3f\xe6\xf8\x61\x16\xb4\x7c\x15\xa8\xb6\xdd\x54\xf2\x5c\xab\xd5\x16\x74\x49\x3c\x3f\x70\x4d\x07\x24\x32\xc0\x3d\x8d\x68\xe3\xa9\x4b\xf7\x8d\x0e\x67\x2c\x8d\xc0\xfa\xa4\x27\x6d\xbd\x5a\xd1\xa0\xde\xd0\x02\xd8\x82\x60\xfb\xa8\xab\x3c\xbd\x74\xd4\x06\xc7\x0b\x68\xb8\x0e\x3c\x19\x01\x18\x57\x17\xd4\x10\x99\x98\xb4\x8c\x8d\xa8\x9b\x59\xbd\xbc\x9e\x6e\x00\x11\xcb\x31\x41\x96\xe1\x22\x30\x2d\x3d\x08\xfc\xa0\xae\xc7\xc5\x78\x24\xf2\xca\xe4\xcc\x22\x38\xd4\x4a\xdd\x9f\xff\x03\x3b\x16\xcc\x73\x00\x5c\xbf\x61\xd8\x9e\x1d\x1a\x46\x1d\xf6\xbc\xe5\x1e\xb9\xf3\x59\xb8\x47\x40\x91\xc1\x1e\xc7\xdf\xf8\xc1\x22\x22\x87\x3f\x04\xd2\x10\x06\x74\x80\x97\xec\x04\xa2\xc1\x04\x5e\xb2\x13\x31\x25\x98\x8c\x6f\x53\x09\x82\xb5\x67\x40\x01\x0c\x7b\xd8\x22\x12\x82\x4b\xcc\xdd\x66\x8f\xc4\x5b\x71\x57\x0d\x7c\x3f\x6c\x2e\x60\xa7\x56\x25\x81\xb2\x25\x52\xdc\x5e\x30\x28\x31\x9e\x9e\xc5\xc6\xa8\x68\x4b\x34\x71\x58\x4f\x64\x6f\x24\xd8\xb7\x58\x52\xc3\xd6\xb8\x30\xb0\xa4\x41\xd1\xe1\xa2\xe1\x9f\x76\x3d\x85\x8a\xaa\x6b\x98\xae\xaa\xaf\x63\x16\x09\x8a\xb0\xbc\xa8\xb1\xba\xa2\xf5\xd1\xa0\x92\x0a\xec\x55\x5d\x6d\xaa\x0d\xf2\x2b\x81\x0b\xfc\x95\x56\x9a\xc1\xcd\x2e\xab\x9b\x7d\xcc\x82\xa2\xb4\xdd\xec\x42\xb2\x00\x89\x06\xe3\x9b\x3c\xab\xb4\x98\xe8\x1e\x1c\x1f\xa7\xb3\xa9\x06\x22\x27\xd6\x67\x1a\xf4\x33\x42\x91\x06\xd4\x16\x0b\x03\x8b\xb7\xec\xaa\x51\x5c\x35\x6f\x03\xb5\x68\x83\x9c\x10\x42\xc5\xdd\xe8\x2a\xc9\xd0\x90\x5c\x35\x65\x29\xfc\x04\x75\x11\xa3\x4a\x3e\xa1\x28\xca\x2c\xd8\x92\x10\x2a\x6b\x40\xc8\x74\x48\x38\xb8\x02\xf6\xdc\x59\xd6\x58\x60\x12\x94\x95\xcc\x4d\x66\x5b\x09\x01\x73\x0d\x7b\xaf\x17\xda\x16\x2f\xc5\x34\x48\x3f\x21\x78\x62\x78\x67\x86\x64\x43\x21\xb3\x60\x5e\x61\xa0\x48\xe8\x9a\xbd\x2d\x30\x61\x98\x88\x78\x72\xc1\x45\x62\xc6\x4e\x48\x41\x7a\x21\x0f\x90\x5f\x16\x60\x1c\x32\x9d\x9e\x13\x8b\x06\x21\xd4\x14\x40\x9a\x6a\xa4\x07\xe8\x4b\xdf\x71\xfc\x0d\x4a\x12\x67\x05\xb2\x81\x22\x02\x9e\xf6\x51\xfc\x84\x12\x17\x91\x8b\x46\xee\x78\x91\x49\xcc\x5b\xd3\xf6\x34\x79\xd9\xa9\xc1\xc0\x59\xd3\x24\x03\x0b\x2e\x18\x29\xab\x7f\x10\xcf\x2c\x2a\x3f\xd6\x45\xb7\x9e\x84\xf7\x5e\x36\xa0\x73\x66\x84\x9c\x6c\x2f\xb7\xdd\x53\xd3\x61\x39\x32\x42\x64\xd6\x7d\xca\x8c\xe2\x4f\xed\xfb\x5e\x08\xda\xde\x9f\x6d\x57\x54\xed\x10\xd5\x5c\xad\x9c\x48\xf5\xbc\xdb\xfc\x15\x9a\x84\xff\x85\x02\x27\x60\x34\xec\x5e\xcd\x4e\xf7\x3f\xa9\x59\xda\xcf\x65\x2e\x6b\x63\x6a\x0f\x49\xa0\xf9\xf7\x9d\x0c\x34\x4c\x04\x1a\x34\x35\xe1\x9a\x41\xda\x81\x9d\xa1\xdb\x25\x47\xad\x76\xa7\x20\x58\x60\xda\x8c\xca\xf9\x55\xe9\xa1\xf2\x29\xde\x6b\x64\xe8\x71\xb3\xf2\x64\x87\xf1\x4b\xc0\xd0\xb1\x56\x34\xa5\x51\x7b\x91\x52\x81\x97\x32\x85\x85\xf2\xb9\x0e\x81\x56\xc6\x87\xa4\xdf\x21\x4f\xcf\xff\xe7\x45\x69\x1f\x1f\x94\xd2\x2c\x93\xf0\xd0\x42\xfa\x18\x26\xe1\xb0\x57\x00\xca\x99\x2a\x95\x90\x2f\x88\x19\x2b\x13\x16\xd7\x25\xea\x13\x7b\xc6\x0e\xa1\x09\xd7\x53\x73\xed\x84\x69\xec\x76\xa3\x63\x98\x02\xaa\x63\xf3\x0d\xe1\xfa\x26\x3f\xe1\x71\xfd\xc6\x79\x43\x88\xd8\xd0\x96\x36\x64\x7a\x89\x6b\xc6\x68\x12\x66\xd6\x24\x39\x5e\x08\x60\x20\x3d\x08\xa0\x80\xa4\x4f\xb6\x27\x93\xd0\xec\x10\xe2\x0d\xc6\xeb\x8d\x9b\x1d\x4d\xa2\x66\x4c\x92\x58\x61\x0f\x1e\x18\x37\x50\xc4\x07\x07\xd5\x17\x4d\x13\x9b\x84\xa8\x7b\xaa\xf6\x8f\x6f\x7b\x75\x69\x2d\x8d\x12\x4b\x55\x18\x2a\xca\xc2\x42\x8d\x51\xf9\x62\xdc\xd2\xd0\x80\xd4\x40\x0d\x4c\x3f\x75\x1b\x17\x64\x88\xed\x22\x88\x52\xe2\x82\x3e\xd8\x96\x98\x07\xc5\x95\xb5\x22\x1c\x8c\x1f\x0c\x76\x49\x86\x82\x26\xef\xcd\xea\xa0\x3f\x1c\x19\xd3\x2d\x0b\xa9\xfb\xa7\x4d\x37\x51\x81\xc3\xcf\x61\x0c\xaf\x0d\xb8\xd9\xc3\xae\x5f\x88\x92\xc3\x10\xf9\xe9\x5f\x7e\x0c\x16\x39\xd9\x73\xd2\x9e\xc6\x6a\x85\xb1\x21\xc0\x33\xb8\xe6\xd1\x0b\x5a\x8e\x1c\x72\x2f\x37\x80\x47\x7e\xc5\x31\xd8\xb0\xd3\xc1\x58\x6a\x31\x12\x97\x6a\xd0\x00\x2c\xb0\xbc\x43\x5d\x70\x37\x35\x1d\xa7\xfe\xaf\x98\x8d\x13\x8a\x00\x4a\x7d\xb3\x98\x26\x86\x1e\x80\x00\x5c\x08\xdb\x03\xdb\x40\xa2\x06\x9f\x84\x0d\x25\x67\xab\x38\x49\xa4\x35\xb2\xd0\x4b\x69\xa4\x47\x5c\xaf\x5b\x37\x22\x7a\x32\x1e\x02\x48\x12\x4e\xb2\x32\xa2\x8e\x4c\x6f\xbd\x34\x2d\xf0\x1a\x1a\xa8\x92\xa3\x35\xb8\x07\x69\xa2\x1e\x49\x0b\x57\x05\x92\x9a\xa5\x41\x31\xad\x28\x12\xf0\x3b\x59\xe3\xaa\xaa\x79\xca\x6a\xcf\x14\xff\x7c\xdf\xaa\x60\x58\xc9\xac\x7f\x07\x89\xd7\x66\x53\x1a\xa0\xa7\xcf\xcc\xdb\x98\x71\x96\x69\x6c\xc6\x2c\x43\x34\x2a\x18\x28\x7f\xae\x90\x1a\x58\x8a\x9f\x6b\x35\x83\xab\xde\x80\xb8\x99\x91\x7c\xc0\x5d\xab\xf9\x7e\x85\xe3\xbc\x4f\xa7\xe7\xf6\x92\x5a\x5b\xcb\xa1\xb8\x63\x06\x50\x36\xd0\xe0\xcf\x88\x68\x66\xbd\x45\x21\xe4\x4e\xfd\x07\x04\x38\x19\x5e\x4c\x23\x8e\x53\x50\xa9\x77\xfb\x1a\xdf\xd4\xb7\x39\xd7\xf4\x51\x6e\x8d\x24\x0c\x39\xa9\x25\xe7\x1f\x85\xa4\x96\x14\x84\x08\x06\x8d\x55\x5d\xaa\xfd\x3e\x10\xaa\xdd\x6a\xa4\xbf\x86\x00\xf4\xc2\x51\xaf\x1f\x9d\x15\x76\xc8\xc9\x49\xa7\xdf\xef\x0c\x06\x1d\x5d\xef\x9c\x9e\x76\x5a\xad\x04\x27\x0c\xb6\x9d\x5c\x49\x6c\x25\x3a\x2a\xd1\xcf\xcb\xbe\x98\xe7\xac\xe6\x37\xde\x54\x57\xf8\x13\x6f\x50\x48\x2f\x04\x85\xce\xd7\x21\x15\x3b\x4e\x89\x38\xd8\x03\xcb\x1b\x25\x8c\x96\x81\x55\xf6\xb3\x8d\x7c\x15\x04\x63\xd7\x9d\x8f\x37\xe8\xfc\x65\x27\x7f\xc5\x82\xa8\x44\x8e\xb4\xbb\xcd\x59\x05\x2d\xcd\xb9\x16\x4c\x73\x19\xf8\x8b\xb5\x15\xe2\xc9\x6b\x07\x12\x4e\x08\x70\x93\x06\xd1\x23\x6b\x93\x76\xeb\x8c\x1c\x5d\x92\xef\xc7\x07\xad\xe6\xf0\xf0\xb8\x45\x82\xf1\xa0\x4f\xf6\x77\x32\x1f\x32\x7f\xb7\x01\x25\xf9\xfe\x0b\xa6\x8b\x44\x29\x31\x1e\xea\x9d\x4f\xe7\x30\xdc\x6c\x25\x85\x3f\xac\x70\xdc\x04\x5e\x63\x50\x1f\x87\xb2\xd3\x4b\x84\x1d\xea\xd5\xdd\x06\xd6\xb4\x6d\xac\x48\x63\x3e\x45\x58\xfc\x59\x90\x47\x6c\x6f\x4d\x8b\x96\x76\x35\xa8\xc2\xa1\xfb\x00\x5a\xa5\xeb\x51\x49\x54\xc9\xb8\x6e\xa1\x3e\xe1\x30\x59\x77\xe0\xe7\x46\x75\x3e\x91\x71\x88\x2b\x6c\x99\xc2\xb5\x07\xdb\xbf\xb3\xdd\x23\x36\xb4\x55\x14\xb6\x7e\xd1\x6a\xa1\x47\x50\x95\x81\x47\xc3\xe8\xca\xc4\x1a\x81\x2c\x6d\xea\x2c\xb8\x3e\xf0\x3c\x07\x9a\x11\x5f\xa2\x06\xbd\xd9\x9d\xf9\xc0\x7b\xbc\xd0\xbc\xa7\x48\x8e\xbf\xbe\x94\x2c\xda\x14\x1a\xcc\x77\x09\x25\x76\x10\x47\x6e\xdd\xfc\xd9\x57\x6a\x7c\x07\x70\x95\x33\xfb\xd6\x9c\xdb\xa1\x52\x61\xcc\x84\x48\x9b\x9c\xcd\x57\x4c\xc6\x66\xb4\x04\x94\x25\x47\x67\xf2\x4f\x6d\xb7\x46\xd0\x18\x01\x95\x16\x19\x21\x99\x62\xc9\x08\x20\x1c\x06\x40\x5e\x80\x39\x43\x2a\x89\x30\xa5\x54\xce\x62\x46\x55\x20\x47\x11\xc8\x51\x35\x08\xc8\x70\x16\xcb\x52\x02\xf3\x5c\xf0\x6d\x3c\x98\x8b\x14\xa0\xdd\xd3\x2d\xab\xf0\x6d\x56\xa1\xe5\xf8\x17\x6b\x5b\x50\xba\x66\x37\xa5\x50\xf3\x80\x9a\xf7\xd9\xce\xa2\x60\x0c\x99\x5a\x95\x03\x44\xde\xce\x47\xb3\xde\xbe\xdc\xc4\x7b\x6d\xc9\xce\x74\x6a\xba\xb6\xb3\x8d\x76\xd1\x0e\x69\x1f\x6b\x2d\xed\xe0\xb7\xea\x64\xf6\x20\xd5\x44\xbb\xa6\xb2\x0c\xaf\x1f\x4f\x66\x42\x94\xb2\xfc\x0f\x33\xb5\x5d\x9a\x88\xf1\xb0\xbf\x4b\x07\x11\x81\xbf\xab\x7d\x88\x71\x7f\x72\xef\x90\x54\x2b\x69\xf6\x46\xcf\x16\x76\xe3\x0d\x40\xb6\xa5\x48\x35\x1b\x63\x02\x62\xfc\xb9\x46\xbd\xf0\x12\x45\x3a\x8c\xe4\x9b\x70\x52\xf9\x14\x32\x59\xbe\x2a\x28\xe4\xf8\x98\x1f\x54\x67\xa6\x25\xca\x32\x33\x3d\xfa\xca\x74\x23\xd9\x9d\xbc\x8c\xc2\x8b\xf5\x5d\x16\x96\x07\x0a\x07\xcb\xed\x08\x1c\xa1\x8c\x7a\x69\xe9\x9c\x0f\xb0\x02\x1e\x8b\xb7\xae\x78\x20\x73\x0e\x9f\x40\xc9\xa5\x66\xfa\x92\xac\xb4\x81\x7e\xa3\x53\xf7\x2f\xaf\x76\x71\xea\x08\xfc\x5d\x4e\x1d\xe3\xfe\x64\xa7\x4e\xf5\xb2\xb3\x57\x27\xa8\xb2\x5b\x17\x5f\xd2\x35\x8a\x08\x05\xaf\xaa\xcc\x7d\xaf\x35\x9f\x15\xbd\x67\x19\xcb\xd4\x39\x6d\x2f\xac\xef\xce\x31\xaa\xfa\xfb\x8e\x6f\xdd\x4f\x39\xad\xca\xd6\xb7\x94\x3f\x7f\xf3\xf8\x03\xfc\xc7\xbc\xfb\xbc\x58\x5e\xc6\x24\xfb\x9c\xe0\x2e\x42\x24\x71\x93\x8c\x64\x02\x27\x85\x93\x23\x47\xbc\x8f\xfe\x91\xa8\x19\x71\x0a\xbb\x04\x4e\x8a\xf1\xae\xd8\x91\xd0\x7f\x62\xf8\xf0\x8b\xd0\xce\xce\xb1\x03\x68\x72\xd4\xe4\xde\x42\xcb\xef\xa5\xdc\x42\xb0\xa8\xfc\x00\xa9\xfc\x3c\x52\x76\xab\x37\x9e\x19\x95\x9d\xda\xec\x46\x4e\x84\x62\x15\x9d\xdc\x62\xf0\xcd\x70\x12\x05\xbb\x33\x9b\x22\x7a\xa3\x22\xda\xf3\xac\x32\x11\xff\x0e\x5e\x22\xca\xdf\xc2\x0c\x7c\x20\x69\xa7\xa8\x9b\x7d\x2d\xcc\x67\xe5\x88\xe2\x9f\x66\x54\x1f\x79\xa4\xed\x3d\x02\x1a\xa8\xb0\x3a\xfe\x31\x6c\xcf\x98\x6f\x43\xca\xe4\x97\xb5\x30\x1e\xad\x2f\x03\xc2\x3f\x24\xc8\x8c\xc4\xf2\x63\xcb\xd8\x6a\x54\x11\x20\x4d\xe8\xce\x0f\x8e\xb2\x97\x42\xfb\x87\x90\x3b\x15\x83\x97\x77\x5b\x66\x5b\xa6\x33\x80\x05\xed\x92\x07\xf2\x78\xef\xca\x06\x05\x22\x3f\x79\x4b\xe5\xe6\xdd\x39\x23\x20\x96\x9c\x12\x32\x9f\x5b\xc8\xb5\x61\x1a\xbb\xef\xdb\x38\xdf\x74\x60\x9c\x11\x6b\xa7\x0c\x24\x0b\xf8\x33\x93\x49\xd6\xdd\x5f\xd0\x84\x8a\xac\xd4\x8a\x54\x32\xf4\x4e\x10\xbf\xa8\x88\x9c\x02\x92\xac\xf5\x7a\x70\xa6\x12\x96\x1e\x7a\xef\x24\x21\xa7\x20\x36\xfe\xd7\x45\x7c\xf9\xd8\x5a\x64\x04\x79\x44\xce\x08\xaa\x9a\x8a\xed\xf8\xe1\xfb\xa4\x05\xc4\xd7\xa5\x44\xa0\x38\x25\xe3\x43\x83\xb7\xfd\xc8\x13\x7b\xf4\x6c\xeb\xc9\x23\x29\xce\xab\xf8\x90\x49\xac\x7c\x56\xce\xab\x25\xdf\xc9\xfd\x48\xd9\x72\xd9\x1f\x0e\xf8\xb9\xf5\x4e\x19\x4b\x46\x7a\x5f\xba\xca\x50\xf8\xc9\xb9\x6a\x69\x09\x3d\xed\x5e\xc0\xc8\x4d\xe4\xce\x59\x48\x19\x50\x66\x81\xf3\xf1\xef\xb4\xab\x5d\x06\x3c\x43\xe5\x46\x25\x91\x51\xd5\xf8\x55\x4e\xca\x5e\x73\xfc\x0d\x7e\x9b\xf5\x4a\x4b\x1c\x2d\x54\xce\xb0\x2f\x7c\xa9\xd7\xc8\xe3\xbd\xa5\x0f\x8e\x75\x19\x3b\x6c\xf4\x9c\xf1\xd9\x18\x26\x72\x5b\xe1\x82\xd2\xdb\x91\x37\xbc\xdd\x7d\xf1\x75\x70\x1e\xfe\x5a\x4d\xba\x61\x2e\xf6\x8b\xaf\x5d\x8a\xc8\x69\x47\x90\x62\x57\x76\xd2\x45\x74\x51\xfa\xa4\xa8\xa5\xad\x44\x11\x8d\x07\x76\x8a\x55\x56\x2e\x15\x91\x4a\x82\x3f\x25\xf1\x7a\x66\xc8\x24\x96\xc2\xdb\x2a\x66\x99\x5e\xfd\x85\xef\xe6\xc0\x51\xc5\x17\x73\x10\xe9\xca\x4b\xef\x8d\xc7\x3e\x19\x5e\x92\xf8\x23\x43\xfc\x7c\x09\x28\x93\x3b\x93\x91\x39\xa5\x5e\xfc\x25\x63\xf2\x75\x09\xd0\x15\x1f\xdc\xbd\x81\x6e\xfa\x01\x59\xfa\xc1\xca\x4b\x84\xd3\x0f\xf6\x76\x23\x9e\xe0\x55\x13\x97\x35\x0b\x26\x10\xdf\x25\x96\xaa\xaf\xc4\xbf\x5f\x8b\x89\xe8\xd8\x02\xf7\x0f\xfc\x2f\x53\xda\x62\xed\xae\x58\x5d\x82\xc7\xef\xc9\x6a\x36\x7e\xf6\x88\x3a\x30\x0c\x5c\xa0\x6a\x00\x3e\x6c\xd1\x86\x2a\x16\x1a\x7d\xdf\xe8\x33\x8d\x7a\x0f\x76\x00\x64\xc0\x4d\xea\xea\xf0\xd2\x98\x5d\x18\xd3\x7e\x6f\xac\xf2\x57\xe4\x1c\x36\xfa\xe4\x31\x0f\x0b\xb9\xb5\x77\xa6\x8f\xf4\xf1\xcc\xb8\x9a\xea\x13\x63\xdc\x1b\xe9\x12\x96\xf4\x3d\xe4\x6b\x98\x97\xbd\xe9\xf4\xaf\x8b\xc9\x40\xc2\xce\x9c\x11\x57\x7b\x1f\x07\x89\x8e\x75\x53\x8b\xe1\x7f\xf4\xa2\xf2\x39\x0d\xea\x8a\x6a\x66\x70\x8b\x6f\x83\xa5\x02\x65\x0b\x82\x3d\xda\x61\xbd\xdd\xa8\xfd\x27\x00\x00\xff\xff\x6b\x33\xfc\x14\x9c\x36\x00\x00")

func idracPyBytes() ([]byte, error) {
	return bindataRead(
		_idracPy,
		"idrac.py",
	)
}

func idracPy() (*asset, error) {
	bytes, err := idracPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idrac.py", size: 13980, mode: os.FileMode(420), modTime: time.Unix(1469690552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _idracToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x49\xcc\x4b\x2f\x4d\x4c\x4f\x55\xb0\x55\x50\x2a\xa8\x2c\xc9\xc8\xcf\x53\xe2\x82\x89\x85\xa5\x16\x15\x67\xe6\xe7\x01\xa5\x8c\xb9\xb8\xa2\xa3\x8b\x52\x0b\x4b\x33\x8b\x52\x73\x53\xf3\x4a\x62\x63\xb9\xf2\x12\x73\xc1\x9a\x40\xa2\xa9\xc5\x25\xc5\x4a\x5c\x65\x70\xe5\x4a\x46\x7a\x86\x06\x7a\x06\x4a\x80\x00\x00\x00\xff\xff\x20\xd2\xea\x2d\x5d\x00\x00\x00")

func idracTomlBytes() ([]byte, error) {
	return bindataRead(
		_idracToml,
		"idrac.toml",
	)
}

func idracToml() (*asset, error) {
	bytes, err := idracTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "idrac.toml", size: 93, mode: os.FileMode(420), modTime: time.Unix(1467192632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _iloPy = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1a\x6b\x6f\xdb\x46\xf2\xbb\x7e\xc5\x56\xf9\x40\x12\xe7\x30\xb2\xec\xe4\x92\x1c\x7c\x80\x2a\xd3\x57\xe1\x2c\xd9\xb0\xd4\x7b\x40\x10\x08\x8a\x5c\x59\x6c\xc8\x25\x8f\xbb\xb4\xa3\x16\xfe\xef\x37\xb3\xcb\xc7\x52\x22\x6d\xb9\x77\x41\x51\xa0\x2c\x9a\x90\xbb\xf3\x7e\xed\xcc\x2a\x6f\xbe\x7b\x97\xf3\xec\xdd\x3a\x64\xef\x28\x7b\x20\xe9\x4e\x6c\x13\xd6\xeb\x85\x71\x9a\x64\x82\xfc\xc4\xe1\xa3\x78\x4f\x78\xf9\xc6\x77\xbc\xb7\xc9\x92\x98\xf8\x49\xba\x23\xc5\x62\x40\x69\x8a\xdf\x15\xee\x36\x0d\xa3\xa4\xd7\xeb\x4d\x47\x63\xf7\xf6\xce\xb9\x9a\xfc\xcb\xfd\xfe\x7a\x34\xfe\xfb\xf5\x64\xbe\x20\x17\x64\xd9\x23\xf0\x18\xef\x07\xf0\xdf\xb9\x71\x42\x8c\xb3\xb3\xf7\x83\x0f\x57\xf8\x36\x18\x7c\xfa\xf8\xe7\x0f\xea\x0d\x9f\xea\x6d\x8c\x6f\xc3\xc1\xf9\xe9\xfb\x33\x7c\x3b\x3d\xff\x74\x3a\x84\x5d\x45\x69\x30\x1c\x14\x94\xae\x9c\x2b\x78\x24\xc4\xe8\xea\xd3\x50\xe1\x0f\x07\x1f\xd5\xdb\xa5\x33\xba\x1c\x8e\x15\xdc\xe8\xf2\xfc\x12\xf0\x57\xbd\xde\xa5\xf3\x8f\xc9\xd8\x71\x27\xb3\xab\x1b\x77\xe1\x4c\x6f\xaf\x47\x0b\x07\xa4\xfc\x45\xd2\x7e\x43\x36\xe1\x3a\xa3\xae\xbf\xf5\x18\xa3\x91\xeb\x7b\x59\xc0\x4f\x48\x10\xf2\x2f\xf0\xd7\x26\xcc\xe2\x47\x0f\xb6\x1f\x68\xc6\xc3\x84\x11\x8f\x05\x64\x1d\x26\xbc\x5e\xc8\x28\xc9\x59\xce\x69\x50\x90\x33\xa5\x71\x48\x90\x50\xce\x0c\x41\xd2\x2c\x79\x08\x03\x4a\x78\xee\x6f\x49\xc8\x36\x89\x25\xe1\xfa\x71\x12\x00\x37\xe6\xc5\xb4\xff\x99\xf4\xfb\x4a\xd1\x3e\x15\x5b\x9a\x31\x2a\x38\x2c\x2e\x57\xc5\x62\x4c\xe3\x24\xdb\xe9\x2b\x2d\x22\xeb\xdb\xc0\xd3\xa7\x9c\x27\x59\x63\x55\xaa\xa4\x2f\x70\x9a\x85\x1e\xc8\x90\xc7\x6b\x9a\xe9\x52\xec\x6b\xad\xef\xe9\xca\x17\xeb\x4f\x3d\x67\xf1\x83\x73\x37\x73\x16\x6d\xe6\xad\x15\x3d\x21\x3c\xa5\x34\x90\x36\x2c\x59\x68\xf6\x7b\xce\x72\x05\x2d\xdd\x7e\x9e\xaf\x8b\xd5\x6e\x4e\xc9\x0f\x17\x72\xf6\x85\x25\x8f\x4c\x09\xf0\x82\x9a\x4f\xbd\xdb\xbb\x9b\xb1\x33\x9f\xdf\xdc\x1d\xea\x73\xc0\x08\xc5\x7a\x59\x7e\x5d\xf2\x52\xa8\x59\xc2\x68\x21\x89\x9f\x64\x94\x57\x4b\x4f\xbd\xa9\x33\xbd\xb9\xfb\xf7\xb7\xe1\x1e\xfe\x4c\x9b\xcc\x9b\xf2\x3c\xf5\x7a\x7e\xe4\x71\x4e\x26\x51\xe2\x64\x59\x92\x99\xce\x57\x9f\xa6\x02\xec\x63\x7d\x96\x08\x29\xec\x42\xee\x07\x74\x43\x58\x92\xc5\x5e\x04\x14\x5d\x70\x87\x09\xff\x17\x20\xf0\x06\x32\xc3\x9f\x76\x9e\xa6\x34\x33\x2d\x3b\xa3\x69\xe4\xf9\xd4\x34\xde\x62\x76\x7e\x36\x94\x34\x19\x15\x79\xc6\x10\xb0\x20\x78\x4f\x85\x0b\x5a\xb8\x21\xe3\xc2\x63\x00\xbf\x4d\xb8\x38\x21\xa0\x60\x76\x22\x19\x3f\x26\x59\x50\x30\x41\x6d\x2f\x54\x29\xb2\x41\x58\x09\x8a\x96\xb9\x50\x38\x51\x72\x1f\xb2\x8b\x26\xe6\x45\x45\x42\x67\xaf\x6a\x19\xb2\x77\x91\x7f\x95\x83\x66\xe6\x3d\xa2\x62\x50\x07\x50\xa6\x22\x4a\x0a\xee\x6f\xc8\x62\x4b\x49\xe0\x09\x8f\x70\x91\xe5\x3e\x50\xa2\x64\x93\x64\x04\x6a\x22\xf1\x82\x00\x1c\xca\x29\x2f\x58\x80\x7b\x64\x51\x55\x2e\x0a\x39\x78\x87\x0a\xb1\x23\xcc\xe3\x62\x77\x52\xd0\xa3\xe0\x08\x1f\x12\x32\xda\x49\x3a\xe1\xf5\xcd\x19\x31\x59\x42\xfc\x88\x7a\x19\x16\x24\x11\x32\x1f\xfd\x40\xd6\x54\x3c\x52\xca\x08\x85\xc4\x0d\x02\x20\x3e\x9b\x8c\xb9\x4c\xab\x70\x3e\x9e\x4f\x0a\x7a\x58\xa9\xb9\x65\xcb\xaf\x4a\x27\xac\xce\x2b\x65\xbe\x8d\xae\x15\x84\x08\x59\x1a\xc0\x73\x88\xfe\x41\xde\xc6\x4a\x29\xda\x02\x7b\x71\x51\x80\xd4\x10\xf8\x80\xcb\x32\x70\x5f\xf0\x15\x98\x0c\xaa\x1d\x1a\x71\xda\x0d\x77\x5a\xed\xa0\xce\x31\x8a\x51\x5a\xbd\x89\xb4\x09\x69\x14\xa0\xf8\xb1\x0d\x4e\x32\x0d\xf5\x0d\xc2\x2e\x57\x56\x13\x10\x6d\xa7\xe8\xb0\x7b\x6a\x56\xcc\x20\x22\x28\x33\x15\x9a\x75\x42\x86\x56\x93\x7e\xa1\xa6\x79\xb0\x58\x33\x5f\x86\xab\xa5\x81\x11\x66\xac\xa4\x09\x6e\xc1\xc0\x06\x9a\xfd\x25\x9c\x07\x2f\xca\x11\xe9\x3b\x65\x37\x03\xbd\xb3\xa6\x51\xc2\xee\x39\x11\x09\x89\xef\x63\x51\x86\xcc\x01\xa9\x16\x31\xf1\x51\x19\xd6\xcc\xbf\x92\x23\xf9\x13\x39\xad\xb9\x5a\xad\xf8\xa0\x2a\xa0\x2c\x3f\x7f\x58\x01\x11\x81\xe6\x6a\x3b\xcb\xdb\x79\xe3\x03\x01\x05\xfc\xcb\xc6\xc0\x3c\x38\x00\xda\xb9\x16\x98\x4b\x03\x78\xa3\x11\x65\xe6\x3f\x03\xa8\x62\xd6\xf6\xa0\x88\xb0\xc0\x84\x05\x45\x96\x46\xed\x01\x79\xae\x05\xe4\x0b\xe1\xf4\x3a\xf9\x1b\x32\xef\x15\xbd\xa5\x01\x96\xdb\x37\x73\xb7\xec\x45\xcd\xa9\x00\xf4\xca\x23\x2b\xb1\xc9\xab\x0a\xc3\xbd\x38\x8d\x68\x89\x22\xfd\xa9\xd5\x11\x38\x00\x86\xef\xcf\xce\xc8\xf4\x87\x9f\xfb\x65\x46\x73\x2c\x2e\xe8\x50\x2c\xe6\xb5\xc6\x98\x37\x21\x13\x26\xb7\x79\x1a\x85\xc2\xec\x93\xbe\xb5\x1c\xac\x1a\x12\x35\x24\xa9\x3b\x08\x59\x04\xf1\x13\xc5\x92\xf0\x65\x9d\x96\x87\x96\xe9\x6b\x01\x7a\x84\xc4\xe7\x24\xd9\x90\x73\x22\x71\xff\x42\x3e\x12\xb1\xcd\xa8\x07\xdd\x8b\x5e\x6a\xfc\x76\x25\xf0\xf1\x0b\x45\xfc\x4a\x11\x20\xa7\xe9\xa2\xe9\xe3\x2b\x69\x6b\x4d\xea\xd2\x87\xb1\x91\x96\xb1\x21\x75\xab\xb9\xe0\xa7\x1e\x19\x87\xbd\x80\xd5\x80\x5d\x1a\xd2\x6b\x32\x32\x34\x27\xa6\xaa\x4a\xcd\xe5\x9e\xb5\x8f\x22\xd5\x97\x28\xb5\x21\x0b\x0c\xe7\x2b\xf5\x73\x59\xe5\x17\xd4\xdf\xb2\x04\x8e\xb1\xdd\x3e\x01\xa5\x4f\x19\x5a\xb8\xd2\xf0\x64\x0d\xa2\xbb\x54\x75\x91\xea\x4c\x93\xaf\xfb\xfe\xc4\xce\xa0\x0a\xbe\x63\xdd\x39\xf8\xf4\x81\x4c\xbf\x6f\xb8\xaf\x23\x06\x5f\x8c\xc3\x66\x2c\xe2\x97\x92\xb3\xe9\xb7\x3a\xa7\xe5\x66\x4d\x1f\xbe\x75\xbf\xed\xb5\x50\x96\x0e\x07\x3e\x03\x5d\x2b\xfb\x4b\xc5\x8b\x63\x65\x8e\x1b\xd6\x3e\x74\x9b\x87\xe3\x0e\x0f\x2b\xb9\x4a\xe7\xc0\x57\xb3\xd5\x91\xbb\xe0\x97\xd6\xfe\xa1\xd9\x2d\xa8\x81\xed\x6d\xd5\x34\xfc\x27\x0f\x05\x05\xf5\xfd\x84\x3d\x50\x16\x52\x26\xa4\x45\x92\x3c\x03\x6a\xd0\xe8\xbc\xf5\x3d\x0e\x4d\x36\x56\x16\x0a\x1b\x94\x3c\x52\xc2\xb0\xe1\x86\x33\x06\x22\x6c\xeb\xa5\x80\x0d\x27\x0d\x01\xd3\x0b\x70\xeb\x3a\x14\xb6\x8a\x0f\x68\x47\x52\x6c\x84\xb1\x73\x72\x51\x26\x19\x26\xd5\x57\x5b\xff\x53\x6d\x56\xcd\x29\x3e\x7d\x98\x1f\x5d\xec\x36\xb5\x39\x43\xae\xb7\x8e\x24\x72\xe7\x60\xba\x51\xab\x9e\xef\x56\x5d\x94\xb6\xf9\xd4\xd6\xbb\x14\xe5\x7f\xb8\x57\xfe\x41\x25\x51\x46\x4b\x25\x6e\x33\x20\x81\x10\x42\x15\xae\xcc\xd7\x3f\x51\x5f\x18\x96\x24\x38\xdf\x71\x01\x21\x35\x01\x65\xa0\xde\x63\x46\x1a\x87\xc7\x61\x45\x16\x82\xa4\xd0\xdc\x58\x55\x89\x09\x94\x0f\xcf\x41\xf0\x1e\x34\x72\x39\x3d\x56\x8e\xdb\xd2\x72\xc7\x8b\x52\x1b\xfb\xd7\x0a\xd3\xde\x06\x75\x89\x38\x55\x79\x7a\x49\x1f\x42\x9f\x76\x77\x44\x1a\xba\xcc\x32\xd9\x0f\xa9\xee\x03\xba\xb4\x28\x82\x3c\x6a\xe0\xb5\xb4\x3e\x9a\x92\x2a\x6e\xfe\x9f\xd6\x86\x14\xc3\x92\x25\x1b\xe9\x7a\xbf\x68\x34\xad\xee\xaa\x86\x4f\xd5\x9e\x1e\x22\x1e\xc2\x82\x2f\xe5\x26\x06\xa7\x82\x6a\x6f\xb4\x3a\x1d\x51\x01\xf0\x6a\x50\x92\x74\xf0\xde\xc2\x17\x56\xa7\x07\x1a\xd2\xd6\x7d\x29\xfa\xd0\xd1\x66\x09\x39\xc3\x8c\x38\x0f\xef\x59\x0c\x55\xc6\x68\xa5\xd5\xd1\x98\xe2\xa3\x7b\x49\xcf\xe3\x97\x9c\x55\x3e\x6b\xe8\x09\xbe\x3c\xef\xca\xae\x16\xf0\xec\x7f\xac\x01\x90\x6d\x01\x14\x63\x32\xc3\x36\xff\x05\x9f\x7f\x93\xe4\x6f\x3f\xff\x8f\x16\xe4\x5b\xa4\x7e\x2d\xdc\xb5\x07\x63\x4b\x53\x9a\xd6\x58\x3b\xc8\xf4\x57\x20\xa8\xf3\x54\xc7\xf8\x0d\x8a\xc2\x91\x49\xaf\x27\xf2\x21\x72\x7b\x82\xa8\x83\x1f\x44\x08\x42\xe9\x66\x3c\xad\x4b\x36\x10\xcc\x3b\x38\xaa\xa3\xd0\x5b\x47\x70\x92\xaf\x73\x01\xa7\xb9\x27\x0c\x8e\x87\x7a\x07\xb5\x84\x45\x3b\xf2\xe8\xed\xf0\x94\x57\x97\x04\xf7\x79\xc8\xb7\xd5\x25\x41\x29\x0e\x9c\xfa\x5c\x48\xd5\xbd\x90\x01\x50\x07\xb9\xe6\x9d\xc2\x03\xb7\xd5\x9d\x02\x57\xd7\x09\xf6\x6f\x52\xa7\x8e\x9e\xb9\xf7\xf0\xf4\xfa\x86\x73\xf7\xef\xa8\x8e\xed\x8f\xb2\x7f\xd4\xb1\x3f\xea\xd8\xaf\xf0\x02\x5e\x4d\x1c\x6f\xf4\x57\x05\xf9\x5e\x18\xeb\x97\x7c\x99\x17\x72\x5a\x5f\x20\xf7\x7f\x64\x3c\x4f\xb1\x7e\x40\x59\x29\xe3\x1c\xe6\x76\x08\x73\x62\xfe\xf2\x64\xd9\x7d\x5b\x35\xb6\xa6\x3e\x66\x58\x65\xa7\x8f\xd7\x76\xad\x21\x6a\x91\xbf\x92\xd3\x4e\xae\x0d\x81\xfb\x0b\x59\x66\xf9\x36\xc9\xa1\x5a\xaf\xa9\xaa\x9a\xe8\xc5\xb6\x3e\x5f\x56\x27\xd2\x6f\x52\x80\xe4\x13\xe5\xc4\x56\xcd\x69\xeb\x5d\x63\x4a\xb3\x6b\x9c\xc6\xd4\x57\x89\x5f\x0c\xe4\xa8\x67\x20\xbb\x65\xa9\x8a\xd4\x3b\xf6\x98\x77\x8f\x97\xd5\x87\xb3\x56\x23\xe9\x71\x84\xae\xc1\xd1\xcd\xda\xd8\x66\x1d\xcc\x66\xaf\x9b\xed\x8a\x0b\x81\x4a\x32\x20\xd0\xf6\xeb\xd9\x3e\x58\x33\x5f\xcb\x39\x59\xbb\x44\xea\xc8\xec\x03\x7e\x4b\xa3\xba\x18\x93\x74\x6a\x2f\xee\x5d\xcd\x77\x47\x6d\x8b\x42\x2d\x6c\x1a\xbf\x7d\xed\xb1\x6a\x0d\xb6\xe5\x60\x55\x24\xb6\xc4\x24\x33\x85\x79\x42\xfa\x7d\xcb\x86\x09\x3e\x4c\xcd\x4e\x66\xf5\xaf\x36\xaf\xe2\xd4\xa8\xe1\x92\x51\x27\x83\xa2\x5c\x94\x96\x2f\xee\x7a\x5a\x0a\x4a\x23\x2c\x35\x1a\x45\x60\x72\xdf\x63\xcf\xfe\xe6\xb2\x91\x32\xe3\x29\xd5\xef\x77\x27\xfc\x2c\x21\x93\xdb\xf2\x4a\x1b\x1b\x13\xa4\x4b\xb6\x1e\x87\xd4\xa3\xac\xfc\x6d\x2a\xb0\xfb\x55\x92\x23\xb7\x23\xa8\xaa\xa8\xc7\xa1\x44\x62\xa0\x49\x9f\x25\x5b\x8a\xff\x3a\xd2\x15\x56\x37\xe9\x3a\x03\x8b\x4b\xa4\x97\x7f\xb5\x52\x57\x58\x8f\xf5\x49\x7f\x90\xc7\xf5\xa6\x59\xb3\xa9\xc1\xeb\x6d\x15\x1f\xb5\xc4\x75\xaa\x19\x6d\x29\x7c\x74\xb9\x29\x6e\x4c\xf1\x8a\x0e\xff\x79\x80\x1d\xe4\x71\xca\x4d\x0d\x15\xaa\x72\xaf\x07\x76\x75\x65\x34\xbb\xae\x6c\x57\x5c\x20\x16\x32\xd7\x35\xea\x9b\x21\x60\x9a\x70\x9b\xb2\x87\x30\x2b\xa5\x9d\xdc\xba\x8b\x1b\x77\x3e\x1e\xcd\xb4\x58\x56\x6e\x3f\x80\x9d\x8e\x66\xa3\xbf\x39\x53\x67\xb6\x70\x7f\x9c\x3b\x77\xee\x6c\x34\x75\x34\xac\xda\xab\x2f\x62\xde\x8e\xe6\xf3\x7f\xde\xdc\x5d\x6a\xd8\x42\xbf\x35\xec\x0e\x78\x75\xae\xc9\xdf\x3e\x89\x59\x46\xca\x49\xfd\x7b\xe3\x38\x89\xe3\x9c\x85\xbe\x3c\x32\xe4\x26\xf4\xb3\x9c\x50\xfd\x2a\x19\x0d\x49\x6d\x2f\xbb\xe7\x8d\x5b\x4e\xc8\x75\x9b\x7e\x0d\x85\x79\x6a\xf5\xfe\x1b\x00\x00\xff\xff\xea\x00\xdd\x31\xa5\x21\x00\x00")

func iloPyBytes() ([]byte, error) {
	return bindataRead(
		_iloPy,
		"ilo.py",
	)
}

func iloPy() (*asset, error) {
	bytes, err := iloPyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ilo.py", size: 8613, mode: os.FileMode(420), modTime: time.Unix(1473167631, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _iloToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xf2\x49\xcc\x4b\x2f\x4d\x4c\x4f\x55\xb0\x55\x50\x2a\xa8\x2c\xc9\xc8\xcf\x53\xe2\x82\x89\x85\xa5\x16\x15\x67\xe6\xe7\x01\xa5\x8c\xb9\xb8\xa2\xa3\x8b\x52\x0b\x4b\x33\x8b\x52\x73\x53\xf3\x4a\x62\x63\xb9\xf2\x12\x73\x91\x34\xe9\x66\x14\x64\xe6\xe4\x2b\x71\x95\xc1\xb5\x28\x19\xeb\x59\x28\x01\x02\x00\x00\xff\xff\x5e\xdd\x05\xf1\x5e\x00\x00\x00")

func iloTomlBytes() ([]byte, error) {
	return bindataRead(
		_iloToml,
		"ilo.toml",
	)
}

func iloToml() (*asset, error) {
	bytes, err := iloTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ilo.toml", size: 94, mode: os.FileMode(420), modTime: time.Unix(1467192632, 0)}
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
	"idrac.py": idracPy,
	"idrac.toml": idracToml,
	"ilo.py": iloPy,
	"ilo.toml": iloToml,
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
	"idrac.py": &bintree{idracPy, map[string]*bintree{}},
	"idrac.toml": &bintree{idracToml, map[string]*bintree{}},
	"ilo.py": &bintree{iloPy, map[string]*bintree{}},
	"ilo.toml": &bintree{iloToml, map[string]*bintree{}},
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

