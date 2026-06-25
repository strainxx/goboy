package main

import (
	"bufio"
	"flag"
	"fmt"
	"goboy/emu"
	"goboy/emu/ioregs"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// const (
// 	screenWidth  = 1000
// 	screenHeight = 500
// )

const (
	screenWidth  = 320
	screenHeight = 288
)

var image *rl.Image
var texture rl.Texture2D

func updateTexture(frame [144][160]uint32) {
	pixels := make([]byte, 160*144*4)

	for y := 0; y < 144; y++ {
		for x := 0; x < 160; x++ {

			v := frame[y][x]

			i := (y*160 + x) * 4

			pixels[i+0] = byte(v >> 16) // R
			pixels[i+1] = byte(v >> 8)  // G
			pixels[i+2] = byte(v)       // B
			pixels[i+3] = 255           // A
		}
	}
	rl.UpdateTexture(texture, pixels)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Application panicked: %v\n", err)
			fmt.Println("Press any key to exit...")

			bufio.NewReader(os.Stdin).ReadBytes('\n')

			os.Exit(1)
		}
	}()
	rl.InitWindow(screenWidth, screenHeight, "goboy")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	image = rl.GenImageColor(160, 144, rl.Black)
	texture = rl.LoadTextureFromImage(image)

	bootromPtr := flag.String("boot_rom", "", "Path to boot rom")
	gbFilePtr := flag.String("gb", "roms/cpu_instrs.gb", "Path to .gb")
	tracePtr := flag.Bool("trace", false, "Print executed instructions")
	logIoPtr := flag.Bool("log-ioregs", false, "Log access to i/o regs")
	gamepadPtr := flag.Bool("use-gamepad", false, "Try to use gamepad if available")
	gamepadNumPtr := int32(*flag.Int("gamepad", 0, "Gamepad number"))
	hostPtr := flag.Bool("host", false, "Host 2 player mode")
	connectPtr := flag.Bool("connect", false, "Connect to 2 player mode")
	addrPtr := flag.String("addr", "localhost:9123", "Address of host")

	flag.Parse()

	rom, err := os.ReadFile(*gbFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ROM size: %d bytes\n", len(rom))
	ppu := &emu.PPU{}
	cpu := &emu.CPU{PPU: ppu}

	cpu.Init(*tracePtr, *logIoPtr)
	err = cpu.Bus.LoadROM(rom, false, *gbFilePtr)
	if err != nil {
		panic(err)
	}

	useSdt := *hostPtr || *connectPtr
	if useSdt {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		if *hostPtr {
			rl.SetWindowTitle("goboy (Host)")
			rl.DrawText("Waiting for second player", 22, screenHeight/2-22, 22, rl.Purple)
		} else {
			rl.SetWindowTitle("goboy (Client)")
			rl.DrawText("Waiting for connection", 0, screenHeight/2-22, 22, rl.Purple)
		}
		rl.EndDrawing()

	}

	sdt := &ioregs.SerialDataTransfer{ShouldConnect: useSdt, Addr: *addrPtr, IsHost: *hostPtr}
	sdt.Init()

	cpu.Bus.RegisterRange(0xFF10, 0xFF3F, &ioregs.Audio{})
	lcd := &ioregs.LCD{}
	cpu.Bus.RegisterRange(0xFF40, 0xFF4B, lcd)
	cpu.Bus.RegisterRange(0xFF47, 0xFF49, &ioregs.Palettes{})
	cpu.Bus.RegisterRange(0xFF42, 0xFF43, &ioregs.Scrolling{})
	// cpu.Bus.RegisterRange(0xFF44, 0xFF45, lcd)
	cpu.Bus.RegisterHook(0xFF50, &ioregs.BootRomMapped{})
	cpu.Bus.RegisterRange(0xFF04, 0xFF07, &ioregs.TimDiv{})
	cpu.Bus.RegisterHook(0xFF0F, &ioregs.InterruptRegs{})
	cpu.Bus.RegisterRange(0xFF01, 0xFF02, sdt)
	cpu.Bus.RegisterRange(0xFF4C, 0xFF4D, &ioregs.KeyRegs{})
	joypad := &ioregs.Joypad{Register: 0xF}
	cpu.Bus.RegisterHook(0xFF00, joypad)
	cpu.Bus.RegisterHook(0xFF46, &ioregs.OAMDMA{})

	if *bootromPtr != "" {
		bootrom, err := os.ReadFile(*bootromPtr)
		if err != nil {
			panic(err)
		}
		cpu.Bus.LoadROM(bootrom, true, *bootromPtr)
		cpu.PC = 0x0
	} else {
		fmt.Printf("Loading without bootrom")
		cpu.Bus.BootRomDisabled = true

		cpu.A = 0x01
		cpu.F = 0xB0
		cpu.B = 0x00
		cpu.C = 0x13
		cpu.D = 0x00
		cpu.E = 0xD8
		cpu.H = 0x01
		cpu.L = 0x4D
		cpu.SP = 0xFFFE
		cpu.PC = 0x100
	}
	defer cpu.Bus.Cart.OnExit()

	fmt.Printf("Title: %s\n", string(cpu.Bus.Memory[0x134:0x144]))

	var useGamepad bool

	if *gamepadPtr {
		if rl.IsGamepadAvailable(gamepadNumPtr) {
			useGamepad = true
		}
	}

	if useSdt && !sdt.IsHost {
		go sdt.ClientLoop(cpu)
	}
	for !rl.WindowShouldClose() {

		if !useGamepad {
			joypad.UpdateAll(cpu,
				rl.IsKeyDown(rl.KeyZ), // A
				rl.IsKeyDown(rl.KeyX), // B
				rl.IsKeyDown(rl.KeyV), // Select
				rl.IsKeyDown(rl.KeyC), // Start
				rl.IsKeyDown(rl.KeyW), // Up
				rl.IsKeyDown(rl.KeyS), // Down
				rl.IsKeyDown(rl.KeyA), // Left
				rl.IsKeyDown(rl.KeyD), // Right
			)
		} else {
			joypad.UpdateAll(cpu,
				rl.IsGamepadButtonDown(gamepadNumPtr, rl.GamepadButtonRightFaceDown),  // A
				rl.IsGamepadButtonDown(gamepadNumPtr, rl.GamepadButtonRightFaceRight), // B
				rl.IsGamepadButtonDown(gamepadNumPtr, rl.GamepadButtonRightFaceLeft),  // Select
				rl.IsGamepadButtonDown(gamepadNumPtr, rl.GamepadButtonRightFaceUp),    // Start
				rl.GetGamepadAxisMovement(gamepadNumPtr, rl.GamepadAxisLeftY) < -0.1,  // Up
				rl.GetGamepadAxisMovement(gamepadNumPtr, rl.GamepadAxisLeftY) > 0.1,   // Down
				rl.GetGamepadAxisMovement(gamepadNumPtr, rl.GamepadAxisLeftX) < -0.1,  // Left
				rl.GetGamepadAxisMovement(gamepadNumPtr, rl.GamepadAxisLeftX) > 0.1,   // Right
			)
		}

		rl.BeginDrawing()
		// rl.DrawText(fmt.Sprintf("Joypad: %#x", joypad.Read()), 10, 200, 20, rl.White)
		rl.ClearBackground(rl.Gray)

		cyclesThisFrame := 0

		for cyclesThisFrame < 70224/4 && !cpu.Panicked {
			cycles := cpu.Step()
			ppu.Step(cycles * 4)

			sdt.Update(cpu)

			cyclesThisFrame += cycles
		}
		if ppu.FrameReady {
			updateTexture(ppu.Framebuffer)
			ppu.FrameReady = false
		}

		sourceRec := rl.NewRectangle(0, 0, 160, 144)
		destRec := rl.NewRectangle(0, 0, 320, 288)

		rl.DrawTexturePro(texture, sourceRec, destRec, rl.NewVector2(0, 0), 0, rl.White)

		if cpu.Panicked {
			rl.DrawText("PANIC", 10, 10, 5, rl.Red)
		}
		// rl.DrawText(fmt.Sprintf("PC: %#x IE: %#x IF: %#x IME: %t Halted: %t", cpu.PC, cpu.IE, cpu.IF, cpu.IME, cpu.Halted), 10, 170, 20, rl.White)
		rl.EndDrawing()
	}
}
