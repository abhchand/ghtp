# ghtp

`ghtp` is a command line utility that provides integration between Github and TargetProcess

<img src="https://raw.githubusercontent.com/abhchand/ghtp/master/meta/how-it-works.png" alt="How it Works" height="350" />

### Opt In

The sync is purely opt-in per Pull Request. It will only consider Pull Requests that have the formatted TargetProcess ID in the title as follows:

<img src="https://raw.githubusercontent.com/abhchand/ghtp/master/meta/pull-request-title.png" alt="How it Works" height="65" />

# Setup

Download the [latest `ghtp` release](https://github.com/abhchand/ghtp/releases)

```
# linux
wget --quiet https://github.com/abhchand/ghtp/releases/download/v0.1-beta/ghtp0.1-beta.linux-amd64.tar.gz

# OSX
wget --quiet https://github.com/abhchand/ghtp/releases/download/v0.1-beta/ghtp0.1-beta.darwin-amd64.tar.gz
```

Extract the file and move the executable to somewhere in your `$PATH`

```
tar -v -xzf ghtp0.1-beta.darwin-amd64.tar.gz
mv ghtp0.1-beta.darwin-amd64/ghtp /usr/local/bin/
```

Fill out a new config file. To get started you can use the example file in the `ghtp` repository:

```
wget --quiet "https://raw.githubusercontent.com/abhchand/ghtp/master/example/config.yml"
```

# Run

Run the sync

```
ghtp sync -config-file ./config.yml -v
```

You can also schedule it as a regular job via `cron` or similar scheduling utility

```
*/5 * * * * /usr/local/bin/ghtp sync -config-file /path/to/config.yml >> /tmp/ghtp.log 2>&1
```
