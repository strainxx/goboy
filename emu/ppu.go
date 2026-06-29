package emu

import "sort"

var palette = []uint32{
	0xFFFFFF,
	0xAAAAAA,
	0x555555,
	0x000000,
}

// var palette = []uint32{
// 	0x9bbc0f,
// 	0x77a112,
// 	0x306230,
// 	0x0f380f,
// } // green

// var palette = []uint32{
// 	0xFFFFFF,
// 	0xff8f84,
// 	0x943a3a,
// 	0x000000,
// } //red

type Sprite struct {
	Y        int
	X        int
	TileID   byte
	Priority bool
	YFlip    bool
	XFlip    bool
	Palette  int
}

type PPU struct {
	Mode  int
	Cycle int

	Scanline int

	VRAM [8192]byte
	OAM  [160]byte

	Cpu *CPU

	Framebuffer [144][160]uint32

	FrameReady bool

	LCDC byte // $FF40 - LCD Control
	STAT byte // $FF41 - STAT
	SCY  byte // $FF42 - Scroll Y
	SCX  byte // $FF43 - Scroll X
	BGP  byte // $FF47 - BG Palette
	OBP0 byte // $FF48 - Object Palette 0
	OBP1 byte // $FF49 - Object Palette 1

	LYC byte // $FF45

	WY byte // $FF4A Window Y position
	WX byte // $FF4B Window X position plus 7

	WindowLine int

	PreviousSTATSignal bool
}

func (p *PPU) CheckSTATInterrupt() {

	lycEqualsLy := p.Scanline == int(p.LYC)
	if lycEqualsLy {
		p.STAT |= 0x04
	} else {
		p.STAT &= 0xFB
	}

	lycInterrupt := (p.STAT&0x40) != 0 && lycEqualsLy
	mode2Interrupt := (p.STAT&0x20) != 0 && p.Mode == 2
	mode1Interrupt := (p.STAT&0x10) != 0 && p.Mode == 1
	mode0Interrupt := (p.STAT&0x08) != 0 && p.Mode == 0

	currentSignal := lycInterrupt || mode2Interrupt || mode1Interrupt || mode0Interrupt

	if currentSignal && !p.PreviousSTATSignal {
		p.Cpu.IF |= 0x02
	}

	p.PreviousSTATSignal = currentSignal
}

func (p *PPU) Step(cycles int) {
	lcdEnabled := (p.LCDC & 0x80) != 0
	if !lcdEnabled {
		p.Cycle = 0
		p.Scanline = 0
		p.Mode = 0
		p.WindowLine = 0

		p.STAT = (p.STAT & 0xFC)
		return
	}
	p.Cycle += cycles

	switch p.Mode {
	case 2: // OAM
		if p.Cycle >= 80 {
			p.Cycle -= 80
			p.Mode = 3
		}

	case 3: // pixel transfer
		if p.Cycle >= 172 {
			p.renderScanline()
			p.Cycle -= 172
			p.Mode = 0
		}

	case 0: // hblank
		if p.Cycle >= 204 {
			p.Cycle -= 204
			p.Scanline++

			if p.Scanline == 144 {
				p.Cpu.IF |= 0x01
				p.Mode = 1 // VBlank
				p.FrameReady = true
				p.WindowLine = 0
			} else {
				p.Mode = 2
			}
		}

	case 1: // vblank
		if p.Cycle >= 456 {
			p.Cycle -= 456
			p.Scanline++

			if p.Scanline > 153 {
				p.Scanline = 0
				p.Mode = 2
			}
		}
	}
	p.STAT = (p.STAT & 0xFC) | uint8(p.Mode&0x03)

	p.CheckSTATInterrupt()
}

