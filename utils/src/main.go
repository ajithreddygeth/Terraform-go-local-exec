package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("testing for go with the terraform")
	var aws_region = ""
	if len(os.Getenv("AWS_REGION")) != 0 {
		aws_region = os.Getenv("AWS_REGION")
	}
	var aws_account = ""
	if len(os.Getenv("AWS_ACCOUNT")) != 0 {
		aws_account = os.Getenv("AWS_Account")
	}
	fmt.Println(fmt.Sprintf("aws_region:%v, aws_account:%v,", aws_region, aws_account))
}
