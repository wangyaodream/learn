import time
import tkinter
import tkinter.messagebox


def download():
    # 模拟下载任务需要花费的时间
    time.sleep(10)
    tkinter.messagebox.showinfo('提示', '下载完成！')


def show_about():
    tkinter.messagebox.showinfo('About', 'author: WangYao')


def main():
    top = tkinter.Tk()
    top.title('单线程')
    top.geometry('200x150')
    top.wm_attributes('-topmost',True)

    panel = tkinter.Frame(top)
    btn1 = tkinter.Button(panel, text='Download', command=download)
    btn1.pack(side='left')
    btn2 = tkinter.Button(panel, text='About', command=show_about)
    btn2.pack(side='right')
    panel.pack(side='bottom')

    tkinter.mainloop()


if __name__ == "__main__":
    main()

