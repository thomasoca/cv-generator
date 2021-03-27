#!/bin/bash

# Declare an array of string with type
declare -a StringArray=("pdfx" "biber" "bibhang" "biblabelsep" "pubtype" "bibsetup" "bibitemsep" )
 
# Iterate the string array using for loop
for val in ${StringArray[@]}; do
   sed -i "/\b$val\b/d" altacv.cls
done
