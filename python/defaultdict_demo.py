from collections import defaultdict

def get_number_with_highest_count(counts):
    max_count = 0
    for number, count in counts.items():
        if count > max_count:
            max_count = count
            number_with_highest_count = number
    return number_with_highest_count

def most_frequent(numbers):
    counts = defaultdict(int)
    for number in numbers:
        counts[number] += 1

    return get_number_with_highest_count(counts)


a = [1,1,3,4,5,6,4,5,3,1,7,8,9,4,4,4,4,9]

res = most_frequent(a)
print(res)
