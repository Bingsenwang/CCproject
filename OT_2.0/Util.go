package main

import (
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	randm "math/rand"
	"time"
)

var curve = elliptic.P256();
var RANseed = randm.New(randm.NewSource(int64(time.Now().Unix())))

func H (m int, ls []byte, l int) []byte{
	x := append(ls,byte(m));
	rt := sha1.Sum(x);
	return rt[0:l/8];
}

func OGen()(*ecdsa.PublicKey){
	k,_ := ecdsa.GenerateKey(curve,rand.Reader);
	return &k.PublicKey;
}

func Gen()(*ecdsa.PrivateKey,*ecdsa.PublicKey){
	k,_ := ecdsa.GenerateKey(curve,rand.Reader);
	return k,&k.PublicKey;
}

func Enc(pt []byte,pk *ecdsa.PublicKey)([]byte){
	pk2 := ecies.ImportECDSAPublic(pk);
	cc,_:= ecies.Encrypt(rand.Reader, pk2, pt, nil, nil);
	return cc;
}

func Dec(y []byte,sk *ecdsa.PrivateKey)([]byte){
	sk2 := ecies.ImportECDSA(sk);
	x,_:= sk2.Decrypt(y,nil,nil);
	return x;
}

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

func bit2byte (x []bool) byte{
	if (len(x)!=8){fmt.Println("exception: length of input to bit2byte() is not 8");}
	rt := 0;
	for i:=0;i<8;i++{
		rt=rt*2;
		if x[i] {rt=rt+1}
	}
	return byte(rt);
}

func appenleftbyte (x []byte, l int) []byte{
	var rt []byte;
	for len(rt)<(l/8-len(x)){
		rt =append(rt,byte(0));
	}
	return append(rt,x...);
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
		if y%2==1 {rt[i]=true;};
		y=y/2;
	}
	return rt[0:];
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

func int2bits(x int) []bool{
	var bs []bool;
	if x==0{
		bs = append(bs,false);
	}
	for ;x!=0;x >>= 1{
		bs = append(bs, x % 2 != 0);
	}
	return bs;
}

func innerProduct(a []bool,b[]bool) int{
	var c = 0;
	for i:=0;i<len(a);i++{
		if a[i] && b[i]{
			c ++;
		}
	}
	return c % 2;

}

func geneMessage(m,n,l int) [][][]byte{
	var M [][][]byte;
	for i:=0;i<m;i++{
		var M_ [][]byte;
		for j:=0;j<n;j++{
			var M__[]byte;
			for k:=0;k<l/8;k++{
				M__ = append(M__,byte(RANseed.Intn(256)));
			}
			M_ = append(M_,M__);
		}
		M = append(M,M_);
	}
	return M;
}
func geneChoice(m,n int) []int{
	var c []int;
	for i:=0;i<m;i++{
		c = append(c,RANseed.Intn(n));
	}
	return c;
}