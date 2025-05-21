package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Point struct{ X, Y int }

func bezier(p0, p1, p2 Point, t float64) Point {
	u := 1 - t
	u2 := u * u
	t2 := t * t
	return Point{
		X: int(u2*float64(p0.X) + 2*u*t*float64(p1.X) + t2*float64(p2.X)),
		Y: int(u2*float64(p0.Y) + 2*u*t*float64(p1.Y) + t2*float64(p2.Y)),
	}
}

func main() {
	fmt.Println("Simulação de ramo de pinheiro iniciado")

	const nPontos = 300
	const mObjetivo = (nPontos * nPontos) / 10

	matriz := make([][]bool, nPontos)
	for i := range matriz {
		matriz[i] = make([]bool, nPontos)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	mid := nPontos / 2
	seeds := make([]Point, 0)
	p0 := Point{mid, 0}
	p2 := Point{mid, nPontos - 1}
	for k := 0; k <= 100; k++ {
		t := float64(k) / 100
		p1 := Point{mid + rng.Intn(21) - 10, nPontos / 2}
		pt := bezier(p0, p1, p2, t)
		seeds = append(seeds, pt)
		matriz[pt.Y][pt.X] = true
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	pSalvos := len(seeds)
	var historico [][2]int
	historico = append(historico, func() [][2]int {
		out := make([][2]int, len(seeds))
		for i, pt := range seeds {
			out[i] = [2]int{pt.Y, pt.X}
		}
		return out
	}()...)

	clamp := func(x int) int {
		if x < 0 {
			return 0
		}
		if x >= nPontos {
			return nPontos - 1
		}
		return x
	}

	for pSalvos < mObjetivo {
		i := rng.Intn(nPontos)
		j := rng.Intn(nPontos)
		if matriz[i][j] {
			continue
		}
		for {
			d := rng.Intn(4)
			ii := clamp(i + dx[d])
			jj := clamp(j + dy[d])
			if matriz[ii][jj] {
				matriz[i][j] = true
				historico = append(historico, [2]int{i, j})
				pSalvos++
				break
			}
			i, j = ii, jj
		}
	}

	fmt.Println("Simulação concluída, total de pontos:", len(historico))

	outFile, err := os.Create("pontos_pinheiro.csv")
	if err != nil {
		log.Fatalf("erro criando CSV: %v", err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	defer writer.Flush()
	writer.Write([]string{"i", "j"})
	for _, pt := range historico {
		writer.Write([]string{strconv.Itoa(pt[0]), strconv.Itoa(pt[1])})
	}

	p := plot.New()
	if err != nil {
		log.Fatalf("erro criando plot: %v", err)
	}
	p.Title.Text = "Ramo de Pinheiro - DLA"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(historico))
	for k, pt := range historico {
		pts[k].X = float64(pt[1])
		pts[k].Y = float64(pt[0])
	}
	scatter, _ := plotter.NewScatter(pts)
	scatter.GlyphStyle.Radius = vg.Points(0.5)
	scatter.GlyphStyle.Color = color.RGBA{0xFF, 0xE9, 0x00, 255}
	p.Add(scatter)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "pinheiro.png"); err != nil {
		log.Fatalf("erro salvando imagem: %v", err)
	}
	fmt.Println("Gráfico salvo: pinheiro.png")
}
