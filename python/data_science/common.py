from collections import defaultdict
from itertools import product

import matplotlib.pyplot as plt


def is_heads_or_tails(outcome):
    return outcome in {'Heads', 'Tails'}


def is_neither(outcome):
    return not is_heads_or_tails(outcome)


def is_heads(outcome):
    return outcome == 'Heads'


def is_tails(outcome):
    return outcome == 'Tails'


def get_matching_event(event_condition, sample_sapce):
    return set([outcome for outcome in sample_sapce if event_condition(outcome)])


def compute_event_probability(event_condition, generic_sample_space):

    event = get_matching_event(event_condition, generic_sample_space)
    if type(generic_sample_space) == type(set()):
        return len(event) / len(generic_sample_space)

    event_size = sum(generic_sample_space[outcome] for outcome in event)
    return event_size / sum(generic_sample_space.values())


def generate_coin_sample_space(num_flips=10):
    weighted_sample_space = defaultdict(int)
    for coin_filps in product(['Heads', 'Tails'], repeat=num_flips):
        heads_count = len([outcome for outcome in coin_filps 
                           if outcome == 'Heads'])
        weighted_sample_space[heads_count] += 1
    return weighted_sample_space


def is_in_interval(number, minimum, maximum):
    return minimum <= number <= maximum


if __name__ == "__main__":
    event_conditions = [is_heads_or_tails, is_heads, is_tails, is_neither]
    weighted_sample_space = generate_coin_sample_space()

    x_10_flips = list(weighted_sample_space.keys())
    sample_space_size = sum(weighted_sample_space.values())
    
    y_10_flips = [weighted_sample_space[key] for key in x_10_flips]
    prob_x_10_flips = [value / sample_space_size for value in y_10_flips]
    plt.plot(x_10_flips, prob_x_10_flips)
    plt.scatter(x_10_flips, prob_x_10_flips)
    # 对概率曲线下的区域进行着色
    where = [is_in_interval(value, 8, 10) for value in x_10_flips]
    plt.fill_between(x_10_flips, prob_x_10_flips, where=where)

    plt.xlabel('Head-count')
    plt.ylabel('Probability')
    plt.show()
