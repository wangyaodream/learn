import numpy as np


red_cards = 26 * [1]
black_cards = 26 * [0]
unshuffled_deck = red_cards + black_cards

np.random.seed(1)
shuffed_deck = np.random.permutation(unshuffled_deck)

remaining_red_cards = 26
for i, card in enumerate(shuffed_deck[:-1]):
    remaining_red_cards -= card
    remaining_total_cards = 52 - i - 1
    if remaining_red_cards / remaining_total_cards > 0.5:
        print(i)
        break


print(f"Stopping the game at index {i}")
final_card = shuffed_deck[i + 1]
print(f"The next card in the deck is {'red' if final_card else 'black'}")
print(f"We have {'Won' if final_card else 'lost'}")

