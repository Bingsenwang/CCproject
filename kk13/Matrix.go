package main

import (
	"fmt"
	"math/big"
)

type bitMatrix struct{
	m,k int;     //size of matrix
	e [][]bool; //both size should divide by 8
	E []bool;
}




func (M *bitMatrix) init(m int, k int){
	if m%8!=0|| k%8 !=0 {fmt.Println("exception: size of the matrix error")}
	M.m=m;
	M.k=k;
	for i:=0;i<m;i++{
		var rn []bool;
		for j:=0;j<k;j++{
			rn = append(rn, int2bool(RANseed.Intn(2)));
		}
		M.e=append(M.e,rn);
	}
}

//get any row or column of this matrix

func (M *bitMatrix) row (x int) []byte{
	return Bit2Byte(M.e[x]);
}


func (M *bitMatrix) col (x int) []byte{
	var rn []bool;
	for i:=0;i<M.m;i++{
		rn = append(rn,M.e[i][x]);
	}
	return Bit2Byte(rn);
}

// related operation over bit and byte
func Byte2Bit (x []byte) []bool{
	var rt []bool;
	for i:=0;i< len(x);i++{
		rt = append(rt,byte2bit(x[i])...);
	}
	return rt;
}


func byte2bit (x byte) []bool{
	y := int(x);
	var rt [8]bool;
	for i:=7;i>=0;i--{
		if y%2==1 {rt[i]=true; };
		y=y/2;
	}
	return rt[0:];
}

func int2byte (x []int)[]byte{
	var rt []byte;
	for i:=0;i<len(x);i++{
		rt = append(rt,byte(x[i]));
	}
	return rt;
}

func Bits2Byte (x []int) []byte{
	var rt []byte;
	for i:=0;i<len(x);i+=8{
		rt=append(rt,bits2byte(x[i:i+8]));
	}
	return rt;
}

func bits2byte (x []int) byte{
	if (len(x)!=8){fmt.Println("exception: length of input to bit2byte() is not 8");}
	rt := 0;
	for i:=0;i<8;i++{
		rt=rt*2;
		if int2bool(x[i]) {rt=rt+1}
	}
	return byte(rt);
}


func bit2byte (x []bool) byte{
	if (len(x)!=8){fmt.Println("exception: length of input to bit2byte() is not 8");}
	rt := 0;
	for i:=0;i<8;i++{
		rt=rt*2;
		if x[i] {rt=rt+1}
	}
	return byte(rt);
}

func Bit2Byte (x []bool) []byte{
	var rt []byte;
	for i:=0;i<len(x);i+=8{
		rt=append(rt,bit2byte(x[i:i+8]));
	}
	return rt;
}

func int2bool (x int) bool{
	if x==0 {return false}
	return true;
}

func zerobyte( l int)[]byte{
	var rt []byte;
	for i:=0;i<l;i++{
		rt = append(rt,byte(0));
	}
	return rt;
}

func byteAnd (x []byte,y []byte) []byte{
	var rt []byte;
	if (len(x)!=len(y)) {fmt.Println("exception: length of input to byteAnd() not equal ") ;fmt.Println(len(x));fmt.Println(len(y))};
	for i:=0;i <len(x);i++{
		rt = append(rt,x[i]&y[i]);
	}
	return rt;
}

func byteXor (x []byte,y []byte) []byte{
	var rt []byte;
	if (len(x)!=len(y)) {fmt.Println("exception: length of input to byteXor() not equal ") ;fmt.Println(len(x));fmt.Println(len(y))};
	for i:=0;i<len(x);i++{
		rt = append(rt,x[i]^y[i]);
	}
	return rt;
}

func byte2Int (x []byte) *big.Int{
	rt := big.NewInt(0);
	return rt.SetBytes(x);
}


func str2Int(str string, l int ) *big.Int{
	by := []byte(str);
	for len(by) < (l/8){
		by=append(by,byte(0))
	}
	rt := big.NewInt(0);
	return rt.SetBytes(by);
}

func Int2str(x *big.Int) string{
	return string(x.Bytes());
}

func appenleftbyte (x []byte, l int) []byte{
	var rt []byte;
	for len(rt)<(l/8-len(x)){
		rt =append(rt,byte(0));
	}
	return append(rt,x...);
}