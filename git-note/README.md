# Git
`Git` 是一个强大的分布式版本控制系统，适用于代码管理和团队协作。大家日常都会使用到`Git`，常用的`IDE`工具都已经集成了一些常用的命令，导致常用的`Git`命令记不住，这里总结了一些常用命令，未来会持续更新。  

Git 的工作流程围绕四个主要区域展开：工作区、暂存区、本地仓库和远程仓库

1. 工作区（Working Directory）进行代码编辑的地方，可以添加新文件、修改现有文件或删除文件。
2. 暂存区（Staging Area/Index）暂存区是一个临时存储区域，用于存放下次提交的快照。执行`git add`命令时，将工作区中的更改添加到暂存区。
3. 本地仓库（Local Repository）本地仓库是项目的历史记录，包含了所有已提交的更改。执行`git commit`命令时，Git 会将暂存区中的内容创建一个新的提交，并将其保存到本地仓库。
4. 远程仓库（Remote Repository）远程仓库是托管在互联网上或网络中的仓库，通常用于团队协作。


## 基础命令操作

命令中加入`--global`就是全局操作，需要在对应项目的目录下执行`git`相关命令操作

##### 配置Git信息  

全局配置使用者信息，用于提交工作时，展示个人信息
```shell
git config --global user.name "ideal"
git config --global user.email "example@ideal.com"
```

##### 初始化仓库（这个可能用到的情况比较少）  

```shell
git init
```

##### 克隆远程仓库  

克隆分为两种`HTTPS`和`SSH`克隆，这里的主要区别是  

| 方式	    | 认证方式	          | 适用场景	                   | 优势	               | 劣势                         |
|--------|----------------|-------------------------|-------------------|----------------------------|
| HTTPS	 | 使用用户名 + 密码	    | 适用于临时访问或不想配置 SSH 密钥的情况	 | 易用，无需额外配置 SSH 密钥	 | 每次操作可能需要输入凭据，或依赖 Git 凭据管理器 |
| SSH	   | 通过 SSH 密钥进行认证	 | 适用于长期开发和高安全性需求	         | 免密码推拉代码，安全性更高	    | 需要配置 SSH 密钥                |

```shell
## 生成密钥
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
git clone https://github.com/coding-w/git-test.git
git clone git@github.com:coding-w/git-test.git
```

##### 查看当前状态

```shell
wx@wxdeMacBook-Pro git-test % git status

On branch main
Your branch is up to date with 'origin/main'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
    new file:   a.txt

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
    modified:   .gitignore
    modified:   a.txt
```

##### 添加文件到暂存区

暂存区是一个临时的区域，用于保存你对文件的更改，可以将它理解为一个“准备提交”的区域  
```shell
# 添加一个指定文件至暂存区
git add 文件名
# 添加所有修改文件至暂存区
git add .
# 取消文件加入暂存区
git restore --staged 文件名
```

##### 提交更改

```shell
git commit -m "提交说明"
```

这里需要补充说明的是，提交说明要按照规范去写，常见规则有  

| 类型       | 类别          | 说明                                    |
|----------|-------------|---------------------------------------|
| feat     | Production  | 新增功能                                  |
| fix      | Production  | Bug 修复                                |
| pref     | Production  | 提高代码性能的变更                             |
| style    | Development | 用于修改代码的样式，例如调整缩进、空格、空行等               |
| refactor | Production  | 用于重构代码，例如修改代码结构、变量名、函数名等但不修改功能逻辑      |
| test     | Development | 用于修改测试用例，例如添加、删除、修改代码的测试用例等           |
| ci       | Development | 用于修改持续集成流程，例如修改 Travis、Jenkins 等工作流配置 |
| docs     | Development | 用于修改文档，例如修改 README 文件、API 文档等         |
| chore    | Development | 用于对非业务性代码进行修改，例如修改构建流程或者工具配置等         |
| build    | Development | 用于修改项目构建系统，例如修改依赖库、外部接口等              |

常见的写法是：`git commit -m "类型: 提交说明"`，例如：`git commit -m "feat: 完成了**需求"`、`git commit -m "fix: 修复了**BUG"`。具体提交规则，参照公司开发规范

已提交，但是未推送，需要取消提交  
1. `git reset --soft HEAD~1`软重置，只是撤销最后一次提交并保留在暂存区
2. `git reset --mixed HEAD~1`混合重置，撤销最后一次提交并将暂存区的文件一通撤销至工作区
3. `git reset --hard HEAD~1`硬重置，撤销最后一次提交的所有内容，**慎用**

已提交，已推送，需要**修改提交信息**
```shell
## 修改最近一次提交信息
git commit --amend -m "..."
## 推送修改后的提交
### --force-with-lease：比 --force 更安全，只有本地分支与远程分支匹配时才会强制推送，避免覆盖他人的提交
### --force：直接覆盖远程提交，不推荐用于多人协作的分支。
git push origin main --force-with-lease
```

##### 查看操作历史
```shell
git log
## 简洁日志
git log --oneline --graph --decorate --all
```

