import pygame
import time
import random

pygame.font.init()

WIDTH, HEIGHT = 1000, 800
WIN = pygame.display.set_mode((WIDTH, HEIGHT))

FPS = 60

# title
pygame.display.set_caption("Sky Blitz")

# BG = pygame.image.load("bg.jpeg")
# resize bg 按照指定大小缩放而不是裁切
BG = pygame.transform.scale(pygame.image.load("bg.jpeg"), (WIDTH, HEIGHT))

# player
PLAYER_WIDTH = 40
PLAYER_HEIGHT = 60
PLAYER_VEL = 5

FONT = pygame.font.SysFont("comicsans", 30)

def draw(player, elapsed_time):
    WIN.blit(BG, (0, 0))

    time_text = FONT.render(f"Time: {round(elapsed_time)}s", 1, "white")
    WIN.blit(time_text, (10, 10))

    pygame.draw.rect(WIN, (255, 0, 0), player)

    pygame.display.update()

def main():
    run = True

    # reactangle 前两个参数决定了在什么位置绘制player，后两个参数决定player的大小
    player = pygame.Rect(200, HEIGHT - PLAYER_HEIGHT, PLAYER_WIDTH, PLAYER_HEIGHT)
    clock = pygame.time.Clock()

    start_time = time.time()
    elapsed_time = 0
    while run:
        # FPS
        clock.tick(FPS)
        elapsed_time = time.time() - start_time

        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                run = False
                break

        keys = pygame.key.get_pressed()
        if keys[pygame.K_LEFT] and player.x - PLAYER_VEL >= 0:
            player.x -= PLAYER_VEL
        elif keys[pygame.K_RIGHT] and player.x + PLAYER_VEL + player.width <= WIDTH:
            player.x += PLAYER_VEL
        draw(player, elapsed_time)

    pygame.quit()


if __name__ == "__main__":
    main()
