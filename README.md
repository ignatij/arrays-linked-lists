# Arrays vs Linked Lists in Go

This repository contains benchmarks and examples comparing **arrays (slices)** and **linked lists** in Go.  
The focus is on understanding performance trade-offs between the two data structures, especially in terms of **traversal, insertions, and memory usage**.

---

## 📊 Benchmarks

All benchmarks are implemented using Go’s [`testing`](https://pkg.go.dev/testing) package.  
To run them:

```bash
make test
```

⚖️ Takeaways

- Traversal: Arrays are significantly faster due to cache efficiency.

- Insert at head: Linked lists dominate with O(1) performance.

- Insert at tail: Arrays usually win unless you maintain a tail pointer in your list.

- Memory: Arrays are compact, while linked lists require extra space for pointers.
