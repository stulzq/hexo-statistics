# compile for version
set -e
make
if [ $? -ne 0 ]; then
    echo "make error"
    exit 1
fi

stat_version=`cat ./VERSION.txt`
echo "build version: $stat_version"

# cross_compiles
make -f ./Makefile.cross-compiles

rm -rf ./release/packages
mkdir -p ./release/packages

os_all='linux windows darwin'
#os_all='windows'
#arch_all='386'
arch_all='386 amd64 arm arm64'

cd ./release

for os in $os_all; do
    for arch in $arch_all; do
        stat_dir_name="hexo_statistics_${stat_version}_${os}_${arch}"
        stat_path="./packages/hexo_statistics_${stat_version}_${os}_${arch}"

        if [ "x${os}" = x"windows" ]; then
            if [ ! -f "./hexo_statistics_${os}_${arch}.exe" ]; then
                continue
            fi
            mkdir ${stat_path}
            mv ./hexo_statistics_${os}_${arch}.exe ${stat_path}/hexo_statistics.exe
        else
            if [ ! -f "./hexo_statistics_${os}_${arch}" ]; then
                continue
            fi
            mkdir ${stat_path}
            mv ./hexo_statistics_${os}_${arch} ${stat_path}/hexo_statistics
        fi
        cp ../LICENSE ${stat_path}
        cp -rf ../conf/* ${stat_path}

        # packages
        cd ./packages
        if [ "x${os}" = x"windows" ]; then
            zip -rq ${stat_dir_name}.zip ${stat_dir_name}
        else
            tar -zcf ${stat_dir_name}.tar.gz ${stat_dir_name}
        fi
        cd ..
        rm -rf ${stat_path}
    done
done

cd -