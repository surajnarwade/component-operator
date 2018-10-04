// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package template

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 8, 16, 8, 21, 38, 253832101, time.UTC),
		},
		"/java": &vfsgen۰DirInfo{
			name:    "java",
			modTime: time.Date(2018, 8, 14, 17, 14, 16, 827264523, time.UTC),
		},
		"/java/imagestream": &vfsgen۰CompressedFileInfo{
			name:             "imagestream",
			modTime:          time.Date(2018, 8, 14, 17, 14, 16, 826950269, time.UTC),
			uncompressedSize: 508,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\xc1\x6a\x33\x31\x0c\x84\xef\xfb\x14\x22\xf7\xc4\xfc\xff\xd1\x39\x95\xf6\xd2\x4b\x29\x0d\xf4\xae\xd8\xda\xd4\x8d\x6d\x09\xcb\x1b\x08\x26\xef\x5e\xe2\xdd\x94\xb4\xf4\x66\x3e\x69\x34\xe3\x41\x09\xef\x54\x34\x70\xb6\x10\x12\x1e\x68\xc3\x42\x59\x3f\xc2\x58\x37\x81\xcd\xe9\xdf\x70\x0c\xd9\x5b\x78\xbe\xce\x76\xb5\x10\xa6\x21\x51\x45\x8f\x15\xed\x00\x10\x71\x4f\x51\xaf\x2f\x00\x14\xb1\xd0\xda\xe6\x05\x13\x5d\x2e\x03\x40\xc6\x44\x9d\x74\xf5\x8d\xab\x90\xeb\x52\xe6\xe3\x24\xaf\x1c\x83\x3b\xcf\x07\x22\x3b\x8c\x16\x46\x8c\x4a\x03\x40\x6b\x61\x84\x45\xfb\xc4\xee\x48\xa5\xbf\xfb\xe9\x8a\x87\xee\xba\xfe\xb1\xf5\x90\x33\x57\xac\x81\xf3\x63\xf2\xda\x17\x01\xf0\x1b\x2e\x39\x01\x5c\xf2\x6a\x61\x55\xa6\xbc\xfe\xc4\x13\x5a\x33\x69\x31\xdd\xdd\xe8\xff\x60\xca\x94\xb7\x8e\x93\x84\x48\x7f\xce\x51\x95\xd2\x3e\xd2\x76\x3f\x85\xe8\xad\xf1\x24\x91\xcf\x89\x72\x55\xd3\x11\x8a\xac\xba\x55\x6b\x94\xfd\x92\x63\x2c\x9c\x6e\x01\xe6\x52\xef\x3e\xb5\xf0\x5f\x8d\xbd\x91\xf0\xa2\x0e\x49\xb8\xd4\xa5\x2d\x68\x33\x9c\xd7\x23\x56\xd2\xda\x41\xa1\x91\x0a\x65\x47\xf7\xb5\x02\xd4\xb3\x90\x85\x1d\x4f\xc5\xcd\xc5\xf6\x54\x5f\x01\x00\x00\xff\xff\xf6\xbc\xc2\x10\xfc\x01\x00\x00"),
		},
		"/java/route": &vfsgen۰CompressedFileInfo{
			name:             "route",
			modTime:          time.Date(2018, 8, 14, 17, 14, 16, 827158760, time.UTC),
			uncompressedSize: 172,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcd\x31\x0e\xc2\x30\x0c\x85\xe1\x3d\xa7\xf0\x09\x82\x58\x73\x08\x06\x90\xd8\x4d\xfb\x10\x16\x4d\x6c\x25\xa6\x4b\xd5\xbb\xa3\x28\x13\xa2\xf3\xff\x3e\x3d\x36\xb9\xa3\x36\xd1\x92\xa8\xea\xc7\x11\xd5\x50\xda\x4b\x9e\x1e\x45\x4f\xeb\x39\xbc\xa5\xcc\x89\xae\xbd\x85\x0c\xe7\x99\x9d\x53\x20\x2a\x9c\x91\x68\xdb\xe2\x85\x33\xf6\x3d\x10\x2d\xfc\xc0\xd2\x7a\x23\x62\xb3\xdf\xf8\x0f\x9a\x61\xea\x63\xd7\x41\xc6\xd1\x0d\x75\x95\x09\x47\xe2\x1b\x00\x00\xff\xff\xdd\xc7\x7e\xf9\xac\x00\x00\x00"),
		},
		"/java/service": &vfsgen۰CompressedFileInfo{
			name:             "service",
			modTime:          time.Date(2018, 8, 14, 17, 14, 16, 827320403, time.UTC),
			uncompressedSize: 247,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xc1\xaa\xc2\x30\x10\x45\xf7\xf9\x8a\xf9\x81\xf7\xc0\x6d\xb6\xee\xa5\xa0\xb8\x1f\xd3\x6b\x09\x26\x99\x30\x19\x0a\x52\xfa\xef\xd2\xe0\x42\x29\xee\x86\x39\xe7\xc0\xe5\x1a\xaf\xd0\x16\xa5\x78\x9a\x0f\xee\x11\xcb\xe8\xe9\x0c\x9d\x63\x80\xcb\x30\x1e\xd9\xd8\x3b\xa2\xc2\x19\x9e\x96\xe5\xff\xc4\x19\xeb\xea\x88\x12\xdf\x90\xda\xc6\x88\xb8\xd6\x6f\xb8\x0f\x5a\x45\xd8\xe4\x2a\x6a\xbd\xfa\xeb\x67\x57\x06\x51\x7b\x67\x55\xc5\x24\x48\xf2\x74\x39\x0e\xfd\x63\xac\x13\x6c\xd8\xb9\x0d\x09\xc1\x44\x7f\x2e\x18\x51\x93\x3c\x33\x8a\x05\x29\xf7\x38\x7d\xf0\x57\x00\x00\x00\xff\xff\x73\x21\x02\x3d\xf7\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/java"].(os.FileInfo),
	}
	fs["/java"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/java/imagestream"].(os.FileInfo),
		fs["/java/route"].(os.FileInfo),
		fs["/java/service"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}