#!/bin/bash
for i in {1..1000000}
do
   curl http://35.167.149.199:30668/ &
done
