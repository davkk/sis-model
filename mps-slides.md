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

# 👥 💬 ℹ️
# Rozprzestrzenianie się informacji w sieci

Dawid Karpiński, 14.06.2024 r.

---

1. Cel projektu
2. Model SIS
3. Użyty algorytm
4. Prezentacja wyników

---

## 1. Cel projektu

- modyfikacja parametrów $\beta=\beta(k)$ i $\gamma=\gamma(k)$
- porównanie Barabasi-Albert vs. Erdos-Renyi

---

## 2. Model SIS

![h:400px](./figures/sis-diagram.png)

---

$$
\Large I(t+\Delta{t}) = I(t) + \beta S(t) I(t) \Delta{t} - \gamma I(t) \Delta{t}
$$

$$
\Large dI(t)/dt = \beta S I - \gamma I
$$

tempo rozprzestrzeniania się:

$$
\Large \lambda = \beta / \gamma
$$

---

## 3. Użyty algorytm

1. Inicjalizacja parametrów symulacji
2. Iteracja po każdym węźle w sieci
    - Wybór węzła $i$ posiadającego informację
    - Przekazanie informacji podatnym sąsiadom $j$ według $\beta$
    - Zapomnienie informacji przez $i$ według $\gamma$

---

## 4. Prezentacja wyników

---

$$
\Large \beta = 0.1 = \gamma
$$

![bg left w:90%](./figures/none/ba_infected_vs_step_beta=0.1_gamma=0.1-1.png)

---

$$
\Large \beta = 0.1
$$

$$
\Large \gamma = 0.8
$$

![bg left w:90%](./figures/none/ba_infected_vs_step_beta=0.1_gamma=0.8-1.png)

---

$$
\Large \beta = 0.8
$$

$$
\Large \gamma = 0.1
$$

![bg left w:90%](./figures/none/ba_infected_vs_step_beta=0.8_gamma=0.1-1.png)

---

## Wpływ stopnia "pacjenta zero"

---

![bg contain](./figures/none/ba_max_min_rnd-1.png)

---

## Modyfikacja parametrów

$$
\Large \beta(k_i) = \beta_0 \cdot k / k_{\max}
$$

$$
\Large \gamma(k_i) = \gamma_0 \cdot k / k_{\max}
$$

---

## Wpływ na próg epidemii

$$
\Large \beta_0 = 0.1 = \gamma_0
$$

---

$\large \beta, \gamma = \text{const}$

![bg w:90% left:66%](./figures/none/ba_infected_vs_ratio-1.png)

---

$$
\Large \beta=\beta(k)
$$

$$
\Large \gamma=\text{const}
$$

![bg w:90% left:66%](./figures/beta/ba_infected_vs_ratio-1.png)

---

$$
\Large \beta=\text{const}
$$

$$
\Large \gamma=\gamma(k)
$$

![bg w:90% left:66%](./figures/gamma/ba_infected_vs_ratio-1.png)

---

$$
\Large \beta=\beta(k)
$$

$$
\Large \gamma=\gamma(k)
$$

![bg w:90% left:66%](./figures/both/ba_infected_vs_ratio-1.png)

---

![bg contain](./figures/none/ba_infected_vs_ratio-1.png)
![bg contain](./figures/both/ba_infected_vs_ratio-1.png)

---

## Wpływ średniego stopnia

---

![bg contain](./figures/none/ba_threshold_vs_k-1.png)
![bg contain](./figures/none/ba_tmax_vs_k-1.png)

---

![bg contain](./figures/beta/ba_threshold_vs_k-1.png)
![bg contain](./figures/beta/ba_tmax_vs_k-1.png)


---

![bg contain](./figures/gamma/ba_threshold_vs_k-1.png)
![bg contain](./figures/gamma/ba_tmax_vs_k-1.png)

---

![bg contain](./figures/both/ba_threshold_vs_k-1.png)
![bg contain](./figures/both/ba_tmax_vs_k-1.png)

---

## BA vs. ER

---

![bg contain](./figures/none/er_threshold_vs_k-1.png)
![bg contain](./figures/both/er_threshold_vs_k-1.png)

---

![bg contain](./figures/beta/er_threshold_vs_k-1.png)
![bg contain](./figures/gamma/er_threshold_vs_k-1.png)

---

## Wnioski

- zamysł modyfikacji parametrów:
    - ufność źródła informacji $\propto$ stopień węzła
- BA:
    - $\beta(k)$ => wyższa wartość $\lambda_c$
    - $\gamma(k)$ => niższa wartość $\lambda_c$
    - $\beta(k)$ i $\gamma(k)$ => bliskie klasycznemu
- ER:
    - brak znaczących różnic

---

<!-- paginate: skip -->

# Dziękuję za uwagę

:)
