

class Car:
    def __init__(self, model, year, color, for_sale):
        self.model = model
        self.year = year
        self.color = color
        self.for_sale = for_sale

    def drive(self):
        print(f"you drive the {self.model}")

    def stop(self):
        print(f"you stop the {self.model}")

    def __str__(self):
        return f"Car({self.model}) - {self.year}"

 


car1 = Car("BMW", 1994, "Blue", True)

print(car1)