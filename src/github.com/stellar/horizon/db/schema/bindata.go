// Code generated by go-bindata.
// sources:
// latest.sql
// migrations/1_initial_schema.sql
// DO NOT EDIT!

package schema

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

var _latestSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5b\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\xc0\xdc\x8b\xe4\x19\xc9\xb5\x2e\x77\xc9\x55\x9e\xbb\x19\x9d\xcd\x34\x9a\xca\x54\xce\xa2\x9b\x64\x3a\x1d\x0e\x44\xc2\x12\x1b\x8a\x60\x08\xca\xb1\xaf\xd3\xef\xde\xe5\x1f\x50\xfc\x03\x10\xa0\x4c\xb9\x79\x89\x45\x2c\x76\xf7\xb7\x58\x2c\x76\x97\xe0\x78\x7c\x36\x1e\xa3\x8f\x94\xc5\x9b\x88\xac\xfe\x58\x20\x17\xc7\x78\x8d\x19\x41\xee\x7e\x17\xc2\xd8\xd9\xd9\xca\xb0\x10\x8b\x71\x4c\x76\x24\x88\xed\xd8\xdb\x11\xba\x8f\xd1\xaf\xe8\xf2\x2a\x1d\xf2\xa9\xf3\xb5\xf9\xd4\xf1\xbd\x84\x9a\x04\x0e\x75\xbd\x60\x03\x03\x83\x7b\xeb\xfd\x2f\x83\x2b\xce\x2e\x70\x71\xe4\xda\x0e\x0d\x1e\x68\xb4\x03\x0a\x9b\xc5\x11\xfc\xc7\x80\x92\x06\x39\x8f\x2d\x01\xd6\x0f\xfb\xc0\x89\x3d\x1a\xd8\x6b\xe0\x44\x92\xf1\x07\xec\x33\x52\x11\x03\x0c\xec\x1d\x61\x0c\x6f\x52\x82\xef\x38\x0a\x80\xd7\x55\xae\x3b\xc1\x91\xb3\xb5\x43\x1c\x6f\x61\x2c\xdc\xaf\x7d\xcf\x19\xa1\x70\x63\x3b\x00\xd5\xa7\x09\xd9\xcd\xdd\xf2\x23\x9a\x9b\x37\xc6\x67\x34\x7f\x8f\x8c\xcf\xf3\x95\xb5\xca\x29\x2f\xe2\x08\xbb\xc4\x26\x0f\x0f\xc4\x89\x99\xbd\x7e\xb6\x69\xe4\x92\x08\xb4\xa1\x5f\xaf\x5a\x27\x7a\x81\x4b\x9e\xec\xad\xc7\x62\x1a\x3d\xdb\xc0\x26\x60\x38\x45\xc2\x6c\x40\xe3\xb9\x5d\x66\xd3\x90\x44\xb8\x98\x1b\x3f\x87\xe4\x05\xb3\x0f\x9a\xbc\x48\x8b\x6e\x73\x7d\xe2\x6e\x48\x94\x4e\x64\xe4\xdb\x1e\x1c\xa3\x13\x84\xd2\xf4\x30\x22\x8f\x1e\xdd\xb3\xfc\x99\xbd\xc5\x6c\x7b\x24\xab\x97\x73\xf0\x76\x21\x8d\x62\xe0\xf1\x08\x0f\xbc\xc4\x73\x8f\x63\x73\xac\x2d\x1d\x9f\x32\xe2\xda\x38\xee\x32\x9f\x3b\xf3\x11\xae\x84\x1d\x87\xee\x83\xf8\x08\xa5\xcb\x33\xb1\xeb\x46\xb0\x5d\xdb\xa7\x6f\x59\xc5\x55\x61\xe7\xa9\x04\x6e\x0b\x9f\xd0\x21\x06\xbd\xec\xf8\xc9\x0e\xf5\x28\x69\xa8\x4b\x49\x74\xc9\x78\x2c\x69\x27\x5e\xf3\xf5\x56\x92\xa9\xdd\x78\x5d\x2c\xc3\xd5\xd9\x6c\x61\x19\x77\xc8\x9a\xfd\xbe\x30\x4a\x84\x4b\x73\xf1\xa5\xac\x66\x2d\x74\x41\x14\x8d\x62\xcf\xf1\x42\x0c\x2b\x89\x52\x51\xd7\x4b\x73\x65\xdd\xcd\xe6\xa6\x55\x62\xa3\x9a\x6a\x87\x5f\xc9\x73\x17\x1d\x8a\xd0\xd3\x55\x03\xf1\x44\x6d\xf9\x1b\x1a\x85\x70\xbc\x6c\xf2\xb8\xd7\x22\xb0\x46\xd9\x2a\x41\xd7\xc0\xd9\xec\xeb\xe5\xe2\xfe\xd6\x44\x9e\x9b\x49\xbf\x31\xde\xcf\xee\x17\x96\x26\x6f\x89\xe1\xda\x39\xa7\xbf\xf4\x95\xe6\x1b\x79\x65\xfc\x71\x6f\x98\xd7\x47\x20\x85\x2d\x93\x1c\x0b\x9d\x25\x57\x98\xe8\xcd\x3e\x1c\x62\xda\x5a\x4b\x7c\xa8\x8b\xce\x62\x16\x7a\x73\xf3\x70\xaf\x47\x9c\xc7\x76\x3d\x62\x1e\x93\xdb\xa9\x6b\x9e\x9d\x13\x1b\x9f\x2d\xc3\x5c\xcd\x97\x66\x79\xd3\x25\x6c\x49\x0b\x41\xe8\x87\x1b\xf6\xcd\xe7\x96\xbf\xfe\x60\xdc\xce\x1a\x02\xaf\xce\xb2\xac\xd4\xc4\x3b\x32\xe5\xcf\x90\x05\x27\xd5\x34\x9f\x72\x85\x56\x90\x1c\xee\xf0\x14\x8d\xaf\xd0\xf2\x7b\x40\x22\xf8\x2b\x4d\x56\xaf\xef\x8c\x99\x65\x70\xce\x9c\xdf\x59\x85\x63\x75\x30\x67\x7c\xbd\xbc\xbd\x35\x4c\xab\x85\x73\x46\x00\xd1\xa1\xca\x00\xcd\x57\x68\xc0\x13\x5a\xfe\x8c\xa5\x4c\x06\x75\xc9\x1c\x7e\x2e\xb3\xb0\x90\x12\x4f\xc5\x96\xe6\xd2\xaa\xd9\x13\x7d\x9a\x5b\x1f\x0a\xb5\xca\x99\x6d\x45\xfc\x81\x4b\x4d\x91\x2e\xe0\x1b\x4c\x52\x03\x7c\x5c\xfc\x25\xdc\x24\xf5\x43\x18\x51\x87\xb8\xfb\x08\xfb\xc8\xc7\xc1\x66\x0f\x29\x79\xc3\x0c\xb9\x9b\xf4\x66\x85\x8c\x5f\xd5\x08\xc2\x95\x3f\x30\xa8\xaa\x70\x1c\xfe\x5c\x6c\x02\x3f\xa9\x9a\x50\x92\x4b\x21\x28\x67\x50\xf2\x3c\x29\x7a\x18\x81\x48\x4b\x1f\xd0\x10\x8e\x82\x11\x7a\xc4\xfe\x9e\x9c\xa3\x10\x7b\x11\x4b\x4d\xa2\x59\x9c\x24\x64\x2e\x79\xc0\x7b\x1f\xf2\x15\xbc\xf6\x09\x0b\xb1\x43\x92\x8a\x6a\x50\x1b\xfd\xee\xc5\x5b\x9b\x7a\x6e\xa9\x48\xaa\xc0\xaf\x6f\xe4\x1c\x7d\xba\xed\x0f\xd8\xf9\xbe\xe0\x06\x00\xb2\x42\xea\x14\x95\xd7\x23\x8b\x17\xf5\x23\x72\x78\x86\xe0\x1f\x9c\x29\x31\x79\x8a\xd3\x65\x32\xef\x17\x8b\x51\xfa\x14\x87\x21\x54\x6c\x49\xbe\x8a\x92\x92\x11\xb6\xcd\x2e\x44\x89\xda\xe9\x4f\xf4\x27\x0d\xc8\xd9\x79\xc3\x5d\xea\xd1\xaa\x1f\xbd\xeb\x6c\x0f\x8a\xaf\xbd\x8d\x17\x34\x54\xcf\x72\x56\xa8\x4a\x71\x04\x87\x0f\x89\x60\x41\xa3\x67\x58\xe6\xe1\xdb\x9f\xce\xe5\x4a\xf3\x78\xdc\xaf\xce\x39\xd7\x5c\xe5\x1a\x12\x5b\x06\xa1\x79\x18\xc9\x28\x7f\x48\x13\xd3\x1f\x10\x8c\x10\x38\x7b\x6a\xa3\xa9\xa3\x8b\x87\x5c\x12\x63\xcf\x67\xe8\xdf\x8c\x06\x6b\xb9\x55\xf8\x91\xd6\xaf\x55\x72\xae\xb9\x55\x78\x99\x29\xd1\xb4\x54\xfb\x89\xd7\xb4\x46\x2f\x2a\x3b\xc5\x13\x73\x23\x95\xb2\x94\x74\x59\x0a\x3d\xf2\x34\x0b\x5d\xd6\x24\x1c\x96\x45\x8f\xbe\xa8\xfd\x6a\x7b\x29\x69\xc4\x14\xdb\xa9\x3e\x27\x22\x38\x56\x4e\xca\x68\xf7\xa1\xab\x4d\x5b\x38\x52\xfe\xb3\x56\x16\x37\xb0\x4c\xea\x2e\x45\x21\xdc\x01\x6e\x0f\x02\x88\xd0\x23\x1f\x08\xb1\x43\x4a\x7d\xf1\x68\xd2\xb1\xb2\x81\x44\xb2\xd6\xe9\x30\xec\x5e\x12\x3d\xca\x48\x76\xf8\x29\xa9\x06\x21\x64\xdb\xcc\xfb\xb3\x49\x25\xf7\x65\x49\x6a\xd7\xaf\x6b\x4b\xd2\xf8\x22\x64\x89\x41\xe9\x6f\x78\x75\x08\xe9\x6a\x00\x9e\x1e\xf3\xc4\x2d\xcf\xb2\xe5\xa6\xa8\xe4\x6f\x3c\x27\xd7\x92\x91\x22\x58\x59\xb3\x3b\x2b\xcb\x01\x26\xe9\x83\xb9\x09\xcc\xd2\x53\xfb\xf7\x2f\xf9\x23\x73\x89\x6e\xe7\xe6\x3f\x66\x8b\x7b\xa3\xf8\x3d\xfb\x7c\xf8\x7d\x3d\x83\xec\x01\x4d\x7a\x01\x8a\x96\x9f\x4c\xe3\x06\x64\x2b\x10\x67\x95\x58\x37\xc0\x05\x6f\x05\xf9\x45\xd2\x89\x50\x60\x39\x99\xa7\xaa\x0e\xd4\x6a\x1f\x52\x72\xe8\x26\xf9\x82\x93\x01\x4b\x8f\xa4\x17\x9e\x48\xd9\x23\x46\xf7\x91\x43\xb8\xaf\x4b\xa2\x3f\x8f\x54\x83\xc1\x74\xda\xa0\xd0\xd8\x15\xd2\x2a\xb5\x5f\x73\x4b\x7b\x07\x9a\xa1\x41\x67\x15\x5e\x12\x1c\x54\x15\x7f\x3f\xe1\x41\x21\xe5\xb5\x02\x44\x47\xb0\x2f\x0c\x11\x0a\x69\xcd\x20\x21\x9b\xd0\x12\x26\x2a\x5d\x9e\x93\x79\x2e\xf7\xd6\xb2\x82\xda\x89\x59\x9e\x8f\x29\xd2\x3d\xdd\x48\xd2\x1e\x14\x84\xb4\x07\xd1\xf2\xcc\x05\x4b\x37\xa2\x2c\xeb\xfb\xbf\xe4\x6d\x90\x01\x91\xe0\x91\xf8\xa0\x94\xa8\x7c\x83\x61\xc8\xa2\xa0\xd4\x94\x0c\xee\x48\x52\x06\x0b\x87\x12\x2b\xc8\x86\x99\xb7\x09\x70\xbc\x07\xd6\x02\xb3\xff\xf5\xed\xf9\x3f\xff\x75\x88\xc6\xff\xf9\xaf\x28\x1e\x03\x45\x2d\x9d\x23\x3b\x9a\xbe\xdc\x68\x72\x3c\xf0\x0a\xc0\x0c\xad\xd1\xfd\xc0\xab\xc9\x26\x47\x06\xe6\xb4\xd7\xb0\x70\x50\x74\x83\x15\x7f\x01\x07\xde\x08\x4a\x58\xd8\x60\xf9\xe6\xe1\x3d\x56\x9d\x1d\x9f\xed\x97\xb4\x1d\xdd\xb1\x9d\x9b\x74\x05\x38\xcc\x00\x0c\xfe\x88\xfd\xe1\x40\x2b\xb5\x00\x7b\x44\x64\xe3\xf8\x98\xb1\xd3\xa1\xd0\x6e\x78\xb7\xe2\x50\xc4\x3f\x31\x92\x9b\xa4\x53\x93\x34\x69\xd4\x2d\x11\x74\x33\xb3\x66\x0a\x88\x73\x73\x65\xc0\xa9\x32\x37\xad\x65\xa3\x11\x92\x1e\x1b\x2b\x34\x1c\x4c\x6c\x2f\xf0\x62\x0f\x0a\x9c\xac\x2f\x78\xc1\xbe\xf9\x83\x11\x1a\xfc\x78\x39\x79\x3b\xbe\x7c\x33\x9e\xbc\x43\x93\x9f\xa7\x97\x93\xe9\x8f\xef\x2e\x2e\x7f\x7e\x77\xf9\xe6\xa7\xf1\xe5\xbb\x81\x5c\xe9\xd6\x86\x88\x8e\xd6\x0a\xbe\xa2\x9e\x45\x0f\x6c\x45\x45\x7f\x0f\x6c\x35\xea\xaf\x2e\x52\x5e\x92\xf2\x83\xbb\x2a\xa4\xac\x8c\x85\x71\x6d\x95\x1a\x7c\x17\x50\x71\x76\xd8\x9d\x23\x34\x19\x65\xbd\x3d\xb5\x7b\x48\x52\xfc\x1e\x4c\xae\x95\xdb\x1e\x6f\xf4\xae\x69\x54\x1f\x66\x57\x05\x93\x2e\x86\x97\x26\x4d\xdd\x4d\x22\x7c\xab\x58\x74\xad\xf9\x6b\xc8\x4e\xc9\x58\x23\x2a\xd7\x64\xa4\xe7\xda\xec\xe6\xa6\xfc\x9a\x53\xa4\x06\xfa\x78\x37\xbf\x9d\xdd\x7d\x41\x7f\x37\xbe\xa0\xa1\xe7\x76\xed\x0f\x9c\x02\x4a\xbb\x48\x11\x32\x0d\x25\xb5\x81\xb6\xbf\xed\x3e\x11\x54\x99\xd0\x36\xb0\xad\x8a\x2a\xe1\x96\x6e\x11\xe4\x98\xd2\xeb\x06\xc7\x54\x04\xd9\x3d\x85\x03\xc3\xe4\x05\x8b\xb0\x3e\xb8\x5f\xcd\xcd\xbf\xa1\x75\x1c\x11\x82\x86\x39\xf1\xa8\x91\x80\x8b\x54\x4d\x6f\x45\xf4\xa6\x67\x5a\x95\x68\x29\x59\xaf\x65\x44\xba\xe5\x17\x3b\x7a\xd3\x2e\xe3\xa7\xa7\x5f\xad\x6c\x1a\x35\x2b\x24\xa1\x9f\x97\xef\xad\xbc\x54\xef\x7b\x73\x0e\x11\x3c\x57\xbf\xc6\xbc\x0c\x82\xbf\xed\xa8\xe8\x2f\xea\x6d\x8e\xf8\x8b\x0b\x99\xea\x87\x0c\xb6\x57\xa5\x21\x53\xd5\x55\xf7\xd0\x43\x19\x09\xdb\xb3\x0a\x08\xfc\x1a\x52\xff\x28\x72\xce\x65\x20\x92\x62\xe3\x28\x5c\x62\x38\xfc\xfe\x55\xff\x70\x72\xce\x92\xbd\x70\x24\xa0\x6a\xb3\xac\x09\xa9\x76\xff\xac\x5f\x54\x55\xe6\x65\x60\xfc\xbd\x57\x05\x87\x58\xbf\xe6\x8d\xba\xbe\x95\x6c\x48\xd0\x0b\x47\x02\x75\xd5\xf7\x07\x7b\xd5\x5d\x29\xae\x0c\xa4\x78\x69\x5c\x3d\x98\x32\xc2\x0e\x48\xfa\x5e\x80\x36\x49\x6a\xfd\x95\x8b\x50\xbf\x39\xda\xcf\xc9\xd5\x2a\x43\x19\x59\x13\x22\x85\xda\xc2\x0b\xb3\xa7\xd0\x5d\x24\x48\xb9\x53\x0b\x4a\x7d\x14\xa7\x75\x9b\x8a\xa0\x63\x02\x8d\xfe\x75\xe9\x13\x2f\x42\xe3\x3d\xb4\x12\x4c\x6d\x82\x3e\xb4\xf2\x5d\xf2\xd7\x59\x9b\xf2\x45\x04\x15\xae\x12\xad\x3e\x24\xe1\x4d\xfb\xd7\xc1\x26\xbc\x6d\xa1\x02\x29\x9a\xa4\x8f\xb6\xf8\x2c\xe1\x75\x10\x16\x2f\x0d\x54\xa8\xa4\xc5\x8d\xe2\xe3\x8c\x13\xc2\xa8\xcb\x12\x26\x8d\x5d\xc3\x44\xeb\x57\x2a\xa7\x88\x13\x6d\x02\x75\x10\x29\xb2\x41\xe5\x17\x3c\xaf\x80\xa9\x76\x7e\x4a\x91\xa8\x8f\x50\xc1\xf7\x4b\x27\x74\xb0\xa6\xb4\xa3\xb3\xc8\xb6\xef\xb7\xfa\x59\x81\x16\x09\xca\xe4\x65\x38\xe4\xf7\x13\xc6\xbf\xfd\x86\x06\x8c\xfa\x90\x08\xb0\xe4\x16\x52\xb2\x26\x83\xe9\x34\x79\x5d\x76\x7e\x3e\x42\x72\x42\x87\xba\x7a\x84\x1e\x63\x7b\x12\xc9\x49\xd7\x74\xbf\xd9\xc6\x5a\xe2\x2b\xa4\xed\x0a\x54\x48\x6b\x2a\x9c\xa3\x4f\x1f\x8c\x3b\x23\x73\x40\xf4\x2b\x7a\xf3\xa6\xb4\x7c\xb2\x8f\x12\x91\x43\x77\xa1\x4f\x62\x92\xae\xc4\xff\x02\x00\x00\xff\xff\x84\x42\x19\x86\xc1\x38\x00\x00")

