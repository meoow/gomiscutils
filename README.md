#Library for personal go programming
---
This library contains a collection of functions for use of personal Go programming.

###1. readline.go
 * **Type**  
 ```
 type fileReader interface {
		ReadString(byte) (string, error)
}
```

 * **Function**  
```
func Readline(fileReader) chan string
```  

###2. fileexists.go
 * **Type**  

 * **Function**  
 ```
 func FileExists(string) bool
 ```  
   
####4. countlines.go  
 * **Type**
 
 * **Function**  
 ```
 func Countlines(f *os.File) uint64
 ```
 
####5. splitext.go
 * **Type**
 
 * **Function**  
 ```
 func SplitExt(s string) (name, ext string)
 ```
 
####6. trimnewline.go
 * **Type**
 
 * **Function**  
 ```
 func TrimNewLineByte(b []byte) []byte
 ```  
 ```
 func TrimRightWhiteSpaceByte(b []byte) []byte
 ```  
 ```  
 func TrimNewLine(s string) string
 ```  
 ```
 func TrimRightWhiteSpace(s string) string
 ```

 
 
 