## Leetcode

### 一.递归

#### 1.计算某个数是否为2的幂

- 使用二进制的方式来考虑

  ~~~c
  //8 -- 1000
  //7 -- 0111
  //有 7 & 8 = 0000
  //即: n & (n-1) == 0 
  bool isPowerOfTwo(int n) {
      return n>0 && (n&(n-1)) == 0;
  }
  
  //8 -- 1000
  //-8 -- 8取反后加1，符号位仍为1
  //11111...0111 + 1 = 11...1000
  //8 & -8 = 01000
  
  bool isPowerOfTwo(int n) {
      return n>0 && (n&(-n)) == n;
  }
  ~~~

- 一般的做法：除法

  ~~~c
  bool isPowerOfTwo(int n) {
     if (n==1){
         return true;
     }
     if (n==0){
         return false;
     }
     if (n%2==0){//可以被2整除
         return isPowerOfTwo(n/2);
     }else{
         return false;
     }
  }
  ~~~

  

_____

#### 2.判断回文链表

- 递归

  `先将尾节点找出然后再于头节点比较大小`

  ~~~c
  struct ListNode* front;
  
  bool Check(struct ListNode* n){
      if (n!=NULL){
          //首先进行这个if条件
          //会使n指向最后一个节点
          if (!Check(n->next)){
              return false;
          }
          if (n->val!=front->val){
              return false;
          }
          //根据栈的特点一个一个往下
          front = front->next;
      }
      return true;
  }
  
  bool Is(struct ListNode* head){
      front = head;
      return Check(head);
  }
  ~~~

#### 3.约瑟夫环

给定一个长度为 `n` 的序列，每次向后数 `m` 个元素并删除，那么最终留下的是第几个元素？

~~~c++
class Solution{
    //计算n个数时，留下的元素序号
    int f(int n,int m){
        if (n==1){
            return 0;
        }
        int x = f(n - 1, m);
        return (m + x) % n;
    }
public:
    int lastRemaining(int n, int m) {
        return f(n, m);
    }
};

~~~



### 二.旋转数组

- 将后面k个元素替换到最前面

- ```
  nums = [1,2,3,4,5,6,7], k = 3
  输出: [5,6,7,1,2,3,4]
  ```

#### 1.数组反转

| 操作                     | 结果    |
| ------------------------ | ------- |
| 原始数组                 | 1234567 |
| 翻转所有元素             | 7654321 |
| 翻转[0,k /n-1]区间的元素 | 5674321 |
| 翻转[k/n,n-1]区间的元素  | 5671234 |

~~~c
// 交换两个值
void swap(int* a, int* b) {
    int t = *a;
    *a = *b, *b = t;
}

// 翻转数组
void reverse(int* nums, int start, int end) {
    while (start < end) {
        swap(&nums[start], &nums[end]);
        start += 1;
        end -= 1;
    }
}

void rotate(int* nums, int numsSize, int k) {
    // 这一步很恶心,当给出的数组长度小于k时,则会引起数组越界,且当k=numsize时,数组未发生改变.所以思考问题的时候得全面
    k %= numsSize;
    // 分三步走
    reverse(nums, 0, numsSize - 1);
    reverse(nums, 0, k - 1);
    reverse(nums, k, numsSize - 1);
}
~~~

#### 2.使用额外的数组

- 我们可以使用额外的数组来将每个元素放至正确的位置。用 n 表示数组的长度，我们遍历原数组，将原数组下标为 ii 的元素放至新数组下标为 (i+k)/n的位置，最后将新数组拷贝至原数组即可

~~~c
void rotate(int* nums, int numsSize, int k) {
    int newArr[numsSize];
    for (int i = 0; i < numsSize; ++i) {
        newArr[(i + k) % numsSize] = nums[i];
    }
    for (int i = 0; i < numsSize; ++i) {
        nums[i] = newArr[i];
    }
}
~~~

### 三.移动零

- 给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

```
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
```

#### 1.双指针法

- 使用双指针，左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部。

  右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移。

  注意到以下性质：

  左指针左边均为非零数；

  右指针左边直到左指针处均为零。

  因此每次交换，都是将左指针的零与右指针的非零数交换，且非零数的相对顺序并未改变。


~~~c
void swap(int* nums,int right,int left){
    int tmp = nums[right];
    nums[right] = nums[left];
    nums[left] = tmp;
}

void moveZeroes(int* nums, int numsSize){
    int left=0,right=0;
    while (right<numsSize){
        if (nums[right] != 0){
            swap(nums,right,left);
            left++;
        }
        right++;
    }
}
~~~

