<p align="center">
<img src="./assets/photos/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/koi/v3?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# lru

<i>lru</i> is a lightweight package that provides [lru](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU) cache algorithm in Go.

# Documentation

## Install

```bash
go get github.com/erfanmomeniii/lru
```   

Next, include it in your application:

```bash
import "github.com/erfanmomeniii/lru"
``` 
## Quick Start
The following example illustrates how to use this package
```go
package main

import (
	"fmt"
	"github.com/erfanmomeniii/lru"
)

func main() {
	l := lru.New(2)
	// 2 is the size of lru (default 2000)
	l.Set("a", 12)
	l.Set("b", 13)
	
	fmt.Println(l.Get("a"))
	// 12
	
	l.Set("c", "hi")
	
	fmt.Println(l.Get("c"))
	// hi
	fmt.Println(l.Get("a"))
	// nil
}

```
## Usage

#### type Lru 
```go
type Lru struct {
q    *queue.Queue
m    map[any]any
size int64
}
```
Lru is an instantiation of the lru.

#### func New
```go
func New(size ...int64) *Lru
```
New creates a new instance of a lru.

#### func Set
```go
func (l *Lru) Set(key any, value any)
```
Set adds or updates with inputted key and value.

#### func Get
```go
func (l *Lru) Get(key any) any
```
Get returns the value associated with the inputted key.

#### func Has
```go
func (l *Lru) Has(key any) bool 
```
Has checks whether the key exists or not.

## Contributing
Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change.
Please make sure to update tests as appropriate.



