#include <stdio.h>
#include <stdbool.h>

#define N 8  // 棋盘大小

int board[N];  // 用于存储皇后的位置

// 检查当前行和列是否安全
bool isSafe(int row, int col) {
    for (int i = 0; i < row; i++) {
        // 检查同一列和对角线
        if (board[i] == col || 
            board[i] - i == col - row || 
            board[i] + i == col + row) {
            return false;
        }
    }
    return true;
}

// 回溯函数
void solveNQueens(int row) {
    if (row == N) {
        // 找到一个解决方案，打印棋盘
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                if (board[i] == j) {
                    printf(" Q ");  // 打印皇后
                } else {
                    printf(" . ");  // 打印空位
                }
            }
            printf("\n");
        }
        printf("\n");
        return;
    }

    for (int col = 0; col < N; col++) {
        if (isSafe(row, col)) {
            board[row] = col;  // 放置皇后
            solveNQueens(row + 1);  // 递归到下一行
            // 回溯，尝试下一个位置
        }
    }
}

int main() {
    solveNQueens(0);  // 从第0行开始解决八皇后问题
    return 0;
}
