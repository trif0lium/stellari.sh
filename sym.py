#!/usr/bin/python3
import os

base = os.path.dirname(__file__)

os.chdir(base)

apps = [
    "./apps/api-gateway",
    "./apps/konachan",
    "./apps/danbooru"
]

def sym(dst):
    os.chdir(base)
    os.chdir(dst)
    depth = len(dst.split("/")) - 1
    relative_path = "../" * depth + "apis"
    os.symlink(relative_path, "apis")

for app in apps:
    sym(app)
