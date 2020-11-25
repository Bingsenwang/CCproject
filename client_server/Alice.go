package main

import (
	"math/big"
)

//Alice as S

type Alice struct {
	sk   []*big.Int; // sk for OTkm
	G    *Group;
	x    [][2]*big.Int;
	s    []int;
	m, k int;
	Q    *bitMatrix;
}

func (A *Alice) init(x *Group, y *[][2]string, m int, k int) {
	A.G = x;
	A.m = m;
	A.k = k;
	A.sk = A.G.Gensk(k);
	for i := 0; i < k; i++ {
		A.s = append(A.s, RANseed.Intn(2));
	}
	for i := 0; i < m; i++ {
		var rn [2]*big.Int;
		rn[0] = str2Int((*y)[i][0], A.k);
		rn[1] = str2Int((*y)[i][1], A.k);
		A.x = append(A.x, rn);
	}
}

func (A *Alice) Send1() [][2]*big.Int {
	return OTkmGenPk(A.G, A.k, A.s, A.sk);
}

func (A *Alice) Receive(c [][2]*Ct) {
	mes := OTkmGetMes(A.G, A.k, A.s, A.sk, c);
	var q bitMatrix;
	q.init(A.m, A.k);
	A.Q = &q;
	for i := 0; i < A.k; i++ {
		col := Byte2Bit(appenleftbyte(mes[i].Bytes(), A.m));
		for j := 0; j < A.m; j++ {
			(*A.Q.e)[j][i] = col[j];
		}
	}

}

func (A *Alice) Send2() *[][2][]byte {
	var y [][2][]byte;
	for i := 0; i < A.m; i++ {
		var rn [2] []byte;
		rn[0] = byteXor(A.x[i][0].Bytes(), H(i, A.Q.row(i), A.k));
		rn[1] = byteXor(A.x[i][1].Bytes(), H(i, byteXor(Bits2Byte(A.s), A.Q.row(i)), A.k))
		y = append(y, rn);
	}
	return &y;
}
