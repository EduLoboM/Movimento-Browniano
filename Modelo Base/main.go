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

func main() {

	fmt.Println("Olá Mundo! Programa iniciado com sucesso")

	const nPontos int = 250
	const mObjetivo int = (nPontos * nPontos) / 10
	ci, cj := nPontos/2, nPontos/2
	var pSalvos int
	matriz := make([][]int, nPontos)
	for i := range matriz {
		matriz[i] = make([]int, nPontos)
	}
	matriz[ci][cj] = 1
	pSalvos++

	var historico [][2]int
	historico = append(historico, [2]int{ci, cj})

	dx := []int{-1, +1, 0, 0}
	dy := []int{0, 0, +1, -1}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	clamp := func(x int) int {
		if x < 0 {
			return 0
		}
		if x >= nPontos {
			return nPontos - 1
		}
		return x
	}

	fmt.Println("Simulação iniciada e variáveis validadas")

	for pSalvos < mObjetivo {
		i, j := rng.Intn(nPontos), rng.Intn(nPontos)

		if matriz[i][j] == 1 {
			continue
		}

		for {
			dir := rng.Intn(4)
			ni := clamp(i + dx[dir])
			nj := clamp(j + dy[dir])

			if matriz[ni][nj] == 1 {
				matriz[i][j] = 1
				pSalvos++
				historico = append(historico, [2]int{i, j})
				break
			}
			i, j = ni, nj
		}
	}

	fmt.Println("Simulação feita")

	outFile, err := os.Create("pontos.csv")

	if err != nil {
		log.Fatalf("não foi possível criar CSV: %v", err)
	}

	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	writer.Write([]string{"i", "j"})
	for _, pt := range historico {
		writer.Write([]string{strconv.Itoa(pt[0]), strconv.Itoa(pt[1])})
	}

	fmt.Println("CSV criado")

	p := plot.New()
	if err != nil {
		log.Fatalf("falha ao criar plot: %v", err)
	}
	p.Title.Text = "Movimento Browniano"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := make(plotter.XYs, len(historico))
	for k, pt := range historico {
		pts[k].X = float64(pt[0])
		pts[k].Y = float64(pt[1])
	}
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatalf("falha ao criar scatter: %v", err)
	}

	scatter.GlyphStyle.Radius = vg.Points(0.5)
	cor := color.RGBA{R: 0, G: 164, B: 153, A: 255}
	scatter.GlyphStyle.Color = cor

	p.Add(scatter)

	if err := p.Save(6*vg.Inch, 6*vg.Inch, "pontos.png"); err != nil {
		log.Fatalf("falha ao salvar gráfico: %v", err)
	}
	fmt.Println("Gráfico salvo em pontos.png")
}
