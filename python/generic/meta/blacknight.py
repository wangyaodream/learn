class Blacknight:

    def __init__(self):
        self.members = ['an arm', 'another arm',
                        'a leg', 'another leg']
        self.phrase = ["'Tis but a scratch.",
                       "It's just a flesh wound.",
                       "I'm invincible!",
                       "All right,we'll call it a draw."]

        @property
        def member(self):
            print('next member is:')
            return self.members[0]

        @member.deleter
        def member(self):
            text = 'BLACK KNIGHT (loss {})\n-- {}'
            print(text.format(self.members.pop(0), self.phrase.pop(0)))


if __name__ == '__main__':
    knight = Blacknight()
    # print(knight.member)
    print(knight.__dict__)
