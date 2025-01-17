package clockpro

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestCache(t *testing.T) {

	// Test data was generated from the python code
	f, err := os.Open("testdata/domains.txt")

	if err != nil {
		t.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	cache := New[string, []byte](200)

	for scanner.Scan() {
		fields := bytes.Fields(scanner.Bytes())

		key := string(fields[0])
		wantHit := fields[1][0] == 'h'

		var hit bool
		v := cache.Get(key)
		if len(v) == 0 {
			cache.Set(key, []byte(key))
		} else {
			hit = true
			if !bytes.Equal(v, []byte(key)) {
				t.Errorf("cache returned bad data: got %+v , want %+v\n", v, key)
			}
		}
		if hit != wantHit {
			t.Errorf("cache hit mismatch on %s: got %v, want %v\n", key, hit, wantHit)
		}
	}
}
