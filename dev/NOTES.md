
docker build . -t tlvmbuilder
 
 
 
docker run -d -v /Users/sa/dev/g5/forks/topolvm:/work --name builder tlvmbuilder 

docker exec -it builder /bin/bash
docker exec -it builder /bin/bash -c "cd /work; /bin/bash"
 
docker exec -it builder /bin/bash -c "ls /"



git remote add upstream https://github.com/cybozu-go/topolvm.git
  