# go_hw_sep2020_repeat

* Download and setup git. 
* Create new repository named **git_homework**. 
* Configure git to use your name and email address.Please use your full name and valid email.
* Add main.go file that prints “Hello world”. Use preferred go version. Console output example:
```
$ bin/main
$ Hello world
```
* Add file to the **master** branch in commit named **initial_commit**. Provide meaningful commit description. For details on commit name and description format please see the [link](https://git-scm.com/docs/git-commit#_discussion).
Create new branch called **feature_owner_name** and switch to it. 
Add possibility to specify name as input string parameter.The specified name should be shown in greeting message returned by the application instead of word “world”. No need to capture  all possible errors, a simple code that works with correct input string (i.e Bob) is acceptable. 
```
$ bin/main
$ please enter your name: Bob 
$ Hello  Bob
```
* Commit the changes in a new commit named **owner_name**. Provide meaningful description. 
Add one more print message with a question of “How is the weather today?”. 
```
$ bin/main
$ please enter your name: Bob 
$ Hello  Bob
$ How is the weather today?
```
* Add these changes to the previous commit(**owner_name**) and modify the commit description accordingly.
Switch to **master** branch. Add another commit named **add_adjective** with changed **main.go** file that adds some const adjectives to the greeting message
```
$ bin/main
$ Hello brave new  world
```
* Merge **feature_owner_name** branch into **master**. Resolve conflicts to keep both code modifications. The application after merge should output something like following: 
```
$ bin/main
$ please enter your name: Bob 
$ Bob
$ Hello brave new  Bob
$ How is the weather today?
```
* Create github account if or use existing one.
* Make a fork from repository [git_homework](https://github.com/Nickolai-Belov/golang_cources_epam_12.19_git_homework)
* Add your fork as a remote to your repo **git_homework**.
* Set newly added remote as tracking branch to the local **master** branch.
* Rebase **master** to fork remote (**origin/master** by default). *Hint:  resulting master branch should include README from upstream repo*.
* Reset **master** to the state pior to the merge with **feature_owner_name**. *Hint: resulting master shouldn’t include any changes from feature_owner_name*.
* Switch to the **feature_owner_name**  branch.
* Rebase **feature_owner_name**  to the **master**.
* Add **.gitignore** file that ignores jpg files as a separate commit.
* Revert the last commit (the one with gitignore changes).
* Merge **feature_owner_name** branch back to **master**.
* Push master to the fork master.
* Send PR from fork master to the upstream [git_homework](https://github.com/Nickolai-Belov/golang_cources_epam_12.19_git_homework) repository master. Provide your name and email in PR description.
* Deadline is December, 5th.


