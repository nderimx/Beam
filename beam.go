package main
import ("net"
		"bufio"
		"fmt"
		"os"
		"strings"
		"io/ioutil"
		"encoding/json")
type Pair struct{
	S string
	F []byte
}
func listen(port string){
	for{
		ln, _ := net.Listen("tcp", ":"+port)
		if ln!=nil{
			conn, _ := ln.Accept()
			message, _:=bufio.NewReader(conn).ReadString('\n')
			var pr Pair
			var msg string
			if err:=json.Unmarshal([]byte(message), &msg); err==nil{
				fmt.Print("\nreceived: "+string(msg)+"\nsend: ")
			}else if err:=json.Unmarshal([]byte(message), &pr); err==nil{
				name:=pr.S
				f:=pr.F
				err:=ioutil.WriteFile(name, f, 0644)
				if err==nil{
					fmt.Print("\n"+name+" received!\nsend: ")
				}
			}
			
		}
		ln.Close()
	}
}
func send(address string, port string) error{
	for{
		fmt.Print("send: ")
		reader:=bufio.NewReader(os.Stdin)
		line, _, _:=reader.ReadLine()
		text:=string(line)
		if strings.HasPrefix(text,"transfer"){
			substring:=text[9:len(text)]
			file, err:=ioutil.ReadFile(substring)
			if err!=nil{
				return err
			}
			conn, err:=net.Dial("tcp", address+":"+port)
			if err==nil{
				paketa:=Pair{substring,file}
				jpaketa, _:=json.Marshal(paketa)
				conn.Write(jpaketa)
				fmt.Println(substring+" sent!")
			}else{
				return err
			}
			conn.Close()
		}else{
			conn, err:=net.Dial("tcp", address+":"+port)
			if err!=nil{
				return err
			}
			jtext, _:=json.Marshal(text)
			conn.Write(jtext)
			conn.Close()
		}
	}
}
func main(){
	address:=string(os.Args[1])
	port:=string(os.Args[2])
	go listen(port)
	send(address, port)
}