
struct A {
    int type; // type 是 Go 语言的关键字
    float i;
};

union B {
    int i;
    float f;
};

enum C {
    ONE,
    TWO,
};

void CSayHello(const char* s);
int Div(int a, int b);

void GoSayHello(const char* s);