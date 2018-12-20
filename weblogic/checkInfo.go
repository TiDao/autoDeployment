/***********************************************************************
# File Name: checkInfo.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-19 16:30:38
*********************************************************************/
package weblogic

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)



func Run(command string) []string {
	cmd := exec.Command("/bin/bash", "-c", command)

	stdout, cmdErr := cmd.StdoutPipe()
	if cmdErr != nil {
		log.Fatal(err)
	}

	startErr := cmd.Start()
	if startErr != nil {
		log.Fatal(err)
	}
	outputBuf := bufio.NewReafer(stdout)
	var result []string
	for {
		output, err := outputBuf.ReadString('\n')
		if err != nil || err.Error() == "EOF" {
			log.Fatal(err)
			break
		}
		result = append(result, output)
	}
	return result
}

//check directory is exist or not,if it's not exist , create it
func IsExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		fmt.Println("can't find directory,so creating...")
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed %v \n", err)
			return err
		}
		fmt.Println("create directory success")
		return nil
	}
	return err
}

// check the weblogic.Data type value 
func (data *Data) CheckInfo() error {
	if err := IsExist(data.Env.JavaHome); err != nil {
		fmt.Prinf("check java home failed, the error is : %v", err)
	}
	if err := IsExist(data.Env.JreHome); err != nil {
		fmt.Printf("check Jre home failed,the error is: %v", err)
	}
	if err := IsExist(data.Env.WlsHome); err != nil {
		fmt.Printf("check Weblogic home failed,the error is:%v", err)
	}
}

func (data *Data) CreateUser() error {
	var command string

	if data.Env.UserName == "" {
		data.Env.UserName = "weblogic"
	}
	command = "useradd -md /home/weblogic/" + data.Env.UserName

	results := Run(command)
	for _, result := range results {
		fmt.Println(result)
	}

	if data.Env.Password == "" {
		data.Env.Password = "weblogic"
	}
	command = "echo " + data.Env.Pssword + " | passwd -stdin " + data.Env.UserName
	results = run(command)
	for _, result := range results {
		fmt.Println(result)
	}

}
