import numpy as np
import matplotlib.pyplot as plt


def frequency_heads(coin_filp_sequence):
    total_heads = len([head for head in coin_filp_sequence if head == 1])
    return total_heads / len(coin_filp_sequence) 


np.random.seed(0)
# print("Let's flip sthe biased coin once.")
# coin_flip = np.random.binomial(1, 0.7)
# print(f"Biased coin landed on {'heads' if coin_flip == 1 else 'tails'}")
#
# print("\nLet's flip the biased coin 10 time")
# number_coin_flips = 10
# head_count = np.random.binomial(number_coin_flips, .7)
# print(f"{head_count} heads were observed out of"
#       f"{number_coin_flips} biased coin flips")
#
coin_filps = []
frequencies = []

for _ in range(1000):
    coin_filps.append(np.random.randint(0, 2))
    frequencies.append(frequency_heads(coin_filps))


plt.plot(list(range(1000)), frequencies)
plt.axhline(0.5, color='red')
plt.xlabel('Number of Coin Filps')
plt.ylabel('Head-Frequency')
plt.show()