##### 拉取和推送分支
```shell
## 推送并设置推送到那一个分支
git push -u origin <branch_name>
## 直接推送
git push

## 拉取远程某一分支
git pull origin branch_name
```

## 分支相关操作

Git 分支管理是版本控制系统中非常重要的一个方面，它允许开发者在不同的分支上独立工作

##### 查看分支
```shell
## 查看本地分支，当前分支有星号（*）
git branch

## 查看所有分支，包括远程分支
git branch -a

## 查看远程分支所有信息
git fetch --all
```

##### 创建分支

```shell
## 在当前提交的基础上创建一个新分支
git branch branch_name
## 创建并切换到新分支
git checkout -b branch_name

## 重命名分支
git branch -m new_branch_name

## 查看分支详细信息
git branch -vv
```
在实际开发中，创建项目新分支，一般是在**代码托管平台上操作**的

##### 切换分支
```shell
## 切换分支
git checkout branch_name
git switch branch_name

## 在本地创建分支并拉取远程分支
git checkout -b 本地分支名 origin/远程分支名

```

##### 合并分支
```shell
## 切换到主分支，将 branch_name 合并至主分支
git merge branch_name
## 如有冲突，手动解决冲突
todo ....

## 使用其他命令合并
git rebase branch_name
```

`git merge`和`git rebase`都是`Git`中用于整合不同分支更改的命令，但它们的工作方式和产生的结果有所不同

###### Git Merge
将两个分支的更改合并成一个新的`合并提交`（merge commit），并且保留两者的历史。  

合并的结果是创建一个新的提交，该提交有`两个父提交`，即原本分支的最后一个提交和目标分支的最后一个提交。

###### Git Rebase
会将当前分支上的提交 “**移动**” 到目标分支的**最前面**，即将当前分支的每个提交“重放”在目标分支的最新提交上  

会改变分支的历史记录，因此它看起来像是当前分支的所有提交是在目标分支的基础上做的

###### 主要区别
| 特点       | Merge                   | Rebase                            |
|----------|-------------------------|-----------------------------------|
| 历史记录     | 保留了两个分支的历史，合并后会生成一个合并提交 | 会修改提交历史，将当前分支的提交“重放”到目标分支的前面      |
| 提交图      | 会生成合并提交，历史呈现出分支结构       | 历史呈现出线性结构，没有合并提交                  |
| 操作后结果    | 生成新的合并提交，并保留每个分支的历史     | 将当前分支的提交合并到目标分支的前面，结果看起来是从目标分支开始的 |
| 适用场景     | 团队协作，需要保留所有分支的提交记录      | 个人分支或者希望清理提交历史，避免合并提交             |
| 对公共分支的影响 | 没有改变历史，适合团队合作           | 改变历史，不推荐在公共分支上使用                  |

###### 什么时候使用merge，什么时候使用 rebase？
`merge`： 当希望保留分支的历史，尤其是多人协作时，merge 是更好的选择。
`rebase`： 当需要让历史保持线性，或者在个人分支上工作时，rebase 会让提交历史更简洁。


##### 删除分支
```shell
## 删除本地分支
git branch -d branch_name
## 强制删除本地分支
git branch -D branch_name
## 删除远程分支
git push origin --delete branch_name
```

##### 查看分支差异
```shell
## 查看两个分支之间的差异
git diff branch_name_a branch_name_b
## 查看分支日志差异
git log branch_name_a branch_name_b
```


## 版本回滚

##### 仅回滚本地代码（不影响历史）
如果只是想让本地代码回到某个版本，但不改变提交历史，可以使用`git checkout`或`git reset --hard`

###### 临时回滚（不影响提交历史）
```shell
git checkout <目标版本号>
```
让工作目录切换到目标版本，但不会修改分支指向。适用于查看历史版本，但不修改当前分支状态。

###### 回滚到指定版本并修改本地提交历史
```shell
## 本地回滚但保留代码
git reset --soft <目标版本号>
```
回滚到<目标版本号>，但 **保留代码和暂存区的内容**，仅撤销后续的提交记录。适用于想撤销提交但保留文件修改，可以重新提交

```shell
## 本地回滚并清除暂存区
git reset --mixed <目标版本号>
```
**撤销提交**，但代码仍然保留在工作区（未 add 但代码还在）。 **适用于**希望回滚提交，并重新`add`需要的文件

```shell
## 彻底回滚（删除提交）
git reset --hard <目标版本号>
```
完全回滚，让代码、暂存区、提交历史全部恢复到**<目标版本号>**，后续提交会被删除。**适用于**彻底恢复到某个版本，不想保留后续的更改

##### 回滚远程仓库(谨慎操作)
```shell
git reset --hard <目标版本号>
## --force 会覆盖远程历史，可能导致团队成员的代码不同步，慎用！
git push origin --force
## 推荐使用 --force-with-lease，以避免意外覆盖其他人的提交：
git push origin --force-with-lease

## 不会删除历史，而是创建一个新的提交来撤销目标版本之后的更改
git revert <目标版本号>
```

上述，是在开发中遇到的会用到一些命令，`Git`是一个强大的版本控制工具，命令很多，也有很复杂的，需要多积累，此篇文章会持续更新。如有不足，多加指点～～