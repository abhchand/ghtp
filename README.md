# ghtp

`ghtp` is a command line utility that provides integration between Github and TargetProcess

- [Setup](#setup)
- [Run](#run)
- [Building from Source](#building-from-source)

<img src="https://raw.githubusercontent.com/abhchand/ghtp/master/meta/how-it-works.png" alt="How it Works" height="350" />

### Opt In

The sync is purely opt-in per Pull Request. It will only consider Pull Requests that have the formatted TargetProcess ID in the title as follows:

<img src="https://raw.githubusercontent.com/abhchand/ghtp/master/meta/pull-request-title.png" alt="How it Works" height="65" />

# <a name="setup"></a>Setup

Download the [latest `ghtp` release](https://github.com/abhchand/ghtp/releases)

```
# linux
wget --quiet https://github.com/abhchand/ghtp/releases/download/v0.1/ghtp0.1.linux-amd64.tar.gz

# OSX
wget --quiet https://github.com/abhchand/ghtp/releases/download/v0.1/ghtp0.1.darwin-amd64.tar.gz
```

Extract the file and move the executable to somewhere in your `$PATH`

```
tar -v -xzf ghtp0.1.darwin-amd64.tar.gz
mv ghtp /usr/local/bin/
```

Fill out a new config file. To get started you can use the example file in the `ghtp` repository:

```
wget --quiet "https://raw.githubusercontent.com/abhchand/ghtp/master/example/config.yml"
```

# <a name="run"></a>Run

Run the sync

```
ghtp sync -config-file ./config.yml -v
```

You can also schedule it as a regular job via `cron` or similar scheduling utility

```
*/5 * * * * /usr/local/bin/ghtp sync -config-file /path/to/config.yml >> /tmp/ghtp.log 2>&1
```

# <a name="building-from-source"></a>Building from Source


Ensure Go is installed (see [Go installation page](https://golang.org/doc/install)) and your `$GOPATH` is set

Clone this repository:

```
mkdir -p $GOPATH/src/github.com/abhchand
git clone https://github.com/abhchand/ghtp.git $GOPATH/src/github.com/abhchand/ghtp
```

Install [`go-dep`](https://golang.github.io/dep/docs/installation.html).

```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

Build the project. If `$GOPATH/bin` is not in your `$PATH` you may have to reference `dep` as `$GOPATH/bin/dep`.

```
dep ensure
go build
```
