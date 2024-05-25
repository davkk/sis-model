import common
import numpy as np
from matplotlib import pyplot as plt
from pandas.api.types import is_numeric_dtype


def mean_str(col):
    if is_numeric_dtype(col):
        return col.mean()
    else:
        return col.unique()


common.set_custom_pyplot_styles()

data, args = common.parse_input()
init, ratio, infected = np.loadtxt(data).T

fig, axs = plt.subplots(nrows=3, ncols=1, sharex=True)

legend = dict(max="maksymalny $k$", min="minimalny $k$", rnd="losowy $k$")

for idx, ax in enumerate(axs):
    ax.plot(
        ratio,
        infected,
        "+",
        markersize=6,
        label="wynik symulacji",
    )
    ax.set_title(legend[np.unique(init)])
    ax.legend()

fig.supxlabel(args.xlabel)
fig.supylabel(args.ylabel)

fig.tight_layout()
common.save_plot(args)
