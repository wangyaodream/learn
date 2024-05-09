import argparse

parser = argparse.ArgumentParser()
# 指定程序能接受哪些命令行选项
parser.add_argument("--echo", help="echo the striing you use here")
parser.add_argument("square", help="display a square of a given number",
                    type=int)
parser.add_argument("--verbose", help="increase output verbosity",
                    action="store_true")


args = parser.parse_args()
print(args.square**2)

# 没有foo选项
if args.echo:
    print("foo")

if args.verbose:
    print("verbosity turned on")

