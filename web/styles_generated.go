// Code generated by go-bindata.
// sources:
// dist/hydrocarbon.min.css
// DO NOT EDIT!

package web

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

var _distHydrocarbonMinCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\x5b\x6e\xe3\x3a\x0f\xde\x4a\x80\xa2\x98\x17\xba\x70\x32\x99\xa2\x95\x81\x7f\x05\xff\xdb\xac\x80\xb2\x68\x9b\x27\xb4\xa4\x4a\x74\x9a\x8c\xe1\xbd\x1f\xf8\xd2\x36\x3d\xed\x3c\x59\xa6\x4c\xfa\xbb\x11\x01\xad\x4d\x80\x75\x0a\xfe\xda\x03\x3a\x97\x28\x67\xc0\x18\x85\x14\x30\x29\xd7\x42\x80\x99\x1d\x01\x0e\x8e\x03\x58\xb0\xdc\x82\x95\x50\x9f\x5e\x86\xa0\x04\x36\xb8\x2b\xd4\xe8\xcf\x98\xa1\xc6\xa8\x1c\x3c\xd4\xe4\x95\x12\xd4\xac\x04\x75\x70\x04\xce\x81\x23\x01\x47\x8a\x2c\x19\x5c\xe3\xc1\xf1\x19\x9c\x80\x53\xa0\x1e\xa8\xb7\xe4\xa0\x61\x12\x97\x49\xa1\xe1\xf6\x6d\x54\xc3\xed\x90\x08\x9a\x10\xe6\x89\x4d\x48\x3d\x74\x7b\xe8\x0e\xd0\xfd\x84\xee\x08\xdd\x2f\xe8\x1e\xa1\x23\x74\x94\xa0\x6b\x53\x18\x22\x74\xda\x0b\x30\x70\x93\xb0\x27\xe0\xbe\x05\xf6\x19\x4e\xd6\x81\xa0\x25\x01\xa1\x96\xbc\x03\x61\xe8\x31\x9d\xa0\x27\x3f\x80\xc7\x33\x04\xfb\x0f\xd5\x0a\x41\x20\x0c\x1a\x07\x85\x08\x31\x11\xbc\x40\x1a\xec\x15\x32\x64\xec\x23\x64\xaa\x17\x5c\xb9\x47\x11\xc8\x11\x3d\x64\x4d\x7c\xa2\xf9\x11\x7c\x0b\x79\xb0\x90\x87\xbe\xc7\x74\x85\x3c\x44\x50\xb4\x42\xa0\x8b\x4c\xea\x40\x67\x26\xa0\x1d\xe8\x0c\x1a\x94\x7b\x02\x4d\xa0\x0a\x03\x0c\x02\x67\x4c\x70\x66\x47\x61\xec\x31\xb5\xec\x4d\x59\x45\x74\x8e\x7d\x6b\xca\xca\x86\xe4\x28\x99\xb2\x3a\xd3\xec\x0c\x4a\x81\xc2\xad\x37\x16\x33\x09\x7b\xaa\x9a\xe0\xd5\xb0\xef\x28\xb1\x2e\x2f\x45\xe6\x3f\x64\xf6\x65\x79\x3f\x7d\x36\xf3\xcd\x89\xbf\x2a\xfd\x59\xd2\x77\x8d\x36\xfa\xa3\xe3\x1c\x05\xaf\x66\x09\xc2\x34\x93\x1b\x67\x04\x45\x47\xdc\x76\x6a\xf6\x53\x10\x18\x64\x14\xce\x5a\x64\xbd\x0a\x19\x1f\x3c\x4d\x37\xb9\x79\x19\x97\x67\xfe\xef\x85\xc1\x66\xfe\xff\x4d\xc1\x52\x13\x66\x1f\xb6\x9b\x97\xad\x30\xd6\xc1\x2b\x79\x35\x3f\x7e\x54\x6f\xc7\x65\xd6\xa2\xf8\xb8\x8a\x55\xe4\x88\xf5\xad\x7a\x45\x1d\x44\x30\x66\x32\x6f\x87\xe9\x61\x56\xc4\x62\x1a\x5f\xd9\x69\x67\x0e\xe5\xfd\xf4\xb0\x0d\xdc\x4a\x4f\xe5\x7d\xb5\xfa\x51\x08\x35\xba\x7e\xb2\x75\x15\x33\xc7\xef\xf4\xb0\x58\x9f\x66\xf5\xbc\x9b\xff\x19\x92\xb9\x6b\x9a\xa6\xda\x8e\xfb\xe3\xfe\xe9\xf0\x73\xf5\xa8\xc1\x9e\xe5\x6a\x7e\xa3\xcf\xc5\x6f\x4a\xdc\x54\x4a\x17\x2d\x12\x79\x47\x69\xc6\x1e\xa2\x72\xcf\x7f\xe8\xff\xd4\xb2\x65\x61\xbd\x4e\x8e\xcf\xef\x18\xb7\xa4\xec\x13\xf5\x3b\x1c\x34\x54\xaf\x21\xb9\xe2\x35\x61\x34\x36\x11\x9e\x8a\xf9\xbd\x5a\x99\x3c\x97\xf7\xb7\xc1\x78\x8c\x97\xea\x93\x6f\x0f\xc7\x65\xf4\xbc\x67\xe3\x82\x62\x4d\xd8\xba\xd0\x13\x8f\x6b\xef\x62\x28\x2b\x0a\xd7\x53\xb7\x1f\x3f\x06\x1e\x8e\xf1\x32\x75\x87\xdb\x4a\x19\x2f\xd3\x32\x2e\x86\xcc\x73\x76\x4c\x22\x41\xe5\x33\x4d\xec\xe3\xa0\x7f\xab\x1b\xf6\x67\x14\x76\x9b\x91\x66\x1f\x2f\xbb\x1c\x84\xdd\xee\xae\xa6\x83\x3b\x96\xd3\x9a\xd4\x8f\x7e\xb4\x39\xc8\xa0\x54\xa5\x85\xcb\xec\xb8\x6a\xe8\x4d\x59\x2d\xa6\x7d\x6c\xd2\xac\x54\xf5\x85\xdd\xad\x2e\x87\x78\x99\xee\xd6\x15\x18\xc3\x99\x52\x23\xe1\xd5\x74\xec\x1c\xf9\xea\xbb\xc5\xfc\x62\xf5\xe1\xf8\xeb\xf9\xf9\xb1\xfa\xc8\x7f\xa1\xd7\xb8\x2d\x81\xf0\x7b\x5c\xd8\x2f\xea\x2f\xa9\xa9\x1a\x09\xa8\x66\xc6\x3a\x09\x3f\x48\x68\xc3\xaa\xe3\xeb\xea\x8d\x0d\xe2\x28\xbd\xc1\xfa\x9f\xf0\xc3\xc2\x73\x5c\xdb\x96\xf3\xcd\xe5\x0e\x3f\x67\xf2\x83\xfc\x31\x5e\x76\x8b\xf1\x37\xa9\xfc\x2a\xc6\x52\x71\x54\x87\x84\x8b\xb8\x1b\xf2\x1d\x9a\x6e\x16\xe4\x9b\x74\xef\x9b\xa3\x7b\xda\x4f\xff\x06\x00\x00\xff\xff\x70\xe6\x6e\x06\x53\x06\x00\x00")

func distHydrocarbonMinCssBytes() ([]byte, error) {
	return bindataRead(
		_distHydrocarbonMinCss,
		"dist/hydrocarbon.min.css",
	)
}

func distHydrocarbonMinCss() (*asset, error) {
	bytes, err := distHydrocarbonMinCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dist/hydrocarbon.min.css", size: 1619, mode: os.FileMode(420), modTime: time.Unix(1485017685, 0)}
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
	"dist/hydrocarbon.min.css": distHydrocarbonMinCss,
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
	"dist": &bintree{nil, map[string]*bintree{
		"hydrocarbon.min.css": &bintree{distHydrocarbonMinCss, map[string]*bintree{}},
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

