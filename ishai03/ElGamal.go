package main

import (
	"math/big"
)

func (G *Group) Gen(x *big.Int) *big.Int{ //given private key, generate a publick key
	rt := big.NewInt(0);
	rt=rt.Exp(G.g,x,G.p);
	return rt;
}


func (G *Group) OGen(x *big.Int) *big.Int{
	rt:=big.NewInt(0);
	rt=rt.Exp(x,big.NewInt(2),G.p);
	return rt;
}

type Ct struct {
	 c1,c2 *big.Int;
}

func (c *Ct)init (a, b *big.Int){
	c.c1 =a;
	c.c2 =b;
}

func (G *Group) Enc (m *big.Int, pk *big.Int) *Ct{
	var rt Ct;
	c1 := big.NewInt(0);
	c2 := big.NewInt(0);
	r:= big.NewInt(0);
	//ran := rand.New(rand.NewSource(time.Now().UnixNano()))


	r= r.Rand(RANseed,G.q);
    c1 = c1.Exp(G.g,r,G.p);
    c1 = c1.Mod(c1,G.p);
    c2 = c2.Exp(pk,r,G.p);
    c2 = c2.Mul(m,c2);
	c2 = c2.Mod(c2,G.p);
    rt.init(c1,c2);
    return &rt;
}

func (G * Group) Dec (c *Ct, sk *big.Int) *big.Int{
	a:= c.c1;
	b:= c.c2;
	a=a.Exp(a,sk,G.p);
	a=a.ModInverse(a,G.p);


	rt := big.NewInt(0);
	rt =rt.Mul(b,a);
    rt = rt.Mod(rt,G.p);
	return  rt;
}

//randomly generate k private key
func (G *Group) Gensk (k int) []*big.Int{
	var rt []*big.Int;
	for i:=0;i<k;i++{
		rn := big.NewInt(0);
		rt=append(rt,rn.Rand(RANseed,G.q));
	}
	return rt;
}














