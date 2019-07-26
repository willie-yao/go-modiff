% go-modiff(8) 
% Sascha Grunert
% July 2019

# NAME

go-modiff - Command line tool for diffing go module dependency changes between versions

# SYNOPSIS

go-modiff

```
[--from|-f]=[value]
[--help|-h]
[--repository|-r]=[value]
[--to|-t]=[value]
[--version|-v]
```

# DESCRIPTION

Command line tool for diffing go module dependency changes between versions

**Usage**:

```
go-modiff [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
```

# GLOBAL OPTIONS

**--from, -f**="": the start of the comparison, any valid git rev (default: master)
**--help, -h**: show help
**--repository, -r**="": repository to be used, like: github.com/owner/repo
**--to, -t**="": the end of the comparison, any valid git rev (default: master)
**--version, -v**: print the version

# COMMANDS

## docs, d

generate the markdown or man page documentation and print it to stdout

**--man**: print the man version
**--markdown**: print the markdown version

## help, h

Shows a list of commands or help for one command
