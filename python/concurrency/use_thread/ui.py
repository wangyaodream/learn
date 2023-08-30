from queue import Queue
from tkinter import Tk
from tkinter import Label, Entry, ttk
from typing import Optional
from stress_test import StressTest


class LoadTester(Tk):

    def __init__(self, loop, *args, **kwargs):
        Tk.__init__(self, *args, **kwargs)
        self._queue = Queue()
        self._refresh_ms = 25

        self._loop = loop
        self._load_test: Optional[StressTest] = None
        self.title('URL Requester')

        self._url_label = Label(self, text='URL:')
        self._url_label.grid(row=0, column=0)

        self._url_entry = Entry(self, width=10)
        self._url_entry.grid(row=1, column=0)

        self._requests_label = Label(self, text='Requests:')
        self._requests_label.grid(row=1, column=0)

        self._requests_field = Entry(self, width=10)
        self._requests_field.grid(row=1, column=1)

        self._submit = ttk.Button(self, text='Submit', command=self._start)
        self._submit.grid(row=1, column=2)

        self._pd_label = Label(self, text='Progress:')
        self._pd_label.grid(row=3, column=0)

        self._pb = ttk.Progressbar(self, orient='horizontal', length=200, mode='determinate')
        self._pb.grid(row=3, column=1, columnspan=2)
    
    def _update_bar(self, pct: int):
        if pct == 100:
            self._load_test = None
            self._submit['text'] = 'Submit'
        else:
            self._pb['value'] = pct
            self.after(self._refresh_ms, self._poll_queue)

    def _queue_update(self, completed_requests: int, total_requests: int):
        self._queue.put((completed_requests / total_requests) * 100)
        
    def _poll_queue(self):
        if not self._queue.empty():
            precent_complete = self._queue.get()

