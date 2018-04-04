// Code generated by go-bindata.
// sources:
// enum.tmpl
// DO NOT EDIT!

package assets

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

var _enumTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x51\x6f\xdb\x36\x10\x7e\x36\x7f\xc5\x57\x21\x05\xc4\xcc\x91\x3b\xec\x6d\x83\x9f\xd6\x2e\xe8\x43\xda\x02\xce\xf6\x12\x14\x05\x63\x9d\x1c\x22\x12\xa5\x51\xb4\xa7\x80\xe0\x7f\x1f\x8e\x92\x6d\x59\xad\x82\xbe\x91\xbc\xfb\xbe\xbb\xef\xee\x78\xde\xdf\x20\xa7\x42\x1b\x42\xf2\x44\x2a\x27\x9b\x84\x20\x56\x2b\xfc\x59\xe7\x84\x1d\x19\xb2\xca\x51\x8e\xc7\x17\xec\xea\x1b\x32\xfb\x8a\x8d\xef\x3f\xe3\xd3\xe7\x7b\x7c\x78\xff\xf1\xfe\x8d\x10\x8d\xda\x3e\xab\x1d\xc1\xfb\x6c\x38\x86\x20\x84\xae\x9a\xda\x3a\xa4\x02\x00\x92\xa2\x72\x89\x90\xc2\x7b\x32\x39\x6e\xd8\x3e\x8e\xcc\xbc\x1c\x77\x5b\x9b\x96\x21\x6c\xbb\xe2\xc7\x4f\xaa\x22\xfc\xbe\x46\xc6\x97\x2c\xde\x18\xec\x3d\xac\x32\x3b\xc2\x95\xfd\x68\x72\xea\x96\xb8\x3a\xa8\x72\x3f\x72\xfd\x87\xaf\x2d\x42\x10\x0b\xef\xa1\x0b\xd0\xbf\x83\x4f\xcf\x92\x7c\x4b\x42\x58\xad\xb0\x79\xd6\x4d\x43\x39\xa2\xc9\x7b\x2a\x5b\x8a\xef\xde\x0f\xde\x5f\x2c\x15\xba\xa3\x9c\x51\x21\x40\xb7\x50\x6c\x3c\x26\x17\x02\xea\x02\xee\xa5\xa1\x33\xa4\x7f\x8f\x52\x43\x88\xf2\xe7\xd8\xce\xa9\xf5\x3a\xf0\x0e\x0c\x1c\xb3\xaf\xa1\x6b\xa7\x06\xb2\x01\x60\x4e\xc2\x8f\x9a\xa2\xdc\x19\x2c\x7e\x81\xf7\xed\xfe\xf1\xd2\xf7\x48\x70\x4e\x94\x8b\xde\x9f\x24\x77\x07\x8e\xaa\xa6\x54\x8e\x90\xb4\xce\x6a\xb3\x23\x9b\x20\xe3\x8a\xf2\x04\x7c\x51\xb6\x25\xef\xcf\x7d\x09\x01\xca\x31\xc4\xb5\x70\x35\xb6\xb5\x39\x90\x75\x50\xe8\xc1\xfc\xc6\x85\x1b\x03\x44\xb1\x37\xdb\x1f\x31\xa5\x86\x5b\xd4\x03\x25\xd2\x4b\xe3\x12\x64\x6d\x6d\x25\xbc\x58\xe8\x02\xdd\x12\xf5\x33\xf7\xfd\xdb\xa5\x5b\x54\xf9\xc0\x44\x5f\xff\x60\x0f\x2f\x16\x0b\x4b\x6e\x6f\xcd\x24\x8b\xb4\x93\x4b\x18\x5d\x8a\x45\x10\x33\x1e\xef\xe4\x12\x45\xe5\xb2\x0f\x1c\xb8\x48\x93\xb7\x2d\xcf\x81\xa9\x59\xde\x41\x95\x3a\x9f\x00\x92\x25\x38\xb0\x14\x71\xcc\xb9\x63\x59\xa5\x6c\xfb\xa4\x4a\x1c\x55\xa7\x1d\xae\x2f\x41\x12\x77\xbd\xcf\x3d\x75\x2e\x95\x48\x1f\xbe\x3e\xbe\x38\x1a\xcb\x1d\xb2\xeb\x0d\x69\x97\x6d\x62\x85\x52\x39\x08\x08\x62\x9e\xfb\x6f\x53\x8d\xd8\x1d\x75\x6e\xa0\x91\x3d\x3f\xd3\x9b\xe1\xb3\xf5\x85\x8f\x4e\x52\x2c\x5c\xd5\xc4\x1c\xd8\x32\xd7\x2b\x19\x5b\xc1\x4e\x6f\xd6\x9c\xca\xb8\xda\x64\x6d\x2c\xed\x75\x87\x35\x5c\xd5\x9c\x64\xf4\x29\x1f\xa7\xef\x58\xa8\xa2\x54\x3b\xf4\x2b\x68\x43\x0e\xba\x6a\x4a\xaa\xc8\xf0\x54\x3d\x11\x6e\xeb\x52\x99\x1d\xd8\x69\x98\x63\x6d\x1c\xd9\x42\x6d\x09\xac\x7d\xbe\x00\x1b\x72\xe9\x41\x95\xa7\xa9\x3a\xa9\x3e\xbc\xaa\xee\xa0\x4a\x39\xe4\x7e\x10\x63\x49\xfd\x37\xb8\x7d\x3d\xc5\x5b\x72\x8e\xec\x4f\xe7\x78\x4b\xdc\xf8\x93\xb7\x0f\xa3\xa6\x5f\x77\x43\xc8\x7b\xde\x35\x93\x98\x3b\xed\x9e\xf6\x8f\xd9\xb6\xae\x56\x6d\x53\xfc\xfa\xdb\xaa\xf9\x8b\xcb\x38\xad\xd0\x7c\x64\x26\x4d\xe5\xf1\xaf\x9e\xa3\x26\x93\xc1\xbe\x6c\xd8\x70\xb8\x58\xe5\xa7\x5d\x71\x5a\xe7\x93\x8f\x19\xd7\xef\x9a\x99\x87\x70\xba\x78\xe9\x97\x36\x38\x80\x38\x28\x3b\x85\xdc\xa9\x06\x6b\xde\x7d\x95\x6a\xc6\xde\xe2\x38\xf0\x1a\xdf\xb5\x7b\xf8\x1a\x23\x49\xba\xe0\xcb\xcc\xba\xb8\x53\xcd\x83\xfe\x6e\x53\xb4\xce\x8e\xd7\x02\xef\x80\x4d\x63\xb5\x71\x45\x3a\xa9\x4c\xfa\x36\x97\xc9\x12\x3a\xfe\xf9\x1f\x48\xe8\x5b\x11\x45\xec\xcd\x85\x8c\xac\xac\xff\x23\xbb\x55\x2d\x61\x54\xd5\xff\x03\x00\x00\xff\xff\x1f\xbf\x0e\xd9\x98\x07\x00\x00")

func enumTmplBytes() ([]byte, error) {
	return bindataRead(
		_enumTmpl,
		"enum.tmpl",
	)
}

func enumTmpl() (*asset, error) {
	bytes, err := enumTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "enum.tmpl", size: 1944, mode: os.FileMode(420), modTime: time.Unix(1522866202, 0)}
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
	"enum.tmpl": enumTmpl,
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
	"enum.tmpl": &bintree{enumTmpl, map[string]*bintree{}},
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

