// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/comparison.tmpl.json (25.576kB)

package report

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _comparisonTmplJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xdd\x73\xdb\xb8\x11\x7f\xf7\x5f\x81\xc1\xb4\x1d\x5f\x6b\xa7\x92\xfc\x75\xf6\x4c\x1f\x7c\xce\x24\xcd\x4c\xd2\xba\x76\x72\x7d\xc8\x79\x34\x10\xb9\xa2\x30\x06\x01\x06\x00\x65\xeb\x3c\xea\xdf\xde\x01\x48\x91\xe0\x87\x25\x59\x91\x6c\xf9\x82\x17\x9b\x58\x80\xe0\xee\x0f\x8b\xc5\x6f\xc1\x0f\x3d\xec\x20\x84\x09\xe7\x42\x13\x4d\x05\x57\xf8\x0c\x19\x11\x42\x98\x51\xa5\xf1\x19\xfa\x6a\x4b\x28\x97\xda\x9a\x41\x4a\x99\xfe\xc0\xf1\x19\xea\xee\x95\xd2\x90\x68\xa2\x44\x2a\x03\xc0\x67\x08\xef\xef\xa3\xf7\x92\x0c\x09\x27\x68\x7f\x1f\x3b\xcd\x80\x93\x01\x33\x4d\xb4\x4c\xc1\x91\x8f\x68\xd8\x22\xa5\x81\xe0\x17\x82\x09\x69\xfa\x94\xd1\x80\xec\x76\xf6\x50\xaf\xdb\xdd\x43\xbd\xa3\xa3\x3d\xd4\xfd\xc9\xed\x9a\x93\xd8\x5e\xfb\xbc\x34\x07\xfd\x05\x9d\x33\x90\x5a\xb9\xed\xf4\x24\xb1\xed\x42\xa2\x46\x03\x41\x64\x88\xf3\xba\xa9\xfd\x7f\xb3\x83\xd0\xd4\x34\xc7\x10\x52\x5d\xd3\x16\x47\x1c\xf4\x87\x10\x9f\xa1\xde\xd1\x61\x2f\x93\x48\x92\x8c\x3e\x0b\xc1\x34\x4d\x66\x98\x60\x6a\x9a\xf0\x94\xb1\xac\xa4\x41\x5a\x85\x4c\xfd\x71\xe7\xe0\xf4\xe4\xe4\xf8\xe8\xf0\xf8\xe8\xd0\xd6\x32\xca\x6f\x0d\xf0\x5f\x6f\x6c\x31\x21\x1c\x98\x2a\xa0\x9f\x01\x8f\x09\xa3\x44\x59\x30\xec\x28\x4d\x67\x16\xe1\x01\xb1\x92\x21\x61\xaa\xc0\xce\x1a\xf7\x11\x78\xa4\x47\xe6\x9a\x9d\x8a\x1c\xda\x9a\xbb\xa3\xc7\x04\x09\x35\x28\xbd\x9f\x8b\x8a\x66\x2d\x88\x64\x72\x29\xed\x18\x55\x3b\x1d\x52\xc6\x5c\x2f\xb1\x82\xf7\x92\x84\x14\xb8\xf1\xad\x52\xab\x48\xd2\xf0\x52\x94\xde\x97\xb9\x04\x3e\x43\xa7\xce\xb8\xdd\x99\xbe\x7a\x8e\xe0\xde\xed\x03\x21\x3c\x31\xe5\xd9\x58\x16\x7d\x8f\x68\x18\x02\xbf\x06\x49\x5b\xec\xb6\xe3\x74\x52\x14\x19\x44\xc0\xc3\xaa\x1a\x64\x1c\x35\x1d\x33\x48\xa5\xcc\x8c\xa8\xf4\x87\x10\x8e\xc9\x7d\xb3\x79\x4c\x79\x53\xa8\x46\xe2\xae\x29\xd5\x42\x13\xd6\xd2\xef\x98\xb0\xd4\x1a\x60\xda\x37\x8c\x64\x94\x17\x95\x15\xe1\x1d\x0d\x33\x17\x70\xa5\x8e\xbb\x65\x33\x27\x65\xec\x52\x50\xae\x3f\x09\x3b\x0b\x71\x20\x38\x87\x40\x43\x58\x0e\xbd\x48\xaa\x01\xa2\x70\x9b\x8f\x45\x7f\x0d\xad\x12\x90\x01\x70\x4d\x22\x68\x00\x9f\x98\xcb\x19\x57\x48\xcd\xb9\x47\x55\x79\x73\x9c\x24\xf0\x10\x24\xd8\x38\x30\x64\x42\x97\x7a\x29\x3b\xb0\xff\x1e\x83\x94\x34\x84\x9a\x61\x2a\x21\x01\xb4\x4d\x03\xa5\x49\x70\xdb\xb8\x8a\xd2\x90\x24\x10\x7e\xa4\xbc\xa9\xb0\x26\x32\x02\xad\x9c\x90\xe8\x06\x45\x33\x07\xee\x13\xab\x9e\x4a\xe3\x5d\x49\x34\xec\xc6\x44\x6b\x90\xb1\x50\xba\x9f\x98\x3f\x76\x64\x1f\x28\x57\x9a\xf0\x00\xfe\xf1\xbf\xdf\xf0\x9f\x14\xc8\x31\xc8\xdf\xf0\xf4\x6b\x37\xbe\x41\x62\x38\x54\xa0\xd1\xc3\xc3\x9b\xec\x68\x3a\x55\x3f\xb9\x11\xce\xcc\x20\x21\x63\x62\xdc\x0e\x6b\x1a\x43\x3f\x33\xbe\xda\x84\x72\x0d\x72\x4c\xd8\x3b\x12\x68\x3b\x29\xbb\x95\xea\xcc\xc5\xdf\x15\xfd\x3c\x3c\xbc\x19\x10\x05\x1f\xc9\x00\xd8\x74\x5a\xed\x2a\x06\x2d\x69\x60\x5a\xb5\xda\x52\x6d\x2c\x61\x68\x03\x23\x3e\xaf\xca\x0d\xa8\x66\x88\x0b\xd9\x74\x6f\x43\x00\x3e\x0b\x58\x1c\xee\xd6\x88\xd5\x2f\x8b\xb0\xca\x8f\x4a\x97\xd6\x23\x09\x6a\x24\x58\x58\x73\x75\x63\xe2\x3b\x29\x62\x67\xdd\x29\xe4\x57\x10\xe5\x73\xb7\x76\xc2\xf5\x88\x0e\x75\xf3\x0c\x6d\x23\x3c\xbe\x14\x4a\x2b\x94\x80\x44\xd7\x10\x08\xee\x04\x03\x5d\x2c\x77\x4e\x30\x88\xd5\x15\x28\xc1\xd2\x7c\xa1\xab\x07\x30\x35\x22\x12\xc2\x96\x30\x28\xa4\xae\x05\x72\x1b\xeb\xfa\xb3\x75\x9a\xf2\x90\x8e\x69\x98\x12\x86\x1b\x11\x66\xd6\xc6\x2e\xc2\xa5\x7e\xf7\xe4\x9e\xd6\x42\xd5\x20\x0d\x6e\xb3\xf9\xeb\x1a\x6b\xd4\xce\x63\x9e\xc1\xa3\x85\x4e\xd4\x5a\xb7\x07\xed\x22\x38\xb7\x04\xc1\x09\xb9\x87\x39\x61\xa3\x74\x52\x35\x32\x48\x54\xfd\xcf\x78\x9a\xa9\xbb\x10\x29\xaf\xd7\x89\xe8\x17\xa2\xa0\xe1\xb3\xd9\x02\x54\x55\xbb\x58\x82\x1a\x62\xc7\x9e\x85\x13\x74\x29\x55\x1b\x57\xd8\xa0\x9e\x8d\xc9\x31\x69\x8e\x3b\x61\x34\x6a\x73\x47\x2b\xff\x08\xe3\x42\xe9\x0a\x09\xcc\x21\xf0\xe4\x6b\x56\x6e\x25\x5f\x15\xc1\xaa\xec\xcb\xc1\xc7\xd3\x2f\x4f\xbf\x36\x41\xbf\x46\x5a\x27\x7d\x09\xdf\x52\x50\x5a\xfd\x51\x78\x98\x35\xca\xce\x7d\xf5\xec\x74\xec\x69\x80\x6e\x1f\x2f\x5b\x16\xbb\xad\xa6\x67\x57\x39\xfe\x9e\xa1\x79\x86\xf6\x14\x55\x3d\x43\xfb\x71\x18\x5a\x7d\x7b\xec\x74\x05\x82\x76\xbc\x04\x3f\xab\x8f\xdd\x22\x82\xd6\x90\xd2\xf6\x80\xb4\x12\x45\xb3\x15\x9e\xa3\x6d\x37\x47\xab\xaf\xc6\x77\x30\x50\xc2\x2e\x00\x73\xf9\x44\x2b\x39\xc3\x2b\x33\x2b\x87\x25\x3d\x89\x10\xad\xa8\xfd\x4b\x12\xa1\xd7\x41\x6a\xfe\x95\xc6\x03\x90\x48\x0c\xd1\xc5\x6c\x82\xa1\xb7\x30\xa6\x01\x28\xb4\xfb\x5f\x18\x5c\x5b\x94\x67\x95\xe6\x0a\x3f\x79\xd2\xe3\x49\x8f\x27\x3d\x9e\xf4\x20\xf4\xd8\xb6\xd4\x2a\xac\xa7\xbb\x04\xeb\xf1\xbb\x52\x9e\xf1\xac\xc4\x78\xc2\x41\x3f\x26\x4a\x83\xec\x07\xe5\x42\xf6\xbd\xbc\x67\x0b\xf6\xa4\xe6\xd8\xb5\xc1\xdd\xa9\xef\x02\x76\xeb\xf6\xa6\x9e\x8e\xe1\xeb\x22\x74\xf6\xb9\x1c\x2d\xd0\x27\x6b\x24\x7a\x4b\x34\x31\x4e\xe6\x59\x9c\x67\x71\x9e\xc5\x79\x16\x87\xd0\x23\x5b\x57\xdd\x9f\x57\x61\x71\x9e\xc6\x79\x1a\x87\xd6\x4a\xe3\xea\xf7\xc2\x48\x42\xfb\x19\x57\x48\xe3\xc7\x6f\x1f\xa2\xbf\xa3\xb9\x27\x07\x26\x06\xaf\xf1\xee\x63\xfe\x3c\x6d\xcd\xd3\x9e\xce\x5e\xc8\x38\x42\x4f\xa6\x81\x2e\x24\xcf\x78\x57\xb2\x3e\x12\x4f\x40\xfd\xe5\x11\x7e\x12\x47\x5c\x02\xe0\x85\xa4\x70\x01\xc0\x23\xaa\xb4\x88\x24\x89\xfb\xdf\x52\xc2\x35\x65\xb0\xdb\x79\x73\x7a\xba\x37\x1f\xd0\x8c\x08\xcd\x99\x06\x83\x09\xda\x65\xb0\x4e\xb8\xeb\x78\x26\xa7\xa7\x73\x3d\xb6\x40\xe8\x62\xf9\x2d\xdf\xef\x05\xe3\x39\x0d\x5f\xb8\xff\xfb\x16\x6f\x47\x66\x70\x7e\xf9\x01\xe5\xf7\xb0\xd1\x67\x1a\x03\xda\xbd\xb6\x49\x99\xdf\xd0\xdd\x70\x2a\xf0\x48\x1a\xf0\x58\x06\xd0\xf3\x19\x40\x09\xb1\xcf\x00\xf2\xc1\x7a\x89\x7d\xdc\x95\x52\x80\x83\x9e\x4f\x01\x7c\x0a\xb0\xfe\x14\x60\x36\x91\x2a\x0f\xc3\xcd\xe1\x9f\x73\xda\x6f\x11\x0d\xdd\x5e\xa2\xbf\x18\xef\xc5\x99\xd7\x32\x63\xb0\x9d\x09\xd8\xeb\x4b\x0f\xe6\x60\xfd\xc3\x66\x09\x4b\x61\xe2\x93\x85\xa5\x92\x85\x0b\x66\xa8\x86\x4f\x16\x7c\xb2\xe0\x93\x85\xac\xe2\x87\x4e\x16\xea\xb7\x0b\x7a\x27\x2b\xe4\x0a\xdd\x03\x9f\x2b\xf8\x5c\x61\x6d\xb9\x02\x19\x47\xcd\x37\x99\xa5\x08\x40\xa9\x7e\x90\xa4\x7d\x65\x5f\xda\x58\xf5\x8d\xa4\xbf\xa2\x6e\xa7\xf3\x92\x8f\x80\xac\x8d\xe3\xaf\x0b\xa7\xe7\xc3\xe4\x75\x3f\x64\x7b\x71\xf9\x05\x7d\xd1\x94\xd1\xdf\xed\x57\x49\xd0\x15\xd1\x80\x76\xff\xec\xb9\xd4\xcb\x3d\x83\x31\x77\x38\xd0\x86\xb9\x8b\xe7\x58\x9e\x63\xad\xb4\x21\xbb\x12\xc9\x3a\xf2\x1c\xcb\x73\xac\xb5\x72\xac\x48\xf4\x63\x88\x95\x26\x5a\xf5\xed\xa5\xfa\x94\xa7\x0a\xfa\x83\x89\x06\xb5\xd1\x07\x6a\xd7\xb5\xcb\x77\x6d\xb4\x3e\x67\x4c\x04\x4b\x6e\xc2\xba\x26\x13\x73\x5e\x66\xec\x46\x9f\xb5\x5d\x15\xea\x6d\x81\x75\xf1\x1e\xea\x2a\xa8\x7e\xef\x4e\x6a\x1d\xd5\x11\x90\xe4\xd5\xf9\xef\x3f\x81\x24\x9b\x76\xdf\x8b\x97\x01\x7a\x4b\x40\xdd\x8c\xf3\xbe\xdd\xe6\x34\xe5\x13\xc4\x42\x4e\xd0\x17\x65\x56\x34\x9f\x9a\x6c\x32\x35\xb1\x9e\xf2\xc4\xad\xde\x1f\x28\x0d\x71\xdf\x28\xf7\x79\xc8\x4c\xbe\x3d\x79\xc8\xc1\xf1\x0a\x79\xc8\xa1\xcf\x43\x7c\x1e\xb2\xb6\x3c\x44\xa5\xb1\x61\x17\x91\x90\x22\xd5\x66\x58\x5e\xdf\x8b\x7c\x15\xf5\x37\xf8\x38\xc7\x32\x30\x6d\xc7\x6b\x79\xcb\x20\xb2\xd5\x3b\xbd\xe5\xdb\x77\xef\x05\xba\xaa\x1b\xe2\xb9\x94\x7f\xd5\xee\x0f\xbe\xaf\xbb\x93\x77\x6b\xe6\xab\x99\x75\x6e\x0f\x58\x05\x23\x88\xc9\xaf\x20\x55\xe6\xea\xbd\xec\x8b\xee\x4a\x4f\x58\xfe\x85\x78\x79\x6b\xd1\xc1\x9a\x44\xe5\xdc\xc3\x1a\xe2\x84\x11\x4d\x79\xb4\xcc\xf7\xf2\x09\x63\xbf\x1a\xbf\x6a\x3a\x61\x49\x27\x2a\xe3\xa3\xe1\x3e\x7b\xf4\x2b\x49\xf6\x3b\x67\x3f\x77\x8e\x4f\xd0\xdf\x90\x29\x74\x6d\xa1\x3a\x5e\xe3\xbc\xeb\xaf\x8e\x10\xb9\xe7\x56\x9a\xe7\x35\x79\x47\x4e\xc5\x4d\x9b\xd7\x2c\xc9\x10\x4d\x43\x18\x52\x4e\xf3\x88\x91\xf9\x50\x3f\x9b\x4c\xbb\xb3\xf8\x5e\xf9\x56\x7f\x9e\x22\xbb\xc1\x82\xf2\x80\xa5\x21\x9c\xb3\x36\x1a\x54\x4c\xa0\x6c\x81\x70\xbb\x8a\x53\xa6\x69\x73\x2a\xcf\x7e\x0d\xa0\x79\x42\x49\x69\xca\x50\x8a\x10\xfe\x96\x82\x9c\x2c\xa7\x7d\xe9\x4b\xdd\x8a\x34\x82\xfb\x5a\xc2\x84\xd5\x2d\x4d\xbe\x48\x76\x3d\xe1\x41\x5b\x2c\x6d\xc6\x4c\x4d\x22\xeb\x2d\xea\x3f\x33\x7d\x70\xb5\xb6\xa1\xb8\x91\xb5\x37\xce\x83\x6a\x66\x99\x53\x91\x2a\xf8\x9c\x75\x54\xf9\x54\xd6\x0e\xaa\xfc\xfc\x81\x0d\xa4\x85\x7f\x0f\xb3\x85\xc9\xac\x9b\xe6\xb0\x58\x33\xb1\x16\xb9\x58\x8b\xe9\x14\x57\xce\x4e\x68\x70\x6b\xc9\x5b\xde\x47\x8e\x5b\x7f\xb6\x54\xbb\xe1\x14\x1f\x39\xeb\x52\xb7\xe3\x14\x0e\xdc\x42\xb7\x7c\xa0\x11\x1f\x39\xc7\x5d\xb7\x70\xd0\x71\x6b\x9c\xf5\xa4\xe7\x1c\x77\xf3\xdf\x7d\xb8\x99\xd9\x61\x88\x85\xe3\x1b\x0b\xaf\xe2\x76\x7c\xec\x76\xec\x5e\xa5\x77\xe8\x16\x1c\x62\x7d\x12\xba\xfa\xce\x74\xa9\xc0\xf7\xbb\xb0\x84\x14\x0f\xa4\xb8\x53\xb9\x0f\x97\x4b\xbb\x41\xdc\x1c\xe7\x23\x81\xd3\xea\x6f\x4e\x8c\x8b\xa8\x76\xb8\x33\xdd\xf9\x7f\x00\x00\x00\xff\xff\x95\x52\x7d\x8c\xe8\x63\x00\x00")

func comparisonTmplJsonBytes() ([]byte, error) {
	return bindataRead(
		_comparisonTmplJson,
		"comparison.tmpl.json",
	)
}

func comparisonTmplJson() (*asset, error) {
	bytes, err := comparisonTmplJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "comparison.tmpl.json", size: 0, mode: os.FileMode(0644), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x62, 0x7b, 0xaa, 0xb1, 0x5d, 0xdb, 0xb8, 0xba, 0xea, 0x5a, 0xad, 0x1d, 0x96, 0x96, 0x5d, 0x5b, 0x5c, 0xa2, 0xae, 0x4d, 0x73, 0x93, 0x3c, 0x79, 0x6c, 0x88, 0x59, 0xb9, 0x8, 0x54, 0x3b, 0x32}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"comparison.tmpl.json": comparisonTmplJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"comparison.tmpl.json": {comparisonTmplJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}