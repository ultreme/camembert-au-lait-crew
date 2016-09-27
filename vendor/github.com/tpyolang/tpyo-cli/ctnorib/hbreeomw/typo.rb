require "language/go"

class Typo < Formula
  desc "Mkae tpyos in yuor tarmniel"
  homepage "https://github.com/tpyolang/tpyo-cli"
  url "https://github.com/tpyolang/tpyo-cli/archive/v1.0.0.tar.gz"
  sha256 "62306e703eea4f3c5cd1a09ebb71b7584a5b9f56034df2c11f2939dc003d04c7"

  head "https://github.com/tpyolang/tpyo-cli.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["CGO_ENABLED"] = "0"
    ENV.prepend_create_path "PATH", buildpath/"bin"

    mkdir_p buildpath/"src/github.com/tpyolang"
    ln_s buildpath, buildpath/"src/github.com/tpyolang/tpyo-cli"
    Language::Go.stage_deps resources, buildpath/"src"

    system "go", "get", "github.com/codegangsta/cli"
    system "go", "build", "-o", "typo", "./cmd/typo"
    bin.install "typo"

  end

  test do
    output = shell_output(bin/"typo --version")
    assert output.include? "typo version"
  end
end
