package main

import (
	"math/big"
)

//Alice as S

type Alice struct {
	sk []*big.Int; // sk for OTkm
	G *Group;
	s []int;
	m, k int;
	n int;
	x [][N]*big.Int;
	Q *bitMatrix;
	WHC [][]bool;
}

func (A *Alice) init ( x *Group, y [][N]string,m int , k int,WHC [][]bool){
	A.G = x;
	A.m = m;
	A.k = k;
	A.n = N;
	A.sk = A.G.Gensk(k);
	A.WHC = WHC;
	for i:=0;i<k;i++{
		A.s = append(A.s,RANseed.Intn(2));
	}
	for i:=0;i<m;i++{
		var rn[N]*big.Int;
		for j:=0;j<N;j++{
			rn[j] = str2Int(y[i][j],A.k);
		}
		A.x = append(A.x,rn);
	}
}



func (A *Alice) Send1 () [][2]*big.Int {
	return OTkmGenPk(A.G,A.k,A.s,A.sk);
}

func(A *Alice) Receive (c [][2]*Ct){
	mes := OTkmGetMes(A.G,A.k,A.s,A.sk,c);
	var q bitMatrix;
	q.init(A.m,A.k);
	A.Q=&q;
	for i:=0; i<A.k;i++{
		col := Byte2Bit(appenleftbyte(mes[i].Bytes(),A.m));
		for j:=0;j<A.m;j++{
			A.Q.e[j][i] = col[j];
		}
	}
}

func (A *Alice) Send2 () [][N][]byte{
	var y [][N][]byte;
    for i:=0;i<A.m;i++{
    	var rn [N] []byte;
    	for j:=0;j<N;j++ {
			rn[j] = byteXor(A.x[i][j].Bytes(), H(i, byteXor(A.Q.row(i),byteAnd(Bits2Byte(A.s),Bit2Byte(A.WHC[j]))), A.k));
		}
		y= append(y,rn);
	}
    return y;
}