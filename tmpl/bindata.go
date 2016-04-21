// Code generated by go-bindata.
// sources:
// tmpl/csharp/server.tmpl
// tmpl/golang/server.gogo
// DO NOT EDIT!

package tmpl

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

var _tmplCsharpServerTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x54\xdf\x6f\xdb\x36\x10\x7e\x96\xfe\x8a\x83\x11\x60\x52\xe0\x32\x7b\xae\xd1\x87\x34\x0d\xb6\x00\x4b\x53\xd4\x06\x06\xec\xa5\xa0\xa9\xb3\xcd\x8d\x22\x05\x92\x72\xa3\x1a\xfa\xdf\x77\xa4\x28\x3b\xfe\xd1\xa2\x1b\x30\xac\x41\x90\x88\x77\xdf\x7d\x1f\xef\x17\x77\xbb\x2b\x5f\x37\xaa\xe2\x9e\xc3\xeb\x37\xc0\xe0\x55\xdf\xe7\x37\xd7\xd7\x39\x5c\xc3\x62\x23\x1d\xac\xa4\x42\xa0\xff\xbc\xf5\x66\x8d\x1a\x2d\xf7\x58\xc1\xb2\x03\xfc\x62\x1b\x11\x60\xef\x8c\xfe\xc9\x83\xd8\x70\xbd\x46\xa8\xb9\x6e\xb9\x52\x1d\x39\x6e\xf2\xbc\x75\x52\xaf\x61\xde\x39\x8f\xf5\xec\xe8\xc4\xee\x8c\x52\x28\xbc\x34\xda\x7d\xdd\xc3\x7e\x09\x8a\x52\x9c\x20\x16\xf8\xec\x4f\x4c\x0f\x4f\xa3\x61\xb1\xb1\x72\xe5\x8f\x4f\xdf\x50\xfb\xd8\x6a\x2f\x6b\x64\x73\xd2\xe1\x4a\x7e\xe1\x01\x74\x12\xfd\xc1\x1a\x6f\x84\x51\x27\xe6\x85\xe5\xda\x35\xc6\xee\xc5\xde\xdf\x2e\xe6\xec\x4e\x49\xd4\x64\xca\x35\xaf\xd1\x35\x5c\x20\xec\x76\xec\xfd\x78\xe8\x7b\xd8\xe5\x99\xf3\xa4\x23\x40\x28\xee\x5c\x70\x93\xfa\x56\x0a\x8c\xb0\xbe\x4f\xa7\x00\xcc\x6e\x6e\x80\xe4\x05\x3a\x67\x6c\xe8\xc3\x51\x0f\xd2\xc1\x58\x02\x36\x56\x6e\xc9\x0e\x89\xfa\x10\xd4\x8c\x5f\x9f\x66\xe7\xb0\x87\x3b\xa3\xf5\x50\x19\x10\xf4\x49\x98\x00\x6a\x97\x8a\x9c\x09\xb3\x35\xb2\x02\xa9\xa5\x2f\x2e\x90\x4e\x8f\x29\xca\x78\xe9\xec\xa0\x09\x6f\x0e\xd8\xa0\x9f\x45\x15\xb2\x8a\xa8\x94\x09\x36\x6f\x97\x4e\x58\xb9\xc4\x5b\xd7\x69\x51\x4c\x9e\x34\x3b\xab\x08\xbb\x9e\x4c\xc1\xe8\xb7\xd6\xf0\x4a\x70\xe7\xcb\xd9\xe5\xd8\xcb\x81\x93\x38\xab\x91\xe1\xd1\xad\x63\x6c\x9f\x9f\xd7\x22\xe6\x19\x21\x85\x59\xfe\x49\x19\x81\x43\x5d\x21\xa5\x48\xa6\x5f\xb9\xae\x14\xda\xfb\x2d\x35\xf7\xd6\xae\x1d\x60\x4a\x75\xcb\x2d\xf8\x30\x09\x94\x93\xc6\xcf\x67\xc3\xc1\x16\x8f\x58\x1b\xdb\xbd\x6d\x57\x2b\xb4\x05\xb2\x47\xaa\x05\x5f\x23\x7b\x47\x3b\x37\xe4\x11\x28\x24\x55\xc9\x27\x8a\xfb\x3f\xe2\xc4\x15\x91\xb6\x9c\x01\xcd\x80\xd4\x01\xe8\x6d\x37\x88\x42\x76\xc4\x0a\xcb\x76\x15\x99\x80\x7e\x52\x34\x98\xc0\xb8\xb7\x52\x0a\x50\xbb\xf5\x50\xf4\xcc\x7d\x96\x5e\x6c\xe0\xc5\x6d\xa8\x92\x21\xe5\x94\x54\xb6\xdb\xbd\x02\x1b\x37\xfa\x2a\xcc\xf1\x14\xae\x6a\xf4\x1b\x53\xc5\x47\x62\x2c\xf1\x63\x34\x39\x7a\x31\xc6\x10\xb9\x02\x4d\x69\x14\x0f\x6e\xdf\xaa\x01\x34\xc6\x97\x09\x4c\x1e\x84\xc9\x8b\xd7\xe7\xb4\x6f\xe4\xd2\xf1\x6b\xf2\x3a\x06\x8c\xf4\x89\x87\x3d\x51\xa1\x78\x97\xd8\x5e\xcc\xdb\x21\xf0\x53\x9a\xd6\xe2\xe7\xe9\x50\xdd\x29\xe8\x56\xa9\xa1\xe4\x03\x1f\x2a\x87\x23\x05\x95\xf0\xfb\x3a\x38\x12\x98\x0b\x1d\x23\x92\xa1\x5f\xa6\xf5\xf9\x3f\xba\x5a\x64\x2b\x53\x7b\x32\xea\x54\x62\x0e\xd3\x38\x2a\x92\x75\xec\x13\x79\x0f\xbd\xfb\x88\x8d\xea\x5e\x60\xc2\x64\x11\x80\x2e\x43\x0f\xa8\x3f\x5c\x7b\x40\x14\xc5\x43\xdc\x96\xb4\x3b\x4d\x58\xdc\x72\x98\xf4\x92\x1d\x76\x99\x7d\x08\x8f\x80\xdb\x14\x75\x5a\x99\xb1\x66\xba\xda\x97\xcc\x22\xff\x6b\x96\x9f\x7b\x8e\x4f\xf1\x6f\x0f\x82\xc7\x91\xbb\x7f\x16\x18\x35\x01\x9f\xc7\x69\x23\x55\x67\x14\xb2\xdf\xad\xf4\xf8\x9b\xd4\x58\xe0\x33\x5b\x98\xb9\xb7\xf4\xac\x16\xe5\x20\xdf\x7f\x7b\x6d\xf7\x03\xf7\xc3\x2f\xef\x7f\xb9\x7d\x5f\x5d\xbc\xa3\xbd\x8b\x4f\xec\xf7\xaf\xde\xbf\xd8\xae\xff\x63\x34\xe8\xb7\xff\x3b\x00\x00\xff\xff\x8c\xaa\x8a\x59\xd0\x08\x00\x00")

func tmplCsharpServerTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplCsharpServerTmpl,
		"tmpl/csharp/server.tmpl",
	)
}

func tmplCsharpServerTmpl() (*asset, error) {
	bytes, err := tmplCsharpServerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/csharp/server.tmpl", size: 2256, mode: os.FileMode(420), modTime: time.Unix(1461204543, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplGolangServerGogo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x56\x4b\x6f\xdb\x38\x10\x3e\x5b\xbf\x62\x56\x08\x36\x92\x61\xcb\x48\x8e\x06\x72\x49\xb2\xc0\xe6\x10\x27\x1b\x07\xd8\x33\x25\x8d\x6c\xb6\x12\xa9\xf2\x51\xc3\x11\xfc\xdf\x3b\xa4\xa4\xc4\xaf\x36\x4e\x0f\x45\xd1\x43\x62\x71\xf4\xf1\xe3\xcc\x37\x0f\xaa\x69\x72\x2c\xb8\x40\x08\x35\xaa\xaf\xa8\x42\x18\x6f\x36\x41\xcd\xb2\xcf\x6c\x81\xd0\x34\xc9\x8c\x55\xa8\x69\x8d\x64\x6e\x9a\x33\x87\xe2\x19\x3a\x2b\x4c\xaf\x20\x99\xb7\x6b\x0f\x23\x04\xaf\x6a\xa9\x0c\x44\xc1\x20\x4c\xd7\x06\x75\x18\xd0\xd3\x82\x9b\xa5\x4d\x93\x4c\x56\x13\xc1\x8c\x1e\x73\xe9\x7f\xc3\xdd\x57\x9a\x55\x16\xcb\xc9\x42\x8e\xcd\x52\xf1\xc2\x4c\xda\x9f\x3d\x14\xbe\xa4\x76\x3d\xd1\x86\xf6\xe7\x61\x10\x07\x81\x59\xd7\x08\xcf\x1e\x3a\x23\xe3\xae\x8b\x9b\xcd\xdc\x47\x05\xda\x28\x9b\x19\x68\x82\x41\x67\x18\x1e\x07\x06\x83\x1b\x29\x04\x00\x0c\x9d\x87\x89\x5b\x04\x9b\x20\x28\xac\xc8\x20\xd2\x30\x7c\xef\xa0\x18\xa4\xb8\xd7\x8b\xa8\xd2\x8b\x8e\x82\x56\xb1\x3b\x57\x39\xb9\xda\x90\x92\x19\xae\x6e\x64\x45\xa2\x9a\x47\x25\x8d\xcc\x64\xf9\x84\x2c\x47\x15\x79\xcd\xdc\xeb\x6e\x4d\x34\xc9\x2d\x33\x2c\xa6\x40\x07\x7a\xc5\x4d\xb6\x04\x67\x9b\xdb\xf4\x13\xfa\x78\x9a\x66\x0c\x8a\x09\xca\xd5\x59\x85\x66\x29\xf3\x9d\xac\xdc\x7b\x93\xf6\xa9\x1b\x03\x2f\x40\x48\x4a\xce\x9d\xbe\x56\x92\xe5\x19\xd3\xa6\x05\xf4\x7b\x63\x02\x0e\xc8\x4c\xd5\xb0\x1f\x5e\x42\x86\x16\xd4\xa5\x3a\x9c\xf6\x9c\xbd\xfd\x41\xe0\x8a\xad\x1d\xc5\xa0\x76\x5e\xfc\xbd\xcf\xb1\x4f\xf1\x84\x5f\x2c\x6a\xd3\xb8\x1d\xa8\xb6\x05\xba\xc5\x4c\xe6\x38\xf7\x49\x8b\xd4\x08\xea\x98\x20\x74\x94\x43\xfd\x75\x05\x82\x97\x4e\x52\x3a\x47\x71\x61\x4a\x11\x91\xdd\x21\x7a\xa2\x2b\xd0\x49\x9b\x8f\x03\xb7\xa3\x93\xa9\xe8\x5f\x5b\x66\xc9\x9d\xc8\x54\xf4\xbe\x22\x54\x9f\x56\x98\x30\xf6\xba\x60\xa9\xd1\x4b\x61\x2e\x5c\x5c\x1d\xd3\x4c\xae\x22\x97\xca\x9f\x10\x48\xa1\x3e\x71\x8f\xae\xa5\xd0\xf8\x8b\x55\x1d\x01\x39\xf8\x01\x69\x53\x5b\xf8\x70\xda\x8a\xbf\xb6\x45\x81\xca\xbb\xbc\xfa\x61\x9f\xfc\xaf\xb8\x71\x7d\x62\x0b\xc7\xd4\xc1\xfe\x11\x5b\x71\xad\x5e\x3d\x69\xdb\x37\x79\xb4\x69\xc9\xf5\xd2\xf7\xd2\x13\xd6\xe5\x7a\x04\xb4\x3d\xb9\x76\x07\x47\xbe\xb1\x06\xe6\xf2\x20\x47\xaf\xc9\x7f\xe6\x15\x17\x8b\x53\xd2\x6f\x3c\x32\x1c\x81\xb9\xa0\xbf\xcb\xae\x0e\x44\xde\x75\xdf\xfe\x13\xe9\xf0\xd1\xc9\xf2\xda\xb6\x87\xf3\xa5\xeb\xc5\xe4\x5f\xb6\xdf\xdc\xee\xa8\xdf\x63\xf8\x7c\x77\xee\xbc\x8d\x9d\x07\x91\x9c\x32\x79\xfe\xcc\x09\x73\x4a\xf0\xbb\x43\xe6\xb0\xa4\xde\x16\x7d\x71\x51\x42\x8f\x57\x54\x44\x97\x35\xb9\xb6\xf7\x6e\x04\x99\xbb\x01\xdf\xee\x3f\x7f\x7b\x1d\x9f\x3d\x2d\x4f\x73\x47\x3c\x58\xa1\xa0\x58\xb8\x14\x53\x70\xbc\x2e\xba\xf6\x8b\xc2\xef\x7c\xaf\xba\x9d\x40\xed\xd3\x14\xf4\x88\x16\xee\xe4\x29\xdd\xc3\xce\x99\x51\xe0\xc5\x9a\x4c\x80\x95\x25\xa4\x7d\x0d\x01\x7d\x98\x68\xfa\x4a\xd1\xa0\x97\xd2\x96\x39\xa4\x08\x56\x50\xe5\x82\xe8\xbf\x59\xe0\xfc\x41\x9c\xf7\x8e\xb4\xe3\x80\xaa\x57\x67\x8a\xa7\x78\x5c\xef\x21\xf5\x6f\x07\xdf\xea\xb7\x78\x97\xe3\x3f\x8b\x16\xb7\x88\x8e\xb2\x84\xf8\xa2\xea\x6c\x9b\xce\xb5\x6a\xb0\x9d\xa2\x6f\x01\x00\x00\xff\xff\xf0\x63\xeb\x10\x7d\x09\x00\x00")

func tmplGolangServerGogoBytes() ([]byte, error) {
	return bindataRead(
		_tmplGolangServerGogo,
		"tmpl/golang/server.gogo",
	)
}

func tmplGolangServerGogo() (*asset, error) {
	bytes, err := tmplGolangServerGogoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/golang/server.gogo", size: 2429, mode: os.FileMode(420), modTime: time.Unix(1461204543, 0)}
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
	"tmpl/csharp/server.tmpl": tmplCsharpServerTmpl,
	"tmpl/golang/server.gogo": tmplGolangServerGogo,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"csharp": &bintree{nil, map[string]*bintree{
			"server.tmpl": &bintree{tmplCsharpServerTmpl, map[string]*bintree{}},
		}},
		"golang": &bintree{nil, map[string]*bintree{
			"server.gogo": &bintree{tmplGolangServerGogo, map[string]*bintree{}},
		}},
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

