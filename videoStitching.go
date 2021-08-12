package main
/**
1024. 视频拼接
你将会获得一系列视频片段，这些片段来自于一项持续时长为 T 秒的体育赛事。这些片段可能有所重叠，也可能长度不一。

视频片段 clips[i] 都用区间进行表示：开始于 clips[i][0] 并于 clips[i][1] 结束。我们甚至可以对这些片段自由地再剪辑，例如片段 [0, 7] 可以剪切成 [0, 1] + [1, 3] + [3, 7] 三部分。

我们需要将这些片段进行再剪辑，并将剪辑后的内容拼接成覆盖整个运动过程的片段（[0, T]）。返回所需片段的最小数目，如果无法完成该任务，则返回 -1 。


输入：clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
输出：3
解释：
我们选中 [0,2], [8,10], [1,9] 这三个片段。
然后，按下面的方案重制比赛片段：
将 [1,9] 再剪辑为 [1,2] + [2,8] + [8,9] 。
现在我们手上有 [0,2] + [2,8] + [8,10]，而这些涵盖了整场比赛 [0, 10]。
 */


import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
	fmt.Println(videoStitching_v2([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
	fmt.Println(videoStitching([][]int{{5, 7}, {1, 8}, {0, 0}, {2, 3}, {4, 5}, {0, 6}, {5, 10}, {7, 10}}, 5))
	fmt.Println(videoStitching_v2([][]int{{5, 7}, {1, 8}, {0, 0}, {2, 3}, {4, 5}, {0, 6}, {5, 10}, {7, 10}}, 5))
	fmt.Println(videoStitching([][]int{{0, 2}, {4, 8}}, 5))
	fmt.Println(videoStitching_v2([][]int{{0, 2}, {4, 8}}, 5))
}

//dp算法，动态规划
func videoStitching(clips [][]int, time int) int {
	//按照开始时间从小到大排序
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	//初始化
	dp := make([]int, time+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0


	//遍历时间切片
	for i := 0; i < len(clips); i++ {
		start, end := clips[i][0], clips[i][1]
		if start > time || dp[start] == -1 { //必须把start>time放在前面，防止越界。start不可达就直接退出
			break
		}
		//为当前时间片的赋值，有两种情况，第一是首次赋值（dp[j]==-1），第二种是之前已经赋值的，那就要作比较（(dp[start]+1) < dp[j]，旧时间片的步骤和当前时间片的步骤[即从start处+1，代表衔接上一时间片]作比较）
		for j := start + 1; j <= end && j <= time; j++ {
			if (dp[start]+1) < dp[j] || dp[j] == -1 {
				dp[j] = dp[start] + 1
			}
		}
	}
	return dp[time]
}

//自己写的解法
//逻辑：从0开始获取，每个从上一个选择的时间片的最大跨度，并将其设置为当前最佳时间片
//例如：从0开始，选了了[0,4]为最佳时间片，那么接下来遍历时间片，计算4到该时间片的结尾，所具有的跨度。如果[2，6]时间片，那么该时间片的跨度就是2；如果[1，8]，那么该时间片跨度就是6，选择最大的跨度[1，8]为当前最佳时间片
func videoStitching_v2(clips [][]int, time int) int {
	//按照开始时间从小到大排序
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})

	//遍历时间切片
	space := 0		//上一个最佳时间片的end到当前时间片末尾的最大跨度
	curr_end := 0	//当前end
	dp := 0
	for i := 0; i < len(clips); i++ {
		start, end := clips[i][0], clips[i][1]
		if start > curr_end {		//如果开始时间超过了上一个最佳时间片的长度，就视为已经筛选出当前最佳时间片，进行拼接
			if space > 0 {
				curr_end += space
				space = 0
				dp += 1
				if curr_end >= time {		//衔接后如果超出了time，则return结束
					return dp
				}
			}else {		//开始时间大于上一个片段，但又没有衔接的部分，出现不可达
				return -1
			}
		}

		if start <= curr_end {		//判断是为当前最佳的时间片
			temp_space := end - curr_end
			if temp_space > space {
				space = temp_space
			}
		}
	}

	if space > 0 && curr_end < time {		//衔接最后一个时间片
		curr_end += space
		dp += 1
	}

	if curr_end < time {		//未满足时长
		return -1
	}

	return dp
}
