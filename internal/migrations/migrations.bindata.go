// Code generated by go-bindata. DO NOT EDIT.
// sources:
// 1569155140_baseline.down.sql (719B)
// 1569155140_baseline.up.sql (4.411kB)

package migrations

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

var __1569155140_baselineDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x0a\x83\x30\x10\x45\xf7\x9e\x22\xb8\xef\x09\x5c\xb5\x98\x42\xa0\x68\xa9\x2e\xdc\x85\xa2\xd3\x76\xa8\x98\x92\x49\xa0\xc7\x2f\x9a\x0a\x0a\x4a\x66\xa7\xe1\xcd\x7f\x99\x4f\xf2\x5b\x79\x15\xaa\xc8\x65\x23\xd4\x59\xc8\x46\x55\x75\x25\x2c\x3c\x2c\xd0\x4b\x3b\xf3\x86\x41\x7b\x1c\x3a\xf8\x66\xc9\x26\x7a\x6f\x5b\x20\xe2\x90\x04\x44\x68\x06\xd2\xd8\xad\xc1\xfa\x78\xba\xc8\x0d\x30\x4b\xf6\x94\x6e\xca\xb1\xa6\x07\xd2\xe1\x4f\x3b\x74\x3d\x4c\x47\xd1\xfc\xd5\xfc\x9e\xc4\x13\xd8\x59\x31\x7e\x8f\xa9\xac\xf4\xc5\x60\x6c\x81\x70\xe7\x48\xbd\x8e\xd5\xd9\x9f\xdb\x33\x86\x3d\x18\xbe\x00\xc6\x6c\x8c\xe6\x7a\xf3\xc4\xc8\x83\x08\x20\xab\xcf\xd9\x25\x9b\x5a\x16\x95\x2a\x8b\x05\x90\x7a\x8f\xdd\xc1\x10\x7d\xd2\x2c\xf9\x05\x00\x00\xff\xff\xb4\xed\x58\x92\xcf\x02\x00\x00")

func _1569155140_baselineDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1569155140_baselineDownSql,
		"1569155140_baseline.down.sql",
	)
}

func _1569155140_baselineDownSql() (*asset, error) {
	bytes, err := _1569155140_baselineDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1569155140_baseline.down.sql", size: 719, mode: os.FileMode(0664), modTime: time.Unix(1570640624, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x48, 0xb4, 0xfd, 0x4f, 0x37, 0x6, 0x30, 0x1, 0x9f, 0x78, 0x91, 0x2b, 0x15, 0x22, 0xf1, 0x17, 0xc8, 0x17, 0x31, 0xdf, 0x4a, 0x81, 0xf1, 0x32, 0xfa, 0x4c, 0x73, 0xbe, 0x1f, 0x3f, 0x71, 0xd1}}
	return a, nil
}

