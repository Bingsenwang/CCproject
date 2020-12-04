package main

import (
	"crypto/ecdsa"
)

type Alice struct {
	sk []*ecdsa.PrivateKey;
	pk [][2]*ecdsa.PublicKey;
	x [][][]byte;
	code [][]byte;
	s []int;
	Q *bitMatrix;
	n,m,k,l int;
}

func (A *Alice) init(x*[][][]byte, n int, m int,l int, k int, code [][]byte){
	A.m = m;
	A.k = k;
	A.x = *x;
	A.n = n;
	A.l = l;
	A.code = code;
	for i:=0; i<k; i++{
		A.s = append(A.s,RANseed.Intn(2));
	}
	A.sk,A.pk = OTkmGenPk(A.k,A.s);
}

func (A *Alice) OTkmsendpk()([][2]*ecdsa.PublicKey){
	return A.pk;
}

func (A *Alice) OTkmReceive(c [][2][]byte){
	mes := OTkmGetMes(A.k,A.s,A.sk,c);
	var q bitMatrix;
	q.init(A.m,A.k);
	A.Q=&q;
	for i:=0;i<A.k;i++{
		col := Byte2Bit(appenleftbyte(mes[i],A.m));
		for j:=0;j<A.m;j++{
			A.Q.e[j][i] = col[j];
		}
	}
}

func (A *Alice) Send() [][][]byte{
	var y[][][]byte;
	for i:=0;i<A.m;i++{
		var rn[][]byte;
		for j:=0;j<A.n;j++ {
			rn = append(rn,byteXor(A.x[i][j], H(i, byteXor(A.Q.row(i),byteAnd(Bits2Byte(A.s),A.code[j])), A.l)));
		}
		y= append(y,rn);
	}
	return y;
}