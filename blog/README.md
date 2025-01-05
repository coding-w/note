在日常程序开发中，排序是经常使用的算法，这里总结了一些常见的排序算法，以及他们的主要原理和思想，这里使用的案例是比较简单的，如果遇到复杂的排序，需要微调。所有排序算法均为从小到大排序

### 冒泡排序（Bubble Sort）

+ **原理：**通过重复遍历待排序的序列，依次比较相邻的元素，如果前一个元素比后一个元素大，则交换两个元素的位置。每一轮遍历会将未排序部分的最大元素“冒泡”到序列的末尾，每次遍历都会确定一个元素的最终位置。
+ **时间复杂度**：$ O\left ( n^2 \right ) $
+ **空间复杂度**：$ O\left ( 1 \right ) $
+ **实现**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 冒泡排序
        System.out.println("bubbleSort = " + Arrays.toString(bubbleSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // bubbleSort 冒泡排序
    public static int[] bubbleSort(int[] nums) {
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = 1; j < nums.length - i; j++) {
                // 判断，交换
                if (nums[j] < nums[j - 1]) {
                    int temp = nums[j];
                    nums[j] = nums[j - 1];
                    nums[j - 1] = temp;
                }
            }
        }
        return nums;
    }
}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(bubbleSort([]int{3, 5, 4, 6, 2, 1}))
}

// bubbleSort 冒泡排序
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		swapped := false
		for j := 1; j < length-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return nums
}
```

### 选择排序（Selection Sort）

+ **原理：**通过重复遍历的序列，每次遍历未排序序列并找到最小值索引，与当前元素进行交换
+ **时间复杂度**：$ O\left ( n^2 \right ) $
+ **实现**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 选择排序
        System.out.println("selectionSort = " + Arrays.toString(selectionSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // selectionSort 选择排序
    public static int[] selectionSort(int[] nums) {
        for (int i = 0; i < nums.length - 1; i++) {
            int minIndex = i;
            for (int j = i + 1; j < nums.length; j++) {
                // 获取最小元素索引
                if (nums[j] < nums[minIndex]) {
                    minIndex = j;
                }
            }
            // 交换
            if (minIndex != i) {
                int temp = nums[minIndex];
                nums[minIndex] = nums[i];
                nums[i] = temp;
            }
        }
        return nums;
    }
}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(selectionSort([]int{3, 5, 4, 6, 2, 1}))
}

// selectionSort 选择排序
func selectionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		minI := i
		for j := i + 1; j < length; j++ {
			if nums[minI] > nums[j] {
				minI = j
			}
		}
		if minI != i {
			nums[minI], nums[i] = nums[i], nums[minI]
		}
	}
	return nums
}
```

### 插入排序（Insertion Sort）

+ **原理：**遍历序列，将第`i`个待排序元素插入到以排序序列中正确位置，如整理乱序的扑克牌
+ **时间复杂度：**$ O\left ( n^2 \right ) $
+ **原理解释**

![](https://cdn.nlark.com/yuque/0/2025/gif/39110424/1735972633538-96d1d1f2-c897-4e0c-8dce-9671d0a78dd5.gif)

+ **实现**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 插入排序
        System.out.println("insertionSort = " + Arrays.toString(insertionSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // insertionSort 插入排序
    public static int[] insertionSort(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            int temp = nums[i];
            int j = i - 1;
            while (j >= 0 && nums[j] > temp) {
                nums[j + 1] = nums[j];
                j--;
            }
            // 插入
            nums[j + 1] = temp;
        }
        return nums;
    }
}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(insertionSort([]int{3, 5, 4, 6, 2, 1}))
}

// insertionSort 插入排序
func insertionSort(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		temp := nums[i]
		j := i - 1
		// 将 nums[i] 插入到已排序部分
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = temp
	}
	return nums
}
```

### 希尔排序（Shell Sort）

+ **原理：**是插入排序的改进版，将待排序的序列按照一定的间隔进行插入排序，逐渐缩小间隔，直到最后间隔为1进行一次普通的插入排序
+ **时间复杂度：**$ O\left (n^{3/2} \sim n^2 \right ) $
+ **实现：**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 希尔排序
        System.out.println("shellSort = " + Arrays.toString(shellSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // shellSort 希尔排序
    public static int[] shellSort(int[] nums) {
        int gap = nums.length / 2;
        while (gap > 0) {
            for (int i = gap; i < nums.length; i++) {
                int temp = nums[i];
                int j = i - gap;
                while (j >= 0 && nums[j] > temp) {
                    nums[j + gap] = nums[j];
                    j -= gap;
                }
                nums[j + gap] = temp;
            }
            gap /= 2;
        }
        return nums;
    }
}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(shellSort([]int{3, 5, 4, 6, 2, 1}))
}

// shellSort 希尔排序
func shellSort(nums []int) []int {
	length := len(nums)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			j := i - gap
			// 将 nums[i] 插入到已排序部分
			for j >= 0 && nums[j] > temp {
				nums[j+gap] = nums[j]
				j = j - gap
			}
			nums[j+gap] = temp
		}
	}
	return nums
}
```

