// file test must be end by _test example "koans_test.go"
// To run test suite execute "go test" command

package koans_test

import(
  "fmt" //
  "testing" // testing package : https://golang.org/pkg/testing/
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

    if x == y {
      t.FailNow()
    }
  }
  {
    x := 3
    y := &x // use & to access to adress of a pointer

    if *y != 3 {
      t.FailNow()
    }

    // y is a pointer of x
    // Use * symbol to de-referencing
    *y = *y + 2

    if x != 5 {
      t.FailNow()
    }
  }
  {
    incr := func(i int) int {
      i = i + 1
      return i
    }

    x := 1
    incr(x)

    if x != 1 {
      t.FailNow()
    }
  }
  {
    incr := func(i *int) *int {
      *i++
      return i
    }

    x := 1
    incr(&x)
    if x != 2 {
      t.FailNow()
    }
  }
  {
    x := 1
    p := &x
    pp := &p
    cp := **pp

    if **pp != x {
      t.FailNow()
    }

    **pp = 4
    if x != 4 {
      t.FailNow()
    }

    if cp != 1 {
      t.FailNow()
    }
  }

}
