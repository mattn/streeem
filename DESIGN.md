All of interface has two connection similar to ReadWriter. Some function or
constant value handle just one in both. For example, seq(100) don't read any
input. STDIN also. STDOUT don't write any thing. But input will be written in
console. Below is a pipeline.

```
seq(100) -> {|x| ... } -> STDOUT
```

# seq(100)
```go
type IO struct {
   rc chan interface{}
   wc chan interface{}
}

type Seq struct {
   IO
   current int
   max int
}

func (s *Seq) ReadWrite() error {
    if s.current < s.max {
        s.current += 1
        s.wc <- s.current
        return nil
    }
    return io.EOF
}
```

# {|x| ...}
```go
type Func struct {
   IO
   f FuncExpr
}

func (f *Func) ReadWrite() (err error) {
    defer func() {
        if e := recover(); e != nil {
            err = e
        }
    }()
    f.wc <- f.Invoke(<- f.rc)
    return nil
}
```

# STDOUT
```go
type OUT struct {
   IO
   w io.Writer
}

var STDOUT = NewOut(os.Stdout)

func (o *OUT) ReadWrite() error {
    _, err := o.w.Write(fmt.Sprint(<-o.rc))
    return err
}
```

# JOIN

`seq(100)` - goroutine - `{|x| ...}` - goroutine - `STDOUT`

```go
var wg sync.WaitGroup
for _, item := range pipeline {
    wg.Add(1)
    go func(item *IO) {
        for {
            err := item.ReadWrite()
            if err != nil {
                ec <- err
                break
            }
	    }
        wg.Done()
    }(item)
}
wg.Wait()
err := <-ec
```
