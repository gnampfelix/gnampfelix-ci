#!/bin/bash

/bin/bash /build/pre.sh
if [ $? -ne 0 ]; then
    exit 1
fi

/bin/bash /build/test.sh
if [ $? -ne 0 ]; then
    exit 2
fi

/bin/bash /build/post.sh
if [ $? -ne 0 ]; then
    exit 3
fi

exit 0
