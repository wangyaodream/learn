import time
import tkinter
from tkinter import ttk



window = tkinter.Tk()
window.title("Hello Tkinter")
window.geometry("300x200")


def say_hello():
    print("Hello there")


hello_button = ttk.Button(text="Say Hello", command=say_hello)
hello_button.pack()


window.mainloop()