# GO LANG CHEAT SHEET

## Variables, Arrays, Slices

```go
var ecampleInt int = 7
var exampleFloat float = 3.1415
var exampleString string = "example"
var exampleArray [5]int = [5]int{1,2,3,4,5}
var exampleSlice []int = []int{1,2,3,4,5}
new_slice := make([]int, length)
slice = append(slice, element)
```

## Struct

```go
type MyStruct{
    field1 string
    field2 float64
}

func (ms MyStruct) MyMethod() string {
    return ms.someField
}
func (ms *MyStruct) MyMethod() {
    ms.someField = "somevalue"
}

mystruct := MyStruct{"x", 1.0}
mystruct := MyStruct{field1:"x", field2:1.0}

```

## Interfaces

```go
type MyInterface interface {
 AMethod(param int, anotherparam string) float64
}
```

## Loop

```go
for i := 0; i < 10; i++ {
    // ...
}
for index, item := range items {
    // ...
}
```

## Errors

```go
func DoWithErrors(...) error {
    errorHappened :=True    
    if errorHappened {
            return errors.New("message") // Error
            return fmt.Errorf("message %s", someData) // Formatted error
    }
    return nil // No error
}
```

## Test, Example, Benchmark

```go
func TestFunction(t *testing.T) {
 want := 1
 got := 2
 if got != want {
  t.Errorf("Failed; want: %d, got; %d", want, got)
 }
}

func ExampleFunction() {
 fmt.Printf("Example Output") // Result of any calculation
 // Output: Example Output
}


func BenchmarkFunction(b *testing.B) {
 for i := 0; i < b.N; i++ {
  // ... (Run the function to benchmark)
 }
}
```

## Useful operators

Test for slice equality

```go
!reflect.DeepEqual(got, want)
```

## Linters and tools

Errcheck

```go
go install github.com/kisielk/errcheck@latest
```

```shell
# Default $GOPATH=~/go
export PATH=$PATH:~/go/bin # export PATH=$PATH:$GOPATH
```
