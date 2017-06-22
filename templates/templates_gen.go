// Code generated by go-bindata.
// sources:
// tpl/issue.tpl
// tpl/merge_request.tpl
// tpl/note_issue.tpl
// tpl/note_merge_request.tpl
// tpl/push.tpl
// tpl/push_file.tpl
// DO NOT EDIT!

package templates

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

var _issueTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7a\xb9\x70\xe7\xf3\xd9\xeb\x14\x6c\x12\x15\x32\x8a\x52\xd3\x6c\x95\xaa\xab\x15\xf4\x02\x8a\xf2\xb3\x52\x93\x4b\xf4\x3c\xf2\x73\x53\x0b\x12\xd3\x53\x15\x6a\x6b\x95\xec\x90\x25\xfc\x12\x73\x41\x82\x36\xfa\x89\x76\x8f\x1b\x26\x90\xa6\x17\x49\x02\xa2\x7f\xa2\xc2\xb3\x39\x9d\x2f\xa7\xaf\x7b\xb9\x68\xc6\xfb\x3d\x1d\x9e\xc1\xc1\xa1\xae\xef\xf7\x74\x3e\xed\x9f\xf1\xb4\x67\xda\xe3\x86\x26\xae\x67\xfd\x13\x9e\xec\x5a\xf2\x64\xd7\xae\xf7\x7b\x66\xa1\xba\xd3\xb3\xb8\xb8\x34\x55\x2f\xb4\x28\x07\x6e\x09\x44\xc4\xb1\xb4\x24\x23\xbf\x08\xcd\x95\x18\xd2\xae\xb9\x89\x99\x20\x9d\x8f\x1b\x26\x72\x3d\xed\x68\x7b\xd9\xda\xfb\x7c\xf7\x4c\x90\x1d\xd5\xd5\x99\x69\x0a\xe9\x25\x70\xe5\xc5\xc5\x99\xe9\x79\xa9\xa9\x9e\x2e\x0a\x06\xb5\xb5\xc4\x39\x00\xaa\x05\xb7\x13\x60\x0a\x90\x1c\x51\x5d\x9d\x9a\x97\x52\x5b\xcb\xf5\x6c\x72\xef\x93\xbd\x73\x88\xf7\x6c\x72\x49\x66\x7e\x1e\xd4\x12\xae\xe7\x5d\xdb\x9e\x35\x34\x12\xad\x39\xb8\x24\xb1\x04\xe6\x40\xae\x67\x0b\xda\xc1\x71\x40\xa4\xde\x90\xcc\x92\x1c\x84\xde\xb5\x8b\x9f\x4d\x6b\x27\x5a\xaf\x4b\x6a\x71\x72\x51\x66\x01\x92\xcb\x01\x01\x00\x00\xff\xff\x60\xcb\x65\x1a\x88\x02\x00\x00")

func issueTplBytes() ([]byte, error) {
	return bindataRead(
		_issueTpl,
		"issue.tpl",
	)
}

