### memfd_create()
From the `memfd_create()` man page: https://man7.org/linux/man-pages/man2/write.2.html

```
memfd_create() creates an anonymous file and returns a file descriptor that refers to it. The file behaves like a regular file, and so can be modified, truncated, memory-mapped, and so on. However, unlike a regular file, it lives in RAM and has a volatile backing storage.
```

We can then use the `fexecve()` to exec the code at the file descriptor returned from `memfd_create()`. From the `fexecve()` man pages: https://man7.org/linux/man-pages/man3/fexecve.3.html

```
fexecve() performs the same task as execve(2), with the
       difference that the file to be executed is specified via a file
       descriptor, fd, rather than via a pathname.
```


There are a number of blogs on the Internet about this method of fileless execution for Linux Malware already, I don't think I need to pile on.
- https://0x00sec.org/t/super-stealthy-droppers/3715/1
- https://blog.sonatype.com/pypi-package-secretslib-drops-fileless-linux-malware-to-mine-monero

### PoC for testing
PoC written in Golang, needs `go 1.19` to build via Makefile
