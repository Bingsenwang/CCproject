package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const K = 128;


func main(){
	address := "127.0.0.1:9090";
	conn,err := net.Dial("tcp",address);
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close();
	f1,_ := os.Create("detail_iknp_l1_k128_n2_bool_b.log");
	f2,_ := os.Create("overall_iknp_l1_k128_n2_bool_b.log");
	defer f1.Close();
	defer f2.Close();
	var M,N,L int;

	N = 2;
	L = 1;
	//code := WHCode(K);
	code := RepetitionCode(K);
	for M = 128;M < 8e6;M *= 2 {
	//for M = 8;M < 10;M *= 2 {
			fmt.Println("");
			fmt.Println("N = ", N,"L = ", L, "M =",M);
			f1.Write([]byte(strconv.Itoa(N)+" "+strconv.Itoa(L)+" "+strconv.Itoa(M)+"\n"));
			var t time.Duration;
			for i := 0; i < 20; i++ {
				//message := geneMessage(M, N, L);
				choice := geneChoice(M, N);
				//fmt.Println(choice);
				//var alice Alice;
				var bob Bob;
				var ts = time.Now();
				//alice.init(&message, N, M, L, K, code);
				bob.init(&choice, N, M, L, K, code);
				t_init := time.Since(ts);
				fmt.Println("  Init Time:", t_init);
				ts = time.Now()
				pks := RecievePks(conn,K);
				//fmt.Println(pks[2][1].X,pks[2][1].Y);
				data := bob.OTkmReceive(pks);
				//fmt.Println(data);
				SendXs(data,conn,K,2,920);
				t_OTkm := time.Since(ts);
				fmt.Println("  OTkm Time:", t_OTkm);
				recievemes := byte3d2bool3d(RecieveXs(conn,N,M,L*8))
				bob.Receive(recievemes);
				//fmt.Println(bob.z);
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