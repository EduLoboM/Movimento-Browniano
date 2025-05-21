package main

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	const nPontos = 300
	const nSeeds = 80
	const maxLen = 400
	const branchProb = 0.2
	const jitter = 0.1
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	crack := make([][]bool, nPontos)
	for i := range crack {
		crack[i] = make([]bool, nPontos)
	}

	clamp := func(x int) int {
		if x < 0 {
			return 0
		}
		if x >= nPontos {
			return nPontos - 1
		}
		return x
	}

	dirs0 := make([]float64, nSeeds)
	for k := range dirs0 {
		dirs0[k] = 2 * math.Pi * float64(k) / float64(nSeeds)
	}

	var propagate func(x, y int, angle float64, length int)
	propagate = func(x, y int, angle float64, length int) {
		if length <= 0 {
			return
		}
		for step := 0; step < length; step++ {
			x = clamp(x)
			y = clamp(y)
			crack[y][x] = true

			angle += (rng.Float64()*2 - 1) * jitter
			nx := x + int(math.Round(math.Cos(angle)))
			ny := y + int(math.Round(math.Sin(angle)))

			if ny >= 0 && ny < nPontos && nx >= 0 && nx < nPontos && crack[ny][nx] {
				crack[ny][nx] = true
				return
			}
			x, y = nx, ny

			if rng.Float64() < branchProb {
				go propagate(x, y, angle+math.Pi/2*(rng.Float64()*2-1), length/2)
			}

			if x < 0 || x >= nPontos || y < 0 || y >= nPontos {
				return
			}
		}
	}

	center := nPontos / 2
	for _, theta0 := range dirs0 {
		x0 := center + int(math.Round(math.Cos(theta0)*float64(center)*0.5))
		y0 := center + int(math.Round(math.Sin(theta0)*float64(center)*0.5))
		go propagate(x0, y0, theta0, maxLen)
	}

	time.Sleep(200 * time.Millisecond)

	historico := make(plotter.XYs, 0)
	for i := 0; i < nPontos; i++ {
		for j := 0; j < nPontos; j++ {
			if crack[i][j] {
				historico = append(historico, plotter.XY{X: float64(j), Y: float64(nPontos - i)})
			}
		}
	}

	file, err := os.Create("fissuras.csv")
	if err != nil {
		log.Fatalf("erro criando CSV: %v", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()
	w.Write([]string{"x", "y"})
	for _, pt := range historico {
		w.Write([]string{strconv.Itoa(int(pt.X)), strconv.Itoa(int(pt.Y))})
	}

	p := plot.New()
	if err != nil {
		log.Fatalf("erro plot: %v", err)
	}
	p.Title.Text = "Solo Craquelado - Geométrico"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	scatter, _ := plotter.NewScatter(historico)
	scatter.GlyphStyle.Radius = vg.Points(0.4)
	scatter.GlyphStyle.Color = color.RGBA{R: 12, G: 00, B: 255, A: 255}
	p.Add(scatter)

	p.X.Min, p.X.Max = 0, float64(nPontos)
	p.Y.Min, p.Y.Max = 0, float64(nPontos)

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "craquelado_geom.png"); err != nil {
		log.Fatalf("erro salvando: %v", err)
	}
	fmt.Println("Imagem salva: craquelado_geom.png")
}
