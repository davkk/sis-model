# Model SIS - Rozprzestrzenianie się informacji w sieci

1. **Cel:**
    - symulacja rozprzestrzeniania się informacji w dwukierunkowych sieciach Barabasi-Albert (BA)
    - zbadanie wpływu węzłów o wysokim stopniu centralności na proces propagacji informacji w sieci BA
    - zbadanie wpływu modyfikacji parametrów na:
        - diagram fazowy
        - rozkład progu epidemii
        - rozkład czasu do osiągnięcia maksymalnej liczby zarażonych, $t_{\max}$
    - porównanie zmodyfikowanego modelu w sieci BA z siecią Erdos-Renyi (ER)

2. **Model:**
    - model SIS
    - agenci mogą znajdować się w jednym z dwóch stanów:
        - S - agent podatny na informację
        - I - agent, który przyjął informację
    - modyfikacje parametrów modelu ($k_i$ = stopień $i$-tego węzła):
        1. $\beta = \text{const}$ i $\gamma = \text{const}$ (klasyczny model)
        2. $\beta \propto k$
            - przekazanie informacji jednemu z dwóch wylosowanych agentów z prawdopodobieństwem $\beta(k_j)$
            $$S_i+I_j\xrightarrow{\beta(k_j)}I_i+I_j$$

        3. $\gamma \propto k$
            - utrata zainteresowania informacją poprzez spotkanie się dwóch zarażonych agentów i przywrócenie podatności na zarażenie jednego z nich, z prawdopodobieństwem $\gamma(k_i)$
            $$I_i+I_j\xrightarrow{\gamma(k_i)}S_i+I_j$$
        4. $\beta \propto k$ i $\gamma \propto k$
            - obie modyfikacje na raz

3. **Topologia:**
    - głównie sieć BA, porównanie z ER
    - sieć dwukierunkowa z asymetrycznymi wagami (założenie, że agenci o stosunkowo wyższym stopniu mają większy wpływ na swoich sąsiadów niż odwrotnie)

4. **Implementacja algorytmu**

5. **Analiza wyników symulacji**