### 归并排序（Merge Sort）

+ **原理：**使用分治的思想，将数组拆分成两部分，然后分别排序再合并。使用递归均分一个数组，直到只有一个元素，通过合并达到排序的目的
+ **时间复杂度：**$ O\left ( n\log{n}  \right )  $
+ **实现：**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 归并排序
        System.out.println("mergeSort = " + Arrays.toString(mergeSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // mergeSort 归并排序
    public static int[] mergeSort(int[] nums) {
        if (nums.length < 2) {
            return nums;
        }
        int mid = nums.length / 2;
        int[] left = new int[mid];
        int[] right = new int[nums.length - mid];
        System.arraycopy(nums, 0, left, 0, left.length);
        System.arraycopy(nums, mid, right, 0, right.length);
        left = mergeSort(left);
        right = mergeSort(right);
        return merge(left, right);
    }

    // 合并
    public static int[] merge(int[] left, int[] right) {
        int[] merged = new int[left.length + right.length];
        int i = 0, j = 0, k = 0;
        while (i < left.length && j < right.length) {
            if (left[i] <= right[j]) {
                merged[k++] = left[i++];
            } else {
                merged[k++] = right[j++];
            }
        }
        while (i < left.length) {
            merged[k++] = left[i++];
        }
        while (j < right.length) {
            merged[k++] = right[j++];
        }
        return merged;
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(mergeSort([]int{3, 5, 4, 6, 2, 1}))
}

// mergeSort 归并排序
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

// merge 合并两个有序数组
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// 将剩余的元素添加到结果中
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	return merged
}
```

### 快速排序（Quick Sort）

+ **原理：**采用分治思想，选择一个元素为基准，将其他元素分为大于基准元素和小于基准元素两组，进行递归排序
+ **时间复杂度：**平均$ O\left ( n\log{n}  \right )  $，最坏$ O\left ( n^2 \right ) $
+ **原理图解：**

![](https://cdn.nlark.com/yuque/0/2025/gif/39110424/1735974131777-52cb5a5c-3ab1-4fe0-bc3d-ec1dabf900f0.gif)

+ **实现：**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 快速排序
        System.out.println("quickSort = " + Arrays.toString(quickSort(new int[]{3, 5, 4, 6, 1, 2}, 0, 5)));
    }

    // quickSort 快速排序
    public static int[] quickSort(int[] nums, int left, int right) {
        if (left < right) {
            int index = partition(nums, left, right);
            quickSort(nums, left, index - 1);
            quickSort(nums, index + 1, right);
        }
        return nums;
    }

    private static int partition(int[] nums, int left, int right) {
        int pivot = nums[left];
        int i = left;
        for (int j = left + 1; j <= right; j++) {
            if (nums[j] <= pivot) {
                i++;
                int temp = nums[j];
                nums[j] = nums[i];
                nums[i] = temp;
            }
        }
        int temp = nums[left];
        nums[left] = nums[i];
        nums[i] = temp;
        return i;
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(quickSort([]int{3, 5, 4, 6, 2, 1}))
}

// quickSort 快速排序
func quickSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	pivot := nums[0]
	var less []int
	var greater []int
	for i := 1; i < length; i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}
```

### 堆排序（Heap Sort）

+ **原理：**构建一个小顶堆，每次取出堆顶元素，将堆顶元素置换为末尾元素，进行堆下虑，保证小顶堆，重复此操作
+ **时间复杂度：**$ O\left ( n\log{n}  \right )  $
+ **实现：**

```java
import java.util.PriorityQueue;
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 堆排序
        System.out.println("heapSort = " + Arrays.toString(heapSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // heapSort 堆排序
    public static int[] heapSort(int[] nums) {
        PriorityQueue<Integer> heap = new PriorityQueue<>();
        for (int num : nums) {
            heap.add(num);
        }
        // 从堆中逐个取出元素，按顺序填充到数组
        for (int i = 0; i < nums.length; i++) {
            nums[i] = heap.poll(); // 取出堆顶元素并移除
        }
        return nums;
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(heapSort([]int{3, 5, 4, 6, 2, 1}))
}

// heapSort 堆排序
func heapSort(nums []int) []int {
	length := len(nums)

	// 构建最大堆
	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, length, i)
	}

	// 一个个从堆中取出元素
	for i := length - 1; i > 0; i-- {
		// 交换当前根节点与末尾元素
		nums[0], nums[i] = nums[i], nums[0]
		// 调整堆
		heapify(nums, i, 0)
	}

	return nums
}

// heapify 调整堆
func heapify(nums []int, heapSize, rootIndex int) {
	largest := rootIndex
	leftChild := 2*rootIndex + 1
	rightChild := 2*rootIndex + 2

	// 如果左子节点大于根节点
	if leftChild < heapSize && nums[leftChild] > nums[largest] {
		largest = leftChild
	}

	// 如果右子节点大于最大节点
	if rightChild < heapSize && nums[rightChild] > nums[largest] {
		largest = rightChild
	}

	// 如果最大节点不是根节点
	if largest != rootIndex {
		nums[rootIndex], nums[largest] = nums[largest], nums[rootIndex]
		// 递归地调整受影响的子树
		heapify(nums, heapSize, largest)
	}
}
```

### 计数排序（Counting Sort）

+ **原理：**统计每个元素出现的次数，然后根据次序生成排序数组。适用于范围数据
+ **时间复杂度：**$ O\left ( n +k  \right )  $****
+ **实现：**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 计数排序
        System.out.println("countingSort = " + Arrays.toString(countingSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // countingSort 计数排序
    public static int[] countingSort(int[] nums) {
        if (nums == null || nums.length == 0) {
            return nums;
        }
        int maxNum = nums[0], minNum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < minNum) {
                minNum = nums[i];
            }
            if (nums[i] > maxNum) {
                maxNum = nums[i];
            }
        }
        int[] count = new int[maxNum -minNum + 1];
        for (int i = 0; i < count.length; i++) {
            count[nums[i]-minNum]++;
        }
        int index = 0;
        for (int i = 0; i < count.length; i++) {
            while (count[i]-- > 0) {
                nums[index] = i+minNum;
                index++;
            }
        }
        return nums;
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(countingSort([]int{3, 5, 4, 6, 2, 1}))
}

// countingSort 计数排序
func countingSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 找到数组中的最大值和最小值
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	// 创建计数数组
	count := make([]int, maxNum-minNum+1)

	// 统计每个元素的出现次数
	for _, num := range nums {
		count[num-minNum]++
	}

	// 生成排序后的数组
	index := 0
	for i, c := range count {
		for c > 0 {
			nums[index] = i + minNum
			index++
			c--
		}
	}

	return nums
}
```

### 桶排序（Bucket Sort）

+ **原理：**将数据分到若干桶中，再对每个桶内的数据进行排序，最后将所有桶中数据合并
+ **时间复杂度：**$ O\left ( n +k  \right )  $****
+ **实现：**

```java
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 堆排序
        System.out.println("bucketSort = " + Arrays.toString(bucketSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // bucketSort 桶排序
    public static int[] bucketSort(int[] nums) {
        if (nums == null || nums.length <= 1) {
            return nums;
        }

        // 找出数组的最大值和最小值
        int maxNum = nums[0], minNum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < minNum) {
                minNum = nums[i];
            }
            if (nums[i] > maxNum) {
                maxNum = nums[i];
            }
        }

        // 创建桶，桶的数量可以自定义，通常取决于数据的分布
        int bucketCount = (maxNum - minNum) / nums.length + 1;  // 保证桶的数量适合数据分布
        ArrayList<Integer>[] buckets = new ArrayList[bucketCount];

        // 将元素分配到桶中
        for (int num : nums) {
            int index = (num - minNum) * (bucketCount - 1) / (maxNum - minNum);
            if (buckets[index] == null) {
                buckets[index] = new ArrayList<>();
            }
            buckets[index].add(num);
        }

        // 对每个桶内的元素进行排序，并合并所有桶中的元素
        List<Integer> list = new ArrayList<>();
        for (int i = 0; i < bucketCount; i++) {
            if (buckets[i] != null && !buckets[i].isEmpty()) {
                Collections.sort(buckets[i]);
                list.addAll(buckets[i]);
            }
        }

        // 将排序后的结果放回原数组
        for (int i = 0; i < nums.length; i++) {
            nums[i] = list.get(i);
        }

        return nums;
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(bucketSort([]int{3, 5, 4, 6, 2, 1}))
}

