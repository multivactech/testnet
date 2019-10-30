package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/multivactech/MultiVAC/model/chaincfg/multivacaddress"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("MultiVAC Address conversion tool")
	fmt.Println("--------------------------------")

	fmt.Print("Please input your old address: ")
	addr, _ := reader.ReadString('\n')
	addr = strings.Replace(addr, "\n", "", -1)

	old, err := hex.DecodeString(addr)
	if err != nil {
		fmt.Println("")
		fmt.Println("ERROR: Please check your address is correct")
		return
	}
	newAddr := multivacaddress.GenerateAddress(old, multivacaddress.UserAddress)
	fmt.Println("New address:", newAddr)
}
