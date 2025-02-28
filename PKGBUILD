pkgname=nyafetch
pkgver=1.0.0
pkgrel=1
pkgdesc="Simple neko fetch ~"
arch=('x86_64')
url="https://github.com/iwnuplynottyan/nyafetch"
license=('MIT')
depends=()
makedepends=('go' 'git')
source=("git+https://github.com/iwnuplynottyan/nyafetch.git")
sha256sums=('SKIP')

build() {
  cd "$srcdir/$pkgname"
  go build -o nyafetch main.go
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 nyafetch "$pkgdir/usr/bin/nyafetch"
}
