// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(
<<<<<<< HEAD
		server.WithHostPorts("127.0.0.1:8080"),
=======
		server.WithHostPorts("0.0.0.0:6770"),
>>>>>>> f3bcb08 (publish完成)
	)

	register(h)
	h.Spin()
<<<<<<< HEAD
}
=======
}
>>>>>>> f3bcb08 (publish完成)