var __1569155140_baselineUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\xc1\x6e\x9c\x30\x10\xbd\xf3\x15\x56\x4e\x8b\x94\x4a\xa9\xd4\x9e\x72\x72\xc0\xab\xa0\x12\x93\x82\x69\x93\x5e\x2c\xb4\x38\x29\x0a\x05\x84\xa1\xe9\xe7\x57\xd8\x38\x31\x89\x97\x10\xd6\x1b\x8e\xe3\xf1\x9b\x19\xbf\x37\x63\xe3\xc5\x08\x12\x04\xd0\x0d\x41\x38\x09\x22\x0c\x82\x2d\xc0\x11\x01\xe8\x26\x48\x48\x02\x4e\xfa\xbe\xc8\x3f\xd5\x9c\x37\x27\xe7\x8e\x33\x3a\x13\x78\x11\xa2\x17\x8e\x3d\x67\x2d\x77\x36\x0e\x00\x00\x14\x39\x50\x5f\x9a\x06\x3e\xd8\xf3\xf9\x68\x0b\xd3\x90\x80\x21\x04\xbd\x67\x15\x6b\xb3\x8e\xd1\xbf\x5f\x36\xae\x00\xc6\x69\x18\x3a\xca\xd7\x8b\x70\x42\x62\x18\x60\x22\x23\xd1\xe6\x01\x5c\xc7\xc1\x15\x8c\x6f\xc1\x37\x74\x7b\x2a\x1c\xcb\xfa\xbe\xa8\xa4\xff\x0f\x18\x7b\x97\x30\xde\x7c\x3d\x73\xf7\x85\x37\x7d\x2a\xae\xc4\xbb\x2b\x5a\xde\xd1\x2a\xfb\xc3\xec\xe0\x95\x99\x82\xb3\x83\xd7\x64\x9c\x3f\xd6\x6d\xae\xd5\xfb\xf9\xec\x5d\x80\xcf\x78\x02\x30\x67\x25\xeb\x98\x60\xef\x22\x8a\x42\x04\xb1\x71\x97\x22\x6e\x0b\xc3\x04\xcd\x02\xee\x5a\x96\x75\x2c\xa7\x59\x07\x48\x70\x85\x12\x02\xaf\xae\xc1\xcf\x80\x5c\x46\x29\x11\x16\xf0\x2b\xc2\xe8\x09\xb0\xaa\x1f\x37\xee\x5c\xc5\x7d\x93\x5b\xc2\x73\xdc\x73\x25\xe7\x14\x07\xdf\x53\x04\x02\xec\xa3\x1b\x93\xaa\x69\x91\xd3\xbe\xa8\x72\xf6\x4f\xe4\x10\x61\x69\x06\x9b\x22\x7f\x07\x88\x10\xa7\x19\x47\x2c\xb9\xf3\xfd\xd5\xd6\x25\x33\xf4\xd7\x31\x1a\x4c\x84\x32\x36\x58\x57\x74\x25\x1b\x37\x58\x51\x70\xce\xf8\xae\x2d\x9a\xae\xa8\xab\xc3\x01\x05\x22\xef\x1b\xd6\x8e\x4b\xd6\x35\x6c\x5d\xc4\x1f\xa3\x62\x49\xe8\x2b\x15\x0b\xf3\x62\x15\x4b\x10\xa1\x00\x33\x8e\x58\x7a\x43\xc5\xd9\x6e\xa0\x5a\xe9\xd8\x24\xa7\xb7\x07\xd8\x9c\x74\x47\x7c\xa3\x78\x4d\x5a\x5b\x1e\xed\xf4\x00\x2d\xd8\xa0\xff\x7d\x8c\xab\x73\x30\xd1\x35\xae\x2d\x23\x4c\xce\x2e\x7d\xf8\x0c\x16\x2a\x27\xd0\xdc\xec\x59\xce\x99\x16\x82\x8e\xe0\xf4\xee\xc1\xd1\xb7\xc7\x68\x8b\x62\x84\x3d\xa4\xde\x19\xfa\x62\x84\x41\x7a\xed\x0f\x05\x78\x30\xf1\xa0\x8f\x06\x8b\x8f\x42\xf4\x6c\x91\x07\x3f\x84\x38\x5a\xe6\x23\xf8\x4c\xe6\xf2\x10\x57\x65\xbe\xea\x1a\x5d\x73\xd3\x4d\x59\x50\x35\x19\x2e\x2d\x3a\x36\xfd\xe8\x7a\xaa\x0e\x77\x59\xff\x4f\x04\x25\x6d\x52\xaa\x96\xc7\x80\x2c\x47\x0f\x30\xc3\x8f\x9a\x4d\x87\x6a\xcb\x96\xba\xa6\x45\x7c\x94\xbe\x8e\x3f\x93\x0c\xa4\x98\x85\x36\xd9\x00\x36\xfa\x8e\xa5\x7a\xe3\x8c\x73\xed\xc2\xd1\x1f\x4e\x6f\x53\xb5\xf2\xf9\xa4\x62\x1a\x2f\x21\x6d\x7a\x2e\xc8\x60\xa5\x72\x9e\x32\x38\xea\x38\xcd\x76\x3b\xc6\x39\xed\xea\x07\xf6\xe2\xa7\x6b\xf5\x4f\xc8\x4b\x58\xda\x75\x25\x00\x43\x49\x07\x9c\xd0\xd8\xa0\xec\xae\x65\xfc\xb7\x96\xae\x95\x6c\x27\xb0\x22\x5d\x2b\xd9\x4e\x1a\x52\x7c\xd6\x5f\x9e\x56\x60\x97\x75\xfe\x93\x1e\x5f\xb5\xb8\x5a\x59\xfc\x0a\x9d\xa8\x63\x1f\x96\xee\xb4\xec\x6d\x3b\x61\x71\x1f\xec\xc4\xcb\x3d\x77\xfe\x07\x00\x00\xff\xff\xe4\xd5\x8c\x92\x3b\x11\x00\x00")

func _1569155140_baselineUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1569155140_baselineUpSql,
		"1569155140_baseline.up.sql",
	)
}

func _1569155140_baselineUpSql() (*asset, error) {
	bytes, err := _1569155140_baselineUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1569155140_baseline.up.sql", size: 4411, mode: os.FileMode(0664), modTime: time.Unix(1571147946, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x19, 0x84, 0x19, 0x42, 0x25, 0xfb, 0x51, 0xce, 0xd8, 0xf, 0xeb, 0x84, 0xaa, 0x6b, 0xc0, 0xfd, 0x2e, 0x8c, 0x8, 0x91, 0xf6, 0x10, 0xe2, 0xc5, 0x8b, 0x2c, 0x2f, 0x39, 0x2b, 0xb9, 0xb8, 0xe6}}
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
	"1569155140_baseline.down.sql": _1569155140_baselineDownSql,

	"1569155140_baseline.up.sql": _1569155140_baselineUpSql,
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
	"1569155140_baseline.down.sql": &bintree{_1569155140_baselineDownSql, map[string]*bintree{}},
	"1569155140_baseline.up.sql":   &bintree{_1569155140_baselineUpSql, map[string]*bintree{}},
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
