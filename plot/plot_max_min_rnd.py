import common
import pandas as pd
from matplotlib import pyplot as plt
from pandas.api.types import is_numeric_dtype


def mean_str(col):
    if is_numeric_dtype(col):
        return col.mean()
    else:
        return col.unique()


common.set_custom_pyplot_styles()

data, args = common.parse_input()

df = pd.read_csv(data, sep="\\s+", names=["init", "ratio", "infected"])
df = df.groupby("init").agg(list).reset_index()

legend = dict(max="maksymalny $k$", min="minimalny $k$", rnd="losowy $k$")

for idx in range(len(df)):
    plt.plot(
        df.ratio[idx],
        df.infected[idx],
        "+-",
        label=legend[df.init[idx]],
        linewidth=0.5,
    )

plt.legend()
plt.tight_layout()

common.annotate_plot(args)
common.save_plot(args)