// bucketSort 桶排序
func bucketSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 找到数组中的最大值和最小值
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	// 创建桶
	bucketCount := len(nums)
	buckets := make([][]int, bucketCount)

	// 将元素分配到桶中
	for _, num := range nums {
		index := (num - minNum) * (bucketCount - 1) / (maxNum - minNum)
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶进行排序并合并结果
	sortedNums := make([]int, 0, len(nums))
	for _, bucket := range buckets {
		sortedNums = append(sortedNums, mergeSort(bucket)...)
	}

	return sortedNums
}
```

### 基数排序（Radix Sort）

+ **原理：**按位对数字进行排序，从最低有效位到最高有效位进行多轮排序（针对浮点数或者负数要进行调整），这里主要描述思想。
+ **时间复杂度：**$ O\left ( n k  \right )  $
+ **原理图解：**

![](https://cdn.nlark.com/yuque/0/2025/gif/39110424/1735975975484-87becdb7-7030-4dc7-8af7-cf9220c8af85.gif)

+ **实现：**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 基数排序
        System.out.println("radixSort = " + Arrays.toString(radixSort(new int[]{113,112,210,6})));
    }

    // radixSort 基数排序
    public static int[] radixSort(int[] nums) {
        if (nums == null || nums.length == 0) {
            return nums;
        }
        int maxNum = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (maxNum < nums[i]) {
                maxNum = nums[i];
            }
        }
        int exp = 1;
        while (maxNum / exp > 0) {
            nums = countingSortByDigit(nums, exp);
            exp *= 10;
        }
        return nums;
    }

    // countingSortByDigit 对数组的某一位进行计数排序
    private static int[] countingSortByDigit(int[] nums, int exp) {
        int[] count = new int[10];
        // 统计每一个数字出现的次数
        for (int num : nums) {
            int index = (num / exp) % 10;
            count[index]++;
        }

        // 计算每个数字的累积计数
        for (int i = 1; i < count.length; i++) {
            count[i] += count[i - 1];
        }
        int[] result = new int[nums.length];
        for (int i = nums.length - 1; i >= 0; i--) {
            int index = (nums[i] / exp) % 10;
            result[count[index] - 1] = nums[i];
            count[index]--;
        }
        return result;
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(radixSort([]int{3, 5, 4, 6, 2, 1}))
}

// radixSort 基数排序
func radixSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 找到数组中的最大值
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// 从个位开始，对每一位进行计数排序
	exp := 1
	for maxNum/exp > 0 {
		nums = countingSortByDigit(nums, exp)
		exp *= 10
	}

	return nums
}

// countingSortByDigit 对数组的某一位进行计数排序
func countingSortByDigit(nums []int, exp int) []int {
	output := make([]int, len(nums))
	count := make([]int, 10)

	// 统计每个数字出现的次数
	for _, num := range nums {
		index := (num / exp) % 10
		count[index]++
	}

	// 计算每个数字的累积计数
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
    // 修改后的 count[i] 表示的是有多少个元素的当前位小于或等于 i

	// 根据当前位的数字，将元素放入输出数组
	for i := len(nums) - 1; i >= 0; i-- {
		index := (nums[i] / exp) % 10
		output[count[index]-1] = nums[i]
		count[index]--
	}

	return output
}
```

