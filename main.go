package main
import ("fmt";"sync";"time")
const workerID = "rate-limiter-a2ec47"
func producer(ch chan<- int, count int){for i:=0;i<count;i++{ch<-i};close(ch)}
func consumer(id int, ch <-chan int, wg *sync.WaitGroup){defer wg.Done();for v:=range ch{fmt.Printf("[%s] worker-%d processed: %d\n",workerID,id,v)}}
func main(){fmt.Printf("[%s] Starting pipeline...\n",workerID);ch:=make(chan int,10);var wg sync.WaitGroup;go producer(ch,20);for i:=0;i<3;i++{wg.Add(1);go consumer(i,ch,&wg)};wg.Wait();fmt.Printf("[%s] All done.\n",workerID);_=time.Now()}
