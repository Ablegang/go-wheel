# make master/message 可以用 . 代替空格
master/%:
# 需要指定当前系统的 sh
	@m=$* sh shell/master.sh

# make master
master:
	@make master/update

# make feature/logpkg_0623 f=master
# b 指新建的分支名 f 指基本分支，默认为 master，则从 master 切出新分支
feature/%:
	@b=$* sh shell/feature.sh

# make push/message
# 打包 push 当前分支
push/%:
	@m=$* sh shell/push.sh

# make rebase/master
rebase/%:
	@b=$* sh shell/rebase.sh

# make quotes
quotes:
	@go run pkg/commands/quotes.go

# make blog
blog/%:
	@m=$* sh shell/blog_master.sh