### Timsort

+ **原理：**是一个复合排序，将一个序列分成多份，每份进行插入排序，然后再进行合并达到排序的效果
+ **时间复杂度：**$ O\left ( n\log{n}  \right )  $
+ **实现：**

```java
import java.util.Arrays;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // timSort
        System.out.println("timSort = " + Arrays.toString(timSort(new int[]{3, 5, 4, 6, 2, 1})));
    }

    // timSort
    public static int[] timSort(int[] nums) {
        int n = nums.length;
        final int run = 32;
        for (int i = 0; i < n; i+=run) {
            insertionSortRange(nums, i, Math.min(i+run, n));
        }
        // 合并
        for (int size = run; size < n; size*=2) {
            for (int left = 0; left < n; left+=2*size) {
                int mid = Math.min(n - 1, left + size - 1);
                int right = Math.min((left + 2 * size - 1), (n - 1));
                if (mid < right) {
                    merge(nums, left, mid, right);
                }
            }
        }
        return nums;
    }

    // 合并两个已排序的区间
    private static void merge(int[] arr, int left, int mid, int right) {
        // 如果两个区间已经是有序的，直接返回
        if (arr[mid] <= arr[mid + 1]) return;

        int len1 = mid - left + 1;
        int len2 = right - mid;

        // 临时数组
        int[] leftArray = new int[len1];
        int[] rightArray = new int[len2];

        System.arraycopy(arr, left, leftArray, 0, len1);
        System.arraycopy(arr, mid + 1, rightArray, 0, len2);

        int i = 0, j = 0, k = left;

        // 合并过程
        while (i < len1 && j < len2) {
            if (leftArray[i] <= rightArray[j]) {
                arr[k++] = leftArray[i++];
            } else {
                arr[k++] = rightArray[j++];
            }
        }

        // 如果左半部分剩余元素
        while (i < len1) {
            arr[k++] = leftArray[i++];
        }

        // 如果右半部分剩余元素
        while (j < len2) {
            arr[k++] = rightArray[j++];
        }
    }

    private static void insertionSortRange(int[] nums, int left, int right) {
        for (int i = left + 1; i < right; i++) {
            int temp = nums[i];
            int j = i-1;
            while (j >= left && nums[j] > temp) {
                nums[j+1] = nums[j--];
            }
            nums[j+1] = temp;
        }
    }

}

```

