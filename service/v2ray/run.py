# -*- coding: utf-8 -*-

"""
    Copyright (C) 2019-2020 KallyDev
    This program under GNU General Public License version 3.0, you
    can redistribute it or modify it under the terms of the, see
    the link below for more details

    https://github.com/kallydev/yogurt/blob/master/LICENSE
"""

import argparse
import json
import os


def load_config():
    with open('config.json', 'r') as file:
        data = json.load(file)
        file.close()
        return data


def save_config(data):
    with open('config.json', 'w') as file:
        json.dump(data, file, indent=2)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--port")
    args = parser.parse_args()
    print(args.port)

    conf = load_config()
    conf["inbounds"][0]["port"] = args.port
    save_config(conf)


if __name__ == '__main__':
    main()
    os.system("./v2ray --config=config.json --format=json")
