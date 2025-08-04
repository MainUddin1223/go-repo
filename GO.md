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

if a variable with same name exist more than one times it will prioritize scopes.
like same variable declared in global and local . it will reflect local value but it will not replace actual value. SO outside of the block we will get global value

### Function types

#### Standard Function

A function with name is called standard function

```
func add(a int,b int){
    fmt.Println(a+b)
}
```

#### Init Function

A Function is invoked immediately while executing the code block is called init function
we must se init after func to declare init function

```
func init(){
    fmt.Println("Hello Init function")
}
```
