#/!bin/sh
### USAGE ###
# ./test.sh test.json [in]
# If you want to debug the input messages and want to see them nicely formatted
# just add "in" to the end of the test command:  
# ./test.sh test.json in
#
# To see the output of your service-core, just run the test command like this:
# ./test.sh test.json

export SERVICE_UUID=WORKER00-1285-4E4C-A44E-AAAABBBB0000 
export SERVICE_NAME=micro-worker-ffmpeg 
export SERVICE_HOST=worker00 
export NAMESPACE_LISTENER=flaneur 
export NAMESPACE_PUBLISHER=flaneur

TESTFILE=$1
IN=$2
case $IN in
  in) PROCESSOR="cat -" ;;
  *) PROCESSOR="/go/bin/worker-ffmpeg" ;;
esac

while read line
do
  case $line in
    sleep*) sleep `echo $line | tr -s ' ' | cut -d ' ' -f 2` ;;
    *) echo $line | CREATED_AT=`date -u +"%FT%T.000Z"` envsubst ;;
  esac
done < $TESTFILE | $PROCESSOR | jq '.'
