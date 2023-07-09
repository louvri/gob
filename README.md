# gob
<hr style="border:1px solid #444; margin-top: -0.5em;">  

`gob` is a library that provides common golang's object process.  
This README is still not fully completed yet, will be updated shortly.
The codes were co-created with [@johnjerrico](https://github.com/johnjerrico), published here so it can be used publicly.

### Installation
<hr style="border:1px solid #444; margin-top: -0.5em;">  

Get the code with:
```
$ go get github.com/louvri/gob
```
### Usage
<hr style="border:1px solid #444; margin-top: -0.5em;">  

- Array/Slice
  ```
    ...
    source := []string{"A", "B", "C", "D"}
    ...
    indexMap := arr.Index(source)
    if indexMap["A"] {
    ...
    }
    ...
  ```
- Object
  ```
    ...
    if !object.IsEmpty(objB) {
        err := object.Assign(objA, prop, object.Get(objB))
        if err != nil {
            fmt.Println(err)
        }
    }
    ...
  ```
- String
  ```
    ...
    if tmp := str.SplitOnNotEmpty("id:desc", ":"); len(tmp) > 0 {
        ...
    }
    ...
  ```
