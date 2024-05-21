from pathlib import Path

import numpy as np
from matplotlib import pyplot as plt

from plot.common import parse_input, set_custom_pyplot_styles

set_custom_pyplot_styles()

data, args = parse_input()

steps, infected, susceptible = np.loadtxt(data).T

plt.plot(steps, infected, label="infected")
plt.plot(steps, susceptible, label="susceptible")

plt.title(args.title)
plt.xlabel(args.xlabel)
plt.ylabel(args.ylabel)

plt.tight_layout()
plt.legend()

if args.out:
    plt.savefig(Path(args.out).with_suffix(".pdf"), dpi=100)
else:
    plt.show()
