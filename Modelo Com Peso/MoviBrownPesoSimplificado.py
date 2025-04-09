import numpy as np
import matplotlib.pyplot as plt
from matplotlib.animation import FuncAnimation
import random
import math
from typing import List

GRID_SIZE = 250
NUM_PARTICLES = 250
TIME_STEP = 1.0
BROWNIAN_STRENGTH = 1.5
COLLISION_DIST_FACTOR = 1.0
REPULSION_STRENGTH = 2.0
ATTRACTION_STRENGTH = 2.0

class Particle:

    def __init__(self, x: float, y: float, mass: float, charge: float):
        self.x = x
        self.y = y
        self.mass = mass
        self.charge = charge
        self.radius = math.sqrt(mass)

    def update_position(self) -> None:

        dx = random.uniform(-BROWNIAN_STRENGTH, BROWNIAN_STRENGTH)
        dy = random.uniform(-BROWNIAN_STRENGTH, BROWNIAN_STRENGTH)
        self.x += dx * TIME_STEP
        self.y += dy * TIME_STEP

        if self.x < 0:
            self.x = -self.x
        elif self.x > GRID_SIZE:
            self.x = 2 * GRID_SIZE - self.x

        if self.y < 0:
            self.y = -self.y
        elif self.y > GRID_SIZE:
            self.y = 2 * GRID_SIZE - self.y

def combine_particles(p1: Particle, p2: Particle) -> Particle:

    new_mass = p1.mass + p2.mass
    new_charge = p1.charge + p2.charge
    new_x = (p1.x * p1.mass + p2.x * p2.mass) / new_mass
    new_y = (p1.y * p1.mass + p2.y * p2.mass) / new_mass
    print("Fusão!")
    return Particle(new_x, new_y, new_mass, new_charge)

def repel_particles(p1: Particle, p2: Particle) -> None:

    dx = p1.x - p2.x
    dy = p1.y - p2.y
    dist = math.hypot(dx, dy)
    if dist == 0:
        return
    force = REPULSION_STRENGTH / dist
    nx, ny = dx / dist, dy / dist
    p1.x += nx * force * TIME_STEP
    p1.y += ny * force * TIME_STEP
    p2.x -= nx * force * TIME_STEP
    p2.y -= ny * force * TIME_STEP

def attract_particles(p1: Particle, p2: Particle) -> None:

    dx = p2.x - p1.x
    dy = p2.y - p1.y
    dist = math.hypot(dx, dy)
    if dist == 0:
        return  
    force = ATTRACTION_STRENGTH / dist  
    nx, ny = dx / dist, dy / dist
    p1.x += nx * force * TIME_STEP
    p1.y += ny * force * TIME_STEP
    p2.x -= nx * force * TIME_STEP
    p2.y -= ny * force * TIME_STEP

def initialize_particles(num_particles: int) -> List[Particle]:

    particles = []
    for _ in range(num_particles):
        x = random.uniform(0, GRID_SIZE)
        y = random.uniform(0, GRID_SIZE)
        mass = random.uniform(1, 5)
        charge = random.uniform(-5, 5)
        particles.append(Particle(x, y, mass, charge))
    return particles

particles = initialize_particles(NUM_PARTICLES)

fig, ax = plt.subplots(figsize=(6, 6))
scatter = ax.scatter([], [], s=[], c=[])

ax.set_xlim(0, GRID_SIZE)
ax.set_ylim(0, GRID_SIZE)
ax.set_title("Simulação de Partículas com Movimento Browniano")
ax.set_xlabel("X")
ax.set_ylabel("Y")

def get_particle_color(charge: float) -> tuple:

    intensity = min(1, abs(charge) / 5)
    return (0, 0, intensity) if charge < 0 else (intensity, 0, 0)

def update(frame: int):

    global particles

    for p in particles:
        p.update_position()

    indices = list(range(len(particles)))
    i = 0
    while i < len(indices):
        idx1 = indices[i]
        p1 = particles[idx1]
        j = i + 1
        while j < len(indices):
            idx2 = indices[j]
            if idx1 == idx2:
                j += 1
                continue

            p2 = particles[idx2]
            dx = p1.x - p2.x
            dy = p1.y - p2.y
            distance = math.hypot(dx, dy)
            interaction_distance = (p1.radius + p2.radius) * COLLISION_DIST_FACTOR

            if distance < interaction_distance * 3:
                if p1.charge * p2.charge < 0:
                    attract_particles(p1, p2)

            if distance < interaction_distance:
                if p1.charge * p2.charge < 0:
                    new_particle = combine_particles(p1, p2)
                    indices.pop(j)
                    indices.pop(i)
                    for index in sorted([idx1, idx2], reverse=True):
                        particles.pop(index)
                    particles.append(new_particle)
                    indices = list(range(len(particles)))
                    i = -1
                    break
                else:
                    repel_particles(p1, p2)
            j += 1
        i += 1

    xs = [p.x for p in particles]
    ys = [p.y for p in particles]
    sizes = [p.mass * 10 for p in particles]
    colors = [get_particle_color(p.charge) for p in particles]
    scatter.set_offsets(np.column_stack((xs, ys)))
    scatter.set_sizes(sizes)
    scatter.set_color(colors)
    ax.set_title(f"Simulação (passo {frame}, partículas: {len(particles)})")
    return scatter,

animation = FuncAnimation(fig, update, frames=200, interval=50, blit=True)

plt.show()
