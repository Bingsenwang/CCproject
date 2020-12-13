package main

import (
	"crypto/ecdsa"
)

func OTkmGenPk(k int, choice []int)([]*ecdsa.PrivateKey, [][2]*ecdsa.PublicKey){
	var sks []*ecdsa.PrivateKey;
	var pks [][2]*ecdsa.PublicKey;
	for i:=0; i<k; i++{
		var rpk [2]*ecdsa.PublicKey;
		sk,pk := Gen()
		rpk[choice[i]] = pk;
		rpk[1-choice[i]] = OGen();

		sks = append(sks,sk);
		pks = append(pks,rpk);
	}
	return sks,pks;
}

func OTkmSend(k int, pk[][2]*ecdsa.PublicKey, message [][2][]byte) [][][]byte{
	var cts [][][]byte;
	for i:=0; i<k; i++{
		var rct [][]byte;
		rct =append(rct,Enc(message[i][0],pk[i][0]));
		rct= append(rct,Enc(message[i][1],pk[i][1]));
		cts = append(cts,rct);
	}
	return cts;
}

func OTkmGetMes(k int, choice []int, sk []*ecdsa.PrivateKey, cts [][][]byte) [][]byte{
	var pts [][]byte;
	for i:=0; i<k; i++{
		pt := Dec(cts[i][choice[i]],sk[i]);
		pts = append(pts,pt);
	}
	return pts;
}

