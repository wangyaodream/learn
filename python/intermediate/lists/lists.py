list_org = ["banana", "cherry", "apple", [1,2,3,4]]

# 这里只是做了一个浅拷贝
list_cpy = list_org[:]

list_cpy[-1][-1] = 10
list_cpy.append("lemon")

print(list_org)
print(list_cpy)
