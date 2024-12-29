EXE := ./main.out
CONFIGS := configs.txt
NOW := $(shell date +%s)

all: build run_classics

build:
	go build -o ${EXE} cmd/main.go

run_classics: ${EXE}
		# --progress \
		# --results output/${NOW}/classic_{} \
	parallel --lb \
		${EXE} -network={} -n=10000 -m=2 -p=0.0004 -init=max -beta=0.01 -gamma=0.01 \
		\| python ./plot/plot_si_vs_steps.py \
			--title "'Sieć {}, \$$N=10000$$, $$\langle k\rangle\approx4$$'" \
			--xlabel "'krok symulacji'" \
			--ylabel "'procent zarażonych'" \
		::: BA ER
