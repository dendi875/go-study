package main

/**
接口类型

Writer 接口类型它定义了所有可以写入字节的类型的抽象，包括文件、内存缓冲区、网络连接
 */
type Writer interface {
	Write(p []byte) (n int, err error)
}

/**
Reader 接口抽象了所有可以读字节的类型，可以读文件、读内存缓冲区、读网络连接
 */
type Reader interface {
	Read(p []byte) (n int, err error)
}

/**
Closer 接口，抽象了所有可以关闭的类型，比如关闭文件，关闭网络练连接
 */
type Closer interface {
	Close() error
}


/**
我们可以通过已有的接口组合得到新的接口，称为嵌入式接口，这样我们就可以直接使用一个接口，而不需要写出这个接口所有的方法
 */
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

/**
把已有的接口和新增的方法混在一起来使用
 */
type ReadWriter2 interface {
	Reader
	Write(p []byte) (n int, err error)
}
