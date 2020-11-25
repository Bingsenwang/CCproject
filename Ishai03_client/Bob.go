package main

import (
	"fmt"
	"math/big"
)

//Bob as R

type Bob struct {
	mes  [][2]*big.Int; // message for primitive OTkm
	G    *Group;
	r    [] int;
	T    *bitMatrix;
	m, k int;
	z    [][]byte;
}

func (B *Bob) init(x *Group, m int, k int, choice *[]int) {
	B.G = x;
	B.m = m;
	B.k = k;
	B.r = *choice;
	B.r = B.r[0:m]
	var t bitMatrix;
	t.init(m, k);
	B.T = &t;
	for i := 0; i < k; i++ {
		var rn [2]*big.Int;
		rn[0] = byte2Int(B.T.col(i));
		rn[1] = byte2Int(byteXor(B.T.col(i), Bits2Byte(B.r)));
		B.mes = append(B.mes, rn);
	}
}

func (B *Bob) Receive(pk [][2]*big.Int) [][2]*Ct {
	return OTkmSend(B.G, B.k, pk, B.mes);
}

func (B *Bob) Receive2(y *[][2][] byte) {
	for i := 0; i < B.m; i++ {
		//fmt.Println("choice")
		//fmt.Println(B.r[i])
		zn := byteXor((*y)[i][B.r[i]], H(i, B.T.row(i), B.k));
		fmt.Println(string(zn));
		B.z = append(B.z, zn);
	}

}
