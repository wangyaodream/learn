

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


if __name__ == "__main__":
    event_conditions = [is_heads_or_tails, is_heads, is_tails, is_neither]
    weighted_sample_space = {"Heads": 4, "Tails": 1}

    for event_condition in event_conditions:
        prob = compute_event_probability(event_condition, weighted_sample_space)
        name = event_condition.__name__
        print(f"Probability of event arising from '{name}' is {prob}")
