import common
import numpy as np
from matplotlib import pyplot as plt

common.set_custom_pyplot_styles()

data, args = common.parse_input()
ratio, infected = np.loadtxt(data).T

plt.plot(
    ratio,
    infected,
    "+",
    label="wynik symulacji",
)

plt.xlabel(r"$\lambda$")
plt.ylabel(r"procent zarażonych")

plt.tight_layout()

common.annotate_plot(args)
common.save_plot(args)
