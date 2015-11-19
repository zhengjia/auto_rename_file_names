package auto_rename_file_names

import (
	"strings"
	"testing"
)

func TestNumericSuffix(t *testing.T) {
	renamer := New()
	path1 := "test1/1.txt"
	path2 := "test1/2.txt"
	path1_0 := renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
	if len(renamer.elems) != 1 {
		t.Errorf("elems length is %d", len(renamer.elems))
	}
	path2_0 := renamer.Get(path2)
	if strings.Compare(path2, path2_0) != 0 {
		t.Errorf("Get returns: %s", path2_0)
	}
	if len(renamer.elems) != 2 {
		t.Errorf("elems length is %d", len(renamer.elems))
	}
	path1_1 := renamer.Get(path1)
	if path1_1 != "test1/1(1).txt" {
		t.Errorf("Get returns: %s", path1_1)
	}
	path1_2 := renamer.Get(path1)
	if path1_2 != "test1/1(2).txt" {
		t.Errorf("Get returns: %s", path1_2)
	}
	path2_1 := renamer.Get(path2)
	if path2_1 != "test1/2(1).txt" {
		t.Errorf("Get returns: %s", path2_1)
	}
	renamer.Reset()
	path1_0 = renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
}

func TestNumericPrefix(t *testing.T) {
	renamer := New()
	renamer.Type = NumericPrefix
	path1 := "test1/1.txt"
	path1_0 := renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
	path1_1 := renamer.Get(path1)
	if strings.Compare("test1/(1)1.txt", path1_1) != 0 {
		t.Errorf("Get returns: %s", path1_1)
	}
	path1_2 := renamer.Get(path1)
	if strings.Compare("test1/(2)1.txt", path1_2) != 0 {
		t.Errorf("Get returns: %s", path1_2)
	}
	renamer.Reset()
	path1_0 = renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
}

func TestNumericConnector(t *testing.T) {
	renamer := New()
	renamer.Connector = "_"
	path1 := "test1/1.txt"
	path1_0 := renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
	path1_1 := renamer.Get(path1)
	if strings.Compare("test1/1_(1).txt", path1_1) != 0 {
		t.Errorf("Get returns: %s", path1_1)
	}
}

func TestStringSuffix(t *testing.T) {
	renamer := New()
	renamer.Type = StringSuffix
	renamer.Connector = "_"
	renamer.StringAffix = "Copy"
	renamer.Seperator = "-"
	path1 := "test1/1.txt"
	path1_0 := renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
	path1_1 := renamer.Get(path1)
	if strings.Compare("test1/1_Copy.txt", path1_1) != 0 {
		t.Errorf("Get returns: %s", path1_1)
	}
	path1_2 := renamer.Get(path1)
	if strings.Compare("test1/1_Copy-Copy.txt", path1_2) != 0 {
		t.Errorf("Get returns: %s", path1_2)
	}
	renamer.Reset()
}

func TestLeadingSlash(t *testing.T) {
	renamer := New()
	path1 := "/test1/1.txt"
	path1_0 := renamer.Get(path1)
	if strings.Compare(path1, path1_0) != 0 {
		t.Errorf("Get returns: %s", path1_0)
	}
}
