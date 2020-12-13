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
	listen,err := net.Listen("tcp","127.0.0.1:9090")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	defer listen.Close();
	fmt.Println("start listen")
	f1,_ := os.Create("detail_iknp_l1_k128_n2_bool_a.log");
	f2,_ := os.Create("overall_iknp_l1_k128_n2_bool_a.log");
	defer f1.Close();
	defer f2.Close();
	var M,N,L int;

	N = 2;
	L = 1;
	//code := WHCode(K);
	code := RepetitionCode(K);
	for {
		conn,err:=listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}

		for M = 128;M < 8e6;M *= 2 {
		//for M = 8;M < 10;M *= 2 {
			fmt.Println("");
			fmt.Println("N = ", N,"L = ", L, "M =",M);
			f1.Write([]byte(strconv.Itoa(N)+" "+strconv.Itoa(L)+" "+strconv.Itoa(M)+"\n"));
			var t time.Duration;
			for i := 0; i < 20; i++ {
				message := geneMessage(M, N, L);
				//fmt.Println(message);
			//choice := geneChoice(M, N);
				var alice Alice;
			//var bob Bob;
				var ts = time.Now();
				alice.init(&message, N, M, L, K, code);
			//bob.init(&choice, M, L, K, code);
				t_init := time.Since(ts);
				fmt.Println("  Init Time:", t_init);
				ts = time.Now()
				pks := alice.OTkmsendpk()
				//fmt.Println(pks[2][1].X,pks[2][1].Y)
				SendPks(pks,conn,K);
				tmp :=RecieveXs(conn,K,2,920);
				//fmt.Println(tmp);
				alice.OTkmReceive(tmp);
				t_OTkm := time.Since(ts);
				fmt.Println("  OTkm Time:", t_OTkm);
				ts = time.Now()
				sendmes := alice.Send()
				SendXs(bool3d2byte3d(sendmes), conn, N, M, L*8);
				//bob.Receive(alice.Send());
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
}