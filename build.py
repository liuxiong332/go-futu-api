# find ./futu-proto -type f -name '*proto' -exec protoc  -I ./futu-proto --go_out=./pb {} \;

import os
import shutil
import subprocess

os.makedirs("proto-temp", exist_ok=True)

for proto_name in os.listdir("./futu-proto"):
    file_path = os.path.join("./futu-proto", proto_name)
    with open(file_path) as f:
        new_content = f.read().replace("github.com/futuopen/ftapi4go/pb", "github.com/liuxiong332/go-futu-api/pb")
    
    with open(os.path.join("./proto-temp", proto_name), "w") as wf:
        wf.write(new_content)
    
for proto_name in os.listdir("./proto-temp"):
    file_path = os.path.join("./proto-temp", proto_name)
    res = subprocess.run(["protoc", "-I", "./proto-temp", "--go_out=./", file_path], capture_output=True, cwd="./", text=True)
    print(res.stderr)

shutil.move("github.com/liuxiong332/go-futu-api/pb", "pb")
