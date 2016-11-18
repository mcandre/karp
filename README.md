# karp - universal app launcher

# EXAMPLES

```
$ go get github.com/mcandre/karp/...

$ karp karp.gif
$ karp https://raw.githubusercontent.com/mcandre/karp/master/karp.gif
```

![magikarp](https://raw.githubusercontent.com/mcandre/karp/master/karp.gif)

# ABOUT

`karp` is a cross-operating system binary for easily launching file-based applications, using the skratchdot/open-golang library to route calls to `xdg-open`, etc.

`karp` will use the default application registered for a given path's protocol, mimetype, or file extension.

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
