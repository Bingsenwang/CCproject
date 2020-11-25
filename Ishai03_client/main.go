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

func Sendbign (conn net.Conn, x *big.Int){
	wb := x.Bytes();
	if len(wb)==0 {wb = append(wb,byte(0))};
	conn.Write(wb);
	reader := bufio.NewReader(conn)
	var buf [1]byte
	reader.Read(buf[:])
	if int(buf[0])!=1 {fmt.Println("Send big integer error")};
}

func Receivepk (conn net.Conn, k int)[][2]*big.Int{
	var rt [][2]*big.Int;
	for i:=0;i<k;i++{
		var rn[2]*big.Int;
		rn[0] = Receivebign(conn);
		rn[1] = Receivebign(conn);
		rt = append(rt,rn);
	}
	return rt;
}

func Sendct (ct [][2]*Ct,conn net.Conn,k int){
  for i:=0;i<k;i++{
  	Sendbign(conn,ct[i][0].c1);
  	Sendbign(conn,ct[i][0].c2);
  	Sendbign(conn,ct[i][1].c1);
  	Sendbign(conn,ct[i][1].c2);
  }
}

func main() {
	m := 256;
	k := 64; //we take l=k for now

	var B Bob;
	G := Group_generator(50);
	y := [][2]string{{"hello", "world"}, {"hello", "world"}};
	choice := []int{0, 1};
	for i:=0;i<15;i++{
		y = append(y,y...);
		choice = append(choice,choice...);
	}

	conn, err := net.Dial("tcp", "84.238.34.142:8080")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	G.p = Receivebign(conn);
	G.q = Receivebign(conn);
	G.g = Receivebign(conn);

	B.init(G, m, k, &choice);



	pk:= Receivepk (conn,k);
    ct := B.Receive(pk);
    Sendct (ct, conn, k);
	var buf0 [204800]byte;
    n0,_ := conn.Read(buf0[:]);
    by := byte_split(buf0[0:n0],m,k/8);
    B.Receive2(&by);

}
