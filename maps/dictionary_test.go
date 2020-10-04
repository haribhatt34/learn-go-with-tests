package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	assertStrings := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q given %q", got, want, "test")
		}
	}

	assertError := func(t *testing.T, got error, want error, given string) {
		t.Helper()
		if got == nil {
			t.Fatal("Should return an error")
		}
		if got != want {
			t.Errorf("got %q want %q given %q", got, want, given)
		}
	}

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		dictionary := Dictionary{"test": "this is just a test"}
		_, got := dictionary.Search("test1")

		assertError(t, got, ErrNotFound, "test1")
	})
}

func TestAdd(t *testing.T) {

	assertDefinition := func(t *testing.T, word, definition string, dictionary Dictionary) {
		t.Helper()
		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word: ", err)
		}
		if got != definition {
			t.Errorf("got %v want %v", got, definition)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if got == nil {
			if want == nil {
				// new word case
				return
			}
			// existing word case
			t.Errorf("Expected to get an error")
		}
	}

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, word, definition, dictionary)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "this is a new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, word, definition, dictionary)
	})
}

// Notes :
/*
 *  1)
 * The key type in map should always be a comparable type, because keys
 * are compared to fetch value.
 *
 * 2)
 * As we can see we can modify map without passing them as an address.
 * Reason - A map value is a pointer to a runtime.hmap structure.
 * So when we pass a map to a function/method, we are actually copying
 * a pointer instead of underlying data.
 *
 * 3) gotchas -
 * Maps can be nil, a nil behave same as an empty map when reading,
 * but writing to a nil map causes runtime panic.
 * Therefore, one should never initialize an empty map
 *
 * Important Read -
 *
 * => golang variable declaration -
 * i) var string name
 * X ii) var string name = "hari"  X
 * ii) var name = "hari"
 * iii) name := "hari"
 *
 * ii cannot be used
 * iii is short notation for ii
 * Short notation cannot be declared outside a func body.
 * we can re-declare shorthand notation, within same block
 *
 *
 * ==> About maps
 *
 * source - https://medium.com/@KeithAlpichi/go-gotcha-nil-maps-66b851c50475
 * Below is a nil map.
 * 'var mp map[string] int'
 * The above syntax will declare a map variable mp.
 *
 * Now if we try below, we will get a runtime panic.
 * mp["hari"] = 34
 *
 * as we know when we create a variable in go it is zero valued.
 * and zero value for map is nil, because it is a pointer.
 * And we are trying to accessing a unallocated address.
 *
 * So we need to initialize a map - 2 methods
 * i) using make
 * 	var mp = make(map[string] int, 1)
 *  mp initialized with size 1, it can grow later.
 *  var mp = make(map[string] int)
 *  mp initialized with size 0.
 *
 * ii) using map literal
 *  var mp := map[string] int {} // an empty map
 *  The above declaration is same as make(mp[string] int)
 *  var mp := map[string] int {"hari", 1} // map initialized with value
 *
 */
