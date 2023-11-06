# Chandy-Lamport Algorithm GoLang Implementation

## Problem Statement


  
	1.Take input 'n' from the user. This represents the number of processes in the distributed system.

	2.Each process should be a go object. Assume each process holds 3 bank accounts initialized to some large amount.

	3.Create "Channels" (Go by Example: Channels) for every process. These channels act as incoming communication edge to the process from other processes. For e.g. assume there are 3 processes in the system: p1, p2, p3. When p1 creates its channel, both p2 and p3 can send messages into this channel to communicate with p1. Go channels are FIFO by default.

	4.For every process, make sure to initialize space to store the local snapshot and state of the communication channel.

	5.Assume that the application has no withdrawals and all transactions are only credit and debit.

	6.Before starting the application, print the total amount in the whole system. Let this value be x.

	7.Each process has a function transaction() that randomly does n*15 transactions to accounts in other processes. Write a function to randomly generate these transactions.

	8.Invoke the transaction() function for all processes as separate Goroutines (A Tour of Go). Thus the transaction function for all processes are running concurrently now.

	9.For every process pn, initiate the Chandy Lamport algorithm after every 'nth' transaction it initiates.

	10.After taking the snapshot, the print the total amount of money in the snapshot. It should be equal to x for every snapshot, for all the processes





## General Overview Of Chandy Lamport Snapshot Algorithm

<h2>
  Marker Sending Rule   State :[P->Process {P1:P2:P3}][channel C1:C2:C3]
  
</h2>
<ul>
  
  <li>P sends a marker along Ck and records its state before P sends other messages</li>
  
</ul>
<h2>
  Marker Recieving Rule   State :[P->Process {P1:P2:P3}][channel C1:C2:C3]
  
</h2>
<ul>
  <h4>On Recieving a marker along Ck</h4>
  <li>
    
    if P2 has not recorded its state then:
        P2 record its state&& record state of Ck->[]
    else:
        P2 record its state && then record channel Ck state before P2 recieved marker along Ck
  </li>
  
</ul>


![image](https://github.com/AjayBadrinath/DistributedComputing/assets/92035508/317a5b34-81b2-4a1a-8af5-19da5c8c9648)

### Retrospecting My Implementation Of The Algorithm:

<ul>
  <li>This implementation is indeed not effective  and is a toy version of chandy-lamport algorithm.</li>
  <li> I had created separate threads for EACH transaction which was not supposed to be done  . </li>
  <li>Works well for n<10 but when n is huge then go limits the number of threads created see:https://softwareengineering.stackexchange.com/questions/222642/are-go-langs-goroutine-pools-just-green-threads/222694#222694</li>
</ul>
<h3>Better Approach</h3>
    <ul>
<li>Another approach discussed could be Each process can be considered a thread / and create child thread for bank accounts .</li><li> Place the transferred
 data  into the channel and read from channel asynchronously .</li><li> This eleminates the use of protecting critical sections of data with mutex</li>
 <li>After n snapshots/ transactions done then exit returning resources to parent threads and safely releasing thread.</li>
 </ul>


    
## References:
https://decomposition.al/blog/2019/04/26/an-example-run-of-the-chandy-lamport-snapshot-algorithm/
