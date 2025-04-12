rm -rf dist
bun run build
cd dist
git init
git checkout -b master
git add -A
git commit -m 'deploy'
git remote add github git@github.com:xushuhui/xushuhui.github.io.git
git push -f github master
cd -