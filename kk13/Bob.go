package main

import (
	"fmt"
	"math/big"
)

//Bob as R

type Bob struct {
	mes [][2]*big.Int;  // message for primitive OTkm
	G *Group;
	r [] int;
	T *bitMatrix;
	m,k int;
	z [][]byte;
}

func (B *Bob) init ( x *Group, m int, k int , choice[]int,WHC [][]bool){
	B.G = x;
	B.m = m;
	B.k = k;
  	B.r = choice;
  	var t bitMatrix;
  	t.init(m,k);
  	B.T=&t;
  	var mesT [][2][]bool;
  	for i:=0;i<m;i++{
  		var rn [2][]bool;

  		rn[0]= Byte2Bit(B.T.row(i));
  		rn[1]= Byte2Bit(byteXor(B.T.row(i),Bit2Byte(WHC[B.r[i]])));

  		mesT= append(mesT,rn);
  	}
  	for i:=0;i<k;i++{
  		var rn [2][]bool;
  		for j:=0;j<m;j++{
  			rn[0] = append(rn[0],mesT[j][0][i]);
  			rn[1] = append(rn[1],mesT[j][1][i]);
		}
		var t [2]*big.Int;
  		t[0] = byte2Int(Bit2Byte(rn[0]));
		t[1] = byte2Int(Bit2Byte(rn[1]));
		B.mes = append(B.mes,t);
	}
}


func (B *Bob) Receive (pk [][2]*big.Int) [][2]*Ct {
	return OTkmSend(B.G,B.k,B.m,pk,B.mes);

}



func (B *Bob) Receive2 (y [][N][] byte)  {
   for i:=0;i<B.m;i++{
   	//fmt.Println("choice")
   	//fmt.Println(B.r[i])
   	zn:=byteXor(y[i][B.r[i]],H(i,B.T.row(i),B.k));
   	fmt.Println(string(zn));
   	B.z = append(B.z,zn);
   }

}