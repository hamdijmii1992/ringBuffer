package ringBuffer

import "testing"


func TestPutOverflow(t *testing.T) {

	rb   :=   GetInstance(5, false)
	data := []byte{1,2, 3,5,4}
	err := rb.Put(data)
	if err != nil {
		t.Error("Expected buffer Overflow")
	}
	rb   =   GetInstance(5, true)
	data = []byte{1,2, 3,5,4,7}
	err  = rb.Put(data)
	if err != nil {
		t.Error("Expected buffer Overflow")
	}
}

func TestGet(t *testing.T) {
	 rb   :=   GetInstance(5, false)
	 data := []byte{1,2, 3,5,4}
	 err := rb.Put(data)
	 if err != nil {
		t.Error("Expected buffer Overflow")
	 }
	 buffer := rb.Get(2)
	 if buffer[0] != 1 && buffer[1] != 0 {
		t.Error("Expected buffer data [1,2] get ", buffer)
	 }
}

func TestPeek(t *testing.T) {
	 rb   :=   GetInstance(5, false)
	 data := []byte{1,2, 3,5,4}
	 err := rb.Put(data)
	 if err != nil {
		t.Error("Expected buffer Overflow")
	 }
	 rb.Peek(3)
	 count := rb.Count()
	 if count != 5{
	 	t.Error("Expected Count data 5 get ", count)
	 }
}