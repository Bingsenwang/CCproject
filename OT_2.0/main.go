package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const K = 128;

func main(){
	f1,_ := os.Create("detail_kk13.log");
	f2,_ := os.Create("overall_kk13.log");
	defer f1.Close();
	defer f2.Close();
	var M,N,L int;

	N = 4;
	L = 32;
	code := WHCode(K);
	//code := RepetitionCode(K);
	for M = 64;M < 8*10e6;M *= 2 {

		fmt.Println("");
		fmt.Println("N = ", N,"L = ", L, "M =",M);
		f1.Write([]byte(strconv.Itoa(N)+" "+strconv.Itoa(L)+" "+strconv.Itoa(M)+"\n"));
		var t time.Duration;
		for i := 0; i < 20; i++ {
			message := geneMessage(M, N, L);
			choice := geneChoice(M, N);
			var alice Alice;
			var bob Bob;
			var ts = time.Now();
			alice.init(&message, N, M, L, K, code);
			bob.init(&choice, M, L, K, code);
			t_init := time.Since(ts);
			fmt.Println("  Init Time:", t_init);
			ts = time.Now()
			alice.OTkmReceive(bob.OTkmReceive(alice.OTkmsendpk()));
			t_OTkm := time.Since(ts);
			fmt.Println("  OTkm Time:", t_OTkm);
			bob.Receive(alice.Send());
			t_OTml := time.Since(ts);
			fmt.Println("  OTml Time:", t_OTml);
			t_t :=t_init+t_OTkm+t_OTml;
			fmt.Println(i, "th Total Time:", t_t);
			t += t_init + t_OTkm + t_OTml;
			t_init.String()
			//fmt.Println(message);
			//fmt.Println(choice);
			//fmt.Println(bob.z);
			f1.Write([]byte(t_init.String()+" "+t_OTkm.String()+" "+t_OTml.String()+" "+t_t.String()+"\n"));
		}

		fmt.Println("Average Time:", t/20);
		f1.Write([]byte("\n"));
		f2.Write([]byte(strconv.Itoa(M)+" "+(t/20).String()+"\n"));

	}
}