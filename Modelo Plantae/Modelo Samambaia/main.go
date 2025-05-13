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

func bezier(p0, p1, p2, p3 Point, t float64) Point {
	u := 1 - t
	u2 := u * u
	u3 := u2 * u
	t2 := t * t
	t3 := t2 * t
	return Point{
		X: int(u3*float64(p0.X) + 3*u2*t*float64(p1.X) + 3*u*t2*float64(p2.X) + t3*float64(p3.X)),
		Y: int(u3*float64(p0.Y) + 3*u2*t*float64(p1.Y) + 3*u*t2*float64(p2.Y) + t3*float64(p3.Y)),
	}
}

func main() {
	fmt.Println("Olá Mundo! Programa iniciado com sucesso")

	const nPontos = 250
	const mObjetivo = (nPontos * nPontos) / 10

	matriz := make([][]int, nPontos)
	for i := range matriz {
		matriz[i] = make([]int, nPontos)
	}
	var historico [][2]int
	energies := make([]int, 0, mObjetivo)
	pSalvos := 0

	mid := nPontos / 2
	ctrl := []Point{{mid, 0}, {mid + 60, nPontos / 3}, {mid - 60, 2 * nPontos / 3}, {mid, nPontos - 1}}
	nSeeds := 200
	for k := 0; k <= nSeeds; k++ {
		t := float64(k) / float64(nSeeds)
		pt := bezier(ctrl[0], ctrl[1], ctrl[2], ctrl[3], t)

		i, j := pt.Y, pt.X
		if i < 0 {
			i = 0
		} else if i >= nPontos {
			i = nPontos - 1
		}
		if j < 0 {
			j = 0
		} else if j >= nPontos {
			j = nPontos - 1
		}

		if matriz[i][j] == 0 {
			matriz[i][j] = 1
			historico = append(historico, [2]int{i, j})
			energies = append(energies, 0)
			pSalvos++
		}
	}
	fmt.Println("Seeds em formato de folha gerados:", pSalvos)

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}
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
		stepCount := 0
		for {
			d := rng.Intn(4)
			ii := clamp(i + dx[d])
			jj := clamp(j + dy[d])
			stepCount++
			if matriz[ii][jj] == 1 {
				matriz[i][j] = 1
				historico = append(historico, [2]int{i, j})
				energies = append(energies, stepCount)
				pSalvos++
				break
			}
			i, j = ii, jj
		}
	}
	fmt.Println("Simulação feita, total de pontos agregados:", len(historico))

	outFile, err := os.Create("pontos.csv")
	if err != nil {
		log.Fatalf("não foi possível criar CSV: %v", err)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	defer writer.Flush()
	writer.Write([]string{"i", "j", "energy"})
	for idx, pt := range historico {
		writer.Write([]string{strconv.Itoa(pt[0]), strconv.Itoa(pt[1]), strconv.Itoa(energies[idx])})
	}
	fmt.Println("CSV criado com energia")

	p := plot.New()
	p.Title.Text = "Movimento Browniano em Nervura Spline De Folha De Samambaia"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	pts := make(plotter.XYs, len(historico))
	for k, pt := range historico {
		pts[k].X = float64(pt[1])
		pts[k].Y = float64(pt[0])
	}
	scatter, _ := plotter.NewScatter(pts)
	scatter.GlyphStyle.Radius = vg.Points(0.5)
	scatter.GlyphStyle.Color = color.RGBA{198, 57, 125, 255}
	p.Add(scatter)
	p.Save(6*vg.Inch, 6*vg.Inch, "pontos.png")

	pe := plot.New()
	pe.Title.Text = "Energia (steps) até convergência"
	pe.X.Label.Text = "Índice do Ponto"
	pe.Y.Label.Text = "Energia (número de passos)"
	enPts := make(plotter.XYs, len(energies))
	for i, e := range energies {
		enPts[i].X = float64(i)
		enPts[i].Y = float64(e)
	}
	line, _ := plotter.NewLine(enPts)
	pe.Add(line)
	pe.Save(6*vg.Inch, 4*vg.Inch, "energia.png")

	fmt.Println("Gráficos gerados: pontos.png e energia.png")
}
