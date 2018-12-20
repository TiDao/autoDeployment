/***********************************************************************
# File Name: test.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-20 15:48:34
*********************************************************************/
package main

import(
    "fmt"
)

func test() []int{
    var test []int
    test = []int{1,2,2,2,3,3341,3,1,4,14,1,4,1,41}
    return test 
}


func main(){
    result := test()
    for _,arg := range result{
        fmt.Println(arg)
    }
    result = []int{1,23123}
    for arg,_ := range result{
        fmt.Println(arg)
    }
    
}
