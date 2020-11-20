#!/usr/bin/env bash
echo "go build -o gin-vue-admin main.go"
go build -o build/gin-vue-admin/gin-vue-admin main.go || exit 1
# 生成版本号
VERSION=1.0.0
version=$VERSION.$(date +%Y%m%d%H%M)-$(git branch -v | grep "^*" | awk '{print $2"-"$3} ' | tr / .)
if [ $(git  status | grep '^\s\+modified:' | wc -l) -ne 0 ]; then
    version=$version-M
fi
if [ -n "$JENKINS_URL" ]; then
        build_branch_tag=$(echo $BUILD_GIT_BRANCH_TAG | tr / .)
        version=$VERSION-$(date +%Y%m%d%H%M)-$build_branch_tag-$(git branch -v | grep "^*" | awk '{print $5} ' | tr -d ")" )
fi
echo $version > version

mkdir -p build/gin-vue-admin
cp version build/gin-vue-admin
cp config.yaml build/gin-vue-admin
cp -r ./resource build/gin-vue-admin
cd build || exit 1
tar -cvf gin-vue-admin.tar.gz gin-vue-admin/*

if [  -d "${HOME}/go-server" ]; then
        if [ "$IS_UNPACK" = true ]; then
            tar -xvf ./gin-vue-admin.tar.gz -C ${HOME}/go-server
        fi
fi