func issueTpl() (*asset, error) {
	bytes, err := issueTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "issue.tpl", size: 648, mode: os.FileMode(436), modTime: time.Unix(1498027020, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _merge_requestTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7a\xb9\x70\xe7\xf3\xd9\xeb\x14\x6c\x12\x15\x32\x8a\x52\xd3\x6c\x95\xaa\xab\x15\xf4\x02\x8a\xf2\xb3\x52\x93\x4b\xf4\x3c\xf2\x73\x53\x0b\x12\xd3\x53\x15\x6a\x6b\x95\xec\x90\x25\xfc\x12\x73\x41\x82\x36\xfa\x89\x76\x8f\x1b\x26\x90\xa6\x17\x49\x02\xa2\x7f\xa2\xc2\xb3\x39\x9d\x01\xa5\x39\x39\x0a\x41\xa9\x85\xa5\xa9\xc5\x25\x4f\xfb\x67\x3c\xed\x99\xf6\xb8\xa1\x89\xeb\x59\xff\x84\x27\xbb\x96\x3c\xd9\xb5\xeb\xfd\x9e\x59\xa8\x4e\xf4\x4d\x2d\x4a\x4f\x85\xaa\xd7\x0b\x2d\xca\x81\x5b\x83\x22\xe1\x58\x5a\x92\x91\x5f\x84\xe6\x5c\x5c\xaa\x5c\x73\x13\x33\x41\xe6\x3c\x6e\x98\xc8\xf5\xb4\xa3\xed\x65\x6b\xef\xf3\xdd\x33\x41\x16\x57\x57\x67\xa6\x29\xa4\x97\xa0\xeb\x2a\x2e\xce\x4c\xcf\x4b\x4d\xf5\x74\x51\x30\xa8\xad\x25\xc7\x71\x50\x03\x08\x3a\x0f\xa6\x0e\xc9\x81\xd5\xd5\xa9\x79\x29\xb5\xb5\x5c\xcf\x26\xf7\x3e\xd9\x3b\x87\xdc\xd0\x49\x2e\xc9\xcc\xcf\x83\xda\xcc\xf5\xbc\x6b\xdb\xb3\x86\x46\x32\x8d\x0a\x2e\x49\x2c\x81\xf9\x81\xeb\xd9\xae\x09\x4f\x3b\xda\x9e\x4d\x59\x4f\xae\x61\xf9\xa5\x45\xc9\xa9\x4e\x45\x89\x79\xc9\x19\x70\xd7\xcd\x5e\xf7\x6c\x41\x3b\x45\xc6\x86\x24\x16\xa5\xa7\x96\xa0\x1a\xfb\x6c\x41\xfb\xcb\x45\x33\xc8\x35\x30\xb3\x24\x07\xe1\xe9\xb5\x8b\x9f\x4d\x6b\x27\xd3\x24\x97\xd4\xe2\xe4\xa2\xcc\x02\xe4\xe8\x00\x04\x00\x00\xff\xff\x1f\x1e\x7e\xd8\x98\x03\x00\x00")

func merge_requestTplBytes() ([]byte, error) {
	return bindataRead(
		_merge_requestTpl,
		"merge_request.tpl",
	)
}

func merge_requestTpl() (*asset, error) {
	bytes, err := merge_requestTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "merge_request.tpl", size: 920, mode: os.FileMode(436), modTime: time.Unix(1498027017, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _note_issueTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7a\xb9\x70\xe7\xf3\xd9\xeb\x14\x6c\x12\x15\x32\x8a\x52\xd3\x6c\x95\xaa\xab\x15\xf4\x02\x8a\xf2\xb3\x52\x93\x4b\xf4\x3c\xf2\x73\x53\x0b\x12\xd3\x53\x15\x6a\x6b\x95\xec\x90\x25\xfc\x12\x73\x41\x82\x36\xfa\x89\x76\x8f\x1b\x26\x90\xa6\x17\x49\x02\xa2\x7f\xa2\xc2\xb3\x39\x9d\x2f\xa7\xaf\x7b\xb9\x68\xc6\xfb\x3d\x1d\x9e\xc1\xc1\xa1\xc1\xef\xf7\x74\x3e\x5d\xd2\xfe\x6c\xf3\x8a\xf7\x7b\x3a\xfc\xf2\x4b\x52\x35\x9f\xf6\xcf\x78\xda\x33\xed\x71\x43\x13\x97\x67\x71\x71\x69\xea\xb3\x05\xed\x60\xd5\xb3\x50\xdd\x0d\x96\xd3\x0b\x2d\xca\x81\x5b\x0a\x11\x09\xc9\x2c\xc9\x81\xd9\xc7\xf5\xac\x7f\xc2\x93\x5d\x4b\x9e\xec\xda\x85\xa1\x1d\x64\x13\x8a\x6e\xb0\x80\x63\x69\x49\x46\x7e\x11\x9a\x97\xd1\x65\x5d\x73\x13\x33\x41\xfa\x1e\x37\x4c\xe4\x82\xb8\xfc\xd9\xda\xc5\xcf\xa6\xb5\x13\x69\x07\x88\x80\xb9\x8f\x0b\x10\x00\x00\xff\xff\xbe\xe0\xa2\x1f\x92\x01\x00\x00")

func note_issueTplBytes() ([]byte, error) {
	return bindataRead(
		_note_issueTpl,
		"note_issue.tpl",
	)
}

func note_issueTpl() (*asset, error) {
	bytes, err := note_issueTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "note_issue.tpl", size: 402, mode: os.FileMode(436), modTime: time.Unix(1498028285, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _note_merge_requestTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7a\xb9\x70\xe7\xf3\xd9\xeb\x14\x6c\x12\x15\x32\x8a\x52\xd3\x6c\x95\xaa\xab\x15\xf4\x02\x8a\xf2\xb3\x52\x93\x4b\xf4\x3c\xf2\x73\x53\x0b\x12\xd3\x53\x15\x6a\x6b\x95\xec\x90\x25\xfc\x12\x73\x41\x82\x36\xfa\x89\x76\x8f\x1b\x26\x90\xa6\x17\x49\x02\xa2\x7f\xa2\xc2\xb3\x39\x9d\x01\xa5\x39\x39\x0a\x41\xa9\x85\xa5\xa9\xc5\x25\x4f\x97\xb4\x3f\xdb\xbc\xe2\xfd\x9e\x0e\xbf\xfc\x92\x54\xcd\xa7\xfd\x33\x9e\xf6\x4c\x7b\xdc\xd0\xc4\x85\xac\xe6\xd9\x82\xf6\x97\x8b\x66\xbc\xdf\x33\x0b\xd5\xe5\xbe\xa9\x45\xe9\xa9\x50\x25\x7a\xa1\x45\x39\x70\xdb\x51\x24\x42\x32\x4b\x72\x60\xf6\x73\x3d\xeb\x9f\xf0\x64\xd7\x92\x27\xbb\x76\x61\x18\x06\xb2\x1e\xc5\x10\xb0\x80\x63\x69\x49\x46\x7e\x11\x5a\x10\xa0\xcb\xba\xe6\x26\x66\x82\xf4\x3d\x6e\x98\xc8\x05\xf1\xce\xb3\xb5\x8b\x9f\x4d\x6b\x27\xd2\x0e\x10\x01\x73\x1f\x17\x20\x00\x00\xff\xff\x08\x02\x85\x7e\xa2\x01\x00\x00")

func note_merge_requestTplBytes() ([]byte, error) {
	return bindataRead(
		_note_merge_requestTpl,
		"note_merge_request.tpl",
	)
}

func note_merge_requestTpl() (*asset, error) {
	bytes, err := note_merge_requestTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "note_merge_request.tpl", size: 418, mode: os.FileMode(436), modTime: time.Unix(1498027769, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pushTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7a\xb9\x70\xe7\xf3\xd9\xeb\x14\x6c\x12\x15\x32\x8a\x52\xd3\x6c\x95\xaa\xab\x15\xf4\x02\x8a\xf2\xb3\x52\x93\x4b\xf4\x3c\xf2\x73\x53\x0b\x12\xd3\x53\x15\x6a\x6b\x95\xec\x90\x25\xfc\x12\x73\x41\x82\x36\xfa\x89\x76\x8f\x1b\x26\x90\xa6\x17\x49\x02\xa2\x7f\xa2\xc2\xb3\x39\x9d\xcf\xa6\x6d\x78\x3e\xab\xe5\x59\xff\x84\x27\xbb\x96\xbc\xdf\xd3\x11\x10\x1a\xec\xf1\x7e\x4f\xe7\xe3\x86\x26\x2e\x88\xd8\x8b\x86\xd6\xf7\x7b\x66\x21\x9c\xa9\x0c\x31\x34\xb4\x38\xb5\x08\xcd\x35\x70\x61\xd7\xdc\xc4\xcc\x1c\x85\xda\xda\xc7\x0d\x13\xb9\x9c\xf3\x73\x73\x33\x4b\x9e\x4d\xdd\xf0\xb2\xbd\x1f\x9b\x39\x10\x79\xbd\x90\xfc\x92\xc4\x1c\x08\xbb\xd8\x39\xbf\x34\xaf\x04\x6a\x2c\xd7\xd3\x8e\xb6\x67\x53\xd6\xbf\xdf\x33\x0b\x87\xc6\xa0\xd4\x34\x98\xd2\xe0\x8c\x44\xdc\xea\x9c\x33\x52\x93\xb3\xf3\x4b\x4b\x82\x33\x12\xa1\xea\x01\x01\x00\x00\xff\xff\x57\x78\x07\x05\x80\x01\x00\x00")

func pushTplBytes() ([]byte, error) {
	return bindataRead(
		_pushTpl,
		"push.tpl",
	)
}

func pushTpl() (*asset, error) {
	bytes, err := pushTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "push.tpl", size: 384, mode: os.FileMode(436), modTime: time.Unix(1498021148, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _push_fileTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x91\x41\x4b\x02\x41\x14\xc7\xef\xfb\x29\xde\xc1\x63\x2c\x9d\x85\x0e\x61\x81\x1d\x0a\x49\xfd\x00\x83\x3e\xdd\xad\x9d\x1d\xd9\x1d\x21\x18\x06\x96\xa8\xd4\x93\x06\x95\x60\x76\x08\x3c\x74\x2a\xa1\x08\x8a\xed\xd3\xa8\x63\x9e\xf6\x2b\x84\x3b\xa9\xdb\x1e\x3a\xed\xe3\xff\xfb\xbf\xff\x63\xff\xb3\x78\xfc\x98\xdf\x3f\x83\x10\x60\x16\x3c\x76\x82\x15\x6e\x1e\x11\x8a\x20\xe5\x24\xe8\x25\xd5\x3c\xa3\xd8\x20\x75\x4d\xae\x41\x0d\x3b\xea\x6e\x3c\x1f\x5c\xa8\x6e\x6f\xfa\x39\x8a\xc2\x76\xa1\x5c\xcc\x47\x61\x67\x12\x9c\x1b\x5a\xfb\x0e\x2e\xa3\x70\x10\x47\x97\x7d\xf4\x52\xb9\xb1\xb4\x4f\x89\xed\xe8\x44\x23\xc7\x28\xb5\xb9\xba\x1d\x2f\x5a\xdd\xd5\x9e\xd6\xcc\x12\xe3\xc4\xd1\xb3\x9f\x63\x4d\x97\x83\x94\xc6\xac\x7d\xa5\x6e\x5e\xa2\x70\x90\x30\x1e\x63\x6d\x89\x8a\x16\xf9\xab\xe7\x2c\xac\x9c\xb2\x26\x2f\x5a\x64\xc9\x85\xb0\x6b\x50\xe7\xff\xe4\x6f\xc7\x2e\x8f\xb8\x75\x84\x8c\xed\x56\xf1\x6c\x0b\x32\xe8\x20\x45\x97\x43\x76\x67\xb5\xe9\x2f\xd3\xf4\x98\x05\x21\x56\x0e\xf3\x60\x4f\x4a\x03\x60\xfa\x35\xdc\x94\xb0\x86\xbb\x4d\x6e\xb1\x54\x1d\x69\x98\x2c\x06\x40\xf5\xdf\x17\xfd\x37\xfd\x4b\x6b\x67\xc9\xa6\xe8\x73\x42\x1b\x10\xdf\x9a\x0d\xc7\xb3\x87\x20\x7d\xab\xec\x39\xbf\x78\xd4\x52\xaf\x4f\x69\x7c\x88\xbe\xaf\xdf\xd4\x10\x02\xdd\xea\xe6\xfb\x13\x00\x00\xff\xff\x58\x2c\x5d\xfa\x18\x02\x00\x00")

func push_fileTplBytes() ([]byte, error) {
	return bindataRead(
		_push_fileTpl,
		"push_file.tpl",
	)
}

func push_fileTpl() (*asset, error) {
	bytes, err := push_fileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "push_file.tpl", size: 536, mode: os.FileMode(436), modTime: time.Unix(1498021079, 0)}
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
	"issue.tpl": issueTpl,
	"merge_request.tpl": merge_requestTpl,
	"note_issue.tpl": note_issueTpl,
	"note_merge_request.tpl": note_merge_requestTpl,
	"push.tpl": pushTpl,
	"push_file.tpl": push_fileTpl,
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
	"issue.tpl": &bintree{issueTpl, map[string]*bintree{}},
	"merge_request.tpl": &bintree{merge_requestTpl, map[string]*bintree{}},
	"note_issue.tpl": &bintree{note_issueTpl, map[string]*bintree{}},
	"note_merge_request.tpl": &bintree{note_merge_requestTpl, map[string]*bintree{}},
	"push.tpl": &bintree{pushTpl, map[string]*bintree{}},
	"push_file.tpl": &bintree{push_fileTpl, map[string]*bintree{}},
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
