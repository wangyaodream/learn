

def calculate_pi(start: int, end: int) -> float:
    result = 0.0
    positive = True if start % 2 == 0 else False
    for i in range(start, end):
        tmp = 1.0 / (float(i * 2) + 1.0)
        if positive:
            result += tmp
        else:
            result -= tmp
        positive = not positive
    result = result * 4.0
    return result

