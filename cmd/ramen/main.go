package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	scriptFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer scriptFile.Close()

	var p Program
	if err := p.Parse(scriptFile); err != nil {
		log.Fatal(err)
	}
	if err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type Program struct {
	Ptr  uint
	Heap []int32
	Code []Command
}

type Command uint8

func (p *Program) Parse(r io.Reader) error {
	reader := bufio.NewReader(r)
	var c *Command

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		switch r {
		case '🍜':
			if c == nil {
				var i Command = 0
				c = &i
			} else {
				*c = *c + 1
			}
		default:
			if c != nil {
				p.Code = append(p.Code, *c)
				c = nil
			}
		}
	}
}

func (p *Program) Run() error {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
L:
	for i := 0; i < len(p.Code); i++ {
		switch p.Code[i] {
		case 2:
			p.Ptr += 1
			continue L

		case 3:
			p.Ptr -= 1
			continue L

		}

		if l := uint(len(p.Heap)); l <= p.Ptr {
			add := make([]int32, 1+p.Ptr-l)
			p.Heap = append(p.Heap, add...)
		}

		switch p.Code[i] {
		// instructionはゼロスタートなので実際の🍜の個数とは1ずれる
		case 0:
			p.Heap[p.Ptr] += 1

		case 1:
			p.Heap[p.Ptr] -= 1

		case 4:
			r, _, err := r.ReadRune()
			if err != nil {
				return err
			}
			p.Heap[p.Ptr] = int32(r)

		case 5:
			_, err := w.WriteRune(rune(p.Heap[p.Ptr]))
			if err != nil {
				return err
			}

		case 6:
			if p.Heap[p.Ptr] == 0 {
				for j := i + 1; j < len(p.Code); j++ {
					if p.Code[j] == 7 {
						i = j
						continue L
					}
				}
				return fmt.Errorf("No closing of the block from Line: %d  was found", i)
			}

		case 7:
			if p.Heap[p.Ptr] != 0 {
				for j := i - 1; j >= 0; j-- {
					if p.Code[j] == 6 {
						i = j - 1
						continue L
					}
				}
				return fmt.Errorf("No beginning of the block from Line: %d  was found", i)
			}
		}
	}

	return w.Flush()
}
