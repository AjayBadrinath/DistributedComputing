// Implementation Of Chandy Lamport Algorithm 

/*

For : CS3001- Distributed Computing - Implementation Of Chandy Lamport Algorithm 

author:Ajay Badrinath

class:IoT-A

rollno:21011102020

date: 30.08.2023

*/
/*
Output Format :
Proc k send m to z
....
..
.
..
Total balance [amt[k0],amt[k1].....]
Channel State of Proc k0: map[int]int
Channel State of Proc k0: map[int]int
Channel State of Proc k0: map[int]int

::::
::::
:::
Sum up the [totalbalance[i]+channel state[i] till k] should be equal to initial sum 
*/
package main

/*

Importing  required libraries

*/

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

/*
declaring 3 bank accounts and 
*/


const numAccounts =3
const numProcesses =3
const buf_len=100 /* Declaring a Buffered FIFO Channel */


/*
	As per Question Since there is no go class i  declare struct 
	each process object has 
	1.process id
	2.account for each bank
	3.channel which is a map[int]int
	4.Mutex for safely writing and reading values in a multi threaded scenario to synchronise access
	5.Snapshot array to store local states once chandy lamport is initiated
	6.Marker to keep track of markers
*/


type process struct{

	pid int 
	
	account [numAccounts] int 
	
	channel []*chan Message 
	
	Mutex sync.Mutex
	
	Snapshot []int
	
	ChannelState []map[int]int
	
	Markers []bool
}

/*
Message struct that keeps track of Sender id,recieverid,amount ,marker or not
*/

type Message struct{

	Senderid int

	Receiverid int 

	amt int

	ismarker bool

}


/*

Function to take snapshot of  the local state and the channel state once marker is sent / recieved

*/

func snap(p*process){

	fmt.Println("Taking Snapshot!")

	for i,c:= range p.channel{

		if i==p.pid{
			continue  							// ignoring the same pid 
		}
		
			for{ 							// iteration through channel

				select{						// wait on multiple channel operation by goroutine

				case msg,is_avl:=<-*c: 			// checking if buffer not closed/ empty

				if(!is_avl){
					return
				}

				/*checking for marker message*/

				if msg.ismarker{
					if !p.Markers[msg.Senderid]{
					p.Markers[msg.Senderid]=true
                    msg.ismarker=true
                    *(p.channel[msg.Senderid]) <- msg
					}


				}else{
					/*update local state and channel state with amount*/
					p.Mutex.Lock()
					p.Snapshot[i]+=msg.amt
					p.ChannelState[i][msg.Senderid]+=msg.amt;
					p.Mutex.Unlock()
				}


			default:


				break
				
				}


			}

		
	}

	/*
	print snapshot of process
	*/

	fmt.Printf("Snapshot of Process %d: %v\n", p.pid, p.Snapshot)

    fmt.Printf("Channel States of Process %d: %v\n", p.pid, p.ChannelState)

}

/*

Function to perform transaction to specific process

As per the question the transact function must be able to randomly choose process.But to keep things simple

,I have done the process selection in the main function that has been documented .

*/

func transact(p*process){

	rand.Seed(time.Now().UnixNano());

	for i:=0;i<10;i++{

		recv:=rand.Intn(numProcesses);
		amt:=rand.Intn(100);

		if(p.pid!=recv){

			p.Mutex.Lock();
			fmt.Printf("Process %d sending %d to Process %d\n",p.pid,amt,recv)
			p.account[p.pid]-=amt
			p.account[recv]+=amt
			p.Mutex.Unlock()

			*(p.channel[recv])<-Message{
				Senderid:p.pid,
				amt:amt,
			}

		}

		time.Sleep(time.Millisecond*100);

	}

}

/*

Function to print balances

*/

func printBalance(p[]*process){

	tot:=make([]int,numAccounts)

	for _,i:=range p{

		for j ,bal := range i.account{
			tot[j]+=bal
		}
	}
	for _,p :=range p{

		for i,_:=range p.ChannelState{

			for j:=0;j<numAccounts;j++{

				amt,isacc:=p.ChannelState[i][j]
				if isacc{

					if(tot[i]-amt>0){

					tot[i]-=amt
					}
				}else{

					continue

				}

			}

		}
		
	}

	fmt.Printf("Total Balance : %v\n=========",tot);

}


/*
Function to peek into a channel to see what value exist there currently
*/


func PrintChannel(p[]*process){

	for _,p1 :=range p{

		fmt.Printf("\nChannel State of Proc %d :\n",p1.pid)

		for j,s := range p1.ChannelState{

			fmt.Printf("Channel %d:%v\n",j,s)

		}

	}
	
}




func main(){

	var p int
	var q int 

	fmt.Println("Enter the number of Processes");
	fmt.Scanln(&p);
	fmt.Println("Enter the number of bank account for each processes");
	fmt.Scanln(&q);

	/*Allocating memory for number of processes*/
	proc:=make([]*process,p);

	/*
	Allocation of memory within each element in the array of process struct
	*/

	for i:=0;i<p;i++{

		proc[i]=&process{pid:i,channel:make([]*chan Message,p),Snapshot:make([]int,p),ChannelState:make([]map[int]int,p),Markers:make([]bool,numProcesses)}
		
		for j:=0;j<p;j++{
		
			ch:=make(chan Message,buf_len)
			proc[i].channel[j]=&ch
			proc[i].ChannelState[j]=make(map[int]int)
		
		}
		
		/*

		Filling each bank account with some money
		
		*/
		
		for m:=0 ;m<numAccounts;m++{
			proc[i].account[m]=3000;
		}

	}
	//initiate  transaction to n random processes
	

	for j:=0;j<len(proc)*15;j++{
		
		n:=rand.Intn(numAccounts);// As per question pick random process to transact 
		
		go transact(proc[n])
		
		if (j%numAccounts==0){
		
			time.Sleep(time.Second*4)
			go snap(proc[n])
		
		printBalance(proc)
		PrintChannel(proc)
		
		time.Sleep(time.Second*4)
		
	}
	
}


	printBalance(proc)
	PrintChannel(proc)


}
