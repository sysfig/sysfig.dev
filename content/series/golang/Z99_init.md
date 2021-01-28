The `init()` function is used to 

1. Initialize imported packages
2. Initialize constants
3. Initialize global variables
4. Run the `init()` function
5. (In an application) Run the `main()` function

There are multiple `init()` functions within the _same_ package, they will be executed in the same order code presented to the compiler:

- If the `init()` functions are from diffrent files, then the files are evaluated using lexical file name order
- If the `init()` functions are from the same file, then the file is evaluated from top-to-bottom

If a package is imported from multiple packages, it will only be initialized once.

![](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/images/2.3.init.png)

Starting from Go v1.16, you can set `GODEBUG=inittrace=1` to debug and profile `init()` functions - see when the init function is run after the start of the program, how long it took to run, how much heap memory was allocated, and how many calls to allocation memory was made. For each `init()` called, Go will emit a line to standard error.

Note that the duration is measured as the _wall-clock time_, which is the actual time it took for the function to run on your computer, and is affected by your machine spec as well as CPU load. This is different from the _user time_, which is how long the CPU spent exclusively running your code. It's also different from kernel time, which is the time the CPU spends on system calls initiated by the user's code; for example, kernel operations like file I/O, forking a new process, etc. are counted as kernel time.
