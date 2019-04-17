package hash

import (
	"testing"
)

func TestHashList(t *testing.T) {
	fakeTestList := map[string][]string{
		"list1": []string{},
		"list2": []string{"list-1", "list-2", "list-3"},
		"list3": []string{"list-4", "list-5", "list-6"},
	}

	for _, test := range fakeTestList {
		fakeHash, err := Hash(test)
		if err != nil {
			t.Errorf("Failed to calculate the hash expected string but got: '%s' Error: '%v'", fakeHash, err)
		}
	}
	fakeHash, err := Hash(fakeTestList)
	if err != nil {
		t.Errorf("Failed to calculate the hash expected string but got: '%s' Error: '%v'", fakeHash, err)
	}
}

func TestHashStruct(t *testing.T) {
	fakeStruct := map[string]struct {
		name string
		list []string
	}{
		"struct1": {
			name: "",
			list: []string{},
		},
		"struct2": {
			name: "abcdefgh",
			list: []string{"abc-1", "abc-2", "abc-3"},
		},
		"struct3": {
			name: "jklmnop",
			list: []string{"abcde", "abcdf", "hash"},
		},
	}
	for _, test := range fakeStruct {
		fakeHash, err := Hash(test)
		if err != nil {
			t.Errorf("Failed to calculate the hash expected string but got: '%s' Error: '%v'", fakeHash, err)
		}
	}
	fakeHash, err := Hash(fakeStruct)
	if err != nil {
		t.Errorf("Failed to calculate the hash expected string but got: '%s' Error: '%v'", fakeHash, err)
	}
}

func TestHashInnerStruct(t *testing.T) {
	type fakeStruct struct {
		name string
		list []string
	}
	fakeComplexStruct := map[int]struct {
		innerStructNum int
		innerStruct    struct {
			name string
			list []string
		}
	}{
		1: {
			innerStructNum: 1,
			innerStruct: fakeStruct{
				name: "",
				list: []string{},
			},
		},
		2: {
			innerStructNum: 2,
			innerStruct: fakeStruct{
				name: "hashInnerStruct",
				list: []string{"hash1", "hash2", "hash3", "hash4"},
			},
		},
	}
	for _, test := range fakeComplexStruct {
		fakeHash, err := Hash(test)
		if err != nil {
			t.Errorf("Failed to calculate the hash expected string but got: '%s' Error: '%v'", fakeHash, err)
		}
	}
	fakeHash, err := Hash(fakeComplexStruct)
	if err != nil {
		t.Errorf("Failed to calculate the hash expected string but got: '%s' Error: '%v'", fakeHash, err)
	}
}