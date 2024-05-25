import common
import numpy as np
from matplotlib import pyplot as plt

common.set_custom_pyplot_styles()

data, args = common.parse_input()
x, y = np.loadtxt(data).T

plt.plot(x, y, "+")

plt.tight_layout()

common.annotate_plot(args)
common.save_plot(args)
