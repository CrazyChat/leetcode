package meituan

/*
小仓的射击练习4
时间限制：C/C++语言 1000MS；其他语言 3000MS
内存限制：C/C++语言 65536KB；其他语言 589824KB
题目描述：
小仓酷爱射击运动。今天的小仓想挑战自我。小仓有N颗子弹，接下来小仓每次会自由选择K颗子弹进行连续射击，全中靶心的概率为p[k]。如果成功小仓将获得a[k]的得分，并且可以使用余下子弹继续射击，否则今天的挑战结束。小仓想知道在最佳策略下，自己能得到的最高期望分数是多少。
输入
第一行一个数N，代表子弹数量。
第二行N个数p[]，第 i 个数代表p[i]。
第三行N个数a[]，第 i 个数代表a[i]。
1<=N<=5000   0<=p[i]<=1   0<=a[i]<=1000
输出
一个数表示最高期望得分，保留两位小数。

样例输入
2
0.80 0.50
1 2
样例输出
1.44
提示
样例1解释
选择用一颗子弹射击：如果命中则再用余下子弹射击（仅剩一颗子弹故只能选择一颗）：0.80 * 1 + 0.80 * 0.80 * 1= 1.44
选择用两颗子弹射击：0.5 * 2 = 1.00
此时最高期望得分为1.44
输入样例2
3
0.90 0.10 0.10
2 1 1
输出样例2
4.88
选择用一颗子弹射击：如果命中则再用一颗子弹进行射击，如果命中则再用一颗子弹进行射击（即3轮均使用了一颗子弹进行）：0.90 * 2 + 0.90 * 0.90 * 2+0.90 * 0.90 * 0.90 * 2= 4.878≈4.88  此种情况的期望得分最高，故为4.88
规则
请尽量在全场考试结束10分钟前调试程序，否则由于密集排队提交，可能查询不到编译结果
点击“调试”亦可保存代码
编程题可以使用本地编译器，此页面不记录跳出次数
 */

//func main() {
//	// 子弹数量
//	count := 0
//	// 连续k课概率p[k]
//	p := []float64{}
//	// 连续k科get分数
//	a := []float64{}
//
//}