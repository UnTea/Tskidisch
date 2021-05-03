package loader

import (
	"bufio"
	"bytes"
	"github.com/UnTea/Tskidisch/linmath"
	"github.com/UnTea/Tskidisch/render"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Header struct {
	Width int
	Height int
}

func LoadHDR(path string) render.Image {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, _ :=  io.ReadAll(file)

	reader := bytes.NewReader(data)
	header, position := parseHeader(reader)

	_, _ =  reader.Seek(position, io.SeekStart)

	image := render.NewImage(header.Width, header.Height)

	for y:= 0; y < image.Height; y++ {
		unpackRLEScanline(y, reader, &image)
	}

	return image
}

func parseHeader(reader *bytes.Reader) (Header, int64) {
	scanner := bufio.NewScanner(reader)
	var pos int64 = 0
	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		pos += int64(advance)
		return
	}
	scanner.Split(scanLines)
	scanner.Scan()
	signature := scanner.Text()

	if signature != "#?RADIANCE" {
		log.Fatal("Is not a .hdr file")
	}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		if !strings.HasPrefix(text, "FORMAT") {
			continue
		}

		parts := strings.Split(text, "=")

		if parts[1] != "32-bit_rle_rgbe" {
			log.Fatal("Does not match 32-bit_rle_rgbe format")
		}
	}

	scanner.Scan()
	resolution := strings.Split(scanner.Text(), " ")

	height, err := strconv.Atoi(resolution[1])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(resolution[3])
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Header{
		Width: width,
		Height: height,
	}, pos
}

func unpackRLEScanline(y int, reader *bytes.Reader, image *render.Image) {
	red := make([]byte, image.Width)
	green := make([]byte, image.Width)
	blue := make([]byte, image.Width)
	exp := make([]byte, image.Width)

	newRLEIndicator := readU16(reader)
	if newRLEIndicator != 0x0202 {
		log.Fatal("hdr::parse(): only New RLE HDRs are supported")
	}

	scanlineWidth := readU16(reader)
	if int(scanlineWidth) != image.Width {
		log.Fatal("hdr::parse(): bad scanline width")
	}

	for i := 0; i < 4; i++ {
	    var x = 0
		color := [][]byte{red, green, blue, exp}

	    for x < image.Width {
	    	count := readU8(reader)

	    	if count > 128 {
	    		count := count & 0X7F
	    		value := readU8(reader)

	    		for j := 0; uint8(j) < count; j++ {
	    			color[i][x] = value
	    			x += 1
				}
			} else {
				for j := 0; uint8(j) < count; j++ {
					color[i][x] = readU8(reader)
					x += 1
				}
			}
		}
	}

	for x := 0; x < image.Width; x++ {
		color := decodeRGBE(red[x], green[x], blue[x], exp[x])
		image.SetPixel(x, y, color)
	}
}

func decodeRGBE(r, g, b, e byte) linmath.Vector {
	diff := 128.0 + 8.0
	exp := math.Pow(2.0, float64(e) - diff)
	rDecoded := float64(r) * exp
	gDecoded := float64(g) * exp
	bDecoded := float64(b) * exp

	return linmath.Vector{
		X: rDecoded,
		Y: gDecoded,
		Z: bDecoded,
	}
}

func readU16(reader *bytes.Reader) uint16 {
	a, err := reader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}

	b, err := reader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}

	return (uint16(a) << 8) | uint16(b)
}

func readU8(reader *bytes.Reader) uint8 {
	a, err := reader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}

	return a
}
