// file test must be end by _test example "koans_test.go"
// To run test suite execute "go test" command

package koans_test

import(
  "fmt" //
  "testing" // testing package : https://golang.org/pkg/testing/

  "strings" // simple functions to manipulate UTF-8 encoded strings.
  "os" // platform-independent interface to operating system functionality
  "strconv" // implements conversions to and from string representations of basic data types
)

// All test must be start with TestXxx(*testing.T)
func TestSample(t *testing.T) {
  // To signal failure use FailNow, Error, Fail, etc. func

  t.Skip("skipping test message") // To skip a test
}
// Benchmarks
// The benchmarks func must use this pattern
// func BenchmarkXxx(*testing.B)
func BenchmarkHello(b *testing.B) {
  for i := 0; i < b.N; i++ {
    fmt.Println("hello world")
  }
}

// Little helper to write test
func Assert(t *testing.T, assertation bool) {
  if assertation {
    fmt.Println("*")
  } else {
    t.FailNow()
  }
}

// Pointer
// Notes
// Go have pointer but go has no pointer arithmetic :(

func TestPointer(t *testing.T) {
  // Use {} to isolate context into go func
  // Relly usefull to have many idependant
  // assertation into func Test
  {
    x := 3
    y := x // b is a copy of a
    // All assignments in golang are copy operation

    y++

    Assert(t, x != y)
  }
  {
    x := 3
    y := &x // use & to access to adress of a pointer

    Assert(t, *y == 3)

    // y is a pointer of x
    // Use * symbol to de-referencing
    *y = *y + 2

    Assert(t, x == 5)
  }
  {
    incr := func(i int) int {
      i = i + 1
      return i
    }

    x := 1
    incr(x)

    Assert(t, x == 1)
  }
  {
    incr := func(i *int) *int {
      *i++
      return i
    }

    x := 1
    incr(&x)

    Assert(t, x == 2)
  }
  {
    x := 1
    p := &x
    pp := &p
    cp := **pp

    Assert(t, **pp == x)

    // Modify value of x
    **pp = 4
    Assert(t, x == 4)
    Assert(t, cp == 1)
  }
  {
    add := func(x,y int) int {
      return x + y
    }

    Assert(t, add(1,1) == 2)

    padd := &add
    fn := *padd
    Assert(t, fn(1,1) == 2)

    *padd = func(x,y int) int {
      return x - y
    }

    Assert(t, add(1, 1) == 0)
  }
}

// Basic
func TestBasic(t *testing.T) {
  Assert(t, true == true)
  Assert(t, true != false)

  var i int = 1
  Assert(t, i == 1.0000000000000000000000000000000000000)

  Assert(t, 5%2 == 1)
  Assert(t, 7*2 == 14)

  var x int
  Assert(t, x == 0)

  var f float32
  Assert(t, f == 0.0)

  var s string
  Assert(t, s == "")

  var c struct { x int }
  Assert(t, c.x == 0)
}

// String
func TestString(t *testing.T) {
  Assert(t, "a" + "bc" == "abc")
  Assert(t, len("abc") == 3)

  Assert(t, "abc"[0] == 'a') // like C (char != char*)

  // Operation
  Assert(t, "gearnode"[:2] == "ge")
  Assert(t, "gearnode"[2:] == "arnode")
  Assert(t, "gearnode"[2:4] == "ar")
  Assert(t, "gearnode"[:] == "gearnode")

  Assert(t, "gearnode" == "gearnode")
  Assert(t, "a" < "b")
  Assert(t, "1" < "2")

  bytes := []byte{'a', 'b', 'c'}
  Assert(t, string(bytes) == "abc")

  bytes[0] = 'q'
  Assert(t, bytes[0] == 'q')
  Assert(t, string(bytes) == "qbc")

  // Compare
  Assert(t, strings.Compare("a", "a") == 0)
  Assert(t, strings.Compare("a", "b") == -1)
  Assert(t, strings.Compare("b", "a") == 1)

  Compare := func(s1,s2 string) int {
    if s1 == s2 { return 0 }
    if s1 < s2 { return -1 }
    return 1
  }

  Assert(t, strings.Compare("a", "a") == Compare("a", "a"))
  Assert(t, strings.Compare("a", "b") == Compare("a", "b"))
  Assert(t, strings.Compare("b", "a") == Compare("b", "a"))

  // Contains
  Assert(t, strings.Contains("gearnode", "node") == true)
  Assert(t, strings.Contains("gearnode", "") == true)
  Assert(t, strings.Contains("gearnode", "b") == false)
  Assert(t, strings.Contains("", "") == true)

  // ContainsAny
  Assert(t, strings.ContainsAny("", "") == false)
  Assert(t, strings.ContainsAny("gearnode", "") == false)
  Assert(t, strings.ContainsAny("gearnode", "g 89 y") == true)
  Assert(t, strings.ContainsAny("gearnode", "z") == false)

  // ContainsRune
  Assert(t, strings.ContainsRune("gearnode", 'g') == true)
  Assert(t, strings.ContainsRune("hel⌘", '⌘') == true)

  // Count
  Assert(t, strings.Count("gearnode", "") == 9)
  Assert(t, strings.Count("gearnode", "z") == 0)
  Assert(t, strings.Count("gearnode", "e") == 2)

  // EqualFold
  Assert(t, strings.EqualFold("Go", "go") == true)

  // HasPrefix
  Assert(t, strings.HasPrefix("gearnode", "ge") == true)
  Assert(t, strings.HasPrefix("gearnode", "node") == false)
  Assert(t, strings.HasPrefix("gearnode", "") == true)

  // HasSuffix
  Assert(t, strings.HasSuffix("gearnode", "node") == true)
  Assert(t, strings.HasSuffix("gearnode", "gear") == false)
  Assert(t, strings.HasSuffix("gearnode", "") == true)
}

func TestEvironmentVariables(t *testing.T) {

  os.Setenv("FOO", "BAR")
  Assert(t, os.Getenv("FOO") == "BAR")
}

func TestFizzBuzz(t *testing.T) {
  FizzBuzz := func(i int) string {
    if i % 3 == 0 && i % 5 == 0 { return "fizzbuz" }
    if i % 3 == 0 { return "fizz" }
    if i % 5 == 0 { return "buzz" }
    return strconv.Itoa(i)
  }

  Assert(t, FizzBuzz(1) == "1")
  Assert(t, FizzBuzz(2) == "2")
  Assert(t, FizzBuzz(3) == "fizz")
  Assert(t, FizzBuzz(4) == "4")
  Assert(t, FizzBuzz(5) == "buzz")
}
