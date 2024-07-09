from datetime import datetime


class Greeter:
    def __init__(self, name):
        self.name = name

    # def _day(self):
    #     # 显示当前时间是星期几
    #     return datetime.now().strftime("%A")
    #
    # def _part_of_day(self):
    #     current_hour = datetime.now().hour
    #     if current_hour < 12:
    #         part_of_day = "morning"
    #     elif 12 <= current_hour< 17:
    #         part_of_day = "afternoon"
    #     else:
    #         part_of_day = "evening"
    #
    #     return part_of_day

    def greet(self, store):
        print(f"Hi, my name is {self.name}, and welcome to {store}")
        print(f"How's your {day()} {part_of_day()} going")
        print("Here's a coupon for 20% off!")


def day():
    # 显示当前时间是星期几
    return datetime.now().strftime("%A")

def part_of_day():
    current_hour = datetime.now().hour
    if current_hour < 12:
        part_of_day = "morning"
    elif 12 <= current_hour< 17:
        part_of_day = "afternoon"
    else:
        part_of_day = "evening"

    return part_of_day


if __name__ == "__main__":
    greeter = Greeter("wangyao")
    greeter.greet("Yonghui")
