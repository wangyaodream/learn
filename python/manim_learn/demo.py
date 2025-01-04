import manim

class Try(manim.Scene):
    def construct(self):
        c = manim.Circle(fill_opacity=0.5)
        s = manim.Square(color=manim.YELLOW, fill_opacity=0.5)
        self.play(manim.FadeIn(c))
        self.wait()
        self.play(manim.ReplacementTransform(c, s))
        self.wait()
        self.play(manim.FadeOut(s))
        self.wait()

