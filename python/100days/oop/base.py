from abc import ABCMeta, abstractmethod


class Employee(metaclass=ABCMeta):
    """ 员工抽象类 """


    def __init__(self, name):
        self.name = name


    @abstractmethod
    def get_salary(self):
        pass



class Manager(Employee):
    """ 部门经理，继承于员工抽象类（经理也是员工）"""


    def get_salary(self):
        return 15000.0


class Programmer(Employee):
    """ 程序员 """


    def __init__(self, name, working_hour=0):
        self.working_hour = working_hour
        super().__init__(name)


    def get_salary(self):
        return 200.0 * self.working_hour


class Salesman(Employee):
    """ 销售员 """

    def __init__(self, name, sales=0.0):
        self.sales = sales
        super().__init__(name)
    
    def get_salary(self):
        return 1800.0 + self.sales * 0.05


class EmployeeFactory:

    @staticmethod
    def create(emp_type, *args, **kwargs):
        """ 创建员工 """
        all_emp_types = {'M': Manager, 'P': Programmer, 'S': Salesman}
        cls = all_emp_types[emp_type.upper()]
        return cls(*args, **kwargs) if cls else None


def main():
    emps = [
            EmployeeFactory.create('M', 'caocoa'),
            EmployeeFactory.create('P', 'xunyu', 120),
            EmployeeFactory.create('P', 'guojia', 85),
            EmployeeFactory.create('S', 'dianwei', 123000)]
    for emp in emps:
        print(f'{emp.name}: {emp.get_salary():.2f}yuan')


if __name__ == "__main__":
    main()


