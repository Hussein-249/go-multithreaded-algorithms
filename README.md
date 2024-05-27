# Multithreaded Algorithms (Go)
![](https://img.shields.io/badge/Tests-Passing-green)

## Purpose
This module is not meant to be a production-grade or professional solution. It is meant to be a proof-of-concept repo, and to demonstrate my skills and understanding of the Go programming language.

## Implemented Structures and Algorithms

Some of these data structures and algorithms may not inherently benefit from multithreading, such as the RootishArrayStack, and thus have not been implemented with goroutines, but 

### Algorithms

- [x] Merge Sort, Counting Sort

### Data Structures

- [ ] None

## Structures and Algorithms Being Developed

### Algorithms

- [ ] Dijkstra's Algorithm

### Data Structures

- [ ] Trie
- [ ] RootishArrayStack

## Tests
 I have currently implemented unit tests for completed algorithms. The unit tests are passing, and you can run them all by running 
 ``` bash
 go test -count=1 ./...
 ``` 
 This can be executed at the module level (targets all tests) or at the package level (targets all tests in the package). When executing multiple iterations of tests where the actual test content may vary, be sure to include -count=1 as this disables test caching.
 <br>

 To test an individual file, run  
 ``` bash
 go test -run <filename>
 ```

 ## Documentation

 Work in progress.
 