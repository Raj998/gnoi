# Copyright 2021 Google LLC
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""Dependencies to build PINS infra."""

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository", "new_git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def gnoi_infra_deps():
    """Declare the third-party dependencies necessary to build GNOI infrastructure"""

    if not native.existing_rule("com_github_grpc_grpc"):
       http_archive(
            name = "com_github_grpc_grpc",
            # Move to newer commit to avoid the use of com_github_google_re2.
            url = "https://github.com/grpc/grpc/archive/565520443bdbda0b8ac28337a4904f3f20276305.zip",
            strip_prefix = "grpc-565520443bdbda0b8ac28337a4904f3f20276305",
            sha256 = "7206cc8e4511620fe70da7234bc98d6a4cd2eb226a7914f68fdbd991ffe38d34",
        )
    if not native.existing_rule("com_google_absl"):
        http_archive(
            name = "com_google_absl",
            url = "https://github.com/abseil/abseil-cpp/archive/refs/tags/20210324.rc1.tar.gz",
            strip_prefix = "abseil-cpp-20210324.rc1",
            sha256 = "7e0cf185ddd0459e8e55a9c51a548e859d98c0d7533de374bf038e4c7434f682",
        )
    if not native.existing_rule("com_google_googletest"):
        http_archive(
            name = "com_google_googletest",
            urls = ["https://github.com/google/googletest/archive/release-1.10.0.tar.gz"],
            strip_prefix = "googletest-release-1.10.0",
            sha256 = "9dc9157a9a1551ec7a7e43daea9a694a0bb5fb8bec81235d8a1e6ef64c716dcb",
        )
    if not native.existing_rule("com_google_protobuf"):
        http_archive(
            name = "com_google_protobuf",
            url = "https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/protobuf-all-3.14.0.tar.gz",
            strip_prefix = "protobuf-3.14.0",
            sha256 = "6dd0f6b20094910fbb7f1f7908688df01af2d4f6c5c21331b9f636048674aebf",
        )
    if not native.existing_rule("com_googlesource_code_re2"):
        http_archive(
            name = "com_googlesource_code_re2",
            # Newest commit on "absl" branch as of 2021-03-25.
            url = "https://github.com/google/re2/archive/72f110e82ccf3a9ae1c9418bfb447c3ba1cf95c2.zip",
            strip_prefix = "re2-72f110e82ccf3a9ae1c9418bfb447c3ba1cf95c2",
            sha256 = "146bf2e8796317843106a90543356c1baa4b48236a572e39971b839172f6270e",
        )
    if not native.existing_rule("com_google_googleapis"):
        http_archive(
            name = "com_google_googleapis",
            url = "https://github.com/googleapis/googleapis/archive/f405c718d60484124808adb7fb5963974d654bb4.zip",
            strip_prefix = "googleapis-f405c718d60484124808adb7fb5963974d654bb4",
            sha256 = "406b64643eede84ce3e0821a1d01f66eaf6254e79cb9c4f53be9054551935e79",
        )
    if not native.existing_rule("rules_cc"):
        git_repository(
            name = "rules_cc",
            commit = "1477dbab59b401daa94acedbeaefe79bf9112167",
            remote = "https://github.com/bazelbuild/rules_cc.git",
            shallow_since = "1595949469 -0700",
        )
    if not native.existing_rule("rules_proto"):
        http_archive(
            name = "rules_proto",
            sha256 = "602e7161d9195e50246177e7c55b2f39950a9cf7366f74ed5f22fd45750cd208",
            strip_prefix = "rules_proto-97d8af4dc474595af3900dd85cb3a29ad28cc313",
            urls = [
                "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
                "https://github.com/bazelbuild/rules_proto/archive/97d8af4dc474595af3900dd85cb3a29ad28cc313.tar.gz",
            ],
        )