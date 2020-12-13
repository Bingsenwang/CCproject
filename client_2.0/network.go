package main

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net"
)

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

func SendPk(Pk *ecdsa.PublicKey,conn net.Conn){
	Sendbign(conn, Pk.X);
	Sendbign(conn, Pk.Y);
}

func SendPks(Pks[][2]*ecdsa.PublicKey, conn net.Conn, k int) {
	for i:=0;i<k;i++{
		SendPk(Pks[i][0],conn);
		SendPk(Pks[i][1],conn);
	}
}

func ReceivePk(conn net.Conn) *ecdsa.PublicKey{
	var pk ecdsa.PublicKey;
	pk.Curve = curve;
	pk.X = Receivebign(conn);
	pk.Y = Receivebign(conn);
	return &pk;
}

func RecievePks(conn net.Conn, k int) [][2]*ecdsa.PublicKey{
	var pks [][2]*ecdsa.PublicKey;
	for i:=0;i<k;i++{
		var pk [2]*ecdsa.PublicKey;
		pk[0] = ReceivePk(conn);
		pk[1] = ReceivePk(conn);
		pks = append(pks,pk);
	}
	return pks;
}


func SendXs(data [][][]byte, conn net.Conn, m int, n int, k int){
	var wb []byte;
	//fmt.Println(len(data[0][0]))
	wb = byte_join(&data, m, n, k/8);
	for i:=0;i<len(wb);i+=20480 {
		if i+20480<len(wb) {
			conn.Write(wb[i : i+20480]);
		} else{
			conn.Write(wb[i : ]);
		}
		reader := bufio.NewReader(conn);
		var buf [1]byte;
		reader.Read(buf[:]);
		if int(buf[0]) != 1 {
			fmt.Println("Send bytes error")
		};
	}

}


func RecieveXs(conn net.Conn, m int, n int,k int)[][][]byte{

	var data []byte;
	var by [][][]byte;

	len_data := m*n*k/8/20480;
	var i int;
	for i = 0; i <= len_data; i++ {
		var buf0 [20480]byte;
		n0, _ := conn.Read(buf0[:]);
		var one [1]byte;
		one[0] = byte(1);
		conn.Write(one[:]);
		data = append(data, buf0[0:n0]...);
	}


	by = byte_split(data, m, n, k/8);


	return by;

}


func byte_split (x []byte,m int, n int,k int) [][][]byte {

	var rt[][][]byte;
	for i:=0;i<m;i++{
		var ri[][]byte;
		for j:=0;j<n;j++{
			var rij[]byte;
			for l:=0;l<k;l++ {
				rij = append(rij, x[i*n*k+j*k+l]);
			}
			ri = append(ri,rij);
		}
		rt = append(rt,ri);
	}

	return rt;
}

func byte_join (x *[][][]byte,m int, n int,k int) []byte {
	if (k == -1){
		k = len((*x)[0][0]);
	}
	var rt[]byte;
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			for l:=0;l<k;l++ {
				rt = append(rt, (*x)[i][j][l]);
			}
		}
	}
	return rt;
}