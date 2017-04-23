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
