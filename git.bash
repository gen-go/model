#!/bin/bash

git config  user.name "liuxuewen"
git config  user.email "liiuxuewen@gmail.com"
git status
git add .
if [[ ! -z $1 ]]; then
  git commit -m "$1"
  echo $1
else
  exit 1;
fi
git push


#
#
#tag=$(git describe --tags `git rev-list --tags --max-count=1`)
#
#echo -e "最新版本tag......"
#echo -e "tag"
#
#echo "最近版本："$tag
##echo $LatestTag
#
#noVTag=${tag:1}
#
#echo "noVTag："$noVTag
##按.分割字符串，变成3段
##按.分割字符串，变成3段
#array=(${noVTag//./ })
#arrayLen=${#array[*]}
#echo "array："$array
#
#
#if [[ $arrayLen != 3 ]];then
#  echo "版本号长度不对，必须为 v1.x.x的格式"
#  exit
#fi
#
##版本号处理
#newVersion="v"
#for (( i = 0; i < $arrayLen; i++ )); do
#    v=${array[i]}
#    if [[ $i == 2 ]];then
#      #第三位小版本+1
#      v=`expr $v + 1`
#    fi
#    newVersion+=$v"."
#done
#newVersion=${newVersion%?}
#
##拼接
#newTag=${tagPre}${newVersion}
#echo "新的版本："$newTag
#
#if git tag -l | grep -q $newTag;then
#    echo "已存在"
#    exit
#fi
#
##打标签
#git tag -a $newVersion -m ""
##推送到远端
#git push origin $newVersion --tags
#
##read -p "按任意键关闭" -n 1