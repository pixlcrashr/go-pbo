package pbo

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unsafe"
)

type HeaderEntry struct {
	FileName                                                   string
	PackingMethod, OriginalSize, Reserved, TimeStamp, DataSize uint32
}

type ProductEntry struct {
	EntryName, ProductName, ProductVersion string
}

type PBO struct {
	Buffer   *bytes.Buffer
	From, To string
	Files    []string
	Prefix   string
	Version  string
}

func (pbo *PBO) WriteProduct(product ProductEntry) error {
	tmpV := make([]byte, len(product.EntryName))
	copy(tmpV[:], product.EntryName)
	err := binary.Write(pbo.Buffer, binary.BigEndian, tmpV)

	if err != nil {
		return err
	}

	tmpV = make([]byte, len(product.ProductName))
	copy(tmpV[:], product.ProductName)
	err = binary.Write(pbo.Buffer, binary.BigEndian, tmpV)

	if err != nil {
		return err
	}

	tmpV = make([]byte, len(product.ProductVersion))
	copy(tmpV[:], product.ProductVersion)
	err = binary.Write(pbo.Buffer, binary.BigEndian, tmpV)

	return err
}

func (pbo *PBO) WriteHeader(header HeaderEntry) error {
	tmpV := make([]byte, len(header.FileName))
	copy(tmpV[:], header.FileName)
	err := binary.Write(pbo.Buffer, binary.BigEndian, tmpV)
	pbo.Buffer.WriteByte(byte('\x00'))

	if err != nil {
		return err
	}

	err = binary.Write(pbo.Buffer, binary.LittleEndian, (*[4]byte)(unsafe.Pointer(&header.PackingMethod)))

	if err != nil {
		return err
	}

	err = binary.Write(pbo.Buffer, binary.LittleEndian, (*[4]byte)(unsafe.Pointer(&header.OriginalSize)))

	if err != nil {
		return err
	}

	err = binary.Write(pbo.Buffer, binary.LittleEndian, (*[4]byte)(unsafe.Pointer(&header.Reserved)))

	if err != nil {
		return err
	}

	err = binary.Write(pbo.Buffer, binary.LittleEndian, (*[4]byte)(unsafe.Pointer(&header.TimeStamp)))

	if err != nil {
		return err
	}

	err = binary.Write(pbo.Buffer, binary.LittleEndian, (*[4]byte)(unsafe.Pointer(&header.DataSize)))

	return err
}

func (pbo *PBO) Generate() error {
	pbo.Buffer.Reset()
	pbo.WriteHeader(HeaderEntry{
		FileName:      "",
		PackingMethod: 0x56657273,
		OriginalSize:  0,
		Reserved:      0,
		TimeStamp:     0,
		DataSize:      0,
	})

	pbo.WriteProduct(ProductEntry{
		ProductName:    pbo.Prefix,
		ProductVersion: pbo.Version,
		EntryName:      "prefix",
	})

	files := pbo.GetFiles()

	for _, f := range files {
		fHandle, err := os.Open(f)

		if err != nil {
			return err
		}

		stat, err := fHandle.Stat()

		if err != nil {
			return err
		}

		size32, err := strconv.ParseUint(strconv.FormatInt(stat.Size(), 10), 10, 32)

		if err != nil {
			return err
		}

		time32, err := strconv.ParseUint(strconv.FormatInt(time.Now().Unix(), 10), 10, 32)

		if err != nil {
			return err
		}

		rP, err := filepath.Rel(pbo.From, f)

		pbo.WriteHeader(HeaderEntry{
			FileName:      rP,
			PackingMethod: 0x0,
			OriginalSize:  uint32(size32),
			Reserved:      0,
			TimeStamp:     uint32(time32),
			DataSize:      uint32(size32),
		})
	}

	pbo.WriteHeader(HeaderEntry{})

	for _, f := range files {
		f, err := os.Open(f)
		defer f.Close()

		if err != nil {
			return err
		}

		_, err = io.Copy(pbo.Buffer, f)

		if err != nil {
			return err
		}
	}

	return nil
}

func (pbo *PBO) Save() error {
	return ioutil.WriteFile(pbo.To, pbo.Buffer.Bytes(), 0644)
}

func (pbo *PBO) SaveTo(to string) error {
	return ioutil.WriteFile(to, pbo.Buffer.Bytes(), 0644)
}

func (pbo *PBO) GetFiles() []string {
	var files []string
	filepath.Walk(
		pbo.From,
		func(file string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				files = append(files, file)
			}

			return nil
		},
	)

	return files
}

func New() *PBO {
	return &PBO{
		Buffer: &bytes.Buffer{},
	}
}