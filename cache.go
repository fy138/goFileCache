package goFileCache

import (
	"errors"
	"time"

	//"path"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xinliangnote/go-util/md5"
)

type FileCache struct {
	BaseDir string
}

func NewFileCache(basedir string) *FileCache {
	return &FileCache{
		BaseDir: basedir,
	}
}
func (fc *FileCache) GetSavePath(key string) string {
	str := md5.MD5(key)
	if len(str) < 6 {
		return fmt.Sprintf("%s/%s.gob", fc.BaseDir, str)
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s.gob", fc.BaseDir, string(str[0]), string(str[1:3]), string(str[3:5]), str)
}
func (fc *FileCache) SetCache(key string, data interface{}) error {
	f := fc.GetSavePath(key)
	dir := filepath.Dir(f)
	if _, err := os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0755)
	}
	filewrite, err := os.Create(f)
	if err != nil {
		return err
	}
	defer filewrite.Close()
	enc := gob.NewEncoder(filewrite)
	return enc.Encode(data)
}
func (fc *FileCache) GetCache(key string, expire int, data interface{}) error {
	f := fc.GetSavePath(key)
	finfo, err := os.Stat(f)
	if err != nil {
		return err
	}
	if finfo.ModTime().Unix() < time.Now().Unix()-int64(expire) {
		return errors.New("expired")
	}

	fileread, err := os.Open(f)
	if err != nil {
		return err
	}

	defer fileread.Close()
	dec := gob.NewDecoder(fileread)
	return dec.Decode(data)
}
func (fc *FileCache) Delete(key string) error {
	return os.Remove(fc.GetSavePath(key))
}
