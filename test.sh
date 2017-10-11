#/!bin/sh
TESTFILE=$1
PROCESSOR="${@:2}"

while read line
do
  if [[ $line == sleep* ]] ; then
    sleep `echo $line | tr -s ' ' | cut -d ' ' -f 2`
  else
    echo $line | CREATED_AT=`date -u +"%FT%T.000Z"` envsubst
  fi
done <$TESTFILE | SERVICE_UUID=WORKER00-1285-4E4C-A44E-AAAABBBB0000 SERVICE_NAME=micro-worker-ffmpeg SERVICE_HOST=worker NAMESPACE_LISTENER=flaneur NAMESPACE_PUBLISHER=flaneur $PROCESSOR
