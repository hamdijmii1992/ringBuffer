
A good explanation of ring buffer can be found in this wiki page :https://en.wikipedia.org/wiki/Circular_buffer

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

This ring buffer allows user to either override data or throw a "Buffer Overflow" error when trying to write in a place that are 
not read

To be done: enable Ring buffer to accept any type of data other than bytes
