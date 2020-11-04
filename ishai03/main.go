package main

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

var RANseed = rand.New(rand.NewSource(int64(time.Now().Unix())))

// the random oracle, note that l should divide by 8
func H (m int, ls []byte, l int) []byte{
  x := append(ls,byte(m));
 rt := sha256.Sum256(x)
 return rt[0:l/8];
}

func main() {
	m:=8;
 k:=128;



 var A Alice;
 var B Bob;


 G := Group_generator(100);
 y := [][2]string{{"hello","world"},{"hello","world"},{"hello","world"},{"hello","world"},{"hello","world"},{"hello","world"},{"hello","world"},{"hello","world"}};

 A.init(G,y,m,k);
 choice := []int {0,1,0,1,0,1,0,1};
 
 B.init(G,m,k,choice[0:]);
 A.Receive(B.Receive(A.Send1()));
 B.Receive2(A.Send2());
}