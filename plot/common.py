import argparse
import io
import sys

import matplotlib.pyplot as plt


def parse_input():
    parser = argparse.ArgumentParser()

    parser.add_argument("--out")
    parser.add_argument("--title")
    parser.add_argument("--xlabel")
    parser.add_argument("--ylabel")

    args = parser.parse_args()

    input_data = sys.stdin.read()
    print(input_data)

    return io.StringIO(input_data), args


def set_custom_pyplot_styles():
    # custom pyplot styles
    SMALL_SIZE = 14
    MEDIUM_SIZE = 16
    BIGGER_SIZE = 22

    plt.rcParams["figure.figsize"] = (10, 7)
    plt.rcParams["figure.dpi"] = 200
    plt.rcParams["font.family"] = "serif"
    plt.rcParams["mathtext.fontset"] = "stix"
    plt.rcParams["xtick.direction"] = "in"
    plt.rcParams["ytick.direction"] = "in"

    plt.rcParams["axes.formatter.limits"] = -3, 3
    plt.rcParams["axes.grid"] = False
    plt.rcParams["grid.color"] = "gainsboro"
    plt.rcParams["axes.formatter.use_mathtext"] = True

    plt.rc("font", size=SMALL_SIZE)  # controls default text sizes
    plt.rc("axes", titlesize=SMALL_SIZE)  # fontsize of the axes title
    plt.rc("axes", labelsize=MEDIUM_SIZE)  # fontsize of the x and y labels
    plt.rc("xtick", labelsize=SMALL_SIZE)  # fontsize of the tick labels
    plt.rc("ytick", labelsize=SMALL_SIZE)  # fontsize of the tick labels
    plt.rc("legend", fontsize=SMALL_SIZE)  # legend fontsize
    plt.rc("figure", titlesize=BIGGER_SIZE)  # fontsize of the figure title

    colors = plt.rcParams["axes.prop_cycle"].by_key()["color"]
    markers = ["*", "1", "+", "2", ".", "3"]
    return colors, markers
