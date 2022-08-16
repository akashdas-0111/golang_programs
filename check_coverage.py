import os
import sys
from statistics import mean

coverage = {}


def get_existing_coverages():
    for path, directory, files in os.walk("internal"):
        for file in files:
            if file.endswith(".go") and not file.endswith("_test.go"):
                command = "go tool cover -func cover.out|grep " + file + "|awk '{print substr($3, 1, length(" \
                                                                     "$3)-1)}'>result.txt "
                os.system(command)
                coverage[file] = mean(map(float, (((open("result.txt").read()).rstrip("\n")).split("\n"))))


commit_flag = False


def check_new_coverages():
    os.system("go test -v -coverprofile cover.out ./...")
    for path, directory, files in os.walk("internal"):
        for file in files:
            if file.endswith(".go") and not file.endswith("_test.go"):
                command = "go tool cover -func cover.out|grep " + file + "|awk '{print substr($3, 1, length(" \
                                                                     "$3)-1)}'>result.txt "
                os.system(command)
                if coverage[file] < mean(map(float, (((open("result.txt").read()).rstrip("\n")).split("\n")))):
                    global commit_flag
                    commit_flag = True
                    coverage[file] = mean(map(float, (((open("result.txt").read()).rstrip("\n")).split("\n"))))
                elif coverage[file] == mean(map(float, (((open("result.txt").read()).rstrip("\n")).split("\n")))):
                    continue
                else:
                    sys.exit("Test Coverage value of file %s reduced from: %.2f to %.2f" % (
                    file, coverage[file], mean(map(float, (((open("result.txt").read()).rstrip("\n")).split("\n"))))))


def cover_commit():
    command1="git config --local user.name "+os.environ['USERNAME']
    os.system(command1)
    os.system('git add cover.out')
    os.system('git commit -m "Updated cover.out"')
    os.system('git push')
    print("Cover.out successfully updated")


get_existing_coverages()
print("Existing coverage")
for key, value in coverage.items():
    print("{0}->".format(key), "{0}%".format(value))
check_new_coverages()
print("New coverage")
for key, value in coverage.items():
    print("{0}->".format(key), "{0}%".format(value))
if os.environ['ACTION_TYPE']=="publish":  
    if commit_flag:
        print("Congratulation! test case enhanced")
        cover_commit()
        print("Updated Successfully")
    else:
        print("Nothing to Update")
print("Success")
print(os.environ['USERNAME'])
print(os.environ['GITHUB_CONTEXT'])