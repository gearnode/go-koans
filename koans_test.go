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

