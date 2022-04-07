package utils

import (
	"crypto/md5"
	"hash"
	"os"

	"golang.org/x/exp/constraints"
)

func GetFileHash[T constraints.Integer](path string, bufSize T, hash hash.Hash) ([]byte, error) {
	bufSizeInt := int(bufSize)

	f, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	fStat, err := f.Stat()
	if err != nil {
		return []byte{}, err
	}
	fSize := int(fStat.Size())

	i := 0
	for {
		buf := make([]byte, Min(bufSizeInt, fSize-(i*bufSizeInt)))
		n, err := f.ReadAt(buf, int64(i*bufSizeInt))
		hash.Write(buf)

		if n != bufSizeInt {
			return hash.Sum(nil), err
		}

		i++
	}
}

func GetFileMD5[T constraints.Integer](path string, bufSize T) ([]byte, error) {
	return GetFileHash(path, bufSize, md5.New())
}
