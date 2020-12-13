package main

import (
	"crypto/ecdsa"
)

type Bob struct{
	mes [][2][]byte;
	pk [][2]*ecdsa.PublicKey;
	r []int;
	T *bitMatrix;
	m,k,l int;
	z [][]bool;
	n int;
}

func (B *Bob) init(choice *[]int,n int,m int,l int,k int, code [][]byte){
	B.m = m;
	B.k = k;
	B.n = n;
	B.r = *choice;
	B.l = l;
	var t bitMatrix;
	t.init(m,k);
	B.T=&t;
	var mesT [][2][]bool;
	for i:=0;i<m;i++{
		var rn [2][]bool;

		rn[0] = Byte2Bit(B.T.row(i));
		rn[1] = Byte2Bit(byteXor(B.T.row(i),code[B.r[i]]));

		mesT=append(mesT,rn);
	}

	for i:=0; i<k; i++{
		var rn[2][]bool;
		for j:=0;j<m;j++{
			rn[0] = append(rn[0],mesT[j][0][i]);
			rn[1] = append(rn[1],mesT[j][1][i]);
		}
		var t [2][]byte;
		t[0] = Bit2Byte(rn[0]);
		t[1] = Bit2Byte(rn[1]);
		B.mes = append(B.mes,t);
	}
}

func (B *Bob) OTkmReceive (pk [][2]*ecdsa.PublicKey) [][][]byte{
	B.pk = pk;
	return OTkmSend(B.k,pk,B.mes);
}

func(B *Bob) Receive(y [][][] bool){
	//fmt.Print(y);
	for i := 0; i < B.m; i++ {
		zn := boolXor(y[B.r[i]][i], Byte2Bit(H(i, B.T.row(i), B.l))[:B.l]);
		B.z = append(B.z, zn);
	}
}