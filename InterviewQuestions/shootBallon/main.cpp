/*
【气球游戏】小Q在进行射击气球的游戏，如果小Q在连续T枪中打爆了所有颜色的气球，将得到一只QQ公仔作为奖励。（每种颜色的气球至少被打爆一只）。这个游戏中有m种不同颜色的气球，编号1到m。小Q一共有n发子弹，然后连续开了n枪。小Q想知道在这n枪中，打爆所有颜色的气球最少用了连续几枪？
输入描述：
第一行两个空格间隔的整数数n，m。n<=1000000 m<=2000
第二行一共n个空格间隔的整数，分别表示每一枪打中的气球的颜色,0表示没打中任何颜色的气球。
输出描述：
一个整数表示小Q打爆所有颜色气球用的最少枪数。如果小Q无法在这n枪打爆所有颜色的气球，则输出-1
示例
输入：
12 5
2 5 3 1 3 2 4 1 0 5 4 3
输出：
6
*/
#include <iostream>
#include <vector>
using namespace std;

bool checkAllHintGuns(int *guns, int kind)
{
    for (int i = 0; i < kind; i++)
    {
        if (guns[i] == 0)
        {
            return false;
        }
    }
    return true;
}

int findStartIndex(int end, int prestart)
{
    return -1;
}

// bullets子弹数
// kind
int minGuns(int bullets, int kind, int *targets)
{
    int *allHintGuns = new int[kind];
    vector<int> *kindVector = new vector<int>[kind]; // 记录每种颜色对应的索引
    int start = 0;
    int end = 0;
    int min = bullets;
    for (int i = 0; i < bullets; i++)
    {
        int k = targets[i];
        if (k < 0 || k > kind)
        {
            return -1;
        }

        if (k > 0)
        {
            allHintGuns[k] = 1;
            bool allhint = checkAllHintGuns(allHintGuns, kind);
            if (allhint)
            {
                end = i;
                start = findStartIndex(end, start);
                if (end - start < min){
                    min = end - start;
                }
            }
        }
    }

    delete []kindVector;
    delete []allHintGuns;
    if (min < bullets){
        return min;
    }
    return -1;
}

int main()
{
    int bullets = 0;
    int kind = 0;
    std::cin >> bullets >> kind;
    if (bullets <= 0)
    {
        return 1;
    }

    int *targets = new int[bullets];
    for (int i = 0; i < bullets; i++)
    {
        cin >> targets[i];
    }

    int num = minGuns(bullets, kind, targets);

    cout << num << endl;
    delete[] targets;
    return 0;
}
