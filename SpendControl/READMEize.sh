# Get pictures

EVILBOSS='<img src="https://upload.wikimedia.org/wikipedia/en/a/ac/Pointy-haired_Boss.png " width="20" />'
SIMON='<img src="https://pbs.twimg.com/profile_images/180727117/Simon_400x400.jpg" width="20" />'
export i=0 ; curl -s https://threadreaderapp.com/thread/1349673271638286336.html | egrep -o "data-src=.*\"" | sed 's/.*"\(.*\)"/\1/' | while read img
do
    curl -so $TMPDIR/bla $img && convert $TMPDIR/bla $i.png && export i=$((i+1)) 
done

export i=0; cat SCRATCH.md | while read line
do
  echo $line | egrep -q "^Image$" && export i=$(( i+1 ))
  echo $line | sed "s/^Image$/\!\[\]($i.png)/" | sed "s@^[Mm][Ee][ :]@$SIMON@" | sed "s@^[X][ :]@$EVILBOSS@"
  echo ""
done > README.md
