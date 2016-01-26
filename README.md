# embd
A tool to embed files into go source code as strings.

```
embd [options] file

```

# Examples

```
$ embd --help
  embed a file into go source code
  Example: 'embd a.sql' creates a.sql.go which which has 'const a = "...contents of a.sql..."'
  Example: 'embd a.sql -n foo' creates a.sql.go with 'const foo = ...'
  Options:
    -n string
      	name of the variable to generate (default <filename>)
```

# More
look at the godoc: https://godoc.org/github.com/tscholl2/embd
