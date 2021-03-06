// Code generated by go-bindata.
// sources:
// schema/init.sql
// DO NOT EDIT!

package hydrocarbon

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

var _schemaInitSQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x96\xcf\x73\xda\x3a\x10\xc7\xcf\xf6\x5f\xb1\x37\x60\x1e\x79\xf3\x5e\xa6\xd3\x4b\x4e\x0e\x28\x53\x4f\x89\x49\x6d\xb9\x43\x7a\xf1\xa8\xd6\x02\x1a\x8c\xe4\xea\x47\x02\xff\x7d\xc7\xbf\x12\x4c\xc2\xb4\xe4\x90\xc9\x05\xc3\xf2\x95\x56\xfb\xfd\x68\x17\x26\x31\x09\x28\x01\xb2\xa0\x24\x4a\xc2\x79\x04\xe1\x0d\x44\x73\x0a\x64\x11\x26\x34\x81\x72\x95\xeb\x7d\x69\xd5\x95\xff\x07\x61\x2e\x2c\xee\xec\x95\xdf\xe9\x68\x70\x3d\x23\xe0\x0c\x6a\x03\x43\xdf\x13\x1c\xd2\x34\x9c\xc2\x5d\x1c\xde\x06\xf1\x3d\x7c\x25\xf7\x30\x25\x37\x41\x3a\xa3\xb0\x42\x99\x69\x26\xb9\xda\x66\xce\x09\x3e\x1c\x8d\x7d\xdf\xcb\x35\x32\x8b\x3c\x63\x16\x68\x78\x4b\x12\x1a\xdc\xde\xd1\x1f\x75\xc6\x28\x9d\xcd\x9e\x16\x4b\xf5\x58\x2d\xf0\x5c\xc9\xcf\xd1\xfb\x9e\xb1\x5a\x94\x98\xe5\xce\x58\xb5\x45\x9d\x09\x0e\x94\x2c\xe8\xd8\xf7\x70\xcb\x44\x01\x93\xb0\xfa\xf8\xb4\x41\xb5\x64\x32\x8f\x12\x1a\x07\x61\x44\xa1\xd6\x64\x4e\x8a\x5f\x90\x46\xe1\xb7\x94\xc0\xb0\x0e\x8d\xfc\xd1\x95\xef\x5f\x5c\x40\xa1\x56\x42\x82\x55\x1b\x94\x06\x98\x46\x50\x12\x2f\xac\xd8\x62\x17\x73\x06\x39\x58\xd5\x0a\xc5\x12\x14\x73\x76\x0d\x97\xff\xfe\x07\xc2\x80\x54\xb6\x56\xf4\xed\xac\xb5\x59\xbb\xc1\xf9\xae\x7a\x15\x8e\xac\x5b\x14\x93\x1b\x12\x93\x68\x42\x92\x16\xd3\x61\xa9\xe7\xda\x8f\xbb\x52\x68\x34\x7f\xa7\x87\x7f\x20\x8c\x28\x89\xbf\x07\x33\x18\x5c\x7e\x82\x2f\xf3\x34\x4e\x06\x55\xda\xfa\x80\x6c\x85\xd2\xc2\x91\xfb\x9e\x28\x61\x12\x4e\xe3\xde\x29\x6b\x27\x1a\x65\x97\x00\x65\xae\x38\x0e\x0f\xaa\xff\xb9\xb7\x68\x86\xff\x7f\x1e\x8d\x61\xb0\xc6\xdd\xa0\x35\x82\xc3\xf5\x7c\x3e\x23\x41\xf4\xf2\x90\x4b\x56\x18\xac\x41\xb6\xee\xb7\x88\xc3\x68\x4a\x16\x3d\x08\xcd\x23\x13\x7c\x07\xf3\xe8\x08\x4f\xfd\x6c\x6f\x83\x41\x63\x84\xea\xdd\x87\x8e\x7f\x45\x1d\xa5\x15\x39\xb3\xd8\xe9\x4c\x1f\x7b\x17\xfd\x58\xc8\xcf\x64\xb5\xc1\xfd\x5b\x48\xb1\xdc\x8a\x07\x3c\xcd\xca\x6a\x77\x1a\x55\x67\x5c\xb6\xc1\x7d\x07\xe9\xd9\xcc\x0d\xee\x5b\x3c\x25\xdb\x6f\x51\xda\x06\xcc\xe3\x1a\xab\x57\x04\x8d\xb9\xd2\x1c\x58\x51\x3c\x0b\xb6\x8c\x23\x3c\x08\x06\xcd\xf8\xe8\x73\x7a\x52\x7d\x2c\x4e\xe7\x4c\xc6\x6e\x7c\x2d\x11\x79\xb6\x54\x05\xaf\xb2\x57\xae\x74\xef\xd5\xb2\xfe\xee\xe8\x86\xf6\xe4\x1f\x6e\xdc\x1f\x16\xd5\x54\x23\x24\x17\x0f\x82\x3b\x56\x9c\xaa\xe6\x4d\x10\x0f\x6c\x78\x8d\x65\xcf\xa5\x77\x43\x5a\x25\x28\x0b\x57\xfd\xca\x1c\xf7\xa9\xd3\xc5\x8b\x98\x15\xb6\xc0\x7e\xf4\x64\x7f\xd5\x05\x35\x7b\x67\x4e\x17\x5d\x8b\xd5\xfe\x0d\x9b\xf8\x18\x9c\x2e\x46\xc7\xff\x0b\x4a\x65\xde\xd6\x26\x6b\x66\xd6\xc7\x47\x7e\x07\xff\x5e\x31\x65\xec\x7b\xb9\x92\xf6\xc5\xf8\xab\xcc\xfa\x1d\x00\x00\xff\xff\x01\xe4\x7d\xdf\x55\x09\x00\x00")

func schemaInitSQLBytes() ([]byte, error) {
	return bindataRead(
		_schemaInitSQL,
		"schema/init.sql",
	)
}

func schemaInitSQL() (*asset, error) {
	bytes, err := schemaInitSQLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/init.sql", size: 2389, mode: os.FileMode(420), modTime: time.Unix(499137600, 0)}
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
	"schema/init.sql": schemaInitSQL,
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
	"schema": {nil, map[string]*bintree{
		"init.sql": {schemaInitSQL, map[string]*bintree{}},
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
