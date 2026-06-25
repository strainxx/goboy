package ioregs

import (
	"fmt"
	"goboy/emu"
	"net"
)

type SerialDataTransfer struct {
	SerialData    byte
	SerialControl byte

	ShouldConnect bool // neither host nor client if false
	IsHost        bool
	Addr          string

	Listener   net.Listener
	Connection net.Conn

	InChan  chan byte
	OutChan chan byte
}

func (r *SerialDataTransfer) Init() {
	if !r.ShouldConnect {
		return
	}

	if r.IsHost {
		fmt.Printf("Hosting SDT at %s\n", r.Addr)
		listener, err := net.Listen("tcp", r.Addr)
		if err != nil {
			fmt.Println("Failed to listen on addr:", err)
			panic("Failed to host")
		}
		r.Listener = listener
		fmt.Printf("Waiting for second player connection\n")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			panic("Failed to accept connection")
		}
		r.Connection = conn
		fmt.Printf("%s connected!\n", conn.RemoteAddr().String())

		r.SerialControl = r.SerialControl | 0x01
	} else {
		r.InChan = make(chan byte, 1)
		r.OutChan = make(chan byte, 1)
		fmt.Printf("Connecting to %s\n", r.Addr)
		conn, err := net.Dial("tcp", r.Addr)
		if err != nil {
			fmt.Println("Failed to connect to server:", err)
			panic("Failed to connect to server")
		}
		r.Connection = conn
		fmt.Printf("Connected!\n")
		r.SerialControl = r.SerialControl &^ (1 << 0)
	}
}

func (r *SerialDataTransfer) Update(cpu *emu.CPU) {
	if !r.ShouldConnect || r.IsHost {
		return
	}
	select {
	case incomingByte := <-r.InChan:
		oldData := r.SerialData

		r.SerialData = incomingByte
		r.SerialControl &= 0x7F
		cpu.IF |= 0x08

		r.OutChan <- oldData

	default:
		// dont freeze
	}
}

func (r *SerialDataTransfer) ClientLoop(cpu *emu.CPU) {
	for {
		buf := make([]byte, 1)
		_, err := r.Connection.Read(buf)
		if err != nil {
			fmt.Println("[SDT] Failed to read byte!")
			return
		}
		fmt.Println("[SDT] Cl recv:", buf[0])

		r.InChan <- buf[0]
		outByte := <-r.OutChan

		r.Connection.Write([]byte{outByte})
	}
}

func (r *SerialDataTransfer) OnRead(cpu *emu.CPU, addr uint16) byte {
	switch addr {
	case 0xFF01:
		return r.SerialData
	case 0xFF02:
		var sc byte

		if r.ShouldConnect {
			if r.IsHost {
				sc = r.SerialControl | 0x01
			} else {
				sc = r.SerialControl &^ 0x01
			}
		} else {
			sc = r.SerialControl
		}

		fmt.Printf(
			"[%s] SC read=%02X\n",
			map[bool]string{true: "HOST", false: "CLIENT"}[r.IsHost],
			sc,
		)

		return sc
	}
	return 0x00
}
func (r *SerialDataTransfer) OnWrite(cpu *emu.CPU, addr uint16, val byte) {
	switch addr {
	case 0xFF01:
		r.SerialData = val
	case 0xFF02:
		r.SerialControl = val
		fmt.Printf(
			"[%s] SC write=%02X read_would_be=%02X\n",
			map[bool]string{true: "HOST", false: "CLIENT"}[r.IsHost],
			val,
			func() byte {
				if r.IsHost {
					return val | 0x01
				}
				return val &^ 0x01
			}(),
		)
		if !r.ShouldConnect {
			if (val & 0x80) != 0 {
				fmt.Printf("%c", r.SerialData)
				r.SerialData = 0xFF
				r.SerialControl &= 0x7F
				// cpu.IF |= 0x08
			}
			return
		} else {
			if r.IsHost && (val&0x80) != 0 {
				fmt.Println("[SDT] Master send:", r.SerialData)
				_, err := r.Connection.Write([]byte{r.SerialData})
				if err != nil {
					fmt.Println("[SDT] Failed to send byte!")
					return
				}

				buf := make([]byte, 1)
				_, err = r.Connection.Read(buf)
				if err != nil {
					fmt.Println("[SDT] Failed to read byte!")
					return
				}
				fmt.Println("[SDT] Master recv:", buf)
				r.SerialData = buf[0]
				r.SerialControl &= 0x7F
				cpu.IF |= 0x08
			}
		}
	}
}