```go
package main

import "fmt"

// main.go golang 实现
func main() {
	fmt.Println(timSort([]int{3, 5, 4, 6, 2, 1}))
}

// timSort
func timSort(nums []int) []int {
	const run = 32
	n := len(nums)

	// 对每个小块使用插入排序
	for i := 0; i < n; i += run {
		insertionSortRange(nums, i, min(i+run, n))
	}

	// 合并块
	for size := run; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := min(left+size, n)
			right := min(left+2*size, n)
			if mid < right {
				mergeRange(nums, left, mid, right)
			}
		}
	}

	return nums
}

// insertionSortRange 对指定范围使用插入排序
func insertionSortRange(nums []int, left, right int) {
	for i := left + 1; i < right; i++ {
		temp := nums[i]
		j := i - 1
		for j >= left && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
}

// mergeRange 合并两个有序的子数组
func mergeRange(nums []int, left, mid, right int) {
	leftPart := append([]int(nil), nums[left:mid]...)
	rightPart := append([]int(nil), nums[mid:right]...)

	i, j, k := 0, 0, left
	for i < len(leftPart) && j < len(rightPart) {
		if leftPart[i] <= rightPart[j] {
			nums[k] = leftPart[i]
			i++
		} else {
			nums[k] = rightPart[j]
			j++
		}
		k++
	}

	for i < len(leftPart) {
		nums[k] = leftPart[i]
		i++
		k++
	}

	for j < len(rightPart) {
		nums[k] = rightPart[j]
		j++
		k++
	}
}
```

### 猴子排序（Bogo Sort）

+ **原理：**一直随机性排序，直到按顺序为止。猴子排序灵感来源于“猴子乱按键盘”这个类比：想象一只猴子随机地打字，直到它正确地写出一个有序的句子。作为一种有趣的排序算法，它展示了“暴力解法”的极限。
+ **时间复杂度：**$ O\left ( n!  \right )  $
+ **实现：**

```java
import java.util.Arrays;
import java.util.Random;

// Main.java java实现
public class Main {
    public static void main(String[] args) {
        // 基数排序
        System.out.println("bogoSort = " + Arrays.toString(bogoSort(new int[]{113,112,210,6})));
    }

    

    // bogoSort 猴子排序
    public static int[] bogoSort(int[] nums) {
        while (!isSorted(nums)) {
            shuffle(nums); // 打乱数组
        }
        return nums;
    }
    // 检查数组是否已排序
    public static boolean isSorted(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] > nums[i]) {
                return false;
            }
        }
        return true;
    }

    // 随机打乱数组
    public static void shuffle(int[] nums) {
        Random random = new Random();
        for (int i = 0; i < nums.length; i++) {
            int j = random.nextInt(nums.length); // 随机选择一个索引
            // 交换 arr[i] 和 arr[j] 的元素
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
        }
    }

}

```

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main.go golang 实现
func main() {
	fmt.Println(bogoSort([]int{3, 5, 4, 6, 2, 1}))
}

