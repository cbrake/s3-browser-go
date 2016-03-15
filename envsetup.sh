
TOPDIR=`pwd`

export GOPATH=$TOPDIR
export PATH=$PATH:$GOPATH/bin

if [ -e local-env.sh ]; then
  source local-env.sh
fi

#if [ ! -e "bin" ]; then
#  echo "creating bin dir"
#  mkdir bin
#fi
#
#if [ ! -e "pkg" ]; then
#  echo "creating pkg dir"
#  mkdir pkg
#fi

function s3b_build() {
  go install github.com/cbrake/s3-browser-go
}

function s3b_build_arm() {
  GOARCH=arm GOARM=5 pw_build
}

function s3b_app() {
  s3b_build && s3-browser-go
}

function s3b_test() {
  go test github.com/cbrake/s3-browser-go/...
}

