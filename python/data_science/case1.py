import numpy as np


red_cards = 26 * [1]
black_cards = 26 * [0]
unshuffled_deck = red_cards + black_cards

np.random.seed(0)
total_cards = 52
total_red_cards = 26

def execute_strategy(min_fraction_red=0.5, shuffled_deck=None, return_index=False):
    if shuffled_deck is None:
        shuffled_deck = np.random.permutation(unshuffled_deck)

    remaining_red_cards = total_red_cards

    for i, card in enumerate(shuffled_deck[:-1]):
        remaining_red_cards -= card 
        fraction_red_cards = remaining_red_cards / (total_red_cards - i - 1)
        if fraction_red_cards > min_fraction_red:
            break

    return (i + 1, shuffled_deck[i + 1]) if return_index else shuffled_deck[i + 1]


observations = np.array([execute_strategy() for _ in range(1000)])

frequency_wins = observations.sum() / 1000
assert frequency_wins == observations.mean()
print(frequency_wins)
