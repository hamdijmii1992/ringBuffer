package ringBuffer

import (
"fmt"
"time"
)


type RingBuffer struct{
	head int
	tail int
	length int
	count int
	data []byte
	override bool
} 

type RingBufferError struct{
	When time.Time
	What string
}


func init(){
}

func GetInstance(length int, override bool) *RingBuffer {
	 return Init(length, override)
}

func (e RingBufferError) Error() string{
	return fmt.Sprintf("Ring Buffer Error: %v, %v", e.What, e.When)
}


func  Init(length int, override bool)   *RingBuffer{
	 rb := new(RingBuffer)
     rb.head = 0
     rb.tail = 0
     rb.count = 0
     rb.length = length
     rb.data = make([]byte, length)
     rb.override = override
     return rb
     
}
func (rb *RingBuffer) Put(buffer []byte) error  {
      bufferLength := len(buffer)
      availableSpace := rb.length - rb.count
      if !rb.override {
      	if bufferLength > availableSpace {	
      	return RingBufferError {
      		time.Now(), "Buffer Overflow",
      	}
      }
      }
      
      for _, val := range buffer{
      	rb.data[rb.head] = val
      	if rb.override && rb.head == rb.tail && rb.count !=0 {
      		rb.tail ++
      		if rb.tail >= rb.length {
      			rb.tail = 0
      		}
      	}
      	rb.head ++
      	if rb.head >= rb.length{
      		rb.head = 0
      	}
      }
      rb.count = rb.count + bufferLength
      return nil

}

func (rb *RingBuffer) Count() int {
     return rb.count
}

func (rb *RingBuffer) Get(length int) []byte {
	if rb.count == 0 {
		return nil
	}
     //if wanted data is greater than data presented in the ring buffer
	 if length > rb.count {
	 	length = rb.count
	 }
	 var filledBuffer []byte = make([]byte, length) 
	 for i:=0; i<length; i++ {
	 	filledBuffer[i] = rb.data[rb.tail]
	 	rb.tail ++
	 	if rb.tail >= rb.length{
	 		rb.tail = 0
	 	}
	 }
	 rb.count = rb.count - length
	 return  filledBuffer
}

func (rb *RingBuffer) Peek(length int)  []byte {
     if rb.count == 0 {
		return nil
	 }
     tmpTail := rb.tail    
     //if read data is greater than data presented in the ring buffer
	 if length > rb.count {
	 	length = rb.count
	 }
	 var filledBuffer []byte = make([]byte, length) 
	 for i:=0; i<length; i++ {
	 	filledBuffer[i] = rb.data[rb.tail]
	 	tmpTail ++
	 	if tmpTail >= rb.length{
	 		tmpTail = 0
	 	}
	 }
	 
	 return  filledBuffer

}

func  (rb *RingBuffer) Reset() {
	  rb.head = 0
	  rb.tail = 0
	  rb.count = 0
}