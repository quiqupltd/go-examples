# Atomic Wrappers

[Blog Article](https://medium.com/@deckarep/the-go-1-19-atomic-wrappers-and-why-to-use-them-ae14c1177ad8)

While coding in Go, any code that gets into multi-threaded or concurrent execution will undoubtedly run into synchronization problems around keeping your data free of the dreaded and evil data-races. Recall that a data-race may occur any time you introduce multiple goroutines accessing shared state where at least one goroutine writes/mutates said shared state along with one or more reads.

This problem is not exclusive or special to Go. In fact, nearly all languages that offer some degree of concurrent programming will suffer from such problems and therefore require the careful and proper synchronization of shared state. (Okay, maybe not the likes of Rust or other purely functional languages)

In Go, we can synchronize with several off-the-shelf options such as channels, mutexes or atomics. Most Go enthusiasts will chant the old mantra: “Do not communicate by sharing memory; instead, share memory by communicating” which largely implies the use of channels because channels coupled with goroutines are practically the true stars of the Go programming language.

Here’s the problem: Although the mantra is ideal in its wisdom, many, many projects do not adhere to it. In most reasonably complex projects you’ll find a myriad of goroutines, channels, atomics and mutexes sprinkled throughout the code. So I say let’s not be afraid to further understand the sync/atomic package and more importantly why the new wrapper API might be useful
