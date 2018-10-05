The ring buffer allows user to either override data that were not read or throw a "Buffer Overflow" error when trying to override a data that are not already processed.

Install

git clone https://github.com/hamdijmii1992/ringBuffer

run tests:

cd ringBuffer
go test

Example:

func main() {
		
	rb := ringBuffer.GetInstance(5, true)
	data := []byte{1, 2, 3, 5, 4, 7}

	err := rb.Put(data)
	if err != nil {
		fmt.Println(err)
	}
	
	buf := rb.Get(6)
	fmt.Println(buf)
	data = []byte{0,0}
	rb.Put(data)
	buf = rb.Peek(3)
	fmt.Println(buf)
	data = []byte{1,1}
}

Youtube : https://www.youtube.com/watch?v=GbBrp6K7IvM

To be done: enable Ring buffer to accept any type of data other than bytes
