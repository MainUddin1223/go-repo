## Package scope

### Package declaration

##### Best practice to declare first package same as folder name

As we know we have to declare package with file name
if we create more than one file we have to use same package with file name

### To call a function from another file

we must need to declare the function with capital latter. Otherwise it remain local scope. If it start with capital latter it become package scope

### Import folder/custom package to other folder

##### Need to create new module

run ` go mod init module name`

### variable shadowing

Variable shadowing occurs when a new variable with the same name is declared in an inner scope, such as inside a function or block.
This new variable "shadows" (hides) the outer variable — it does not modify the original variable, and no redeclaration error occurs because it's a different scope.

```
var a = 100 // global variable

func main() {
    a := 50 // this shadows the global 'a'
    fmt.Println(a) // prints 50, not 100
}

```

### Function types

#### Standard Function

A function with name is called standard function

```
func add(a int,b int){
    fmt.Println(a+b)
}
```

#### Init Function

A Function is invoked immediately while executing the code is called init function
we must se init after func to declare init function

```
func init(){
    fmt.Println("Hello Init function")
}
```

#### Anonymous Function

A function without name is called anonymous function and it must called immediately

```
func main(){
    (a int,b int){
    c :=a+b
    fmt.Println(c)
    }
}(5,7) //It also called Immediately Invoked Function Expression
```

#### Function Expression and Assign function in a variable

```
add : = func(a int , b int){
    c:= a+b
    fmt.Println(c)
}
add(1,3)
```

#### Interview question link

https://docs.google.com/document/d/1g7UHRVzjVGdeKdbyqv6HwJE2k7SoDIvwgvQa58VzvBs/edit?fbclid=IwY2xjawL9VfhleHRuA2FlbQIxMQABHvyrNqPxpkYoWP11rO_tW6MV17Dysvav5wzTuXCktJVlumSzcOGxIcf9WGUM_aem_0u_u3ff1N3sk_QbxldLtpg&tab=t.0
