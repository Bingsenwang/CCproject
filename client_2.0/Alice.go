package main

import (
	"crypto/ecdsa"
)

type Alice struct {
	sk []*ecdsa.PrivateKey;
	pk [][2]*ecdsa.PublicKey;
	x [][][]bool;
	code [][]byte;
	s []int;
	Q *bitMatrix;
	n,m,k,l int;
}

func (A *Alice) init(x*[][][]bool, n int, m int,l int, k int, code [][]byte){
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

func (A *Alice) OTkmReceive(c [][][]byte){
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

func (A *Alice) Send() [][][]bool{
	var y[][][]bool;
	//for i:=0;i<A.m;i++{
	//	for j:=0;j<A.n;j++ {
	//		fmt.Print(byteXor(A.Q.row(i), byteAnd(Bits2Byte(A.s), A.code[j])));
	//	}
	//	fmt.Println();
	//}
	for i:=0;i<A.n;i++{
		var rn[][]bool;
		for j := 0; j < A.m; j++ {
			rn = append(rn, boolXor(A.x[j][i], Byte2Bit(H(j, byteXor(A.Q.row(j), byteAnd(Bits2Byte(A.s), A.code[i])), A.l))[:A.l]));
			}
		y= append(y,rn);
	}
	return y;
}