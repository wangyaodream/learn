

class Student:

    class_year = 2024
    num_students = 0
    
    def __init__(self, name, age):
        self.name = name
        self.age = age
        Student.num_students += 1



student1 = Student('wangyao', 23)
student2 = Student('zhangman', 20)

# 两个不同的对象所使用的class_year是指向同一个内存地址
print("After:")
print(id(student1.class_year))
print(id(student2.class_year))


student1.class_year = 2999
student2.class_year = 2888

print(student1.class_year)
print(student2.class_year)

# student1修改后将变成私有值,所指向的地址也不会变化

print("Before:")
print(id(student1.class_year))
print(id(student2.class_year))

# 但是类本身的class_year还依然存在，所指向的地址还是最初的地址
print(Student.class_year)
print(id(Student.class_year))

print(Student.num_students)
