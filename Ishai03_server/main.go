package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"time"
)

var RANseed = rand.New(rand.NewSource(int64(time.Now().Unix())))

// the random oracle, note that l should divide by 8
func H(m int, ls []byte, l int) []byte {
	x := append(ls, byte(m));
	rt := sha256.Sum256(x)
	return rt[0 : l/8];
}
func Sendbign (conn net.Conn, x *big.Int){
	wb := x.Bytes();
	if len(wb)==0 {wb = append(wb,byte(0))};
	conn.Write(wb);
	reader := bufio.NewReader(conn)
	var buf [1]byte
	reader.Read(buf[:])
	if int(buf[0])!=1 {fmt.Println("Send big integer error")};
}

func Receivebign (conn net.Conn) *big.Int{
	rt := big.NewInt(0);
	var buf0 [20480]byte;
	n0,_ := conn.Read(buf0[:]);
	var one [1]byte;
	one[0]= byte(1);
	conn.Write(one[:]);
	rt = rt.SetBytes(buf0[0:n0]);
	return rt;
}

func Sendpk (pk [][2]*big.Int, conn net.Conn, k int){
 for i:=0;i<k;i++{
	 Sendbign(conn,pk[i][0]);
	 Sendbign(conn,pk[i][1]);
 }
}

func Receivect (conn net.Conn, k int) [][2]*Ct{
	var rt [][2]*Ct;
	for i:=0;i<k;i++{
		var rn [2]*Ct;
		var rn0 Ct;
		var rn1 Ct;
		rn0.init(Receivebign(conn),Receivebign(conn));
		rn1.init(Receivebign(conn),Receivebign(conn));
		rn[0]=&rn0;
		rn[1]=&rn1;
		rt = append(rt,rn);
	}
	return rt;
}


func main() {
	m := 256;
	k := 64; //we take l=k for now
	var A Alice;
	//var B Bob;
	G := Group_generator(200);
	y := [][2]string{{"hello", "world"}, {"hello", "world"}};
	choice := []int{0, 1};
	for i:=0;i<15;i++{
		y = append(y,y...);
		choice = append(choice,choice...);
	}

	A.init(G, &y, m, k);

	listen, err := net.Listen("tcp", "192.168.8.110:9090")

	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	fmt.Println("start listen")

	for {
		// wait for connection request
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
	    Sendbign(conn,G.p);
		Sendbign(conn,G.q);
		Sendbign(conn,G.g);

     pk := A.Send1();
     Sendpk(pk,conn,k);
     ct := Receivect(conn,k);
     A.Receive(ct);

     by := byte_join(A.Send2(),m,k/8);
     conn.Write(by);
	}



	/*B.init(G, m, k, &choice);
	A.Receive(B.Receive(A.Send1()));
	B.Receive2(A.Send2());*/
}
