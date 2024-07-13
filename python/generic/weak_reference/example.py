import weakref

players = weakref.WeakValueDictionary()


class Player:
    def __init__(self):
        for i in range(1000):
            if i not in players:
                self.id = i
                break
        players[self.id] = self


def game():
    p1 = Player()
    p2 = Player()

    # Do something


for _ in range(2):
    game()
    print(dict(players))