// bogoSort 猴子排序
func bogoSort(nums []int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for !isSorted(nums) {
		shuffle(nums)
	}

	return nums
}

// isSorted 检查数组是否有序
func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

// shuffle 随机打乱数组
func shuffle(nums []int) {
	for i := range nums {
		j := rand.Intn(len(nums))
		nums[i], nums[j] = nums[j], nums[i]
	}
}

```

### 总结

时间复杂度对比

| **排序算法**              | **最佳时间复杂度**       | **最差时间复杂度**     | **平均时间复杂度**       |
|-----------------------|-------------------|-----------------|-------------------|
| 冒泡排序 (Bubble Sort)    | $ O(n) $          | $ O(n^2) $      | $ O(n^2) $        |
| 选择排序 (Selection Sort) | $ O(n^2) $        | $ O(n^2) $      | $ O(n^2) $        |
| 插入排序 (Insertion Sort) | $ O(n) $          | $ O(n^2) $      | $ O(n^2) $        |
| 希尔排序 (Shell Sort)     | $ O(n \log^2 n) $ | $ O(n^2) $      | $ O(n \log^2 n) $ |
| 归并排序 (Merge Sort)     | $ O(n \log n) $   | $ O(n \log n) $ | $ O(n \log n) $   |
| 快速排序 (Quick Sort)     | $ O(n \log n) $   | $ O(n^2) $      | $ O(n \log n) $   |
| 堆排序 (Heap Sort)       | $ O(n \log n) $   | $ O(n \log n) $ | $ O(n \log n) $   |
| 计数排序 (Counting Sort)  | $ O(n + k) $      | $ O(n + k) $    | $ O(n + k) $      |
| 桶排序 (Bucket Sort)     | $ O(n + k) $      | $ O(n^2) $      | $ O(n + k) $      |
| 基数排序 (Radix Sort)     | $ O(nk) $         | $ O(nk) $       | $ O(nk) $         |
| Timsort               | $ O(n) $          | $ O(n \log n) $ | $ O(n \log n) $   |
| 猴子排序 (Bogo Sort)      | $ O(1) $          | $ O((n+1)!) $   | $ O((n+1)!) $     |


稳定性对比

| **排序算法**              | **稳定性** | **说明**                            |
|-----------------------|---------|-----------------------------------|
| 冒泡排序 (Bubble Sort)    | 稳定      | 相等元素不会交换顺序，保持原始顺序。                |
| 选择排序 (Selection Sort) | 不稳定     | 在交换最小值和当前元素时，可能会改变相等元素的相对顺序。      |
| 插入排序 (Insertion Sort) | 稳定      | 插入时不会改变相等元素的顺序。                   |
| 希尔排序 (Shell Sort)     | 不稳定     | 在分组排序时，相等元素可能跨组移动，导致相对顺序被打乱。      |
| 归并排序 (Merge Sort)     | 稳定      | 在合并两个有序子数组时，相等元素的相对顺序保持不变。        |
| 快速排序 (Quick Sort)     | 不稳定     | 在分区时，可能导致相等元素的相对顺序被改变。            |
| 堆排序 (Heap Sort)       | 不稳定     | 堆调整过程中，相等元素的相对顺序可能被打乱。            |
| 计数排序 (Counting Sort)  | 稳定      | 使用额外的空间记录元素出现的顺序，保证了稳定性。          |
| 桶排序 (Bucket Sort)     | 稳定 ？    | 如果每个桶内使用稳定排序算法，则是稳定的，否则**可能不稳定**。 |
| 基数排序 (Radix Sort)     | 稳定      | 基于每个位排序时，使用稳定排序算法（如计数排序），保证了稳定性。  |
| Timsort               | 稳定      | 基于归并排序和插入排序的优化，天然具有稳定性。           |
| 猴子排序 (Bogo Sort)      | 不确定     | 随机排序算法，顺序完全依赖随机结果，稳定性无法保证。        |


上述是对12种排序算法的原理介绍和实现，比较有意思的是猴子排序，随机性很强，并不实用。

下一章介绍`Java`和`Golang`默认使用的排序算法。

这篇的文章分享就到这里啦，<font style="color:rgb(77, 77, 77);">如果文章对你有帮助，点赞+收藏～～</font>

  

