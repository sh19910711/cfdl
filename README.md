[![Build Status](https://travis-ci.org/sh19910711/cfurl.svg?branch=master)](https://travis-ci.org/sh19910711/cfurl)

### The CLI tool for transferring data from codeforces with URLs

```
$ go get github.com/sh19910711/cfurl
$ cfurl http://codeforces.com/contest/912/submission/33932854 > source.code
```

### TODO

- [ ] cfurl http://codeforces.com/contest/912/problem/A
	- It should return the problem statement.
- [ ] cfurl --input http://codeforces.com/contest/912/problem/A
	- It should return the example inputs.
- [ ] cfurl --output http://codeforces.com/contest/912/problem/A
	- It should return the example outputs.
- [x] cfurl http://codeforces.com/contest/912/submission/33932854
	- It should return the source code of the submission.
- [ ] cfurl --input http://codeforces.com/contest/912/submission/33932854
	- It should return the inputs.
- [ ] cfurl --answer http://codeforces.com/contest/912/submission/33932854
	- It should return the answers.
