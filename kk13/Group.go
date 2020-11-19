package main

import (
	"math/big"
	"runtime"
)

type Group struct{
	p,q,g *big.Int;
}

func (G *Group) init (x,y,z *big.Int){
	G.p=x;
	G.q=y;
	G.g=z;
}


func Group_generator (base int) *Group{
	c1 := make(chan *big.Int);
	c2 := make(chan *big.Int);
	//use goroutine to save time
	for i:=0;i<runtime.NumCPU();i++{
		go gG(base+i,c1,c2);
	}
	P1 := <-c1;
	Q1 := <-c2;
	var G Group;
	G.init(P1,Q1,find_g(P1,Q1));
	return &G;
}

func Factorial ( x int) *big.Int{
	rt := big.NewInt(1);
	if rt.ProbablyPrime(10) ==true {};
	if x>0 {rt= rt.Mul(big.NewInt(int64(x)),Factorial(x-1))};
	return rt;
}

// start from Fact(i) to find q and p=2q+1 such that both q and p are prime
func gG (i int, c1 chan (*big.Int),c2 chan (*big.Int)) {
	P:=big.NewInt(0);
	Q:=big.NewInt(0);
	base := Factorial (i);
	for i:=0;i<=1000000;i++ {
		q:=big.NewInt(0);;
		q = q.Add(base,big.NewInt(int64(i)));
		if q.ProbablyPrime(10) == false {
			continue;
		}
		p:=big.NewInt(0);
		p = p.Add(p.Mul(q,big.NewInt(int64(2))),big.NewInt(1)); // p=2q+1
		if p.ProbablyPrime(10) == false {
			continue;
		}
		P=p; Q=q;
		break;

	}

	c1<-P;
	c2<-Q;
}

func find_g(p , q *big.Int) *big.Int{
	return big.NewInt(4); //Since both p, q are big and (p-1)/q =2, q is the only prime divide p-1
}
