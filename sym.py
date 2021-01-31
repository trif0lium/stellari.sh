#!/usr/bin/python3
import os

base = os.path.dirname(__file__)

os.chdir(base)

apps = [
    "./apps/api-gateway",
    "./apps/konachan",
    "./apps/danbooru",
    "./apps/nhentai/src/main"
]

def sym(dst):
    os.chdir(base)
    os.chdir(dst)
    depth = len(dst.split("/")) - 1
    relative_path = "../" * depth + "apis"
    os.symlink(relative_path, "apis")

for app in apps:
    try:
        sym(app)
    except:
        pass
