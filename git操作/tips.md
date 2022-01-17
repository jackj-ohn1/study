### 远程分支与本地分支的区别

> 远程分支其实就是**远程代码仓库当中的分支**，比如我们的repo如果是存在github的，那么这个远程仓库就是github，如果是存在gitlab的，那么这个仓库就是gitlab，如果是其他的仓库也是一样的道理。

> 当我们在使用git clone的时候，git会自动地将这个远程的repo命名为origin，拉取它所有的数据之后，创建一个指向它master的指针，命名为origin/master，**之后会在本地创建一个指向同样位置的指针**，命名为master，和远程的master作为区分。



git clone 连接（去GitHub上找）

克隆到本地之后，可以向这个文件夹中添加要上传的文件

git add . （提交所有文件）

git add 文件名               提交单个文件

git commit -m "注释"

git push origin 远程分支名 -f

以上就是从本地将代码上传到github的步骤



git checkout -b 新建本地分支名 (同时也会自动切换到这个分支上)

git branch -a （查看本地分支与远程分支）

git push origin 新建分支名:新建分支名

这时候再使用

git branch 分支名    就会发现github上新增了一个分支。 



使用：

git push origin  :远程分支名

来删除远程分支

git branch -d (本地分支名),不过得先不在这个本地分支上

来删除本地分支

#### 合并分支

git merge 分支名   （是将该分支合并到main/master分支上）

#### 合并冲突

当一个分支上的文件发生改动，然后主分支上的同名文件发生**了不同的改动**，**当把这个分支合并时，会出现合并冲突**。这个冲突只能自己手动更改内容进行解决

使用git diff 来查看分支文件中的冲突

1. 克隆代码

```text
git clone https://github.com/master-dev.git  
# 这个git路径是无效的，示例而已
```

2. 查看所有分支

```text
git branch --all  
# 默认只有master分支，所以会看到如下两个分支
# master[本地主分支] origin/master[远程主分支]
# 新克隆下来的代码默认master和origin/master是关联的，也就是他们的代码保持同步
```

3. 创建本地新的dev分支

```text
git branch dev  # 创建本地分支
git branch  # 查看分支
# 这是会看到master和dev，而且master上会有一个星号
# 这个时候dev是一个本地分支，远程仓库不知道它的存在
# 本地分支可以不同步到远程仓库，我们可以在dev开发，然后merge到master，使用master同步代码，当然也可以同步
```

4. 发布dev分支
   发布dev分支指的是同步dev分支的代码到远程服务器

```text
git push origin dev:dev  # 这样远程仓库也有一个dev分支了
```

5. 在dev分支开发代码

```text
git checkout dev  # 切换到dev分支进行开发
# 开发代码之后，我们有两个选择
# 第一个：如果功能开发完成了，可以合并主分支
git checkout master  # 切换到主分支
git merge dev  # 把dev分支的更改和master合并
git push  # 提交主分支代码远程
git checkout dev  # 切换到dev远程分支
git push  # 提交dev分支到远程
# 第二个：如果功能没有完成，可以直接推送
git push  # 提交到dev远程分支
# 注意：在分支切换之前最好先commit全部的改变，除非你真的知道自己在做什么
```

6. 删除分支

```text
git push origin :dev  # 删除远程dev分支，危险命令哦
# 下面两条是删除本地分支
git checkout master  # 切换到master分支
git branch -d dev  # 删除本地dev分
```
