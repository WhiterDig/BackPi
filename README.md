# BackPi

Before using this application, cd to the project directory and run a 

```
go get -u github.com/kardianos/govendor
```

We use this to manage our dependencies.


Afterwards, run 
```
govendor init
```
to initialize it

finally, to get all necessary dependencies, run

```
govendor fetch +external
govendor fetch +missing
```

You should be able to just build and run using
```
go build
BackPi.exe (or ./BackPi on Unix)
```