func latestSqlBytes() ([]byte, error) {
	return bindataRead(
		_latestSql,
		"latest.sql",
	)
}

func latestSql() (*asset, error) {
	bytes, err := latestSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "latest.sql", size: 14529, mode: os.FileMode(420), modTime: time.Unix(1458252087, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\x6d\x6f\xdb\x46\x12\xfe\xee\x5f\xb1\xc8\x17\xc9\x38\xf9\x2e\x41\x0e\x41\xce\x46\x02\x28\x36\x73\x11\x2a\x53\x89\x44\x35\x09\x8a\x82\x58\x91\x2b\x8a\x35\xc9\x65\x76\x49\xbf\xa4\xe8\x7f\xef\x2c\xdf\xdf\x96\xa4\x6c\xd2\x2d\x0a\xb4\xe2\xce\xce\xcc\x33\x33\xfb\xcc\x70\xe9\xb3\x33\xf4\x2f\xd7\xb6\x18\x0e\x08\xda\xfa\x27\x67\x67\xf0\x2f\xfa\x4c\x79\x60\x31\xb2\xf9\xb2\x44\x26\x0e\xf0\x0e\x73\x82\xcc\xd0\x8d\x96\x4f\x36\x8a\x86\x78\x00\xf2\x2e\xf1\x02\x3d\xb0\x5d\x42\xc3\x00\xbd\x43\x2f\x2f\xa2\x25\x87\x1a\x37\xf5\xa7\x86\x63\x0b\x69\xe2\x19\xd4\xb4\x3d\x0b\x16\x26\x5b\xed\xe3\xdb\xc9\x45\xaa\xce\x33\x31\x33\x75\x83\x7a\x7b\xca\x5c\x90\xd0\x79\xc0\xe0\x3f\x1c\x24\xa9\x97\xe8\x38\x10\x50\xbd\x0f\x3d\x23\xb0\xa9\xa7\xef\x40\x13\x11\xeb\x7b\xec\x70\x52\x32\x03\x0a\x74\x97\x70\x8e\xad\x48\xe0\x0e\x33\x0f\x74\x5d\x9c\x24\xf0\x54\xec\x92\x73\xe4\x3b\xbe\xc5\x7f\x38\x17\x48\x7b\xf0\xe1\xa7\xf2\x4d\x53\xd4\xcd\x62\xa5\x5e\xa0\x0d\x58\x72\xf1\x39\x3a\xbb\x40\xab\x3b\x8f\x30\xf8\xbf\x08\xf9\xe5\x5a\x99\x6b\x4a\x2e\x89\x16\x1f\x91\xba\xd2\xe0\xc1\x62\xa3\x6d\x52\x85\xe8\xeb\x42\xfb\x84\x36\x97\x9f\x94\xeb\x39\xf2\x2d\xdd\x80\x08\x3a\x54\x58\x2f\x99\xcf\xb5\x54\x1c\xb9\x5c\x5d\x5f\x2b\xaa\xd6\xe2\x46\x2c\x80\x60\x6b\x4d\x09\x5a\x6c\xd0\xe4\xf3\xf2\x3f\xbe\x25\x92\xe7\x33\x6a\x10\x33\x64\xd8\x41\x0e\xf6\xac\x10\xe2\x31\xa9\xfa\x71\xe0\x01\x65\x64\xb8\x28\xc4\xfa\xca\x41\x08\x77\x8e\x6d\xc8\x03\x50\x76\xe1\x71\xf8\x13\xb3\x02\xbe\x28\x59\x14\x80\x2e\x04\xb5\x84\xc4\x73\x51\x71\x9c\x04\x1c\xd1\x3d\x9a\xde\x90\x87\x19\xba\xc5\x4e\x48\x4e\x91\x8f\x6d\xc6\xa3\x90\x44\x65\x48\x30\x33\x0e\xba\x8f\x83\x03\x54\x4d\xec\xf5\xac\x9c\x42\x21\x66\x92\x3d\x0e\x1d\x28\x7d\xbc\x73\x08\xf7\xb1\x41\x44\x39\x4f\x2a\xab\x77\x76\x70\xd0\xa9\x6d\x16\x2a\xb4\x1c\x77\x5b\x78\xf6\xa0\x63\xc3\xa0\xa1\x17\xf0\x14\xbe\x36\xff\xb0\x54\x72\xf0\x49\xec\xb2\x08\x80\x58\x66\xf6\xbc\x98\x8f\x68\x5f\x4d\x2b\x9a\x9e\x20\xf8\xc7\x36\xd1\xce\xb6\x6c\x2f\x88\x32\xa5\x6e\x97\xcb\x59\xf4\x1c\x9b\x26\x83\x73\x02\x47\x0b\x33\x6c\x04\x84\x41\x60\xd8\x03\x84\x6b\xfa\xe6\xbf\xa7\x27\xa7\xb5\x5a\x49\xb4\x93\xfd\x9e\x18\x43\xbb\x9c\x28\x4d\x3c\xae\x00\xd1\x65\x08\x52\x39\xea\x13\xe0\x30\xc1\x0b\x32\xc9\x17\x94\x99\x84\xbd\x40\xb0\x42\x2c\x40\x5a\x5e\x8d\xea\xa5\x79\xc9\x24\x01\xb6\x1d\x8e\xfe\xe0\xd4\xdb\xc9\x83\xe2\x10\x13\xf6\x0e\x1c\x94\x44\x69\x12\x14\x4e\x7e\x84\x40\xa1\x32\x47\x63\x61\xfd\x80\xf9\xa1\x39\xa3\x15\x79\x9f\x91\x5b\x9b\x86\x5c\xef\xdc\x98\xc4\x88\x61\x8f\xe3\x98\x7d\xa3\xac\x64\x7e\x5c\x29\x1f\xe7\xdb\xa5\x86\x5e\x56\x2c\xe4\x59\xe9\x27\x6f\x38\x94\x13\x53\xc7\x01\x12\x1d\x04\xda\x82\xeb\x23\x71\x90\x44\x2f\x11\x4f\xd0\x4f\xea\x91\xea\x1e\x46\xa0\x19\x75\x6d\x8a\x65\x43\xdf\xec\x2d\x9b\xd5\x51\xf2\xd3\xf5\x29\x83\xb0\xe8\xb7\x90\x0f\x40\x54\xc3\xf2\xaa\x5a\x51\x14\x48\x03\x70\xdb\x1e\x6f\x2e\xc8\x3d\x21\xba\x4f\xa9\xd3\xbc\x2a\x9a\xae\x0e\x22\x92\x5c\x47\xcb\x70\x76\x09\xbb\x95\x89\xb8\xf8\x5e\x0f\xee\x75\x20\x3e\x9d\xdb\x3f\xeb\x52\xf2\x52\xce\xd3\xe6\x63\x16\xd8\x86\xed\xe3\xc1\x19\xaa\xd9\x46\xce\x57\xcd\x98\xfa\x1f\xf7\x6e\x02\x39\x16\x3f\xa8\x80\x60\xfe\x48\xc3\xb0\x51\xbe\x6c\x15\xf5\xb2\x25\x12\x45\xf0\xa9\x74\x3f\x1b\x11\x82\x8d\x36\x5f\x6b\x71\x23\x7d\x15\x3d\x58\xa8\xa0\x2c\x6a\x7d\x1f\xbe\x27\x8f\xd4\x15\xba\x5e\xa8\xbf\xce\x97\x5b\x25\xfb\x3d\xff\x96\xff\xbe\x9c\x43\x0b\x46\xaf\x06\x01\x8a\x56\x5f\x55\xe5\x0a\x6c\x77\x20\x9e\x2f\x35\x65\x7d\x24\xe0\x4c\x77\x87\xf8\xbf\x6d\xb3\x13\xcb\x58\x85\xda\xd5\x4c\x8b\xf4\x28\x6d\xb8\xbe\x0f\x3e\xc4\xb8\xa2\x7e\xf4\xc4\x76\x14\x3f\xe2\x34\x64\x06\x49\x4b\x5d\xc2\xfd\x29\x4f\x4d\x26\xe7\xe7\x35\x89\x1e\x87\xa2\x08\x6f\x3c\x5a\x90\x59\x89\x62\x2f\xa1\x85\xa6\xbd\xcd\x09\x78\x0a\x29\xc8\x3c\x1b\x96\x16\x3a\xac\x3c\x17\x31\x1c\x09\xf6\x89\xd4\xd0\x61\xad\x4e\x0e\xb2\x0d\x2d\xf4\x50\xd8\x32\x5e\xc9\xa6\x14\x51\xf4\xaf\xf7\x38\x96\x4c\x61\x1d\x43\x5e\x5f\x06\x69\x27\x83\x46\xd9\xdc\xb4\x7c\x5e\xc1\xd2\xd6\x2c\x9b\xf5\xfe\x91\x69\x0d\xe6\x1e\xe2\xdd\x12\x07\x9c\x42\x01\xb9\xaf\x51\xf5\xbd\x98\x9d\xe0\x35\x4d\xb2\xe8\x12\xf1\x0a\xd9\xb8\x24\xa2\x20\x5b\xe6\xb6\xe5\xe1\x20\x04\xd5\x0d\x61\xff\xdf\x9b\xd3\xdf\x7e\xcf\x59\xf8\xcf\xbf\x9a\x78\x18\x24\x2a\x43\x1c\x71\xa9\x1e\x75\x83\x3a\x67\x67\xba\x3c\x08\x43\x2b\xab\xe7\xba\xea\x6a\x12\x64\x10\x4e\x7d\x07\x89\x83\x17\x56\x88\xe2\x5b\x28\x60\x8b\x44\x64\x58\x3c\x4c\x70\xbc\x92\xa3\x93\xd8\xee\x75\xde\xe3\xe3\xb2\x52\x97\x5d\xdd\x1d\xc5\xf2\x97\xab\xe5\xf6\x5a\x15\x29\x15\x2f\xd4\x29\x4a\x0f\xe2\x0d\xaf\xed\xd3\x49\xaf\x81\x02\xc2\xc1\x88\x65\x38\x98\xf3\x1a\xa3\x0f\x86\x42\xda\xac\x8e\xc2\xd1\xc1\x7e\x6d\x48\x3a\x42\xe1\xdf\x90\x87\xfc\x5a\x45\xdd\x68\xeb\xf9\x42\x6d\x41\x5b\x27\xbc\x23\x13\x18\x95\xd2\xfc\xea\xaa\x60\xad\x8f\x8f\xe8\xf3\x7a\x71\x3d\x5f\x7f\x47\xbf\x28\xdf\xd1\xd4\x36\x8f\xef\xc1\x23\x22\x95\xd9\x6c\xc3\xda\xea\x67\x27\xda\x5d\x36\xa0\xa4\x90\x16\xea\x95\xf2\xed\x11\x8d\x2a\xda\x57\xd0\x27\xee\xcc\x1a\xdb\xd6\x76\xb3\x50\xff\x8f\x76\x01\x83\x17\xce\x69\x22\x3c\xab\xf5\x85\x26\x4f\x45\x7b\x1b\xcc\xcd\xa8\x57\xf6\xf2\xb1\xda\x61\x9b\x5c\x8b\x1b\xea\x60\xce\xc5\xea\xfa\xb9\x57\xe9\xe5\xb3\x7a\xdb\x6e\xac\x71\x1d\x38\xf8\x21\x5e\x7f\xaa\xdb\x5b\x75\x01\x53\x56\xe2\x7d\x45\x77\x11\x43\x7a\xed\x56\x72\xbf\xe9\x35\x7b\x96\xde\xa0\xc9\x3c\xcf\x69\x75\x48\x9f\x81\x3d\xfb\x7a\x9b\x4f\xf5\xb3\xc6\x8b\x82\x0e\x04\xd4\xd7\xfd\x51\x40\x24\x8a\x8b\x38\x24\xfd\xef\x51\xb0\xea\x68\xb2\x1b\x3d\x48\xf8\xd0\x80\xca\xba\x8b\x98\xd2\xbb\xca\x12\x88\x66\xf7\x8a\xa7\x77\x14\x1f\x6b\x06\xfa\x1d\xdb\x06\x6f\x6d\xcf\x24\xf7\x7a\xf5\x5e\x5d\x07\xbd\xc9\xe5\xf9\xa0\xae\x77\x5a\x2b\xe2\xc8\x2e\xf9\xcb\xec\x1d\x0b\x1e\x01\x64\xe0\xf0\xb7\x19\xea\x76\xbf\x33\x05\x09\x05\x08\x7d\x62\x2e\x1e\x86\xde\x5b\x4d\x74\x12\x90\x10\xea\xf0\x3a\x39\x1c\x42\x65\x76\xc9\x3d\x86\xeb\x4d\x76\x3a\x0f\x69\x26\xd9\x1f\xc4\xa8\x35\x53\xb2\xf3\x18\x8a\x91\xab\xab\xdc\xe2\x8f\x9c\x82\xda\x47\x83\x4e\x2c\x95\x0d\xfd\x91\x15\xbe\xe1\x3c\x4f\x66\x8a\x1f\x8d\xba\x60\x15\x64\xfb\x23\x6a\xfa\x3c\xf5\x3c\xd0\x1a\x3f\x8c\x75\x61\x6c\xda\xd4\x1f\x6c\x3a\x29\x3e\x0f\xc0\xec\xa2\xa7\x0b\x94\x74\xf2\x2f\xab\xce\xef\xc8\x47\xe7\x86\xaa\xa9\xc6\xa9\xea\x58\x86\x28\x2b\x2d\xdf\x23\x8f\x41\x11\x6d\xf6\xfa\x00\x2a\xef\x38\x0e\xdc\x48\x3d\xb3\x6e\xa5\x17\x90\xa6\xce\x19\x0d\xcd\xc1\xfd\x48\xd3\x78\xa2\x58\x32\x10\x3e\x72\x1e\xaf\x27\x44\x9e\x8f\xe2\xf8\x39\xfa\x71\xa9\x1b\x7b\xf4\x24\x0c\xc2\x26\xc9\x66\xa3\xf4\x5d\x52\xdf\x51\x7a\x33\x4c\x41\xb5\x18\xe8\x1c\xc1\xa6\xd3\xf4\xbb\xd8\xd9\xfb\xf7\x68\xc2\xa9\x03\xf3\x0c\x17\xdf\xbe\x45\x89\x4d\xce\xcf\xc5\x75\xed\xe9\xe9\x0c\xc9\x05\x0d\x6a\xf6\x13\xb4\x39\x0f\x09\x93\x8b\xee\x68\x68\x1d\x82\x5e\xe6\x4b\xa2\xed\x0e\x94\x44\x2b\x2e\x9c\xa2\xaf\x9f\x94\xb5\x12\x9f\x27\xf4\x0e\xbd\x7e\x5d\xc8\x9e\xec\xaf\xf9\x90\x41\x5d\xdf\x21\x01\x89\x32\x51\xfc\x43\xc0\x2b\x7a\xe7\x9d\x98\x8c\xfa\x28\xfa\x1b\xa7\xe6\x72\x31\x30\x37\x20\x5f\x17\x1d\x82\xe5\x03\xd5\xb6\xa9\xc0\x11\xbd\xc4\xfa\x6b\x4e\x5b\x5b\x9b\x4c\x5a\x55\x6d\x32\xd9\x1b\x4b\x26\xf4\x77\x00\x00\x00\xff\xff\x5d\xb2\x1f\x7d\x3f\x29\x00\x00")

func migrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_initial_schemaSql,
		"migrations/1_initial_schema.sql",
	)
}

func migrations1_initial_schemaSql() (*asset, error) {
	bytes, err := migrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_initial_schema.sql", size: 10559, mode: os.FileMode(420), modTime: time.Unix(1457459961, 0)}
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
	"latest.sql": latestSql,
	"migrations/1_initial_schema.sql": migrations1_initial_schemaSql,
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
	"latest.sql": &bintree{latestSql, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"1_initial_schema.sql": &bintree{migrations1_initial_schemaSql, map[string]*bintree{}},
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

