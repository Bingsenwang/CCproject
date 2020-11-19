package main

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

var RANseed = rand.New(rand.NewSource(int64(time.Now().Unix())))

const N = 8;
var S = [8]string{"a","b","c","d","e","f","g","h"};
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
 var y [][N] string;

 for i:=0;i<8;i++{
 	yy := [N]string{"-","-","-","-","-","-","-","-"};
 	y = append(y,yy);
 	for j:=0;j<N;j++{
 		y[i][j] = string(S[j]);
	}
 }
 var WHC = geneWHCode(k);
 //fmt.Print(len(WHC),"\n");
 //fmt.Print(len(WHC[0]),"\n");
 A.init(G,y,m,k,WHC);
 choice := []int {1,2,3,4,5,6,7,0};
 
 B.init(G,m,k,choice[0:],WHC);
 A.Receive(B.Receive(A.Send1()));
 B.Receive2(A.Send2());
}