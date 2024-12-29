---
marp: true
filetype: pdf
theme: uncover
class: invert
paginate: true
_paginate: false
style: |
    * {
        text-align: left;
        margin-left: 0 !important;
        margin-right: 0 !important;
        width: fit-content;
    }
    .MathJax {
        margin-top: 1em !important;
        margin-bottom: 1em !important;
    }
    pre {
        width: 100%;
    }
    img[alt~="center"] {
      display: block;
      margin: 0 auto;
    }
---

# 🦠 😷 👥
# Symulacja epidemii<br> w sieciach złożonych

Dawid Karpiński, 13.06.2024 r.

---

# Model SIS

---

![bg w:90%](./figures/sis-diagram.png)

---

$$
\Large I(t+\Delta{t}) = I(t) + \beta S(t) I(t) \Delta{t} - \gamma I(t) \Delta{t}
$$

$$
\Large dI(t)/dt = \beta S I - \gamma I
$$

tempo epidemii:

$$
\Large \lambda = \beta / \gamma
$$

---

# Sieć Barabasi-Albert

---

![bg left:40%](./figures/graph.png)

$$
\Large N, m
$$

$$
\large P(k)\sim k^{-3}
$$

- potęgowy rozkład węzłów
- preferential attachment

---

## Algorytm

1. Inicjalizacja sieci BA
2. Wybór zarażonego agenta $i$
3. Dla każdego niezarażonego sąsiada $j$:
    - zarażenie $j$ z prawdopodobieństwem $\beta$
4. Wyzdrowienie agenta $i$ z prawdopodobieństwem $\gamma$

---

# Wyniki

---

$$
\Large \beta = 0.1 = \gamma
$$

![bg left contain](./figures/none/ba_infected_vs_step_beta=0.1_gamma=0.1-1.png)

---

$$
\Large \beta = 0.1
$$

$$
\Large \gamma = 0.2, 0.4, 0.8
$$

![bg left vertical contain](./figures/none/ba_infected_vs_step_beta=0.1_gamma=0.2-1.png)
![bg contain](./figures/none/ba_infected_vs_step_beta=0.1_gamma=0.4-1.png)
![bg contain](./figures/none/ba_infected_vs_step_beta=0.1_gamma=0.8-1.png)

---

$$
\Large \gamma = 0.1
$$

$$
\Large \beta = 0.2, 0.4, 0.8
$$

![bg left vertical contain](./figures/none/ba_infected_vs_step_beta=0.2_gamma=0.1-1.png)
![bg contain](./figures/none/ba_infected_vs_step_beta=0.4_gamma=0.1-1.png)
![bg contain](./figures/none/ba_infected_vs_step_beta=0.8_gamma=0.1-1.png)

---

# Próg epidemii

---

dla sieci BA ($n=1000$, $m=2$)

$$
\LARGE \lambda_c = \braket{k}/\braket{k^2} \approx 0.1006962
$$

---

![bg contain](./figures/none/ba_infected_vs_ratio-1.png)

---

# Jak na epidemię wpływa średni stopień w sieci?

---

![bg contain](./figures/none/ba_threshold_vs_k-1.png)

---

![bg contain](./figures/none/ba_tmax_vs_k-1.png)

---

# Jaki wpływ na rozwój epidemii ma stopień pacjenta zero?

---

![bg contain](./figures/none/ba_max_min_rnd-1.png)

---

<!-- paginate: skip -->

# Dziękuję za uwagę

:)
