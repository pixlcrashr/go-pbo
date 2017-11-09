package pbo

import (
	"path/filepath"
	"testing"
)

const testPath = "test/"

func TestPBOWriteProduct(t *testing.T) {
	pbo := New()

	if err := pbo.WriteProduct(ProductEntry{
		EntryName:      "testPbo",
		ProductName:    "prefix",
		ProductVersion: "",
	}); err != nil {
		t.Fatal(err.Error())
	}
}

func TestPBOWriteHeader(t *testing.T) {
	pbo := New()

	if err := pbo.WriteHeader(HeaderEntry{
		FileName:      "config.cpp",
		PackingMethod: 0x0,
		OriginalSize:  24,
		Reserved:      0,
		TimeStamp:     0,
		DataSize:      24,
	}); err != nil {
		t.Fatal(err.Error())
	}
}

func TestPBOSave(t *testing.T) {
	pbo := New()
	path, _ := filepath.Abs(testPath)

	pbo.From = path + "/testPbo"
	pbo.To = path + "/test.pbo"
	pbo.Prefix = "testPbo"

	if err := pbo.Generate(); err != nil {
		t.Fatal(err.Error())
	}

	if err := pbo.Save(); err != nil {
		t.Fatal(err.Error())
	}
}

func TestPBOSaveTo(t *testing.T) {
	pbo := New()
	path, _ := filepath.Abs(testPath)

	pbo.From = path + "/testPbo"
	pbo.Prefix = "testPbo"

	if err := pbo.Generate(); err != nil {
		t.Fatal(err.Error())
	}

	if err := pbo.SaveTo(path + "/test.pbo"); err != nil {
		t.Fatal(err.Error())
	}
}

func TestPBOGenerate(t *testing.T) {
	pbo := New()
	path, _ := filepath.Abs(testPath)

	pbo.From = path + "/testPbo"
	pbo.Prefix = "testPbo"

	if err := pbo.Generate(); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetFiles(t *testing.T) {
	pbo := New()
	path, _ := filepath.Abs(testPath)

	pbo.From = path + "/testPbo"

	files := pbo.GetFiles()

	if len(files) <= 0 {
		t.Fatal("there have to be more than 0 files")
	}
}
