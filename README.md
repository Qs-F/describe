# package describe

give description of directory.

## infomation

This is originally private repository.  
Use with `Qs-F/gondom` and `Qs-F/akdir` at the same time.

Like this:

```
cd $GOPATH/localhost/tmp
gondom 5 | akdir
describe "Golang struct and pointer try"
```

or when I want to find,

```
cd $GOPATH/localhost/tmp
describe -l
# print all described directories in list in "tmp" directory
```

```
cd $GOPATH/localhost/tmp
describe -s "Golang"
# print search result in list
```
