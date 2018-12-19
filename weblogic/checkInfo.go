/***********************************************************************
# File Name: checkInfo.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-19 16:30:38
*********************************************************************/
package weblogic

import(
    "fmt"
    "bufio"
    "os/exec"
    "log"
)

func errcheck(err error){
    if err != nil{
        panic(err)
    }
}

func Run(command string) []string{
    cmd := exec.Command("/bin/bash","-c",command)

    stdout,cmdErr := cmd.StdoutPipe() 
    errcheck(cmdErr)

    
    startErr := cmd.Start()
    errcheck(startErr)
    outputBuf := bufio.NewReafer(stdout)
    var result []string
    for {
        output,err := outputBuf.ReadString('\n')
        if err != nil || err.Error() == "EOF"{
            log.Fatal(err)
            break
        } 
        result = append(result,output)
    }
    return result
}

func CheckInfo() bool{
     
}
