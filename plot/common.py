import argparse
import io
import sys

import scienceplots
from matplotlib import pyplot as plt


def parse_input():
    parser = argparse.ArgumentParser()

    parser.add_argument("--out")
    parser.add_argument("--title")
    parser.add_argument("--xlabel")
    parser.add_argument("--ylabel")

    args = parser.parse_args()

    input_data = sys.stdin.read()
    return io.StringIO(input_data), args


def annotate_plot(args):
    if args.title:
        plt.title(args.title)
    if args.xlabel:
        plt.xlabel(args.xlabel)
    if args.ylabel:
        plt.ylabel(args.ylabel)


def save_plot(args):
    if args.out:
        plt.savefig(args.out)
    else:
        plt.show()


def set_custom_pyplot_styles():
    plt.style.use(["science", "ieee"])

    # custom font
    plt.rcParams[
        "text.latex.preamble"
    ] += r"""\usepackage[T1]{fontenc}
\usepackage[T1]{fontenc}
\usepackage[utf8]{inputenc}
%\usepackage[sfdefault,scale=0.95]{FiraSans}
\usepackage[lf]{Baskervaldx} % lining figures
\usepackage[bigdelims,vvarbb]{newtxmath} % math italic letters from nimbus Roman
\usepackage[cal=boondoxo]{mathalfa} % mathcal from STIX, unslanted a bit
\renewcommand*\oldstylenums[1]{\textosf{#1}}"""

    plt.rcParams["figure.figsize"] = (5, 3.5)
    plt.rcParams["figure.dpi"] = 300

    plt.rcParams["axes.formatter.limits"] = -3, 3
    plt.rcParams["axes.grid"] = False
    plt.rcParams["axes.formatter.use_mathtext"] = True

    SMALL_SIZE = 10
    MEDIUM_SIZE = 12
    BIGGER_SIZE = 14

    plt.rc("font", size=SMALL_SIZE)  # controls default text sizes
    plt.rc("axes", titlesize=BIGGER_SIZE)  # fontsize of the axes title
    plt.rc("axes", labelsize=MEDIUM_SIZE)  # fontsize of the x and y labels
    plt.rc("xtick", labelsize=SMALL_SIZE)  # fontsize of the tick labels
    plt.rc("ytick", labelsize=SMALL_SIZE)  # fontsize of the tick labels
    plt.rc("legend", fontsize=SMALL_SIZE)  # legend fontsize
    plt.rc("figure", titlesize=BIGGER_SIZE)  # fontsize of the figure title

    colors = plt.rcParams["axes.prop_cycle"].by_key()["color"]
    markers = ["*", "1", "+", "2", ".", "3"]
    return colors, markers
