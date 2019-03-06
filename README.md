# karp - universal app launcher

# EXAMPLES

```console
$ karp karp.gif
...

$ karp https://raw.githubusercontent.com/mcandre/karp/master/karp.txt
...

$ karp
Usage: karp <path>

-v Show version information
```

![magikarp](https://raw.githubusercontent.com/mcandre/karp/master/karp.gif)

# ABOUT

`karp` is a cross-operating system binary for easily launching file-based applications, using the skratchdot/open-golang library to route calls to `xdg-open`, etc.

`karp` will use the default application registered for a given path's protocol, mimetype, or file extension.

# DOWNLOAD

https://github.com/mcandre/karp/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/karp

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.11+

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/karp/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/karp.git $GOPATH/src/github.com/mcandre/karp
$ cd $GOPATH/src/github.com/mcandre/karp
$ git submodule update --init --recursive
$ sh -c 'cd cmd/karp && go install'
```

# PORT

```
$ mage port
```

# LINT

Keep the code tidy:

```
$ mage lint
```

# LICENSE

FreeBSD

# CREDIT

Powered by [open-golang](https://github.com/skratchdot/open-golang)
