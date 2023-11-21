golang 
static typed 
compiled
powerful concurrency using Goroutines
has a garbage collector

package:
    collection of go files
module:
    collection of package

go mod init


go is not oop lang
instead of creating classes with methods and vars
we create structs with vars and attach functions which receive that struct

package types:
- executable (main)
- reusable

value types:
int, float, string, bool, struct
passed by value to funcs, so use pointer to change underlying value

reference types:
slice, maps, channels, pointers, functions
passed by reference by default

## Learn about these
goroutines
channels: like queue to pass data btw goroutines
waitgroup
generics
interfaces
concurrency patterns like pipeline, fan in, fan out, semaphores, workerpool

## Resources
- https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit
- https://go.dev/learn/
- https://exercism.org/tracks/go/concepts
- https://gobyexample.com/
- https://www.karanpratapsingh.com/blog/learn-go-the-complete-course
