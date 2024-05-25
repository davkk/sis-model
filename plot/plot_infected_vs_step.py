import common
import numpy as np
from matplotlib import pyplot as plt

common.set_custom_pyplot_styles()

data, args = common.parse_input()

_, steps, infected = np.loadtxt(data).T

plt.plot(
    steps,
    infected,
    "+",
    label="zarażony",
    markersize=4,
)

plt.plot(
    steps,
    1 - infected,
    "+",
    label="podatny",
    markersize=4,
)

plt.xlabel("krok symulacji")
plt.ylabel("liczność stanu")

plt.xlim(0, 100)
plt.legend()

common.annotate_plot(args)
common.save_plot(args)
