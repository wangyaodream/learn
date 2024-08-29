#define QUIT    5
#define HOTEL1  180.00    
#define HOTEL2  225.0
#define HOTEL3  255.0
#define HOTEL4  355.00
#define DISCOUNT  0.95
#define STARS "********************"


// 选择列表
int menu(void);

// 返回预定天数
int getnights(void);


// 计费
void showprice(double rate, int nights);


