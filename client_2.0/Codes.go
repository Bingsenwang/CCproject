package main

import (
	"fmt"
	"math"
)

const eps = 1e-10;

func RepetitionCode(k int) [][]byte{
	var C[][]byte;
	var cn0[]int;
	var cn1[]int;
	for i:=0;i<k;i++{
		cn0 = append(cn0,0);
		cn1 = append(cn1,1);
	}
	C = append(append(C,Bits2Byte(cn0)),Bits2Byte(cn1));
	return C;
}

func WH(a int,q int) []int{
	var c []int;
	var a_b = int2bits(a);
	for ;len(a_b) < q;{
		a_b = append(a_b,false);
	}
	for i:=0;i<(1<<q);i++{
		var t_b = int2bits(i);
		for ;len(t_b) < q;{
			t_b = append(t_b,false);
		}
		c=append(c,innerProduct(a_b, t_b));
	}
	return c;
}

func WHCode(k int) [][]byte{
	var q_ = math.Log2(float64(k));
	var q int;
	if math.Abs(q_ - math.Floor(q_)) <= eps{
		q = int(math.Floor(q_));
	} else{
		fmt.Println("exception: k size error");
	}
	var C [][]byte;
	for i:=0;i<k;i++{
		C=append(C,Bits2Byte(WH(i,q)));
	}
	return C;
}