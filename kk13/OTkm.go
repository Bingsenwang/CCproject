package main

import "math/big"
func b2int (x bool) int{
	if x {return 1};
	return 0;
}
//note that l is the length of message and should be divide by 8
//transfer a string into big integer for encryption


func OTkmGenPk (G *Group, k int, choice []int, sk []*big.Int) [][2]*big.Int{
	var rt [][2]*big.Int
	r:= big.NewInt(0);
	for i:=0;i<k;i++{
		var rn [2]*big.Int;
        rn[0]= G.OGen(r.Rand(RANseed,G.q));
        rn[1]= G.OGen(r.Rand(RANseed,G.q));
        rn[choice[i]]= G.Gen(sk[i]);
        rt = append(rt,rn);
	}
	return rt;
}


func OTkmSend (G *Group, k int, m int, pk [][2]*big.Int , message [][2]*big.Int) [][2]*Ct{
	var rt [][2]*Ct;
	for i:=0;i<k;i++{
		var rn [2]*Ct;
		rn[0]=G.Enc(message[i][0],pk[i][0]);
		rn[1]=G.Enc(message[i][1],pk[i][1]);
		rt = append(rt,rn);
	}
 return rt;
}

func OTkmGetMes (G *Group, k int, choice []int, sk []*big.Int, ci [][2]*Ct) []*big.Int{
	var rt []*big.Int;
	for i:=0;i<k;i++{
		rn := G.Dec(ci[i][choice[i]],sk[i]);
		rt = append(rt,rn);
	}
	return rt;
}