package main

import (
	"reflect"
	"testing"
)

var mockDB = map[string]Man{
	"John Doe": {
		Name:     "John",
		LastName: "Doe",
		Age:      30,
		Gender:   "M",
		Crimes:   1,
	}, "Jane Smith": {
		Name:     "Jane",
		LastName: "Smith",
		Age:      25,
		Gender:   "M",
		Crimes:   0,
	}, "Michael Johnson": {
		Name:     "Michael",
		LastName: "Johnson",
		Age:      45,
		Gender:   "M",
		Crimes:   5,
	}, "Emily Williams": {
		Name:     "Emily",
		LastName: "Williams",
		Age:      32,
		Gender:   "F",
		Crimes:   5,
	}, "Daniel Garcia": {
		Name:     "Daniel",
		LastName: "Garcia",
		Age:      35,
		Gender:   "M",
		Crimes:   3,
	},
}

func TestMostSuspiciousEmptySusp(t *testing.T) {
	var want []string
	susp := make([]string, 0)
	if ans := mostSuspicious(susp, mockDB); !reflect.DeepEqual(ans, want) {
		t.Errorf("want: %v, got %v", want, ans)
	}
}

func TestMostSuspiciousTwoEqualCrimes(t *testing.T) {
	var want = []string{"Michael Johnson", "Emily Williams"}
	var susp = []string{"Michael Johnson", "Emily Williams"}
	if ans := mostSuspicious(susp, mockDB); !reflect.DeepEqual(ans, want) {
		t.Errorf("want: %v, got %v", want, ans)
	}
}

func TestMostSuspiciousNoCrimes(t *testing.T) {
	// var want []string
	var want = []string{"Jane Smith"}
	var susp = []string{"Jane Smith"}
	if ans := mostSuspicious(susp, mockDB); !reflect.DeepEqual(ans, want) {
		t.Errorf("want: %v, got %v", want, ans)
	}
}

func TestMostSuspiciousMultSuspDiffCrimes(t *testing.T) {
	var want = []string{"Michael Johnson", "Emily Williams"}
	var susp = []string{"Michael Johnson", "Emily Williams", "Daniel Garcia"}
	if ans := mostSuspicious(susp, mockDB); !reflect.DeepEqual(ans, want) {
		t.Errorf("want: %v, got %v", want, ans)
	}
}

func TestMostSuspiciousNoSuspectsInDB(t *testing.T) {
	var want []string
	var susp = []string{"Alex Rodriguez", "Sarah Connor"}
	if ans := mostSuspicious(susp, mockDB); !reflect.DeepEqual(ans, want) {
		t.Errorf("want: %v, got %v", want, ans)
	}
}

func TestMostSuspiciousMixedSuspects(t *testing.T) {
	var want = []string{"Michael Johnson", "Emily Williams"}
	var susp = []string{"Michael Johnson", "Emily Williams", "Alex Rodriguez", "Sarah Connor"}
	if ans := mostSuspicious(susp, mockDB); !reflect.DeepEqual(ans, want) {
		t.Errorf("want: %v, got %v", want, ans)
	}
}
