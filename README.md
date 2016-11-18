# karp - universal app launcher

# EXAMPLE

```
$ go get github.com/mcandre/karp/...
$ karp karp.gif
```

# ABOUT

Operating systems vary on syntax for launching GUI applications from command line interfaces. macOS uses `open`, Linux uses `xdg-open`, and Windows uses `start`. karp provides a small, self-sufficient binary for helping scripts launch graphical applications, in a way that supports a wide variety of different environments.

# REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [Git](https://git-scm.com)
* [Make](https://www.gnu.org/software/make/)
* [Bash](https://www.gnu.org/software/bash/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/karp/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/karp.git $GOPATH/src/github.com/mcandre/karp
$ cd $GOPATH/src/github.com/mcandre/karp
$ git submodule update --init --recursive
$ sh -c 'cd cmd/karp && go install'
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
