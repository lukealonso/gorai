http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.8.1/rules_go-0.8.1.tar.gz",
    sha256 = "90bb270d0a92ed5c83558b2797346917c46547f6f7103e648941ecdb6b9d0e72",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
go_rules_dependencies()
go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.8/bazel-gazelle-0.8.tar.gz",
    sha256 = "e3dadf036c769d1f40603b86ae1f0f90d11837116022d9b06e4cd88cae786676",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()


go_repository(
    name = "com_github_golang_crypto",
    commit = "95a4943f35d008beabde8c11e5075a1b714e6419",
    importpath = "github.com/golang/crypto",
)

go_repository(
    name = "com_github_stretchr_testify",
    commit = "b91bfb9ebec76498946beb6af7c0230c7cc7ba6c",
    importpath = "github.com/stretchr/testify",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "ecdeabc65495df2dec95d7c4a4c3e021903035e5",
    importpath = "github.com/davecgh/go-spew",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    commit = "792786c7400a136282c1664665ae0a8db921c6c2",
    importpath = "github.com/pmezard/go-difflib",
)

go_repository(
    name = "com_github_lukealonso_ed25519",
    commit = "17f813b763327a448c682941924c7cdae3fd1b38",
    importpath = "github.com/lukealonso/ed25519",
)