func (p *PPU) renderScanline() {
	var bgMapBase uint16 = 0x1800
	if (p.LCDC & 0x08) != 0 {
		bgMapBase = 0x1C00
	}

	var winMapBase uint16 = 0x1800
	if (p.LCDC & 0x40) != 0 {
		winMapBase = 0x1C00
	}

	y := p.Scanline

	bgWinEnabled := (p.LCDC & 0x01) != 0

	windowEnabled := (p.LCDC&0x20) != 0 && y >= int(p.WY) && bgWinEnabled
	windowDrawnOnThisLine := false

	for x := 0; x < 160; x++ {

		var mapBase uint16
		var tileX, tileY int
		var pixelX, pixelY int

		isWindowPixel := windowEnabled && x >= (int(p.WX)-7)
		if !bgWinEnabled {
			pixelColor := palette[0]
			p.Framebuffer[y][x] = pixelColor
			continue
		}

		if isWindowPixel {
			windowDrawnOnThisLine = true
			mapBase = winMapBase
			pixelX = x - (int(p.WX) - 7)
			pixelY = p.WindowLine
		} else {
			mapBase = bgMapBase
			pixelX = (x + int(p.SCX)) % 256
			pixelY = (y + int(p.SCY)) % 256
		}

		tileX = pixelX / 8
		tileY = pixelY / 8

		mapAddr := mapBase + uint16(tileY*32) + uint16(tileX)
		tileIndex := p.VRAM[mapAddr]

		color := p.getTilePixel(int(tileIndex), pixelX%8, pixelY%8)

		p.Framebuffer[y][x] = color
	}

	if windowDrawnOnThisLine {
		p.WindowLine++
	}

	// oam scan
	spriteHeight := 8
	if (p.LCDC & 0x04) != 0 {
		spriteHeight = 16
	}
	var scanlineSprites []Sprite

	for i := 0; i < 40; i++ {
		baseAddr := i * 4
		y := int(p.OAM[baseAddr]) - 16
		x := int(p.OAM[baseAddr+1]) - 8
		tileID := p.OAM[baseAddr+2]
		flags := p.OAM[baseAddr+3]

		if p.Scanline >= y && p.Scanline < y+spriteHeight {
			scanlineSprites = append(scanlineSprites, Sprite{
				Y:        y,
				X:        x,
				TileID:   tileID,
				Priority: (flags & 0x80) != 0,
				YFlip:    (flags & 0x40) != 0,
				XFlip:    (flags & 0x20) != 0,
				Palette:  int((flags >> 4) & 1),
			})
		}

		if len(scanlineSprites) == 10 {
			break
		}
	}

	sort.SliceStable(scanlineSprites, func(i, j int) bool {
		if scanlineSprites[i].X != scanlineSprites[j].X {
			return scanlineSprites[i].X > scanlineSprites[j].X
		}
		return i > j
	})

	// sprites render
	if (p.LCDC & 0x02) != 0 {
		for _, sprite := range scanlineSprites {
			for tileX := 0; tileX < 8; tileX++ {
				pixelX := sprite.X + tileX
				if pixelX < 0 || pixelX >= 160 {
					continue
				}

				line := p.Scanline - sprite.Y
				if sprite.YFlip {
					line = spriteHeight - 1 - line
				}

				tileID := sprite.TileID
				if spriteHeight == 16 {
					tileID &= 0xFE
					if line >= 8 {
						tileID |= 0x01
						line -= 8
					}
				}

				bitX := tileX
				if sprite.XFlip {
					bitX = 7 - bitX
				}

				tileAddr := uint16(tileID)*16 + uint16(line*2)
				lo := p.VRAM[tileAddr]
				hi := p.VRAM[tileAddr+1]

				bit := 7 - bitX
				colorNum := ((hi>>bit)&1)<<1 | ((lo >> bit) & 1)

				if colorNum == 0 {
					continue
				}

				palReg := p.OBP0
				if sprite.Palette == 1 {
					palReg = p.OBP1
				}

				paletteColor := (palReg >> (colorNum * 2)) & 0x03
				if sprite.Priority {
					if p.Framebuffer[p.Scanline][pixelX] != palette[0] {
						continue
					}
				}

				p.Framebuffer[p.Scanline][pixelX] = palette[paletteColor]
			}
		}
	}
}

func (p *PPU) getTilePixel(tile int, x, y int) uint32 {
	var tileAddr uint16
	if (p.LCDC & 0x10) != 0 {
		tileAddr = uint16(uint8(tile))*16 + uint16(y*2)
	} else {
		tileAddr = uint16(int32(0x1000) + int32(int8(tile))*16 + int32(y*2))
	}

	lo := p.VRAM[tileAddr]
	hi := p.VRAM[tileAddr+1]

	bit := 7 - x

	color := ((hi>>bit)&1)<<1 | ((lo >> bit) & 1)
	paletteColor := (p.BGP >> (color * 2)) & 0x03

	return palette[paletteColor]
}
