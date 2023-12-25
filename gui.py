import pygame
import chess
import chess.svg
import pyperclip
import sys

# Initialize Pygame
pygame.init()

# Constants
WIDTH, HEIGHT = 600, 600
BOARD_SIZE = 8
SQUARE_SIZE = WIDTH // BOARD_SIZE
FPS = 120
WHITE = (255, 255, 255)
BLACK = (120, 49, 46)
FONT_SIZE = 30

# Initialize screen
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Chess UI")

# Load font
font = pygame.font.Font(None, FONT_SIZE)

# Drag and drop variables
dragging = False
dragged_piece = None
start_square = None
offset_x, offset_y = 0, 0  # Offset to adjust the position of the dragged piece

def draw_board():
    for row in range(BOARD_SIZE):
        for col in range(BOARD_SIZE):
            color = WHITE if (row + col) % 2 == 0 else BLACK
            pygame.draw.rect(screen, color, (col * SQUARE_SIZE, row * SQUARE_SIZE, SQUARE_SIZE, SQUARE_SIZE))

def piece_image_filename(piece):
    color_prefix = "white_" if piece.color == chess.WHITE else "black_"
    piece_name = piece.symbol().lower()
    return f"images/{color_prefix}{piece_name}.png"

def copy_fen_to_clipboard(board):
    fen_string = board.fen().split(' ')[0]  # Split and take the first part (before the first space)
    pyperclip.copy(fen_string)
    print("FEN copied to clipboard:", fen_string)

def draw_pieces(board):
    for square in chess.SQUARES:
        piece = board.piece_at(square)
        if piece is not None:
            img = pygame.image.load(piece_image_filename(piece))
            img = pygame.transform.scale(img, (SQUARE_SIZE, SQUARE_SIZE))
            screen.blit(img, (chess.square_file(square) * SQUARE_SIZE, (7 - chess.square_rank(square)) * SQUARE_SIZE))

def get_square_from_coords(coords):
    col = coords[0] // SQUARE_SIZE
    row = 7 - (coords[1] // SQUARE_SIZE)  # Invert row for chess board
    return chess.square(col, row)

def main(fen_string):
    global dragging, dragged_piece, start_square, offset_x, offset_y

    # Initialize chess board with the provided FEN string
    board = chess.Board(fen_string)

    clock = pygame.time.Clock()

    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                return
            elif event.type == pygame.MOUSEBUTTONDOWN:
                start_square = get_square_from_coords(event.pos)
                piece = board.piece_at(start_square)
                if piece is not None:
                    dragging = True
                    dragged_piece = piece
                    offset_x, offset_y = event.pos[0] - chess.square_file(start_square) * SQUARE_SIZE, event.pos[1] - (7 - chess.square_rank(start_square)) * SQUARE_SIZE
            elif event.type == pygame.MOUSEBUTTONUP:
                if dragging:
                    end_square = get_square_from_coords(event.pos)
                    move_input = f"{chess.square_name(start_square)}{chess.square_name(end_square)}"
                    legal_moves = [move.uci() for move in board.legal_moves]
                    if move_input in legal_moves:
                        board.push_uci(move_input)
                        copy_fen_to_clipboard(board)
                        return

        screen.fill(WHITE)
        draw_board()
        draw_pieces(board)
        pygame.display.flip()
        clock.tick(FPS)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python gui.py <FEN_STRING>")
    else:
        initial_fen = sys.argv[1]
        main(initial_fen)
