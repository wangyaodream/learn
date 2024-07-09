class SpeakMixin:
    def speak(self):
        name = self.__class__.__name__.lower()
        print(f"The {name} says, 'Hello!'")


class RollOverMixin:
    def roll_over(self):
        print('Did a barrel roll')


class Dog(SpeakMixin, RollOverMixin):
    # 狗子会说话会翻身
    pass
