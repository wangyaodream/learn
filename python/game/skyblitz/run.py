import os

import pygame


pygame.display.set_caption("skyblitz pro")
WIDTH, HEIGHT = 900, 500
WIN = pygame.display.set_mode((WIDTH, HEIGHT))

BLACK = (0, 0, 0)

BORDER = pygame.Rect(WIDTH/2 - 5, 0, 10, HEIGHT)

FPS = 60
VEL = 5

SPACESHIP_WIDTH, SPACESHIP_HEIGHT = 55, 40

YELLOW_SPACESHIP_IMAGE = pygame.image.load(os.path.join('Assets', 'spaceship_yellow.png'))
RED_SPACESHIP_IMAGE = pygame.image.load(os.path.join('Assets', 'spaceship_red.png'))

YELLOW_SPACESHIP = pygame.transform.rotate(pygame.transform.scale(YELLOW_SPACESHIP_IMAGE, (SPACESHIP_WIDTH, SPACESHIP_HEIGHT)), 90)
RED_SPACESHIP = pygame.transform.rotate(pygame.transform.scale(RED_SPACESHIP_IMAGE, (SPACESHIP_WIDTH, SPACESHIP_HEIGHT)), 270)


def draw_window(red, yellow):
    """绘制的顺序会相互覆盖"""
    WIN.fill((255,255,255))
    pygame.draw.rect(WIN, BLACK, BORDER)

    WIN.blit(YELLOW_SPACESHIP, (yellow.x, yellow.y))
    WIN.blit(RED_SPACESHIP, (red.x, red.y))
    pygame.display.update()


def red_handle_movement(keys_pressed, red):
    if keys_pressed[pygame.K_a] and red.x - VEL > 0:
        red.x -= VEL
    if keys_pressed[pygame.K_d] and red.x + VEL + red.width < BORDER.x:
        red.x += VEL
    if keys_pressed[pygame.K_w] and red.y - VEL > 0:
        red.y -= VEL
    if keys_pressed[pygame.K_s] and red.y + VEL + red.height < HEIGHT - 15:
        red.y += VEL


def yellow_handle_movement(keys_pressed, yellow):
    if keys_pressed[pygame.K_LEFT] and yellow.x - VEL - yellow.width > BORDER.x - 15:
        yellow.x -= VEL
    if keys_pressed[pygame.K_RIGHT]:
        yellow.x += VEL
    if keys_pressed[pygame.K_UP]:
        yellow.y -= VEL
    if keys_pressed[pygame.K_DOWN]:
        yellow.y += VEL


def handle_movement(keys_pressed, red, yellow):
    pass


def main():
    red = pygame.Rect(100, 300, SPACESHIP_WIDTH, SPACESHIP_HEIGHT)
    yellow = pygame.Rect(700, 300, SPACESHIP_WIDTH, SPACESHIP_HEIGHT)

    clock = pygame.time.Clock()
    run = True
    while run:
        clock.tick(FPS)
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                run = False
        keys_pressed = pygame.key.get_pressed()
        red_handle_movement(keys_pressed, red)
        yellow_handle_movement(keys_pressed, yellow)
        draw_window(red, yellow)

    pygame.quit()


if __name__ == "__main__":
    main()
