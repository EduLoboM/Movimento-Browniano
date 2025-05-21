<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-blue.svg" alt="Go">
  <img src="https://img.shields.io/badge/Language-Python-yellow.svg" alt="Python">
  <img src="https://img.shields.io/badge/Particle%20Sim-Python-green.svg" alt="Particle Simulation">
  <img src="https://img.shields.io/badge/Erosion%20Engine-Go-red.svg" alt="Erosion Engine">
</p>

# Simulador de Movimento Browniano 🧮

Repositório contendo sistemas em Go e Python para:

- Simulação de movimento Browniano (simplificado e complexo).
- Modelos paramétricos de vegetação.
- Simulação de processos erosivos em mapas digitais.

---

## ⚙️ Recursos Principais

- **Modelo Base**: implementação em **Go**, responsável pela orquestração das simulações.
- **Modelos de Vegetação**:
  - **Goiabeira** e **Vitória-régia** em **Python**.
  - **Samambaia** e **Ramos de pinheiro** em **Go**.
- **Simulação de Partículas** (Python): versões **simplificada** e **complexa** com constantes físicas normalizadas, saída em GIF e CSV.
- **Erosão de Terreno** (Go): engine para simular processos erosivos em mapas de altura.

---

## 🚀 Tech Stack

| Componente                   | Linguagem | Descrição                                                                |
|------------------------------|-----------|--------------------------------------------------------------------------|
| Modelo Base                          | Go        | Orquestração, leitura de parâmetros e sequenciamento             |
| Vegetação (goiabeira, vitória-régia) | Python    | L-systems e geração de imagens de folhas                         |
| Vegetação (samambaia, pinheiro)      | Go      | Algoritmos de ramificação e estrutura botânica                     |
| Simulação de Partículas              | Python    | Movimento Browniano, geração de GIFs e registros em CSV          |
| Erosão de Terreno                     | Go        | Simulação de desgaste e sedimentação em mapas digitais          |

---

## 🔧 Pré-requisitos

- **Go** 1.18+
- **Python** 3.8+
  - Pacotes Python: `numpy`, `matplotlib`, `imageio`

---

## ⚙️ Instalação

Basta clonar o repositório:

```bash
git clone https://github.com/EduLoboM/Movimento-Browniano.git
````

---

## 🎯 Exemplos de Uso

* **Executar Modelo Base (Go):**

  ```bash
  cd Movimento-Browniano
  go run main.go
  ```

* **Gerar Vegetação:**

  * Goiabeira (Python):

    ```bash
    python vegetation_goiabeira.py --output output/goiabeira.png
    ```
  * Vitória-régia (Python):

    ```bash
    python vegetation_vitoria.py --output output/vitoria.png
    ```
  * Samambaia (Go):

    ```bash
    go run vegetation.go --type=samambaia --output output/samambaia.png
    ```
  * Ramos de pinheiro (Go):

    ```bash
    go run vegetation.go --type=pinheiro --output output/pinheiro.png
    ```

* **Simulação de Partículas (Python):**

  ```bash
  python particle_simple.py --steps 5000 --output output/simple.gif
  python particle_complex.py --steps 5000 --output output/complex.gif
  ```

* **Simulação de Erosão (Go):**

  ```bash
  go run erosion.go --map maps/terrain.dat --iterations 5000 --output output/erosion.png
  ```

---

## 👥 Integrantes

* **Eduardo Lobo**
* **Matheus Barreto**

---

## 🔗 Referências

* Einstein, A. (1905). "Über die molekularkinetische Theorie der Wärme geforderte Bewegung von in ruhenden Flüssigkeiten suspendierten Teilchen." *Annalen der Physik*.
* Mandelbrot, B. (1982). *The Fractal Geometry of Nature*. W. H. Freeman.
* Howard, A. (1997). *Números Aleatórios e Modelagem de Processos Físicos*.

```
```
