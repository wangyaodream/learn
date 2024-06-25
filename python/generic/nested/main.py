

class Company:
    class Employee:
        def __init__(self, name, position):
            self.name = name
            self.position = position


        def get_details(self):
            return f"{self.name} {self.position}"

    
    def __init__(self, company_name):
        self.company_name = company_name
        self.emplpyees = []

    
    def add_employee(self, name, position):
        new_employee = self.Employee(name, position)
        self.emplpyees.append(new_employee)

    def list_employees(self):
        return [employee.get_details() for employee in self.emplpyees]


def make_multiplier_of(x):
    def multiplier(n):
        return n * x
    return multiplier


company = Company("Apple")
another_company = Company("Amazon")

company.add_employee("Wangyao", "Dev")
company.add_employee("Alice", "Manager")
company.add_employee("Bob", "Cashier")

another_company.add_employee("Maria", "worker")
another_company.add_employee("Loili", "CEO")



e_list = company.list_employees()
print(e_list)

print(another_company.list_employees())

times_three = make_multiplier_of(3)
print(times_three(4))
print(times_three(5))