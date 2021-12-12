# gopath e goroot

export GOROOT=/usr/local/go
export GOPATH=$HOME/go

## Run

go run ccompiler/ main.ct

https://norasandler.com/2017/11/29/Write-a-Compiler.html

go run ccompiler testing/stage_1/valid/return_2.c

instalar a biblioteca:

```python
sudo apt-get install gcc-multilib
```

```python
./ccompiler testing/stage_1/valid/return_2.c
```

```python
gcc -m32 return_2.s -o return_2
```
