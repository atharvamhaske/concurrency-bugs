# My implemetation notes of **"Understanding Real-World Concurrency Bugs in Go"**

This repository contains my hands-on learnings, notes, and runnable code while studying and implementing concepts from the research paper **"Understanding Real-World Concurrency Bugs in Go"**. The goal of this repo is not theory repetition, but practical exploration: reproducing real concurrency bugs described in the paper, understanding why they happen, and verifying correct fixes by running the code.

The code here focuses heavily on Go concurrency primitives such as goroutines, channels, select statements, timeouts, and cancellation. Each bug pattern is implemented in isolation so it can be run, observed, and reasoned about without abstraction or frameworks. Wherever possible, both the buggy version and the corrected version are included to clearly show the difference in behavior.

This repo acts as my personal lab notebook while reading the paper — documenting what breaks, why it breaks, and what actually works in practice. Many examples are inspired by real-world systems like Kubernetes and etcd, where these bugs caused production issues. The emphasis is on understanding runtime behavior, not just writing “correct-looking” code.

If you are reading the research paper and want concrete, runnable examples instead of pseudocode, this repository is meant to be used alongside it. Every file is intended to be executed, modified, and experimented with to build an intuitive understanding of Go concurrency pitfalls and safe patterns.

This is a learning and experimentation repository, not a library.
