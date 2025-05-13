import random
import math
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.animation as animation

def brownian_motion(n):

    grid = [[0 for _ in range(n)] for _ in range(n)]
    centro = n // 2
    grid[centro][centro] = 1
    pontos = [(centro, centro)]
    particle_info = []
    
    direcoes = [(0, 1), (0, -1), (1, 0), (-1, 0)]
    
    alvo = int(0.5 * n * n)
    
    while len(pontos) < alvo:

        theta = random.uniform(0, 2 * math.pi)
        a = 125
        b = 70
        x_start = centro + int(a * math.cos(theta))
        y_start = centro + int(b * math.sin(theta))
        
        steps = 0
        x, y = x_start, y_start
        attached = False
        
        while True:
            dx, dy = random.choice(direcoes)
            novo_x = x + dx
            novo_y = y + dy
            steps += 1

            if not (0 <= novo_x < n and 0 <= novo_y < n):
                break

            vizinhos = [(novo_x + dx, novo_y + dy) for dx, dy in direcoes]
            if any(0 <= vx < n and 0 <= vy < n and grid[vx][vy] == 1 
                   for vx, vy in vizinhos):
                grid[novo_x][novo_y] = 1
                pontos.append((novo_x, novo_y))
                particle_info.append((x_start, y_start, steps))
                attached = True
                break

            x, y = novo_x, novo_y

        if attached:
            continue

    return pontos, particle_info

def animacao_browniano(pontos, n):

    fig, ax = plt.subplots(figsize=(8, 8))
    ax.set_xlim(0, n - 1)
    ax.set_ylim(0, n - 1)
    ax.set_title("Movimento Browniano")
    ax.set_aspect('equal')

    scat = ax.scatter([], [], s=1, c='green', marker='.')

    passo = 5
    frames_indices = range(0, len(pontos), passo)

    def init():
        scat.set_offsets(np.empty((0, 2)))
        return scat,

    def update(frame_idx):
        dados = [(p[1], p[0]) for p in pontos[:frame_idx + 1]]
        scat.set_offsets(dados)
        return scat,

    ani = animation.FuncAnimation(
        fig,
        update,
        frames=frames_indices,
        init_func=init,
        blit=True,
        interval=20
    )

    ani.save("movimento_browniano_oval.gif", writer="pillow", fps=30)
    plt.close()

def plot_passos_from_file(arquivo="dados_browniano_oval.txt"):

    passos = []
    with open(arquivo, "r") as f:
        next(f)
        for line in f:
            campos = line.strip().split(",")
            if len(campos) == 3:
                passos.append(int(campos[2]))
    
    plt.rcParams.update({'font.size': 12})
    plt.figure(figsize=(12, 7))
    
    x_values = range(1, len(passos) + 1)
    
    plt.plot(x_values, passos, linestyle='-', linewidth=1.0, 
             color='darkblue', alpha=0.8, label='Linha de conexão')
    plt.scatter(x_values, passos, s=10, color='orange', alpha=0.7, 
                label='Pontos individuais')
    
    plt.xlim(min(x_values), max(x_values))
    y_min = min(passos)
    y_max = max(passos)
    margin = (y_max - y_min) * 0.1 if (y_max - y_min) != 0 else 1
    plt.ylim(y_min - margin, y_max + margin)
    
    plt.xlabel("Número do ponto", fontsize=14)
    plt.ylabel("Quantidade de movimentos", fontsize=14)
    plt.title("Movimentos para completar cada ponto", fontsize=16)
    
    plt.grid(True, linestyle='--', alpha=0.6)
    plt.legend()
    
    plt.savefig("grafico_passos_oval.png")
    plt.show()

if __name__ == "__main__":
    n = 250
    pontos, particle_info = brownian_motion(n)
    animacao_browniano(pontos, n)

    with open("dados_browniano_oval.txt", "w") as f:
        f.write("inicial_x,inicial_y,passos\n")
        for info in particle_info:
            x_start, y_start, steps = info
            f.write(f"{x_start},{y_start},{steps}\n")
    
    plot_passos_from_file("dados_browniano_oval.txt")
