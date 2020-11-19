package main

import (
	"fmt"
	"math"
)

var eps float64 = 1e-10;

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

func innerProduct(a []bool,b[]bool) bool{
	var c = 0;
	for i:=0;i<len(a);i++{
		if a[i] && b[i]{
			c ++;
		}
	}
	return (c % 2) != 0;

}

func WH(a int,q int) []bool{
	var c []bool;
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

func geneWHCode(k int) [][]bool{
	var q_ = math.Log2(float64(k));
	var q int;
	if math.Abs(q_ - math.Floor(q_)) <= eps{
		q = int(math.Floor(q_));
	} else{
		fmt.Println("exception: k size error");
	}
	var C [][]bool;
	for i:=0;i<k;i++{
		C=append(C,WH(i,q));
	}
	return C;
}